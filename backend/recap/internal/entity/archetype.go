package entity

type Archetype struct {
	UserArchetype ArchetypeName
	Title         string
	Description   string
	Reasons       []ArchetypeReason
}

type ArchetypeName string

const (
	ArchetypeCollector  ArchetypeName = "collector"
	ArchetypeDealmaker  ArchetypeName = "dealmaker"
	ArchetypeNegotiator ArchetypeName = "negotiator"
	ArchetypeExplorer   ArchetypeName = "explorer"
)

func (a ArchetypeName) Valid() bool {
	switch a {
	case ArchetypeCollector,
		ArchetypeDealmaker,
		ArchetypeNegotiator,
		ArchetypeExplorer:
		return true
	default:
		return false
	}
}

// Metric is a scoring input. The values mirror the MetricCode enum in openapi.yaml.
type Metric string

const (
	MetricActiveDays Metric = "active_days"
	MetricViews      Metric = "views"
	MetricFavorites  Metric = "favorites"
	MetricPurchases  Metric = "purchases"
	MetricSales      Metric = "sales"
	MetricMessages   Metric = "messages"
	MetricCategories Metric = "categories"
	MetricListings   Metric = "listings"
)

func (m Metric) Valid() bool {
	switch m {
	case MetricActiveDays,
		MetricViews,
		MetricFavorites,
		MetricPurchases,
		MetricSales,
		MetricMessages,
		MetricCategories,
		MetricListings:
		return true
	default:
		return false
	}
}

type ArchetypeReason struct {
	Metric      Metric
	Value       string
	Explanation string
}
