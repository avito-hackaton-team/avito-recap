package recap

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

var (
	electronics = uuid.MustParse("11111111-1111-1111-1111-111111111111")
	hobby       = uuid.MustParse("22222222-2222-2222-2222-222222222222")
	phones      = uuid.MustParse("33333333-3333-3333-3333-333333333333")
	laptops     = uuid.MustParse("44444444-4444-4444-4444-444444444444")
)

func TestScoreCategoriesPrefersIntentOverViews(t *testing.T) {
	t.Parallel()

	activities := []entity.CategoryActivity{
		{CategoryID: electronics, CategoryTitle: "Электроника", Views: 20},
		{CategoryID: hobby, CategoryTitle: "Хобби", Views: 2, Purchases: 2},
	}

	scores := ScoreCategories(activities, DefaultCategoryWeights)
	require.NotEmpty(t, scores)

	assert.Equal(t, hobby, scores[0].CategoryID, "intent must outweigh browsing")
}

func TestScoreCategoriesPicksTopSubcategoryAndShare(t *testing.T) {
	t.Parallel()

	activities := []entity.CategoryActivity{
		{
			CategoryID: electronics, CategoryTitle: "Электроника",
			SubcategoryID: &phones, SubcategoryTitle: "Телефоны",
			Views: 10, Favorites: 5,
		},
		{
			CategoryID: electronics, CategoryTitle: "Электроника",
			SubcategoryID: &laptops, SubcategoryTitle: "Ноутбуки",
			Views: 10,
		},
		{CategoryID: hobby, CategoryTitle: "Хобби", Views: 10},
	}

	scores := ScoreCategories(activities, DefaultCategoryWeights)
	require.Len(t, scores, 2)

	favorite := scores[0]
	require.Equal(t, electronics, favorite.CategoryID)

	// Both subcategories have 10 views, favorites decide: 30 against 10.
	require.NotNil(t, favorite.Subcategory, "the winning category must name its subcategory")
	assert.Equal(t, phones, favorite.Subcategory.ID)

	// electronics = 40 of 50 total.
	assert.Equal(t, int32(80), CategoryShare(scores, favorite))
}

func TestScoreCategoriesIsDeterministicOnTies(t *testing.T) {
	t.Parallel()

	activities := []entity.CategoryActivity{
		{CategoryID: hobby, CategoryTitle: "Хобби", Views: 10},
		{CategoryID: electronics, CategoryTitle: "Электроника", Views: 10},
	}

	first := ScoreCategories(activities, DefaultCategoryWeights)
	second := ScoreCategories(activities, DefaultCategoryWeights)

	require.Len(t, first, 2)
	assert.Equal(t, "Хобби", first[0].Title, "a tie must be broken by title")
	assert.Equal(t, first, second, "scoring is not reproducible")
}

func TestScoreCategoriesSkipsUntouchedCategories(t *testing.T) {
	t.Parallel()

	activities := []entity.CategoryActivity{
		{CategoryID: electronics, CategoryTitle: "Электроника"},
	}

	assert.Empty(t, ScoreCategories(activities, DefaultCategoryWeights))
}

func TestScoreCategoriesWithoutActivity(t *testing.T) {
	t.Parallel()

	assert.Empty(t, ScoreCategories(nil, DefaultCategoryWeights),
		"an empty period must not produce a category slide")
}
