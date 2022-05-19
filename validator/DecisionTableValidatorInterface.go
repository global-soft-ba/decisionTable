package validator

import (
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type DecisionTableValidatorInterface interface {
	Validate(decisionTable decisionTable.DecisionTable, standard standard.Standard) error
}
