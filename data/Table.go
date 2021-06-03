package data

type Table struct {
	Key              string
	Name             string
	HitPolicy        HitPolicy
	CollectOperator  CollectOperator
	NotationStandard DTableStandard
	Interferences    bool

	InputFields  []Field
	OutputFields []Field
	Rules        []Rule
}