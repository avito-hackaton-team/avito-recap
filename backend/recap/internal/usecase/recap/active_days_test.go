package recap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

func day(year int, month time.Month, dayOfMonth int, actions int64) entity.DayActivity {
	return entity.DayActivity{
		Date:    time.Date(year, month, dayOfMonth, 0, 0, 0, 0, time.UTC),
		Actions: actions,
	}
}

func TestActiveDaysSlideCarriesTheYearGrid(t *testing.T) {
	t.Parallel()

	input := slideInput{
		activity: entity.UserActivity{ActiveDays: 3},
		days: []entity.DayActivity{
			day(2026, time.January, 4, 7),
			day(2026, time.March, 14, 23),
			day(2026, time.December, 31, 2),
		},
	}

	built, ok := buildActiveDaysSlide(input)
	require.True(t, ok)

	slide, isActiveDays := built.(activeDaysSlide)
	require.True(t, isActiveDays)

	assert.Equal(t, int32(3), slide.ActiveDays)
	require.Len(t, slide.Days, 3)
	assert.Equal(t, "2026-01-04", slide.Days[0].Date)
	assert.Equal(t, int32(23), slide.Days[1].Actions)
	assert.Equal(t, "2026-12-31", slide.Days[2].Date)

	require.NotNil(t, slide.Peak)
	assert.Equal(t, "2026-03-14", slide.Peak.Date)
	assert.Equal(t, int32(23), slide.Peak.Actions)
}

func TestActiveDaysSlideSurvivesWithoutDailyBreakdown(t *testing.T) {
	t.Parallel()

	built, ok := buildActiveDaysSlide(slideInput{
		activity: entity.UserActivity{ActiveDays: 12},
	})
	require.True(t, ok)

	slide, isActiveDays := built.(activeDaysSlide)
	require.True(t, isActiveDays)

	assert.Equal(t, int32(12), slide.ActiveDays)
	assert.Nil(t, slide.Days)
	assert.Nil(t, slide.Peak)
}

func TestPeakDayPrefersTheEarlierDateOnATie(t *testing.T) {
	t.Parallel()

	days := dayActivityRefs([]entity.DayActivity{
		day(2026, time.February, 2, 9),
		day(2026, time.May, 5, 14),
		day(2026, time.August, 8, 14),
	})

	peak := peakDay(days)

	require.NotNil(t, peak)
	assert.Equal(t, "2026-05-05", peak.Date)
}

func TestDayActivityRefsDropDaysWithoutActions(t *testing.T) {
	t.Parallel()

	refs := dayActivityRefs([]entity.DayActivity{
		day(2026, time.June, 1, 0),
		day(2026, time.June, 2, 5),
		day(2026, time.June, 3, -1),
	})

	require.Len(t, refs, 1)
	assert.Equal(t, "2026-06-02", refs[0].Date)
	assert.Equal(t, int32(5), refs[0].Actions)
}
