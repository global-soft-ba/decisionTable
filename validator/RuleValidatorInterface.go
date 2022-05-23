package validator

import (
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/rule"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type RuleValidatorInterface interface {
	Validate(rule rule.Rule, decisionTable decisionTable.DecisionTable, standard standard.Standard) error
}
