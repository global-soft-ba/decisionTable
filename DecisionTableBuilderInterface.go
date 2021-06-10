package main

import (
	"github.com/global-soft-ba/decisionTable/model"
)

type DecisionTableBuilderInterface interface {
	Build() (DecisionTable, []error)
	SetDefinitionKey(key string) DecisionTableBuilderInterface
	SetName(name string) DecisionTableBuilderInterface
	SetNotationStandard(lang model.DTableStandard) DecisionTableBuilderInterface
	SetHitPolicy(policy model.HitPolicy) DecisionTableBuilderInterface
	SetCollectOperator(collector model.CollectOperator) DecisionTableBuilderInterface
	AddInputField(name string, label string, typ model.VariableTyp) DecisionTableBuilderInterface
	AddOutputField(name string, label string, typ model.VariableTyp) DecisionTableBuilderInterface
	AddRule(description string) DecisionTableRuleBuilderInterface
}

type DecisionTableRuleBuilderInterface interface {
	AddInputEntry(expr string, exprLang model.ExpressionLanguage) DecisionTableRuleBuilderInterface
	AddOutputEntry(expr string, exprLang model.ExpressionLanguage) DecisionTableRuleBuilderInterface
	BuildRule() DecisionTableBuilderInterface
}
