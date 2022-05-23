package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/rule"
)

type Rule struct {
	data rule.Rule
}

func (r Rule) Annotation() string {
	return r.data.Annotation
}

func (r Rule) InputEntries() []string {
	return r.data.InputEntries
}

func (r Rule) OutputEntries() []string {
	return r.data.OutputEntries
}
