# Decision Table Representation
* simple data structure
* support of DMN and GRULE table standardization
* no execution - only representation of decision tables 
* value object characteristics
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

We assume that the frontend will represent a decision table as a kind of table. In case, that a user changes something in the frontend table, we drop the old decision table representation and rebuild from the new 
(frontend) table (which was manipulated by the user). By doing so, we don't need any "insert function" for rules/inputs/outputs etc and so it keeps the DecisionTable-Type simple. 

## Hit Policies and Collect Operators
A decision table consists of several rules, typically represented as rows. When reading such a row we look at certain input values and deduct a certain result represented by output values. 

When using the simplest hit policy "unique" (U), such rules do not overlap: only a single rule must match. 

Now consider that we build a decision table with overlapping rules. In other words that means more than one rule may match a given set of input values. We then need one of the alternative hit policy indicators to unambiguously understand the decision logic according to which such rules are interpreted.

### Single Decision Tables 
Such tables either return the output of only one rule or aggregate the output of many rules into one result. The hit policies to be considered are
* **Unique:** Rules do not overlap. Only a single rule can match.
* **First:** Rules are evaluated from top to bottom. Rules may overlap, but only the first match counts.
* **Priority:** Rule outputs are prioritized. Rules may overlap, but only the match with the highest output priority counts.
* **Any:** Multiple matching rules must not make a difference: all matching rules must lead to the same output.
* **Collect:** The output of all matching rules is aggregated by means of an operator:
    * **Sum:** Sums/Add all outputs of the matching rule’s distinct outputs.
    * **Minimum:** Take the smallest value of all the matching rule’s outputs.
    * **Maximum:** Take the largest value of all the matching rule’s outputs.
    * **Number:** Return the number of all the matching rule’s distinct outputs. 

### Multiple Result DecisionTable
Multiple result tables may return the output of multiple rules. The hit policies for such tables are:

* **Collect:** All outputs of the matching rules are combined in an arbitrarily ordered list of all the output entries.
  * **List:** Collect all outputs of all matching rules in a list of output entries
* **Rule Order:** All outputs of the matching rules are combined in a list of outputs ordered by the sequence of those rules in the decision table.
* **Output Order:** All outputs of the matching rules are combined in a list of outputs ordered by their (decreasing) output priority.

### Examples
Examples can be found here: [Camunda Page](https://camunda.com/best-practices/choosing-the-dmn-hit-policy/#_knowing_the_dmn_hit_policy_strong_basics_strong)
