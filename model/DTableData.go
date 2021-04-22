package model

type DTableData struct {
	Key             string
	Name            string
	HitPolicy       string
	CollectOperator string
	DTableStandard  string

	InputFields  []Field
	OutputFields []Field
	Rules        []Rule
}
