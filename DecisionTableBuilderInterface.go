package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
)

type DecisionTableBuilderInterface interface {
	SetID(id string) DecisionTableBuilderInterface
	SetName(name string) DecisionTableBuilderInterface
	SetHitPolicy(hitPolicy data.HitPolicy) DecisionTableBuilderInterface
	SetCollectOperator(collectOperator data.CollectOperator) DecisionTableBuilderInterface
	SetExpressionLanguage(expressionLanguage data.ExpressionLanguage) DecisionTableBuilderInterface
	SetStandard(standard data.Standard) DecisionTableBuilderInterface
	AddInputField(inputField data.FieldInterface) DecisionTableBuilderInterface
	AddOutputField(outputField data.FieldInterface) DecisionTableBuilderInterface
	AddRule(rule Rule) DecisionTableBuilderInterface
	Build() (DecisionTable, error)
	BuildWithoutValidation() DecisionTable
}
