package validator

import "github.com/global-soft-ba/decisionTable/data"

type DecisionTableValidatorInterface interface {
	Validate(table data.DecisionTable, standard data.Standard) error
}
