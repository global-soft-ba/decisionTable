package main

import (
	"decisionTable/convert/interfaces"
	"decisionTable/data"
	"errors"
)

var (
	ErrDTableNotValid = errors.New("decision table must be valid before converting")
)

func CreateDecisionTable() DecisionTableBuilderInterface {
	d := DecisionTableBuilder{}
	return d
}

type DecisionTable struct {
	data  data.Table
	valid bool
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

func (d DecisionTable) Valid() bool {
	return d.valid
}

func (d DecisionTable) InputFields() []data.Field {
	return d.data.InputFields
}

func (d DecisionTable) OutputFields() []data.Field {
	return d.data.OutputFields
}

func (d DecisionTable) Rules() []data.Rule {
	return d.data.Rules
}

func (d DecisionTable) Interferences() bool {
	return d.data.Interferences
}

func (d DecisionTable) Convert(converter interfaces.ConverterInterface) (interface{}, error) {

	if !d.valid {
		return []string{}, ErrDTableNotValid
	}
	return converter.Convert(d.data)
}
