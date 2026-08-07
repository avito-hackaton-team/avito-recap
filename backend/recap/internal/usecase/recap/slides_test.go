package recap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

func TestFinalStatsCollectTheNumbersOfTheYear(t *testing.T) {
	t.Parallel()

	input := slideInput{
		activity: entity.UserActivity{
			ActiveDays:       128,
			Views:            127,
			Favorites:        21,
			MessagesAsBuyer:  100,
			MessagesAsSeller: 28,
		},
		seasons: []seasonLeader{
			{season: entity.SeasonWinter},
			{season: entity.SeasonSpring},
			{season: entity.SeasonSummer},
			{season: entity.SeasonAutumn},
		},
	}

	stats := finalStats(input)

	expected := []struct {
		code  string
		value int64
	}{
		{statActiveDays, 128},
		{statViews, 127},
		{statFavorites, 21},
		{statMessages, 128},
		{statSeasons, 4},
	}

	require.Len(t, stats, len(expected))

	for i, want := range expected {
		assert.Equalf(t, want.code, stats[i].Code, "unexpected code at position %d", i)
		assert.Equalf(t, want.value, stats[i].Value, "unexpected value for %s", want.code)
		assert.NotEmptyf(t, stats[i].Label, "tile %s has no label", want.code)
	}
}

func TestFinalStatsSkipEmptyCounters(t *testing.T) {
	t.Parallel()

	stats := finalStats(slideInput{
		activity: entity.UserActivity{ActiveDays: 12, Views: 40},
	})

	require.Len(t, stats, 2)

	for _, tile := range stats {
		assert.NotZerof(t, tile.Value, "tile %s has nothing to show", tile.Code)
	}
}
