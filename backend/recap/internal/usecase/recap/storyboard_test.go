package recap

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

// Порядок экранов — продуктовое решение, а не деталь реализации: сначала «что
// вы смотрели», потом «что сделали», в конце выводы.
func TestStoryboardOrder(t *testing.T) {
	t.Parallel()

	category := entity.CategoryScore{
		CategoryID: uuid.New(),
		Title:      "Электроника",
		Score:      100,
	}

	raw, err := buildSlides(slideInput{
		year: 2025,
		activity: entity.UserActivity{
			ActiveDays:      128,
			Views:           2847,
			Favorites:       56,
			FavoritesActive: 12,
			Purchases:       23,
			Sales:           42,
			SalesAmount:     7_777_000,
			MessagesAsBuyer: 128,
		},
		categories: []entity.CategoryScore{category},
		seasons:    []seasonLeader{{season: entity.SeasonWinter, category: category, share: 40}},
		archetype: entity.Archetype{
			UserArchetype: entity.ArchetypeExplorer,
			Title:         "Исследователь",
			Description:   "Смотрит много и вширь.",
			Reasons: []entity.ArchetypeReason{
				{Metric: entity.MetricViews, Value: "2847", Explanation: "Просмотры — 90% действий"},
			},
		},
	})
	require.NoError(t, err)

	var slides []struct {
		Type string `json:"type"`
	}
	require.NoError(t, json.Unmarshal(raw, &slides))

	types := make([]string, 0, len(slides))
	for _, slide := range slides {
		types = append(types, slide.Type)
	}

	require.Equal(t, []string{
		slideIntro,
		slideActiveDays,
		slideViews,
		slideFavorites,
		slideFavoriteCategory,
		slidePurchases,
		slideSales,
		slideMessages,
		slideInterests,
		slideArchetype,
		slideFinal,
	}, types)
}
