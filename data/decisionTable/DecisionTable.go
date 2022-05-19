package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/rule"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type DecisionTable struct {
	ID                 string                                `json:"id"`
	Name               string                                `json:"name"`
	HitPolicy          hitPolicy.HitPolicy                   `json:"hitPolicy"`
	CollectOperator    collectOperator.CollectOperator       `json:"collectOperator"`
	ExpressionLanguage expressionLanguage.ExpressionLanguage `json:"expressionLanguage"`
	Standard           standard.Standard                     `json:"standard"`

	InputFields  []field.Field `json:"inputFields"`
	OutputFields []field.Field `json:"outputFields"`
	Rules        []rule.Rule   `json:"rules"`
}
