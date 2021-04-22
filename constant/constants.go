package constant

const (
	CollectOperatorPolicy      = "COLLECT"
	DecisionTableStandardGrule = "GRL"
	DecisionTableStandardDMN   = "DMN"
)

var DTableStandards = map[string]string{
	DecisionTableStandardGrule: "GRULE Engine Gopher Holds The Rules",
	DecisionTableStandardDMN:   "Decision Model and Notation - DMN (OMG)",
}

type DecisionTableConfig struct {
	ExpressionLanguage map[string]string
	HitPolicies        map[string]string
	CollectOperators   map[string]string
	VariableType       map[string]string
}

var AllowedConfigForStandard = map[string]DecisionTableConfig{
	DecisionTableStandardDMN:   dmnStandard,
	DecisionTableStandardGrule: gruleStandard,
}

var dmnStandard = DecisionTableConfig{
	ExpressionLanguage: map[string]string{
		"FEEL":       "feel",
		"JAVASCRIPT": "javascript",
		"PYTHON":     "python",
		"GROOVY":     "groovy",
		"JRUBY":      "jruby",
		"JUEL":       "juel",
	},

	HitPolicies: map[string]string{
		"UNIQUE":              "Unique Hit Policy ",
		"FIRST":               "First Hit Policy ",
		"PRIORITY":            "Priority Hit Policy ",
		"ANY":                 "Any Hit Policy ",
		"RULE_ORDER":          "Rule Order Hit Policy ",
		"OUTPUT_ORDER":        "Output Order Hit Policy ",
		CollectOperatorPolicy: "Collect Hit Policy ",
	},

	CollectOperators: map[string]string{
		"LIST":  "List Operator ",
		"SUM":   "Sum Operator ",
		"MIN":   "Min Operator ",
		"MAX":   "Max Operator ",
		"COUNT": "Count Operator ",
	},

	VariableType: map[string]string{
		"STRING":  "string",
		"BOOLEAN": "boolean",
		"INTEGER": "integer",
		"LONG":    "long",
		"DOUBLE":  "double",
		"DATE":    "date",
	}}

var gruleStandard = DecisionTableConfig{
	ExpressionLanguage: map[string]string{
		"GRL": "grl",
	},

	HitPolicies: map[string]string{
		"FIRST":    "Unique Hit Policy ",
		"PRIORITY": "Priority Hit Policy ",
	},

	CollectOperators: map[string]string{},

	VariableType: map[string]string{
		"STRING":  "string",
		"BOOLEAN": "boolean",
		"INTEGER": "integer",
		"FLOAT":   "float",
		"DATE":    "date",
	}}
