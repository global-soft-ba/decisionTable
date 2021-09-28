package types

import "github.com/global-soft-ba/decisionTable/data"

type TableConfigStandard struct {
	ExpressionLanguage map[data.ExpressionLanguage]string
	HitPolicies        map[data.HitPolicy]string
	CollectOperators   map[data.CollectOperator]string
	VariableType       map[data.DataTyp]string
}

var DecisionTableStandards = map[data.DTableStandard]TableConfigStandard{
	data.GRULE: grule,
	data.DMN:   dmn,
}

var dmn = TableConfigStandard{
	ExpressionLanguage: map[data.ExpressionLanguage]string{
		data.FEEL: "",
	},

	HitPolicies: map[data.HitPolicy]string{
		data.Unique:      "",
		data.First:       "",
		data.Priority:    "",
		data.Any:         "",
		data.RuleOrder:   "",
		data.OutputOrder: "",
		data.Collect:     "",
	},

	CollectOperators: map[data.CollectOperator]string{
		data.List:  "",
		data.Sum:   "",
		data.Min:   "",
		data.Max:   "",
		data.Count: "",
	},

	VariableType: map[data.DataTyp]string{
		data.String:   "",
		data.Boolean:  "",
		data.Integer:  "",
		data.Long:     "",
		data.Double:   "",
		data.DateTime: "",
	}}

var grule = TableConfigStandard{
	ExpressionLanguage: map[data.ExpressionLanguage]string{
		data.SFEEL: "",
	},

	HitPolicies: map[data.HitPolicy]string{
		data.Unique:   "",
		data.First:    "",
		data.Priority: "",
	},

	CollectOperators: map[data.CollectOperator]string{},
	//TODO IN GRL Float = REAL
	VariableType: map[data.DataTyp]string{
		data.String:   "",
		data.Boolean:  "",
		data.Integer:  "",
		data.Float:    "",
		data.DateTime: "",
	}}
