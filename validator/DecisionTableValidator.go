package validator

import (
	"errors"
	conf "github.com/global-soft-ba/decisionTable/config" // TODO: Rename package?
	"github.com/global-soft-ba/decisionTable/data"
	"go.uber.org/multierr"
)

var (
	ErrDecisionTableIdIsRequired                = "decision table ID is required"
	ErrDecisionTableNameIsRequired              = "decision table name is required"
	ErrDecisionTableHitPolicyIsInvalid          = "decision table hit policy is invalid"
	ErrDecisionTableCollectOperatorIsInvalid    = "decision table collect operator is invalid"
	ErrDecisionTableExpressionLanguageIsInvalid = "decision table expression language is invalid"
	ErrDecisionTableInputFieldIsRequired        = "at least one decision table input field is required"
	ErrDecisionTableOutputFieldIsRequired       = "at least one decision table output field is required"

	ErrInputCountMismatch  = "the number of input entries does not match the number of input fields"
	ErrOutputCountMismatch = "the number of output entries does not match the number of output fields"
)

type DecisionTableValidator struct {
	decisionTable data.DecisionTable
	standard      data.Standard
	errors        []error
}

func NewDecisionTableValidator() DecisionTableValidatorInterface {
	return DecisionTableValidator{}
}

func (v DecisionTableValidator) Validate(decisionTable data.DecisionTable, standard data.Standard) error {
	v.decisionTable = decisionTable
	v.standard = standard

	v.executeValidation(v.validateID)
	v.executeValidation(v.validateName)
	v.executeValidation(v.validateHitPolicy)
	v.executeValidation(v.validateCollectOperator)
	v.executeValidation(v.validateExpressionLanguage)
	v.executeValidation(v.validateInputFields)
	v.executeValidation(v.validateOutputFields)
	v.executeValidation(v.validateRuleSchema)
	v.executeValidation(v.validateRules)

	if len(v.errors) > 0 {
		return multierr.Combine(v.errors...)
	}

	return nil
}

func (v *DecisionTableValidator) executeValidation(validate func() error) {
	if err := validate(); err != nil {
		v.errors = append(v.errors, err)
	}
}

func (v DecisionTableValidator) validateID() error {
	if len(v.decisionTable.ID) == 0 {
		return errors.New(ErrDecisionTableIdIsRequired)
	}

	return nil
}

func (v DecisionTableValidator) validateName() error {
	if len(v.decisionTable.Name) == 0 {
		return errors.New(ErrDecisionTableNameIsRequired)
	}

	return nil
}

func (v DecisionTableValidator) validateHitPolicy() error {
	if _, ok := conf.DecisionTableStandards[v.standard].HitPolicies[v.decisionTable.HitPolicy]; !ok {
		return errors.New(ErrDecisionTableHitPolicyIsInvalid)
	}

	return nil
}

func (v DecisionTableValidator) validateCollectOperator() error {
	if v.decisionTable.HitPolicy == data.Collect {
		if _, ok := conf.DecisionTableStandards[v.standard].CollectOperators[v.decisionTable.CollectOperator]; !ok {
			return errors.New(ErrDecisionTableCollectOperatorIsInvalid)
		}
	}

	return nil
}

func (v DecisionTableValidator) validateExpressionLanguage() error {
	if _, ok := conf.DecisionTableStandards[v.standard].ExpressionLanguages[v.decisionTable.ExpressionLanguage]; !ok {
		return errors.New(ErrDecisionTableExpressionLanguageIsInvalid)
	}

	return nil
}

func (v DecisionTableValidator) validateInputFields() error {
	if len(v.decisionTable.InputFields) == 0 {
		return errors.New(ErrDecisionTableInputFieldIsRequired)
	}

	return v.validateFields(v.decisionTable.InputFields)
}

func (v DecisionTableValidator) validateOutputFields() error {
	if len(v.decisionTable.OutputFields) == 0 {
		return errors.New(ErrDecisionTableOutputFieldIsRequired)
	}

	return v.validateFields(v.decisionTable.OutputFields)
}

func (v DecisionTableValidator) validateFields(fields []data.FieldInterface) error {
	var validationErrors []error

	fieldValidator := NewFieldValidator()

	for _, field := range fields {
		if err := fieldValidator.Validate(field, v.standard); err != nil {
			validationErrors = append(validationErrors, err)
		}
	}

	if len(validationErrors) > 0 {
		return multierr.Combine(validationErrors...)
	}

	return nil
}

func (v DecisionTableValidator) validateRuleSchema() error {
	var validationErrors []error

	numberOfInputFields := len(v.decisionTable.InputFields)
	numberOfOutputFields := len(v.decisionTable.OutputFields)

	inputCountMismatch := false
	outputCountMismatch := false

	for _, rule := range v.decisionTable.Rules {
		if len(rule.InputEntries) != numberOfInputFields {
			inputCountMismatch = true
		}

		if len(rule.OutputEntries) != numberOfOutputFields {
			outputCountMismatch = true
		}
	}

	if inputCountMismatch {
		validationErrors = append(validationErrors, errors.New(ErrInputCountMismatch))
	}

	if outputCountMismatch {
		validationErrors = append(validationErrors, errors.New(ErrOutputCountMismatch))
	}

	if len(validationErrors) > 0 {
		return multierr.Combine(validationErrors...)
	}

	return nil
}

func (v DecisionTableValidator) validateRules() error {
	var validationErrors []error

	ruleValidator := NewRuleValidator()

	for _, rule := range v.decisionTable.Rules {
		if err := ruleValidator.Validate(rule, v.decisionTable, v.standard); err != nil {
			validationErrors = append(validationErrors, err)
		}
	}

	if len(validationErrors) > 0 {
		return multierr.Combine(validationErrors...)
	}

	return nil
}
