package validators

import (
	conf "decisionTable/config"
	"decisionTable/model"
	"decisionTable/validators/expressionlanguages"
	"errors"
)

var (
	ErrDTableNameEmpty       = errors.New("name of decision table is required")
	ErrDTableKeyEmpty        = errors.New("definition key of decision table is required")
	ErrDTableHitPolicy       = errors.New("hit policy of decision table is invalid")
	ErrDTableCollectOperator = errors.New("collect operator of decision table is invalid")
	ErrDTableInputEmpty      = errors.New("at least one input is required")
	ErrDTableOutputEmpty     = errors.New("at least one output is required")

	ErrDTableFieldNameEmpty  = errors.New("field name is empty")
	ErrDTableFieldLabelEmpty = errors.New("field label is empty")
	ErrDTableFieldTypInvalid = errors.New("field type is invalid")

	ErrRuleHaveDifferentAmountOfInputFields  = errors.New("amount of input entries does not match input fields of decision table")
	ErrRuleHaveDifferentAmountOfOutputFields = errors.New("amount of output entries does not match input fields of decision table")
	ErrDTableEntryExpressionLangInvalid      = errors.New("entry expression language of the table standard is invalid")
)

func CreateDecisionTableValidator(dTable model.TableData, prsFac expressionlanguages.ExpressionValidatorFactory) ValidatorInterface {
	r := Validator{dTable: dTable, prs: prsFac, valid: false, errs: []error{}}
	return r
}

type Validator struct {
	dTable model.TableData
	prs    expressionlanguages.ExpressionValidatorFactory
	valid  bool
	errs   []error
}

func (d Validator) Validate() (bool, []error) {
	var v = d
	v = v.executeValidation(v.validateName)
	v = v.executeValidation(v.validateKey)
	v = v.executeValidation(v.validateHitPolicy)
	v = v.executeValidation(v.validateCollectOperator)
	v = v.executeValidation(v.validateInput)
	v = v.executeValidation(v.validateOutput)
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
	if d.dTable.HitPolicy == model.Collect {
		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].CollectOperators[d.dTable.CollectOperator]; !ok {
			return false, []error{ErrDTableCollectOperator}
		}
	}

	return true, nil
}

func (d Validator) validateInput() (bool, []error) {
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

func (d Validator) validateOutput() (bool, []error) {
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
		if ok, err := d.checkInputRuleEntries(r); !ok {
			errResult = append(errResult, err...)
		}

		if ok, err := d.checkOutputRuleEntries(r); !ok {
			errResult = append(errResult, err...)
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) checkFields(f model.Field) (bool, error) {
	if len(f.Name) == 0 {
		return false, ErrDTableFieldNameEmpty
	}

	if len(f.Label) == 0 {
		return false, ErrDTableFieldLabelEmpty
	}

	if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].VariableType[f.Typ]; !ok {
		return false, ErrDTableFieldTypInvalid
	}

	return true, nil
}

func (d Validator) checkInputRuleEntries(r model.Rule) (bool, []error) {
	var errResult []error

	for i, v := range r.InputEntries {
		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].ExpressionLanguage[v.ExpressionLanguage()]; !ok {
			errResult = append(errResult, ErrDTableEntryExpressionLangInvalid)
		}

		if i < len(d.dTable.InputFields) {
			if ok, err := d.validateInputEntry(d.dTable.InputFields[i], v); !ok {
				errResult = append(errResult, err...)
			}
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) checkOutputRuleEntries(r model.Rule) (bool, []error) {
	var errResult []error

	for i, v := range r.OutputEntries {
		if _, ok := conf.DecisionTableStandards[d.dTable.NotationStandard].ExpressionLanguage[v.ExpressionLanguage()]; !ok {
			errResult = append(errResult, ErrDTableEntryExpressionLangInvalid)
		}

		if i < len(d.dTable.OutputFields) {
			if ok, err := d.validateOutputEntry(d.dTable.OutputFields[i], v); !ok {
				errResult = append(errResult, err...)
			}
		}
	}

	if len(errResult) != 0 {
		return false, errResult
	}

	return true, nil
}

func (d Validator) validateInputEntry(f model.Field, v model.Entry) (bool, []error) {
	prs, err := d.prs.GetParser(v.ExpressionLanguage())
	if err != nil {
		return false, []error{err}
	}

	if ok, errs := prs.ValidateInputEntry(f, v); !ok {
		return false, errs
	}

	return true, nil
}

func (d Validator) validateOutputEntry(f model.Field, v model.Entry) (bool, []error) {
	prs, err := d.prs.GetParser(v.ExpressionLanguage())
	if err != nil {
		return false, []error{err}
	}

	if ok, errs := prs.ValidateOutputEntry(f, v); !ok {
		return false, errs
	}

	return true, nil
}

func (d Validator) ValidateContainsInterferences() bool {
	output := d.dTable.OutputFields

	for _, val := range d.dTable.InputFields {
		if d.fieldIsContained(val, output) {
			return true
		}
	}

	return false
}

func (d Validator) fieldIsContained(field model.Field, setOfFields []model.Field) bool {
	for _, val := range setOfFields {
		if (val.Label == field.Label) && (val.Name == field.Name) {
			return true
		}
	}
	return false
}
