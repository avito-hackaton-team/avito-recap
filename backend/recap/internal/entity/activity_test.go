package entity_test

import (
	"testing"
	"time"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

func TestYearPeriodIsHalfOpen(t *testing.T) {
	t.Parallel()

	period := entity.YearPeriod(2025)

	if !period.From.Equal(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("unexpected start: %s", period.From)
	}

	if !period.To.Equal(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("unexpected end: %s", period.To)
	}
}

func TestPeriodValid(t *testing.T) {
	t.Parallel()

	from := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(1, 0, 0)

	tests := []struct {
		name   string
		period entity.Period
		want   bool
	}{
		{
			name:   "valid period",
			period: entity.Period{From: from, To: to},
			want:   true,
		},
		{
			name:   "equal boundaries",
			period: entity.Period{From: from, To: from},
			want:   false,
		},
		{
			name:   "reversed boundaries",
			period: entity.Period{From: to, To: from},
			want:   false,
		},
		{
			name:   "zero start",
			period: entity.Period{To: to},
			want:   false,
		},
		{
			name:   "zero end",
			period: entity.Period{From: from},
			want:   false,
		},
		{
			name: "same instant in different locations",
			period: entity.Period{
				From: from,
				To:   from.In(time.FixedZone("UTC+5", 5*60*60)),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.period.Valid(); got != tt.want {
				t.Errorf("Period.Valid() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestSeasonsCoverTheWholeYearOnce(t *testing.T) {
	t.Parallel()

	var days int

	for _, window := range entity.Seasons(2025) {
		for _, period := range window.Ranges {
			days += int(period.To.Sub(period.From).Hours() / 24)
		}
	}

	if days != 365 {
		t.Fatalf("seasons must cover 365 days of 2025 exactly once, got %d", days)
	}
}

func TestWinterIsSplitInTwoRanges(t *testing.T) {
	t.Parallel()

	windows := entity.Seasons(2025)

	if windows[0].Season != entity.SeasonWinter {
		t.Fatalf("expected winter first, got %s", windows[0].Season)
	}

	if len(windows[0].Ranges) != 2 {
		t.Fatalf("winter must hold January-February and December, got %d ranges", len(windows[0].Ranges))
	}

	if windows[0].Ranges[1].From.Month() != time.December {
		t.Errorf("expected the second winter range to start in December, got %s", windows[0].Ranges[1].From)
	}
}

func TestUserActivityTotalActionsExcludesDerivedMetrics(t *testing.T) {
	t.Parallel()

	activity := entity.UserActivity{
		ActiveDays:         100,
		Views:              1,
		UniqueListingsSeen: 200,
		Favorites:          2,
		Purchases:          3,
		Sales:              4,
		MessagesAsBuyer:    5,
		MessagesAsSeller:   6,
		CategoriesTouched:  300,
		ListingsCreated:    7,
	}

	const want int64 = 28
	if got := activity.TotalActions(); got != want {
		t.Errorf("UserActivity.TotalActions() = %d, want %d", got, want)
	}
}
