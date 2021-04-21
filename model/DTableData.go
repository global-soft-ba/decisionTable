package model

type DTableData struct {
	Key                string
	Name               string
	HitPolicy          string
	CollectOperator    string
	ExpressionLanguage string

	InputFields  []Field
	OutputFields []Field
	Rules        []Rule
}
