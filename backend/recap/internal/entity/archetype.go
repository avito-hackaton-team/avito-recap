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

type ArchetypeReason struct {
	Metric      Metric
	Value       string
	Explanation string
}
