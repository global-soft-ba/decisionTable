package validator

import (
	"errors"
	conf "github.com/global-soft-ba/decisionTable/config"
	"github.com/global-soft-ba/decisionTable/data"
	"go.uber.org/multierr"
)

var (
	ErrDecisionTableFieldIdIsRequired  = "field id is required"
	ErrDecisionTableFieldTypeIsInvalid = "field type is invalid"
)

type FieldValidator struct {
	field    data.FieldInterface
	standard data.Standard
	errors   []error
}

func NewFieldValidator() FieldValidatorInterface {
	return FieldValidator{}
}

func (v FieldValidator) Validate(field data.FieldInterface, standard data.Standard) error {
	v.field = field
	v.standard = standard

	v.executeValidation(v.validateID)
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

func (v FieldValidator) validateID() error {
	if len(v.field.ID()) == 0 {
		return errors.New(ErrDecisionTableFieldIdIsRequired)
	}

	return nil
}

func (v FieldValidator) validateType() error {
	if _, ok := conf.DecisionTableStandards[v.standard].VariableTypes[v.field.DataType()]; !ok {
		return errors.New(ErrDecisionTableFieldTypeIsInvalid)
	}

	return nil
}
