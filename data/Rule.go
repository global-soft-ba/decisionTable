package data

type Rule struct {
	Annotation        string   `json:"annotation"`
	InputExpressions  []string `json:"inputEntries"`
	OutputExpressions []string `json:"outputEntries"`

	InputEntries  []EntryInterface `json:"-"`
	OutputEntries []EntryInterface `json:"-"`
}
