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

type ArchetypeReason struct {
	Metric      string
	Value       *string
	Explanation string
}
