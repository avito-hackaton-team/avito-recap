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
