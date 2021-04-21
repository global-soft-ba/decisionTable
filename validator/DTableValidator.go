package validator

import (
	"decisionTable/model"
	"errors"
)

var (
	ErrRuleHaveDifferentAmountOfInputFields  = errors.New("amount of input fields does not match fields of decisiontable")
	ErrRuleHaveDifferentAmountOfOutputFields = errors.New("amount of output fields does not match fields of decisiontable")
)

func CreateDTableValidator(input []model.Field, output []model.Field, rules []model.Rule) DTableValidatorInterface {
	r := DTableValidator{valid: false, errors: []error{}}
	return r
}

type DTableValidator struct {
	valid  bool
	errors []error

	input  []model.Field
	output []model.Field
	rules  []model.Rule
}

func (d DTableValidator) Validate() (bool, []error) {
	panic("implement me")

	/*
	   	if len(d.dTable.inputs) != len(input) {return d, exceptions.ErrRuleHaveDifferentAmountOfInputFields}
	   	if len(d.dTable.outputs) != len(output) {return d, exceptions.ErrRuleHaveDifferentAmountOfOutputFields}

	       ValidateRules(rules []model.Rule) (bool, []error)
	      	ValidateExpressionLanguage(rules []model.Rule) (bool, []error)
	      	ValidateInput(input []model.Field) (bool, []error)
	      	ValidateOutput(output []model.Field) (bool, []error)
	*/
}
