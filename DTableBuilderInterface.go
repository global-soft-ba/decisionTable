package main

import (
	"decisionTable/model"
)

type DTableBuilderInterface interface {
	Build() (DecisionTable, []error)
	SetExpressionLanguage(lang string) DTableBuilderInterface
	SetDefinitionKey(key string) DTableBuilderInterface
	SetName(name string) DTableBuilderInterface
	SetHitPolicy(policy string) DTableBuilderInterface
	SetCollectOperator(collector string) DTableBuilderInterface
	AddInputField(name string, label string, typ string) DTableBuilderInterface
	AddOutputField(name string, label string, typ string) DTableBuilderInterface
	AddRule(input []model.Entry, output []model.Entry, description string) DTableBuilderInterface
}
