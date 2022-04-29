package data

type DecisionTable struct {
	ID                 string
	Name               string
	HitPolicy          HitPolicy
	CollectOperator    CollectOperator
	ExpressionLanguage ExpressionLanguage
	Standard           Standard

	InputFields  []FieldInterface
	OutputFields []FieldInterface
	Rules        []Rule

	Interferences bool
}
