package decisionTable

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/conv"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/valid"
)

var (
	ErrDTableNotValid = errors.New("decision table must be valid before converting")
)

func CreateDecisionTable() DecisionTableBuilderInterface {
	d := DecisionTableBuilder{}
	return d
}

type DecisionTable struct {
	data data.Table
}

func (d DecisionTable) Key() string {
	return d.data.Key
}

func (d DecisionTable) Name() string {
	return d.data.Name
}

func (d DecisionTable) HitPolicy() data.HitPolicy {
	return d.data.HitPolicy
}

func (d DecisionTable) CollectOperator() data.CollectOperator {
	return d.data.CollectOperator
}

func (d DecisionTable) NotationStandard() data.DTableStandard {
	return d.data.NotationStandard
}

func (d DecisionTable) InputFields() []data.FieldInterface {
	return d.data.InputFields
}

func (d DecisionTable) OutputFields() []data.FieldInterface {
	return d.data.OutputFields
}

func (d DecisionTable) Rules() []data.Rule {
	return d.data.Rules
}

func (d DecisionTable) Interferences() bool {
	return d.data.Interferences
}

//ToDo define explicit output format type (instead of string) into engine standard (notation standard)
func (d DecisionTable) Convert(format string) (interface{}, error) {
	if ok, _ := d.Validate(); !ok {
		return []string{}, ErrDTableNotValid
	}

	tableConverter := conv.CreateConverter()

	res, err := tableConverter.Convert(d.data, format)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *DecisionTable) Validate() (bool, []error) {
	validator := valid.CreateDecisionTableValidator()
	val, err := validator.Validate(d.data)
	if val != true {
		return false, err
	}
	d.data.Interferences = validator.ValidateContainsInterferences(d.data)
	return true, []error(nil)
}
