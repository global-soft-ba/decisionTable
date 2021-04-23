package types

import "decisionTable/model"

type DTableConfig struct {
	ExpressionLanguage map[model.ExpressionLanguage]string
	HitPolicies        map[model.HitPolicy]string
	CollectOperators   map[model.CollectOperator]string
	VariableType       map[model.VariableTyp]string
}

var NotationStandards = map[model.DTableStandard]DTableConfig{
	model.GRULE: gruleStandard,
	model.DMN:   dmnStandard,
}

var dmnStandard = DTableConfig{
	ExpressionLanguage: map[model.ExpressionLanguage]string{
		model.FEEL:       "",
		model.Javascript: "",
		model.Python:     "",
		model.Groovy:     "",
		model.JRuby:      "",
		model.Juel:       "",
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

var gruleStandard = DTableConfig{
	ExpressionLanguage: map[model.ExpressionLanguage]string{
		model.GRL: "",
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
