package valid

import (
	"errors"
	conf "github.com/global-soft-ba/decisionTable/config"
	"github.com/global-soft-ba/decisionTable/data"
)

var (
	ErrDTableNameEmpty       = errors.New("name of decision table is required")
	ErrDTableKeyEmpty        = errors.New("definition key of decision table is required")
	ErrDTableHitPolicy       = errors.New("hit policy of decision table is invalid")
	ErrDTableCollectOperator = errors.New("collect operator of decision table is invalid")
	ErrDTableInputEmpty      = errors.New("at least one input is required")
	ErrDTableOutputEmpty     = errors.New("at least one output is required")

	ErrDTableFieldIdIsEmpty  = errors.New("field id is empty")
	ErrDTableFieldTypInvalid = errors.New("field type is invalid")

	ErrRuleHaveDifferentAmountOfInputFields  = errors.New("amount of input entries does not match input fields of decision table")
	ErrRuleHaveDifferentAmountOfOutputFields = errors.New("amount of output entries does not match input fields of decision table")
	ErrDTableEntryExpressionLangInvalid      = errors.New("entry expression language of the table standard is invalid")
	ErrDTableEntryReferencedFieldTypInvalid  = errors.New("referenced field type does not match field type")
)

func CreateDecisionTableValidator() ValidatorInterface {
	return Validator{}
}

type Validator struct {
	dTable data.Table
	valid  bool
	errs   []error
}

func (d Validator) ValidateContainsInterferences(table data.Table) bool {
	d.dTable = table
	output := d.dTable.OutputFields

	for _, val := range d.dTable.InputFields {
		if d.fieldIsContained(val, output) {
			return true
		}
	}

	return false
}

func (d Validator) fieldIsContained(field data.FieldInterface, setOfFields []data.FieldInterface) bool {
	for _, val := range setOfFields {
		if val.Id() == field.Id() {
			return true
		}
	}
	return false
}

func (d Validator) Validate(table data.Table) (bool, []error) {
	d.dTable = table
	var v = d
	v = v.executeValidation(v.validateName)
	v = v.executeValidation(v.validateKey)
	v = v.executeValidation(v.validateHitPolicy)
	v = v.executeValidation(v.validateCollectOperator)
	v = v.executeValidation(v.validateInputFields)
	v = v.executeValidation(v.validateOutputFields)
	v = v.executeValidation(v.validateRuleSchema)
	v = v.executeValidation(v.validateRules)

	if len(v.errs) == 0 {
		v.valid = true
	}
	return v.valid, v.errs
}

func (d Validator) executeValidation(v func() (bool, []error)) Validator {
	if ok, err := v(); !ok {
		d.valid = false
		d.errs = append(d.errs, err...)
	}

	return d
}

func (d Validator) validateName() (bool, []error) {
	if len(d.dTable.Name) == 0 {
		return false, []error{ErrDTableNameEmpty}
	}
	return true, nil
}

func (d Validator) validateKey() (bool, []error) {
	if len(d.dTable.Key) == 0 {
		return false, []error{ErrDTableKeyEmpty}
	}
	return true, nil
}

func (d Validator) validateHitPolicy() (bool, []error) {

	if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].HitPolicies[d.dTable.HitPolicy]; !ok {
		return false, []error{ErrDTableHitPolicy}
	}

	return true, nil
}

func (d Validator) validateCollectOperator() (bool, []error) {
	if d.dTable.HitPolicy == data.Collect {
		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].CollectOperators[d.dTable.CollectOperator]; !ok {
			return false, []error{ErrDTableCollectOperator}
		}
	}

	return true, nil
}

func (d Validator) validateInputFields() (bool, []error) {
	if len(d.dTable.InputFields) == 0 {
		return false, []error{ErrDTableInputEmpty}
	}
	return d.validateFields(d.dTable.InputFields)
}

func (d Validator) validateOutputFields() (bool, []error) {
	if len(d.dTable.OutputFields) == 0 {
		return false, []error{ErrDTableOutputEmpty}
	}
	return d.validateFields(d.dTable.OutputFields)
}

func (d Validator) validateFields(fields []data.FieldInterface) (bool, []error) {
	var errResult []error

	for _, v := range fields {

		if v.Id() == "" {
			errResult = append(errResult, ErrDTableFieldIdIsEmpty)
		}

		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].VariableType[v.DataTyp()]; !ok {
			errResult = append(errResult, ErrDTableFieldTypInvalid)
		}

	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) validateRuleSchema() (bool, []error) {
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

func (d Validator) validateRules() (bool, []error) {

	var errResult []error

	for _, r := range d.dTable.Rules {
		if ok, err := d.validateInputRuleEntries(r); !ok {
			errResult = append(errResult, err...)
		}

		if ok, err := d.validateOutputRuleEntries(r); !ok {
			errResult = append(errResult, err...)
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) validateInputRuleEntries(r data.Rule) (bool, []error) {
	var errResult []error

	for i, v := range r.InputEntries {
		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].ExpressionLanguage[v.ExpressionLanguage()]; !ok {
			errResult = append(errResult, ErrDTableEntryExpressionLangInvalid)
		}

		if i < len(d.dTable.InputFields) {
			if ok, err := d.validateEntry(v, d.dTable.InputFields[i]); !ok {
				errResult = append(errResult, err...)
			}
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) validateOutputRuleEntries(r data.Rule) (bool, []error) {
	var errResult []error

	for i, v := range r.OutputEntries {
		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].ExpressionLanguage[v.ExpressionLanguage()]; !ok {
			errResult = append(errResult, ErrDTableEntryExpressionLangInvalid)
		}

		if i < len(d.dTable.OutputFields) {
			if ok, err := d.validateEntry(v, d.dTable.OutputFields[i]); !ok {
				errResult = append(errResult, err...)
			}
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) validateEntry(entry data.EntryInterface, field data.FieldInterface) (bool, []error) {
	if ok, err := entry.Validate(); !ok {
		return false, err
	}
	if ok, err := entry.ValidateDataTypeOfExpression(field.DataTyp()); !ok {
		return false, []error{err}
	}

	fields, err := entry.ValidateExistenceOfFieldReferencesInExpression(append(d.dTable.InputFields, d.dTable.OutputFields...))
	if err != nil {
		return false, err
	}

	for _, val := range fields {
		if val.DataTyp() != field.DataTyp() {
			return false, []error{ErrDTableEntryReferencedFieldTypInvalid}
		}
	}

	return true, nil
}
