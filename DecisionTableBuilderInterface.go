package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type DecisionTableBuilderInterface interface {
	SetID(id string) DecisionTableBuilderInterface
	SetName(name string) DecisionTableBuilderInterface
	SetHitPolicy(hitPolicy hitPolicy.HitPolicy) DecisionTableBuilderInterface
	SetCollectOperator(collectOperator collectOperator.CollectOperator) DecisionTableBuilderInterface
	SetExpressionLanguage(expressionLanguage expressionLanguage.ExpressionLanguage) DecisionTableBuilderInterface
	SetStandard(standard standard.Standard) DecisionTableBuilderInterface
	AddInputField(inputField field.Field) DecisionTableBuilderInterface
	AddOutputField(outputField field.Field) DecisionTableBuilderInterface
	AddRule(rule Rule) DecisionTableBuilderInterface
	Build() (DecisionTable, error)
	BuildWithoutValidation() DecisionTable
}
