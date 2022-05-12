package validator

import (
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type FieldValidatorInterface interface {
	Validate(field field.Field, standard standard.Standard) error
}
