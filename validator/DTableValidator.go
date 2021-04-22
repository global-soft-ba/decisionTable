package validator

import (
	"decisionTable/constant"
	"decisionTable/model"
	"errors"
)

var (
	ErrDTableNameEmpty            = errors.New("name of decision table is required")
	ErrDTableKeyEmpty             = errors.New("definition key of decision table is required")
	ErrDTableStandardInvalid      = errors.New("missing or invalid table standard")
	ErrDTableHitPolicy            = errors.New("hit policy of decision table is invalid")
	ErrDTableCollectOperator      = errors.New("collect operator of decision table is invalid")
	ErrDTableEmptyCollectOperator = errors.New("collect operator of decision table cannot be empty for the hit policy")
	ErrDTableInputEmpty           = errors.New("at least one input is required")
	ErrDTableOutputEmpty          = errors.New("at least one output is required")

	ErrDTableFieldNameEmpty  = errors.New("field name is empty")
	ErrDTableFieldLabelEmpty = errors.New("field label is empty")
	ErrDTableFieldTypInvalid = errors.New("field type is invalid")

	ErrRuleHaveDifferentAmountOfInputFields  = errors.New("amount of input entries does not match input fields of decision table")
	ErrRuleHaveDifferentAmountOfOutputFields = errors.New("amount of output entries does not match input fields of decision table")
	ErrDTableEntryExpressionLangInvalid      = errors.New("entry expression language of the table standard is invalid")
)

func CreateDTableValidator(dTable model.DTableData) DTableValidatorInterface {
	r := DTableValidator{dTable: dTable, valid: false, errs: []error{}}
	return r
}

type DTableValidator struct {
	dTable model.DTableData
	valid  bool
	errs   []error
}

func (d DTableValidator) Validate() (bool, []error) {
	var v = d
	v = v.collectValidationResult(v.validateName)
	v = v.collectValidationResult(v.validateKey)
	v = v.collectValidationResult(v.validateTableStandard)
	v = v.collectValidationResult(v.validateHitPolicy)
	v = v.collectValidationResults(v.validateInput)
	v = v.collectValidationResults(v.validateOutput)
	v = v.collectValidationResults(v.validateRuleSchema)
	v = v.collectValidationResults(v.validateRules)

	if len(v.errs) == 0 {
		v.valid = true
	}
	return v.valid, v.errs
}

func (d DTableValidator) collectValidationResult(v func() (bool, error)) DTableValidator {
	if ok, err := v(); !ok {
		d.valid = false
		d.errs = append(d.errs, err)
	}

	return d
}

func (d DTableValidator) collectValidationResults(v func() (bool, []error)) DTableValidator {
	if ok, err := v(); !ok {
		d.valid = false
		d.errs = append(d.errs, err...)
	}

	return d
}

func (d DTableValidator) validateName() (bool, error) {
	if len(d.dTable.Name) == 0 {
		return false, ErrDTableNameEmpty
	}
	return true, nil
}

func (d DTableValidator) validateKey() (bool, error) {
	if len(d.dTable.Key) == 0 {
		return false, ErrDTableKeyEmpty
	}
	return true, nil
}

func (d DTableValidator) validateTableStandard() (bool, error) {
	if _, ok := constant.DTableStandards[d.dTable.DTableStandard]; !ok {
		return false, ErrDTableStandardInvalid
	}
	return true, nil
}

func (d DTableValidator) validateHitPolicy() (bool, error) {
	if _, ok := constant.AllowedConfigForStandard[d.dTable.DTableStandard].HitPolicies[d.dTable.HitPolicy]; !ok {
		return false, ErrDTableHitPolicy
	}

	if d.dTable.HitPolicy == constant.CollectOperatorPolicy {
		if len(d.dTable.CollectOperator) == 0 {
			return false, ErrDTableEmptyCollectOperator
		}

		if _, ok := constant.AllowedConfigForStandard[d.dTable.DTableStandard].CollectOperators[d.dTable.CollectOperator]; !ok {
			return false, ErrDTableCollectOperator
		}
	}

	return true, nil
}

func (d DTableValidator) validateInput() (bool, []error) {
	if len(d.dTable.InputFields) == 0 {
		return false, []error{ErrDTableInputEmpty}
	}

	var errResult []error

	for _, v := range d.dTable.InputFields {
		if ok, err := d.checkFields(v); !ok {
			errResult = append(errResult, err)
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d DTableValidator) validateOutput() (bool, []error) {
	if len(d.dTable.OutputFields) == 0 {
		return false, []error{ErrDTableOutputEmpty}
	}

	var errResult []error

	for _, v := range d.dTable.OutputFields {
		if ok, err := d.checkFields(v); !ok {
			errResult = append(errResult, err)
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d DTableValidator) validateRuleSchema() (bool, []error) {
	lengthInput := len(d.dTable.InputFields)
	lengthOutput := len(d.dTable.OutputFields)

	var errResult []error

	var inputResult = false
	var outputResult = false
	for _, v := range d.dTable.Rules {
		if len(v.InputEntries) != lengthInput {
			inputResult = true
		}

		if len(v.OutputEntries) != lengthOutput {
			outputResult = true
		}

	}
	if inputResult {
		errResult = append(errResult, ErrRuleHaveDifferentAmountOfInputFields)
	}
	if outputResult {
		errResult = append(errResult, ErrRuleHaveDifferentAmountOfOutputFields)
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d DTableValidator) validateRules() (bool, []error) {

	var errResult []error

	for _, r := range d.dTable.Rules {
		if ok, err := d.checkRule(r); !ok {
			errResult = append(errResult, err...)
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d DTableValidator) checkFields(f model.Field) (bool, error) {
	if len(f.Name) == 0 {
		return false, ErrDTableFieldNameEmpty
	}

	if len(f.Label) == 0 {
		return false, ErrDTableFieldLabelEmpty
	}

	if _, ok := constant.AllowedConfigForStandard[d.dTable.DTableStandard].VariableType[f.Typ]; !ok {
		return false, ErrDTableFieldTypInvalid
	}

	return true, nil
}

func (d DTableValidator) checkRule(r model.Rule) (bool, []error) {
	var errResult []error

	for _, v := range r.InputEntries {
		if _, ok := constant.AllowedConfigForStandard[d.dTable.DTableStandard].ExpressionLanguage[v.ExpressionLanguage]; !ok {
			errResult = append(errResult, ErrDTableEntryExpressionLangInvalid)
		}
	}

	for _, v := range r.OutputEntries {
		if _, ok := constant.AllowedConfigForStandard[d.dTable.DTableStandard].ExpressionLanguage[v.ExpressionLanguage]; !ok {
			errResult = append(errResult, ErrDTableEntryExpressionLangInvalid)
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}
