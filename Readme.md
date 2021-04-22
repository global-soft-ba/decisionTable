# Decision Table Representation
* simple data structure
* support of DMN and GRULE table standardization
* no execution - only representation of decision tables 
```
CreateDecisionTable().
		SetName("Salutation").
		SetDefinitionKey("4711-ZXY-ZXXX-293").
		SetDTableStandard(model.GRULE).
		SetHitPolicy(model.First).
		AddInputField("age", "person", model.String).
		AddOutputField("salutation", "Header", model.String).
		AddRule(
			[]model.Entry{model.Entry{Expression: ">30", ExpressionLanguage:model.GRL}}, //In
			[]model.Entry{model.Entry{Expression: "Dear Customer", ExpressionLanguage:model.GRL}}, //Out
			"salutation formal").
		AddRule(
		[]model.Entry{model.Entry{Expression: ">18 && <=30", ExpressionLanguage: model.GRL}}, //In
		[]model.Entry{model.Entry{Expression: "Hi", ExpressionLanguage: model.GRL}}, //Out
		"salutation informal").
		AddRule(
		[]model.Entry{model.Entry{Expression: "<=18", ExpressionLanguage: model.GRL}}, //In
		[]model.Entry{model.Entry{Expression: "YOLO", ExpressionLanguage: model.GRL}}, //Out
		"salutation teenHipp").
		Build()
```
## Simple Interface
We assume that the frontend will represent a decision table as a kind of table. In case, that a user changes something in the frontend table, we drop the old decision table representation and rebuild from the new 
(frontend) table (which was manipulated by the user). By doing so, we don't need any "insert function" for rules/inputs/outputs etc and so it keeps the DecisionTable-Type simple. 
