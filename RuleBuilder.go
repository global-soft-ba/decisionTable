package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
)

type RuleBuilder struct {
	data data.Rule
}

func NewRuleBuilder() RuleBuilderInterface {
	return RuleBuilder{}
}

func (r RuleBuilder) SetAnnotation(annotation string) RuleBuilderInterface {
	r.data.Annotation = annotation
	return r
}

func (r RuleBuilder) AddInputEntry(expression string) RuleBuilderInterface {
	r.data.InputExpressions = append(r.data.InputExpressions, expression)
	return r
}

func (r RuleBuilder) AddOutputEntry(expression string) RuleBuilderInterface {
	r.data.OutputExpressions = append(r.data.OutputExpressions, expression)
	return r
}

func (r RuleBuilder) Build() Rule {
	return Rule{data: r.data}
}
