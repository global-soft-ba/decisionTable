package validator

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type DecisionTableValidatorInterface interface {
	Validate(decisionTable data.DecisionTable, standard standard.Standard) error
}
