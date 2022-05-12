package decisionTable

import "github.com/global-soft-ba/decisionTable/data"

type Rule struct {
	data data.Rule
}

func (r Rule) Annotation() string {
	return r.data.Annotation
}

func (r Rule) InputEntries() []string {
	return r.data.InputExpressions
}

func (r Rule) OutputEntries() []string {
	return r.data.OutputExpressions
}
