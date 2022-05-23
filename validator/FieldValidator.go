package validator

import (
	"errors"
	"fmt"
	conf "github.com/global-soft-ba/decisionTable/config"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"go.uber.org/multierr"
)

var (
	ErrDecisionTableFieldNameIsRequired = "field name is required"
	ErrDecisionTableFieldTypeIsInvalid  = "invalid field type"
)

type FieldValidator struct {
	field    field.Field
	standard standard.Standard
	errors   []error
}

func NewFieldValidator() FieldValidatorInterface {
	return FieldValidator{}
}

func (v FieldValidator) Validate(field field.Field, standard standard.Standard) error {
	v.field = field
	v.standard = standard

	v.executeValidation(v.validateName)
	v.executeValidation(v.validateType)

	if len(v.errors) > 0 {
		return multierr.Combine(v.errors...)
	}

	return nil
}

func (v *FieldValidator) executeValidation(validate func() error) {
	if err := validate(); err != nil {
		v.errors = append(v.errors, err)
	}
}

func (v FieldValidator) validateName() error {
	if len(v.field.Name) == 0 {
		return errors.New(ErrDecisionTableFieldNameIsRequired)
	}

	return nil
}

func (v FieldValidator) validateType() error {
	if _, ok := conf.DecisionTableStandards[v.standard].VariableTypes[v.field.Type]; !ok {
		return fmt.Errorf("%s \"%s\"", ErrDecisionTableFieldTypeIsInvalid, v.field.Type)
	}

	return nil
}
