package constant

var ExpressionLanguage = map[string]string{
	"GRULE": "Gopher Hold The Rules Engine Language",
	"DMN":   "Decision Model and Notation (OMG)",
}

var HitPolicies = map[string]string{
	"UNIQUE":              "Unique Hit Policy ",
	"FIRST":               "First Hit Policy ",
	"PRIORITY":            "Priority Hit Policy ",
	"ANY":                 "Any Hit Policy ",
	CollectOperatorPolicy: "Collect Hit Policy ",
	"RULE_ORDER":          "Rule Order Hit Policy ",
	"OUTPUT_ORDER":        "Output Order Hit Policy ",
}

const CollectOperatorPolicy = "COLLECT"

var CollectOperators = map[string]string{
	"LIST":  "List Operator ",
	"SUM":   "Sum Operator ",
	"MIN":   "Min Operator ",
	"MAX":   "Max Operator ",
	"COUNT": "Count Operator ",
}
