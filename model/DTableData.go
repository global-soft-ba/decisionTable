package model

type DTableData struct {
	Key             string
	Name            string
	HitPolicy       HitPolicy
	CollectOperator CollectOperator
	DTableStandard  DTableStandard

	InputFields  []Field
	OutputFields []Field
	Rules        []Rule
}
