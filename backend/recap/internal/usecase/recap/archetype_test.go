package recap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

const categoriesOnPlatform = 32

func TestDetectArchetype(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		activity entity.UserActivity
		expected entity.ArchetypeName
	}{
		{
			name: "browsing wide across the catalog makes an explorer",
			activity: entity.UserActivity{
				ActiveDays: 200, Views: 2000, Favorites: 10, CategoriesTouched: 20,
			},
			expected: entity.ArchetypeExplorer,
		},
		{
			name: "favorites dominate the mix for a collector",
			activity: entity.UserActivity{
				ActiveDays: 90, Views: 300, Favorites: 250, Purchases: 8, CategoriesTouched: 3,
			},
			expected: entity.ArchetypeCollector,
		},
		{
			name: "selling dominates the mix for a dealmaker",
			activity: entity.UserActivity{
				ActiveDays: 120, Views: 60, Sales: 40, ListingsCreated: 45,
				Purchases: 3, CategoriesTouched: 5,
			},
			expected: entity.ArchetypeDealmaker,
		},
		{
			name: "messages dominate the mix for a negotiator",
			activity: entity.UserActivity{
				ActiveDays: 100, Views: 200, MessagesAsBuyer: 700, MessagesAsSeller: 300,
				Purchases: 6, Sales: 6, CategoriesTouched: 4,
			},
			expected: entity.ArchetypeNegotiator,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			archetype := DetectArchetype(test.activity, categoriesOnPlatform)

			require.Equal(t, test.expected, archetype.UserArchetype)
			assert.True(t, archetype.UserArchetype.Valid(), "archetype is not part of the contract")
			assert.NotEmpty(t, archetype.Title, "archetype must carry a title")
			assert.NotEmpty(t, archetype.Description, "archetype must carry a description")
		})
	}
}

func TestDetectArchetypeIgnoresVolume(t *testing.T) {
	t.Parallel()

	small := entity.UserActivity{
		ActiveDays: 30, Views: 20, Favorites: 50, Purchases: 2, CategoriesTouched: 2,
	}
	large := entity.UserActivity{
		ActiveDays: 30, Views: 200, Favorites: 500, Purchases: 20, CategoriesTouched: 2,
	}

	assert.Equal(t,
		DetectArchetype(small, categoriesOnPlatform).UserArchetype,
		DetectArchetype(large, categoriesOnPlatform).UserArchetype,
		"the same action mix must yield the same archetype at any scale",
	)
}

func TestDetectArchetypeExplainsItselfWithFacts(t *testing.T) {
	t.Parallel()

	activity := entity.UserActivity{
		ActiveDays: 100, Views: 200, MessagesAsBuyer: 900, Purchases: 6, Sales: 6,
	}

	archetype := DetectArchetype(activity, categoriesOnPlatform)
	require.Len(t, archetype.Reasons, reasonsPerArchetype)

	leading := archetype.Reasons[0]
	assert.Equal(t, entity.MetricMessages, leading.Metric, "messages must lead the explanation")
	assert.Equal(t, "900", leading.Value, "the reason must carry the raw value")

	for _, reason := range archetype.Reasons {
		assert.Truef(t, reason.Metric.Valid(), "metric %s is not part of the contract", reason.Metric)
		assert.NotEmptyf(t, reason.Explanation, "reason %s is not explained", reason.Metric)
	}
}

func TestReasonExplanationCoversEveryMetric(t *testing.T) {
	t.Parallel()

	metrics := []entity.Metric{
		entity.MetricActiveDays,
		entity.MetricViews,
		entity.MetricFavorites,
		entity.MetricPurchases,
		entity.MetricSales,
		entity.MetricMessages,
		entity.MetricCategories,
		entity.MetricListings,
	}

	for _, metric := range metrics {
		explanation := reasonExplanation(metric, 42, 0.42, categoriesOnPlatform)

		assert.NotEmptyf(t, explanation, "metric %s has no explanation", metric)
		assert.NotContainsf(t, explanation, "%!", "metric %s has a broken format", metric)
		assert.Containsf(t, explanation, "42", "metric %s does not mention its value", metric)
	}
}

func TestDetectArchetypeIsReproducible(t *testing.T) {
	t.Parallel()

	activity := entity.UserActivity{
		ActiveDays: 150, Views: 400, Favorites: 40, Purchases: 5,
		Sales: 5, MessagesAsBuyer: 120, CategoriesTouched: 9, ListingsCreated: 5,
	}

	first := DetectArchetype(activity, categoriesOnPlatform)
	require.NotEmpty(t, first.Reasons)

	for range 10 {
		next := DetectArchetype(activity, categoriesOnPlatform)

		require.Equal(t, first.UserArchetype, next.UserArchetype, "archetype is not stable")
		require.Equal(t, first.Reasons, next.Reasons, "reasons are not stable between runs")
	}
}

func TestDetectArchetypeAlwaysExplainsMinimalActivity(t *testing.T) {
	t.Parallel()

	archetype := DetectArchetype(entity.UserActivity{ActiveDays: 1, Views: 5}, categoriesOnPlatform)

	assert.NotEmpty(t, archetype.Reasons, "a user with activity must get an explanation")
}

func TestDetectArchetypeSurvivesUnknownCatalog(t *testing.T) {
	t.Parallel()

	activity := entity.UserActivity{ActiveDays: 50, Views: 100, CategoriesTouched: 7}

	archetype := DetectArchetype(activity, 0)
	require.True(t, archetype.UserArchetype.Valid())

	for _, reason := range archetype.Reasons {
		assert.NotEqual(t, entity.MetricCategories, reason.Metric,
			"breadth must not be used as a reason without a catalog size")
	}
}
