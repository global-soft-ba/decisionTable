package data

type Rule struct {
	Description   string
	InputEntries  []EntryInterface
	OutputEntries []EntryInterface

	InputExpressions  []string
	OutputExpressions []string
}
