package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
)

type RuleBuilder struct {
	rule data.Rule
}

func NewRuleBuilder() RuleBuilderInterface {
	return RuleBuilder{}
}

func (r RuleBuilder) SetAnnotation(annotation string) RuleBuilderInterface {
	r.rule.Description = annotation
	return r
}

func (r RuleBuilder) AddInputEntry(expression string) RuleBuilderInterface {
	r.rule.InputExpressions = append(r.rule.InputExpressions, expression)
	return r
}

func (r RuleBuilder) AddOutputEntry(expression string) RuleBuilderInterface {
	r.rule.OutputExpressions = append(r.rule.OutputExpressions, expression)
	return r
}

func (r RuleBuilder) Build() Rule {
	return Rule{data: r.rule}
}
