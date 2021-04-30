package types

import "decisionTable/model"

type DecisionTableConfig struct {
	ExpressionLanguage map[model.ExpressionLanguage]string
	HitPolicies        map[model.HitPolicy]string
	CollectOperators   map[model.CollectOperator]string
	VariableType       map[model.VariableTyp]string
}

var DecisionTableStandards = map[model.DTableStandard]DecisionTableConfig{
	model.GRULE: grule,
	model.DMN:   dmn,
}

var dmn = DecisionTableConfig{
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

	VariableType: map[model.VariableTyp]string{
		model.String:  "",
		model.Boolean: "",
		model.Integer: "",
		model.Long:    "",
		model.Double:  "",
		model.Date:    "",
	}}

var grule = DecisionTableConfig{
	ExpressionLanguage: map[model.ExpressionLanguage]string{
		model.SFEEL: "",
	},

	HitPolicies: map[model.HitPolicy]string{
		model.Unique:   "",
		model.First:    "",
		model.Priority: "",
	},

	CollectOperators: map[model.CollectOperator]string{},

	VariableType: map[model.VariableTyp]string{
		model.String:  "",
		model.Boolean: "",
		model.Integer: "",
		model.Float:   "",
		model.Date:    "",
	}}
