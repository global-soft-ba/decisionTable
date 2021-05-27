package types

import "github.com/global-soft-ba/decisionTable/model"

type TableConfigStandard struct {
	ExpressionLanguage map[model.ExpressionLanguage]string
	HitPolicies        map[model.HitPolicy]string
	CollectOperators   map[model.CollectOperator]string
	VariableType       map[model.DataTyp]string
}

var DecisionTableStandards = map[model.DTableStandard]TableConfigStandard{
	model.GRULE: grule,
	model.DMN:   dmn,
}

var dmn = TableConfigStandard{
	ExpressionLanguage: map[model.ExpressionLanguage]string{
		model.FEEL: "",
	},

	HitPolicies: map[model.HitPolicy]string{
		model.Unique:      "",
		model.First:       "",
		model.Priority:    "",
		model.Any:         "",
		model.RuleOrder:   "",
		model.OutputOrder: "",
		model.Collect:     "",
	},

	CollectOperators: map[model.CollectOperator]string{
		model.List:  "",
		model.Sum:   "",
		model.Min:   "",
		model.Max:   "",
		model.Count: "",
	},

	VariableType: map[model.DataTyp]string{
		model.String:   "",
		model.Boolean:  "",
		model.Integer:  "",
		model.Long:     "",
		model.Double:   "",
		model.DateTime: "",
	}}

var grule = TableConfigStandard{
	ExpressionLanguage: map[model.ExpressionLanguage]string{
		model.SFEEL: "",
	},

	HitPolicies: map[model.HitPolicy]string{
		model.Unique:   "",
		model.First:    "",
		model.Priority: "",
	},

	CollectOperators: map[model.CollectOperator]string{},
	//TODO IN GRL Float = REAL
	VariableType: map[model.DataTyp]string{
		model.String:   "",
		model.Boolean:  "",
		model.Integer:  "",
		model.Float:    "",
		model.DateTime: "",
	}}
