package model

type DTableData struct {
	Key              string
	Name             string
	HitPolicy        HitPolicy
	CollectOperator  CollectOperator
	NotationStandard DTableStandard

	InputFields  []Field
	OutputFields []Field
	Rules        []Rule
}
