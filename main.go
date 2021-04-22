package main

import (
	"decisionTable/constant"
	"decisionTable/model"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	t, err := CreateDecisionTable().
		SetName("test").SetDTableStandard(constant.DecisionTableStandardGrule).
		SetDefinitionKey("Key1").
		SetHitPolicy("FIRST").
		AddInputField("name", "person", "STRING").
		AddOutputField("out2", "properties", "STRING").
		AddRule(
			[]model.Entry{model.Entry{Expression: "=\"Max\"", ExpressionLanguage: "GRL"}}, //In
			[]model.Entry{model.Entry{Expression: "Dear Max", ExpressionLanguage: "GRL"}}, //Out
			"salutation").
		Build()
	fmt.Println(t)
	fmt.Println(err)
}
