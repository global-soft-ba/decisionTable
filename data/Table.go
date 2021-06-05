package data

type Table struct {
	Key              string
	Name             string
	HitPolicy        HitPolicy
	CollectOperator  CollectOperator
	NotationStandard DTableStandard
	Interferences    bool

	InputFields  []FieldInterface
	OutputFields []FieldInterface
	Rules        []Rule
}
