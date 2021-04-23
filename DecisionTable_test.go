package main

import (
	"decisionTable/model"
	"testing"
)

func TestDecisionTablee(t *testing.T) {
	testTable, err := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("Type of claim", "claim", model.String).
		AddInputField("Expenditure of claim", "claim", model.Integer).
		AddOutputField("Responsible employee", "Employee", model.String).
		AddOutputField("4 eyes principle", "Employee", model.Boolean).
		AddRule(
			[]model.Entry{model.Entry{Expression: "Car Accident", ExpressionLanguage: model.GRL},
				model.Entry{Expression: "<1000", ExpressionLanguage: model.GRL}},
			[]model.Entry{model.Entry{Expression: "Müller", ExpressionLanguage: model.GRL},
				model.Entry{Expression: "false", ExpressionLanguage: model.GRL}},
			"").
		AddRule(
			[]model.Entry{model.Entry{Expression: "Car Accident", ExpressionLanguage: model.GRL},
				model.Entry{Expression: "[1000..10000]", ExpressionLanguage: model.GRL}},
			[]model.Entry{model.Entry{Expression: "Meier", ExpressionLanguage: model.GRL},
				model.Entry{Expression: "false", ExpressionLanguage: model.GRL}},
			"").
		AddRule(
			[]model.Entry{model.Entry{Expression: "Car Accident", ExpressionLanguage: model.GRL},
				model.Entry{Expression: ">=10000", ExpressionLanguage: model.GRL}},
			[]model.Entry{model.Entry{Expression: "Schmidt", ExpressionLanguage: model.GRL},
				model.Entry{Expression: "true", ExpressionLanguage: model.GRL}},
			"").
		Build()

	if got := testTable.valid; got != true {
		t.Errorf("Example Decision Table is invalid, got = %v, want true, with error %v", got, err)
	}
}
