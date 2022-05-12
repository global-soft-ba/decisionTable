package validator

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type RuleValidatorInterface interface {
	Validate(rule data.Rule, decisionTable data.DecisionTable, standard standard.Standard) error
}
