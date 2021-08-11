package decisionTable

type ValidatorInterface interface {
	Validate(table DecisionTable) (bool, []error)
	ValidateContainsInterferences(table DecisionTable) bool
}
