package recap

import (
	"math"
	"sort"

	"github.com/google/uuid"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

type CategoryWeights struct {
	Views     float64
	Favorites float64
	Purchases float64
	Sales     float64
}

var DefaultCategoryWeights = CategoryWeights{
	Views:     1,
	Favorites: 4,
	Purchases: 12,
	Sales:     8,
}

func (w CategoryWeights) Score(activity entity.CategoryActivity) float64 {
	return w.Views*float64(activity.Views) +
		w.Favorites*float64(activity.Favorites) +
		w.Purchases*float64(activity.Purchases) +
		w.Sales*float64(activity.Sales)
}

type subcategoryAggregate struct {
	id    uuid.UUID
	title string
	score float64
}

type categoryAggregate struct {
	id            uuid.UUID
	title         string
	score         float64
	subcategories []*subcategoryAggregate
}

func ScoreCategories(
	activities []entity.CategoryActivity,
	weights CategoryWeights,
) []entity.CategoryScore {
	return rankCategories(aggregateCategories(activities, weights))
}

// CategoryShare is how much of the whole activity the category holds, in percent.
func CategoryShare(scores []entity.CategoryScore, category entity.CategoryScore) int32 {
	var total float64
	for _, score := range scores {
		total += score.Score
	}

	if total == 0 {
		return 0
	}

	return int32(math.Round(category.Score / total * 100))
}

func aggregateCategories(
	activities []entity.CategoryActivity,
	weights CategoryWeights,
) []*categoryAggregate {
	byCategory := make(map[uuid.UUID]*categoryAggregate, len(activities))
	categories := make([]*categoryAggregate, 0, len(activities))

	for _, activity := range activities {
		score := weights.Score(activity)

		category, known := byCategory[activity.CategoryID]
		if !known {
			category = &categoryAggregate{id: activity.CategoryID, title: activity.CategoryTitle}
			byCategory[activity.CategoryID] = category
			categories = append(categories, category)
		}

		category.score += score
		addSubcategory(category, activity, score)
	}

	return categories
}

func addSubcategory(category *categoryAggregate, activity entity.CategoryActivity, score float64) {
	if activity.SubcategoryID == nil {
		return
	}

	for _, subcategory := range category.subcategories {
		if subcategory.id == *activity.SubcategoryID {
			subcategory.score += score

			return
		}
	}

	category.subcategories = append(category.subcategories, &subcategoryAggregate{
		id:    *activity.SubcategoryID,
		title: activity.SubcategoryTitle,
		score: score,
	})
}

func rankCategories(aggregates []*categoryAggregate) []entity.CategoryScore {
	scores := make([]entity.CategoryScore, 0, len(aggregates))

	for _, aggregate := range aggregates {
		if aggregate.score == 0 {
			continue
		}

		scores = append(scores, entity.CategoryScore{
			CategoryID:  aggregate.id,
			Title:       aggregate.title,
			Score:       aggregate.score,
			Subcategory: topSubcategory(aggregate.subcategories),
		})
	}

	sort.SliceStable(scores, func(i, j int) bool {
		return lessByScore(
			scores[i].Score, scores[j].Score,
			scores[i].Title, scores[j].Title,
			scores[i].CategoryID, scores[j].CategoryID,
		)
	})

	return scores
}

func topSubcategory(subcategories []*subcategoryAggregate) *entity.SubcategoryScore {
	var best *subcategoryAggregate

	for _, subcategory := range subcategories {
		if subcategory.score == 0 {
			continue
		}

		if best == nil || lessByScore(
			subcategory.score, best.score,
			subcategory.title, best.title,
			subcategory.id, best.id,
		) {
			best = subcategory
		}
	}

	if best == nil {
		return nil
	}

	return &entity.SubcategoryScore{ID: best.id, Title: best.title, Score: best.score}
}

func lessByScore(
	leftScore, rightScore float64,
	leftTitle, rightTitle string,
	leftID, rightID uuid.UUID,
) bool {
	if leftScore != rightScore {
		return leftScore > rightScore
	}

	if leftTitle != rightTitle {
		return leftTitle < rightTitle
	}

	return leftID.String() < rightID.String()
}
