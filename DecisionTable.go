package main

import (
	"decisionTable/converter"
	"decisionTable/model"
	"errors"
)

var (
	ErrDTableNotValid = errors.New("decision table must be valid before converting")
)

func CreateDecisionTable() DTableBuilderInterface {
	d := DTableBuilder{}
	return d
}

type DecisionTable struct {
	key              string
	name             string
	hitPolicy        model.HitPolicy
	collectOperator  model.CollectOperator
	notationStandard model.DTableStandard
	valid            bool

	inputFields  []model.Field
	outputFields []model.Field
	rules        []model.Rule
}

func (d DecisionTable) Key() string {
	return d.key
}

func (d DecisionTable) Name() string {
	return d.name
}

func (d DecisionTable) HitPolicy() model.HitPolicy {
	return d.hitPolicy
}

func (d DecisionTable) CollectOperator() model.CollectOperator {
	return d.collectOperator
}

func (d DecisionTable) NotationStandard() model.DTableStandard {
	return d.notationStandard
}

func (d DecisionTable) Valid() bool {
	return d.valid
}

func (d DecisionTable) InputFields() []model.Field {
	return d.inputFields
}

func (d DecisionTable) OutputFields() []model.Field {
	return d.outputFields
}

func (d DecisionTable) Rules() []model.Rule {
	return d.rules
}

func (d DecisionTable) Convert(converter converter.DTableConverterInterface) ([]byte, error) {

	if !d.valid {
		return []byte{}, ErrDTableNotValid
	}

	dTable := model.DTableData{
		Key:              d.key,
		Name:             d.name,
		HitPolicy:        d.hitPolicy,
		CollectOperator:  d.collectOperator,
		NotationStandard: d.notationStandard,
		InputFields:      d.inputFields,
		OutputFields:     d.outputFields,
		Rules:            d.rules,
	}
	return converter.Convert(dTable)
}
