package main

import (
	"github.com/global-soft-ba/decisionTable/data"
)

type DecisionTableBuilderInterface interface {
	Build() (DecisionTable, []error)
	SetDefinitionKey(key string) DecisionTableBuilderInterface
	SetName(name string) DecisionTableBuilderInterface
	SetNotationStandard(lang data.DTableStandard) DecisionTableBuilderInterface
	SetHitPolicy(policy data.HitPolicy) DecisionTableBuilderInterface
	SetCollectOperator(collector data.CollectOperator) DecisionTableBuilderInterface
	AddInputField(field data.FieldInterface) DecisionTableBuilderInterface
	AddOutputField(field data.FieldInterface) DecisionTableBuilderInterface
	AddRule(description string) DecisionTableRuleBuilderInterface
}

type DecisionTableRuleBuilderInterface interface {
	AddInputEntry(expr string, exprLang data.ExpressionLanguage) DecisionTableRuleBuilderInterface
	AddOutputEntry(expr string, exprLang data.ExpressionLanguage) DecisionTableRuleBuilderInterface
	BuildRule() DecisionTableBuilderInterface
}
