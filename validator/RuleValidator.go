package validator

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/rule"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"go.uber.org/multierr"
)

var (
	ErrDecisionTableRuleAnnotationIsRequired          = "decision table rule annotation is required"
	ErrDecisionTableEntryReferencedFieldTypeIsInvalid = "referenced field type does not match field type"
)

type RuleValidator struct {
	rule          rule.Rule
	decisionTable decisionTable.DecisionTable
	standard      standard.Standard
	errors        []error
}

func NewRuleValidator() RuleValidatorInterface {
	return RuleValidator{}
}

func (v RuleValidator) Validate(rule rule.Rule, decisionTable decisionTable.DecisionTable, standard standard.Standard) error {
	v.rule = rule
	v.decisionTable = decisionTable
	v.standard = standard

	v.executeValidation(v.validateAnnotation)
	v.executeValidation(v.validateInputEntries)
	v.executeValidation(v.validateOutputEntries)

	if len(v.errors) > 0 {
		return multierr.Combine(v.errors...)
	}

	return nil
}

func (v *RuleValidator) executeValidation(validate func() error) {
	if err := validate(); err != nil {
		v.errors = append(v.errors, err)
	}
}

func (v RuleValidator) validateAnnotation() error {
	if len(v.rule.Annotation) == 0 {
		return errors.New(ErrDecisionTableRuleAnnotationIsRequired)
	}

	return nil
}

func (v RuleValidator) validateInputEntries() error {
	var validationErrors []error

	for i, inputEntry := range v.rule.InputEntries {
		if i < len(v.decisionTable.InputFields) {
			if err := v.validateInputEntry(inputEntry, v.decisionTable.InputFields[i]); err != nil {
				validationErrors = append(validationErrors, err)
			}
		}
	}

	if len(validationErrors) > 0 {
		return multierr.Combine(validationErrors...)
	}

	return nil
}

func (v RuleValidator) validateOutputEntries() error {
	var validationErrors []error

	for i, outputEntry := range v.rule.OutputEntries {
		if i < len(v.decisionTable.OutputFields) {
			if err := v.validateOutputEntry(outputEntry, v.decisionTable.OutputFields[i]); err != nil {
				validationErrors = append(validationErrors, err)
			}
		}
	}

	if len(validationErrors) > 0 {
		return multierr.Combine(validationErrors...)
	}

	return nil
}

func (v RuleValidator) validateInputEntry(entry string, field field.Field) error {
	validator, err := CreateInputEntryValidator(v.decisionTable.ExpressionLanguage, entry)
	if err != nil {
		return err
	}

	return v.validateEntry(validator, entry, field)
}

func (v RuleValidator) validateOutputEntry(entry string, field field.Field) error {
	validator, err := CreateOutputEntryValidator(v.decisionTable.ExpressionLanguage, entry)
	if err != nil {
		return err
	}

	return v.validateEntry(validator, entry, field)
}

func (v RuleValidator) validateEntry(validator data.EntryValidatorInterface, entry string, field field.Field) error {
	if ok, err := validator.Validate(); !ok {
		return multierr.Combine(err...)
	}

	if ok, err := validator.ValidateDataTypeOfExpression(field.Type); !ok {
		return err
	}

	referencedFields, err := validator.ValidateExistenceOfFieldReferencesInExpression(append(v.decisionTable.InputFields, v.decisionTable.OutputFields...))
	if err != nil {
		return multierr.Combine(err...)
	}

	for _, referencedField := range referencedFields {
		if referencedField.Type != field.Type {
			return errors.New(ErrDecisionTableEntryReferencedFieldTypeIsInvalid)
		}
	}

	return nil
}
