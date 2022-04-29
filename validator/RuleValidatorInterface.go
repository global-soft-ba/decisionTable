package validator

import "github.com/global-soft-ba/decisionTable/data"

type RuleValidatorInterface interface {
	Validate(rule data.Rule, decisionTable data.DecisionTable, standard data.Standard) error
}
