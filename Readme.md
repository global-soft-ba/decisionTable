# Decision Table Representation
A representation of a decision table with: 
* simple data structure
* no execution - only representation
* support of DMN and GRULE notation standards

## Table Example
![Image of Decision Table](img.png)

## Code Example
```
CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("TypeOfClaim", "claim", model.String).
		AddInputField("ExpenditureOfClaim", "claim", model.Integer).
		AddOutputField("ResponsibleEmployee", "Employee", model.String).
		AddOutputField("4EyesPrinciple", "Employee", model.Boolean).
		AddRule("R1").
          AddInputEntry(`"Car Accident"`, model.SFEEL).
          AddInputEntry("<1000", model.SFEEL).
          AddOutputEntry(`"Müller"`, model.SFEEL).
          AddOutputEntry("false", model.SFEEL).BuildRule().
		AddRule("R2").
          AddInputEntry(`"Car Accident"`, model.SFEEL).
          AddInputEntry("[1000..10000]", model.SFEEL).
          AddOutputEntry(`"Meier"`, model.SFEEL).
          AddOutputEntry("false", model.SFEEL).BuildRule().
		AddRule("R3").
          AddInputEntry("-", model.SFEEL).
          AddInputEntry(">=10000", model.SFEEL).
          AddOutputEntry("-", model.SFEEL).
          AddOutputEntry("true", model.SFEEL).BuildRule().
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

## Variable Types
The allowed types of input or output variables differs on the selected notation standard of a decision table:

Variable Types | DMN Notation | GRL Notation
------------ | ---------------|--------------
Boolean|X|X
String|X|X
Integer|X|X
Float||X   
Long|X|    
Double|X|  
Date|X|X    

## Term Languages
The expresion language defined the functions and comparisons of a single rule entry. It depends on the chosen table notations standard.

Term Language | DMN Notation | GRL Notation   
------------ | ---------------|--------------  	
SFEEL| |X
FEEL|X|       
Javascript|X|
Python|X|     
Groovy|X|     
JRuby|X|      
Juel|X|       
       
More Details can be found here:
* [GRULE](http://hyperjumptech.viewdocs.io/grule-rule-engine/GRL_en/)
* [Decision Model and Notation DMN (OMG)](https://www.omg.org/spec/DMN/1.2/PDF)

# SFEEL INPUT EXPRESSIONS
* [Simple Friendly Enough Term Language (SFEEL)](https://docs.camunda.org/manual/7.4/reference/dmn11/feel/language-elements/)
