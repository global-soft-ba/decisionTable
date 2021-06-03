package main

import (
	"decisionTable/data"
)

type DecisionTableBuilderInterface interface {
	Build() (DecisionTable, []error)
	SetDefinitionKey(key string) DecisionTableBuilderInterface
	SetName(name string) DecisionTableBuilderInterface
	SetNotationStandard(lang data.DTableStandard) DecisionTableBuilderInterface
	SetHitPolicy(policy data.HitPolicy) DecisionTableBuilderInterface
	SetCollectOperator(collector data.CollectOperator) DecisionTableBuilderInterface
	AddInputField(name string, label string, typ data.DataTyp) DecisionTableBuilderInterface
	AddOutputField(name string, label string, typ data.DataTyp) DecisionTableBuilderInterface
	AddRule(description string) DecisionTableRuleBuilderInterface
}

type DecisionTableRuleBuilderInterface interface {
	AddInputEntry(expr string, exprLang data.ExpressionLanguage) DecisionTableRuleBuilderInterface
	AddOutputEntry(expr string, exprLang data.ExpressionLanguage) DecisionTableRuleBuilderInterface
	BuildRule() DecisionTableBuilderInterface
}
