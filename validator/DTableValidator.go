package validator

import (
	"decisionTable/constant"
	"decisionTable/model"
	"errors"
)

var (
	ErrDTableNameEmpty            = errors.New("name of decision table is required")
	ErrDTableKeyEmpty             = errors.New("definition key of decision table is required")
	ErrDTableHitPolicy            = errors.New("hit policy of decision table is not supported")
	ErrDTableCollectOperator      = errors.New("collect operator of decision table is not supported")
	ErrDTableEmptyCollectOperator = errors.New("collect operator of decision table cannot be empty for the hit policy")

	ErrRuleHaveDifferentAmountOfInputFields  = errors.New("amount of input fields does not match fields of decisiontable")
	ErrRuleHaveDifferentAmountOfOutputFields = errors.New("amount of output fields does not match fields of decisiontable")
)

func CreateDTableValidator(dTable model.DTableData) DTableValidatorInterface {
	r := DTableValidator{dTable: dTable}
	return r
}

type DTableValidator struct {
	dTable model.DTableData
}

func (d DTableValidator) Validate() (bool, []error) {
	return true, nil
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

func (d DTableValidator) validateExpresionLanguage() (bool, error) {
	if _, ok := constant.ExpressionLanguage[d.dTable.ExpressionLanguage]; !ok {
		return false, ErrDTableKeyEmpty
	}
	return true, nil
}

func (d DTableValidator) validateHitPolicy() (bool, error) {
	if _, ok := constant.HitPolicies[d.dTable.HitPolicy]; !ok {
		return false, ErrDTableHitPolicy
	}

	if d.dTable.HitPolicy == constant.CollectOperatorPolicy {
		if len(d.dTable.CollectOperator) == 0 {
			return false, ErrDTableEmptyCollectOperator
		}

		if ok, err := d.validateCollectOperator(); !ok {
			return false, err
		}
	}

	return true, nil
}

func (d DTableValidator) validateCollectOperator() (bool, error) {
	if _, ok := constant.CollectOperators[d.dTable.CollectOperator]; !ok {
		return false, ErrDTableCollectOperator
	}
	return true, nil
}
