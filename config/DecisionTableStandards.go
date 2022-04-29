package types

import "github.com/global-soft-ba/decisionTable/data"

type DecisionTableStandard struct {
	ExpressionLanguages map[data.ExpressionLanguage]bool
	HitPolicies         map[data.HitPolicy]bool
	CollectOperators    map[data.CollectOperator]bool
	VariableTypes       map[data.DataType]bool
}

var DecisionTableStandards = map[data.Standard]DecisionTableStandard{
	data.GRULE: grule,
	data.DMN:   dmn,
}

var dmn = DecisionTableStandard{
	ExpressionLanguages: map[data.ExpressionLanguage]bool{
		data.FEEL: true,
	},

	HitPolicies: map[data.HitPolicy]bool{
		data.Unique:      true,
		data.First:       true,
		data.Priority:    true,
		data.Any:         true,
		data.Collect:     true,
		data.RuleOrder:   true,
		data.OutputOrder: true,
	},

	CollectOperators: map[data.CollectOperator]bool{
		data.List:  true,
		data.Sum:   true,
		data.Min:   true,
		data.Max:   true,
		data.Count: true,
	},

	VariableTypes: map[data.DataType]bool{
		data.Boolean:  true,
		data.Integer:  true,
		data.Long:     true,
		data.Double:   true,
		data.String:   true,
		data.DateTime: true,
	},
}

var grule = DecisionTableStandard{
	ExpressionLanguages: map[data.ExpressionLanguage]bool{
		data.SFEEL: true,
	},

	HitPolicies: map[data.HitPolicy]bool{
		data.Unique:   true,
		data.First:    true,
		data.Priority: true,
	},

	CollectOperators: map[data.CollectOperator]bool{},

	VariableTypes: map[data.DataType]bool{
		data.Boolean:  true,
		data.Integer:  true,
		data.Float:    true, // TODO: In GRL, Float is REAL
		data.String:   true,
		data.DateTime: true,
	},
}
