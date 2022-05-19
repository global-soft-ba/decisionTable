package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/rule"
)

type RuleBuilder struct {
	data rule.Rule
}

func NewRuleBuilder() RuleBuilderInterface {
	return RuleBuilder{}
}

func (r RuleBuilder) SetAnnotation(annotation string) RuleBuilderInterface {
	r.data.Annotation = annotation
	return r
}

func (r RuleBuilder) AddInputEntry(expression string) RuleBuilderInterface {
	r.data.InputEntries = append(r.data.InputEntries, expression)
	return r
}

func (r RuleBuilder) AddOutputEntry(expression string) RuleBuilderInterface {
	r.data.OutputEntries = append(r.data.OutputEntries, expression)
	return r
}

func (r RuleBuilder) Build() Rule {
	return Rule{data: r.data}
}
