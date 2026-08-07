package recap

import (
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

func TestInterestsSummary(t *testing.T) {
	t.Parallel()

	electronicsScore := entity.CategoryScore{CategoryID: electronics, Title: "Электроника"}
	hobbyScore := entity.CategoryScore{CategoryID: hobby, Title: "Хобби"}

	t.Run("names every leading category when interests moved", func(t *testing.T) {
		t.Parallel()

		summary := interestsSummary([]seasonLeader{
			{season: entity.SeasonWinter, category: electronicsScore},
			{season: entity.SeasonSummer, category: hobbyScore},
		})

		assert.Contains(t, summary, "Электроника")
		assert.Contains(t, summary, "Хобби")
	})

	t.Run("names the single category when it never changed", func(t *testing.T) {
		t.Parallel()

		summary := interestsSummary([]seasonLeader{
			{season: entity.SeasonWinter, category: electronicsScore},
			{season: entity.SeasonSummer, category: electronicsScore},
		})

		assert.Contains(t, summary, "Электроника")

		// One category all year long must not be listed season by season.
		assert.NotContains(t, summary, "зимой")
		assert.NotContains(t, summary, "летом")
	})

	t.Run("stays empty without seasons", func(t *testing.T) {
		t.Parallel()

		assert.Empty(t, interestsSummary(nil))
	})
}

func TestInterestsSummaryStartsWithACapitalLetter(t *testing.T) {
	t.Parallel()

	summary := interestsSummary([]seasonLeader{
		{
			season:   entity.SeasonWinter,
			category: entity.CategoryScore{CategoryID: electronics, Title: "Электроника"},
		},
		{
			season:   entity.SeasonSummer,
			category: entity.CategoryScore{CategoryID: hobby, Title: "Хобби"},
		},
	})

	require.True(t, utf8.ValidString(summary), "summary is not valid UTF-8: %q", summary)

	first, _ := utf8.DecodeRuneInString(summary)
	assert.Truef(t, unicode.IsUpper(first), "summary must start with a capital letter: %q", summary)
}
