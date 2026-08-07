package recap

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

const reasonsPerArchetype = 2

type archetypeProfile struct {
	name        entity.ArchetypeName
	title       string
	description string
	weights     map[entity.Metric]float64
}

// These numbers are a product definition ("what we call a dealmaker"), not a
// statistical claim about a typical user.
var archetypeProfiles = []archetypeProfile{
	{
		name:        entity.ArchetypeExplorer,
		title:       "Исследователь",
		description: "Вы изучаете площадку вширь: смотрите много и заглядываете в разные категории.",
		weights: map[entity.Metric]float64{
			entity.MetricCategories: 0.60,
			entity.MetricActiveDays: 0.25,
			entity.MetricViews:      0.15,
		},
	},
	{
		name:        entity.ArchetypeCollector,
		title:       "Коллекционер",
		description: "Вы копите то, что понравилось, и возвращаетесь к отложенному, когда приходит время.",
		weights: map[entity.Metric]float64{
			entity.MetricFavorites:  0.70,
			entity.MetricPurchases:  0.20,
			entity.MetricActiveDays: 0.10,
		},
	},
	{
		name:        entity.ArchetypeDealmaker,
		title:       "Делец",
		description: "Вы не только покупаете, но и продаёте: ваши объявления находят новых владельцев.",
		weights: map[entity.Metric]float64{
			entity.MetricSales:     0.45,
			entity.MetricListings:  0.45,
			entity.MetricPurchases: 0.10,
		},
	},
	{
		name:        entity.ArchetypeNegotiator,
		title:       "Переговорщик",
		description: "Вы всегда пишете первым: обсуждаете детали, торгуетесь и договариваетесь.",
		weights: map[entity.Metric]float64{
			entity.MetricMessages:  0.80,
			entity.MetricPurchases: 0.10,
			entity.MetricSales:     0.10,
		},
	},
}

var metricLabels = map[entity.Metric]string{
	entity.MetricViews:     "Просмотры",
	entity.MetricFavorites: "Избранное",
	entity.MetricPurchases: "Покупки",
	entity.MetricSales:     "Продажи",
	entity.MetricMessages:  "Сообщения",
	entity.MetricListings:  "Публикации",
}

func DetectArchetype(activity entity.UserActivity, totalCategories int64) entity.Archetype {
	shares := metricShares(activity, totalCategories)

	var (
		best      archetypeProfile
		bestScore float64
	)

	for index, profile := range archetypeProfiles {
		score := profileScore(profile, shares)
		if index == 0 || score > bestScore {
			best, bestScore = profile, score
		}
	}

	return entity.Archetype{
		UserArchetype: best.name,
		Title:         best.title,
		Description:   best.description,
		Reasons:       archetypeReasons(best, activity, shares, totalCategories),
	}
}

func metricShares(activity entity.UserActivity, totalCategories int64) map[entity.Metric]float64 {
	actions := activity.TotalActions()

	return map[entity.Metric]float64{
		entity.MetricViews:      ratio(activity.Views, actions),
		entity.MetricFavorites:  ratio(activity.Favorites, actions),
		entity.MetricPurchases:  ratio(activity.Purchases, actions),
		entity.MetricSales:      ratio(activity.Sales, actions),
		entity.MetricMessages:   ratio(activity.Messages(), actions),
		entity.MetricListings:   ratio(activity.ListingsCreated, actions),
		entity.MetricActiveDays: ratio(activity.ActiveDays, daysInYear),
		entity.MetricCategories: ratio(activity.CategoriesTouched, totalCategories),
	}
}

func ratio(value, total int64) float64 {
	if value <= 0 || total <= 0 {
		return 0
	}

	if value >= total {
		return 1
	}

	return float64(value) / float64(total)
}

func metricValues(activity entity.UserActivity) map[entity.Metric]int64 {
	return map[entity.Metric]int64{
		entity.MetricActiveDays: activity.ActiveDays,
		entity.MetricViews:      activity.Views,
		entity.MetricFavorites:  activity.Favorites,
		entity.MetricPurchases:  activity.Purchases,
		entity.MetricSales:      activity.Sales,
		entity.MetricMessages:   activity.Messages(),
		entity.MetricCategories: activity.CategoriesTouched,
		entity.MetricListings:   activity.ListingsCreated,
	}
}

func profileScore(profile archetypeProfile, shares map[entity.Metric]float64) float64 {
	var score float64
	for metric, weight := range profile.weights {
		score += weight * shares[metric]
	}

	return score
}

func archetypeReasons(
	profile archetypeProfile,
	activity entity.UserActivity,
	shares map[entity.Metric]float64,
	totalCategories int64,
) []entity.ArchetypeReason {
	type contribution struct {
		metric entity.Metric
		amount float64
	}

	contributions := make([]contribution, 0, len(profile.weights))

	for metric, weight := range profile.weights {
		amount := weight * shares[metric]
		if amount == 0 {
			continue
		}

		contributions = append(contributions, contribution{metric: metric, amount: amount})
	}

	sort.Slice(contributions, func(i, j int) bool {
		if contributions[i].amount != contributions[j].amount {
			return contributions[i].amount > contributions[j].amount
		}

		return contributions[i].metric < contributions[j].metric
	})

	if len(contributions) > reasonsPerArchetype {
		contributions = contributions[:reasonsPerArchetype]
	}

	values := metricValues(activity)
	reasons := make([]entity.ArchetypeReason, 0, len(contributions))

	for _, item := range contributions {
		value := values[item.metric]

		reasons = append(reasons, entity.ArchetypeReason{
			Metric:      item.metric,
			Value:       strconv.FormatInt(value, 10),
			Explanation: reasonExplanation(item.metric, value, shares[item.metric], totalCategories),
		})
	}

	return reasons
}

func reasonExplanation(
	metric entity.Metric,
	value int64,
	share float64,
	totalCategories int64,
) string {
	switch metric {
	case entity.MetricActiveDays:
		return fmt.Sprintf("вы были с авито %d %s из %d", value, plural(value, "день", "дня", "дней"), daysInYear)

	case entity.MetricCategories:
		return fmt.Sprintf("вы посмотрели %d %s из %d на площадке",
			value,
			plural(value, "категорию", "категории", "категорий"),
			totalCategories,
		)

	default:
		return fmt.Sprintf("%s — %d%% всех ваших действий",
			metricLabels[metric],
			int64(math.Round(share*100)),
		)
	}
}
