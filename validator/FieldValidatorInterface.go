package validator

import "github.com/global-soft-ba/decisionTable/data"

type FieldValidatorInterface interface {
	Validate(field data.FieldInterface, standard data.Standard) error
}
