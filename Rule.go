package decisionTable

import "github.com/global-soft-ba/decisionTable/data"

type Rule struct {
	data data.Rule
}

func (r Rule) Annotation() string {
	return r.data.Description
}

func (r Rule) InputEntries() []data.EntryInterface {
	return r.data.InputEntries
}

func (r Rule) OutputEntries() []data.EntryInterface {
	return r.data.OutputEntries
}
