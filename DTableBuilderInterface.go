package main

import (
	"decisionTable/model"
)

type DTableBuilderInterface interface {
	Build() (DecisionTable, []error)
	SetDefinitionKey(key string) DTableBuilderInterface
	SetName(name string) DTableBuilderInterface
	SetNotationStandard(lang model.DTableStandard) DTableBuilderInterface
	SetHitPolicy(policy model.HitPolicy) DTableBuilderInterface
	SetCollectOperator(collector model.CollectOperator) DTableBuilderInterface
	AddInputField(name string, label string, typ model.VariableTyp) DTableBuilderInterface
	AddOutputField(name string, label string, typ model.VariableTyp) DTableBuilderInterface
	AddRule(input []model.Entry, output []model.Entry, description string) DTableBuilderInterface
}
