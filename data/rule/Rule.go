package rule

type Rule struct {
	Annotation    string   `json:"annotation"`
	InputEntries  []string `json:"inputEntries"`
	OutputEntries []string `json:"outputEntries"`
}
