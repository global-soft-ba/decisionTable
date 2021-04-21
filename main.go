package main

import (
	"decisionTable/model"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	t, _ := CreateDecisionTable().
		SetName("test").SetExpressionLanguage("GRULE").
		AddInputField("name", "person", "string").
		AddOutputField("out2", "properties", "string").
		AddRule(
			[]model.Entry{model.Entry{Expression: "Markus", ExpressionLanguage: "GRULE"}},    //In
			[]model.Entry{model.Entry{Expression: "Du Kalbei", ExpressionLanguage: "GRULE"}}, //Out
			"salutation").
		Build()
	fmt.Println(t)
}
