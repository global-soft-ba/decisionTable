package main

import "decisionTable/model"

type Builder struct{
	dtable DecisionTable
}

type DecisionTable struct{
	key string
	name string
	hitPolicy string
	collectOperator string
	inputs []model.Input
	outputs []model.Output
	rules []model.Rule
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

func (d DecisionTable) Inputs() []model.Input {
	return d.inputs
}

func (d DecisionTable) Outputs() []model.Output {
	return d.outputs
}

func (d DecisionTable) Rules() []model.Rule {
	return d.rules
}
