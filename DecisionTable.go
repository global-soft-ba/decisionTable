package main

import (
	"decisionTable/model"
)

func CreateDecisionTable() DTableBuilderInterface {
	d := DTableBuilder{}
	return d
}

type DecisionTable struct {
	key             string
	name            string
	hitPolicy       string
	collectOperator string
	DTableStandard  string
	valid           bool

	inputFields  []model.Field
	outputFields []model.Field
	rules        []model.Rule
}

func (d DecisionTable) ExpressionLanguage() string {
	return d.DTableStandard
}

func (d DecisionTable) Valid() bool {
	return d.valid
}

func (d DecisionTable) Key() string {
	return d.key
}

func (d DecisionTable) Name() string {
	return d.name
}

func (d DecisionTable) HitPolicy() string {
	return d.hitPolicy
}

func (d DecisionTable) CollectOperator() string {
	return d.collectOperator
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
