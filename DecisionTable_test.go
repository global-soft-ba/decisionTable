package main

import (
	"decisionTable/model"
	"testing"
)

func TestDecisionTable(t *testing.T) {
	testTable, err := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("Type of claim", "claim", model.String).
		AddInputField("Expenditure of claim", "claim", model.Integer).
		AddOutputField("Responsible employee", "Employee", model.String).
		AddOutputField("4 eyes principle", "Employee", model.Boolean).
		AddRule("R1").
		AddInputEntry("Car Accident", model.GRL).
		AddInputEntry("<1000", model.GRL).
		AddOutputEntry("Müller", model.GRL).
		AddOutputEntry("false", model.GRL).BuildRule().
		AddRule("R2").
		AddInputEntry("Car Accident", model.GRL).
		AddInputEntry("[1000..10000]", model.GRL).
		AddOutputEntry("Meier", model.GRL).
		AddOutputEntry("false", model.GRL).BuildRule().
		AddRule("R3").
		AddInputEntry("Car Accident", model.GRL).
		AddInputEntry(">=10000", model.GRL).
		AddOutputEntry("Schmidt", model.GRL).
		AddOutputEntry("true", model.GRL).BuildRule().
		Build()

	if got := testTable.valid; got != true {
		t.Errorf("Example Decision Table is invalid, got = %v, want true, with error %v", got, err)
	}
}

func TestDecisionTableEmptyEntry(t *testing.T) {
	testTable, err := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("Type of claim", "claim", model.String).
		AddOutputField("Responsible employee", "Employee", model.String).
		AddRule("R1").
		AddInputEntry("Car Accident", model.GRL).
		AddOutputEntry("-", model.GRL).
		BuildRule().
		Build()

	if got := testTable.Rules()[0].OutputEntries[0].EmptyExpression(); got != true {
		t.Errorf("EmptyEntry in decision table is invalid, got = %v, want = %v true, with error %v", got, true, err)
	}
}
