package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/lang/sfeel"
)

type DecisionTableBuilder struct {
	decisionTable data.DecisionTable
}

func NewDecisionTableBuilder() DecisionTableBuilderInterface {
	return DecisionTableBuilder{}
}

func (d DecisionTableBuilder) SetID(id string) DecisionTableBuilderInterface {
	d.decisionTable.ID = id
	return d
}

func (d DecisionTableBuilder) SetName(name string) DecisionTableBuilderInterface {
	d.decisionTable.Name = name
	return d
}

func (d DecisionTableBuilder) SetHitPolicy(hitPolicy data.HitPolicy) DecisionTableBuilderInterface {
	d.decisionTable.HitPolicy = hitPolicy
	return d
}

func (d DecisionTableBuilder) SetCollectOperator(collectOperator data.CollectOperator) DecisionTableBuilderInterface {
	d.decisionTable.CollectOperator = collectOperator
	return d
}

func (d DecisionTableBuilder) SetExpressionLanguage(expressionLanguage data.ExpressionLanguage) DecisionTableBuilderInterface {
	d.decisionTable.ExpressionLanguage = expressionLanguage
	return d
}

func (d DecisionTableBuilder) SetStandard(standard data.Standard) DecisionTableBuilderInterface {
	d.decisionTable.Standard = standard
	return d
}

func (d DecisionTableBuilder) AddInputField(inputField data.FieldInterface) DecisionTableBuilderInterface {
	d.decisionTable.InputFields = append(d.decisionTable.InputFields, inputField)
	return d
}

func (d DecisionTableBuilder) AddOutputField(outputField data.FieldInterface) DecisionTableBuilderInterface {
	d.decisionTable.OutputFields = append(d.decisionTable.OutputFields, outputField)
	return d
}

func (d DecisionTableBuilder) AddRule(rule Rule) DecisionTableBuilderInterface {
	d.decisionTable.Rules = append(d.decisionTable.Rules, rule.data)
	return d
}

func (d DecisionTableBuilder) covertExpressionsIntoEntries(decisionTable data.DecisionTable, expressionLanguage data.ExpressionLanguage) data.DecisionTable {
	for i, rule := range decisionTable.Rules {
		for _, inputExpression := range rule.InputExpressions {
			switch expressionLanguage {
			case data.SFEEL:
				inputEntry := sfeel.CreateInputEntry(inputExpression)
				decisionTable.Rules[i].InputEntries = append(decisionTable.Rules[i].InputEntries, inputEntry)
			}
		}

		for _, outputExpression := range rule.OutputExpressions {
			switch expressionLanguage {
			case data.SFEEL:
				outputEntry := sfeel.CreateOutputEntry(outputExpression)
				decisionTable.Rules[i].OutputEntries = append(decisionTable.Rules[i].OutputEntries, outputEntry)
			}
		}
	}

	return decisionTable
}

func (d DecisionTableBuilder) Build() (DecisionTable, error) {
	d.decisionTable = d.covertExpressionsIntoEntries(d.decisionTable, d.decisionTable.ExpressionLanguage)
	decisionTable := DecisionTable{data: d.decisionTable}

	if err := decisionTable.Validate(decisionTable.Standard()); err != nil {
		return DecisionTable{}, err
	}

	return decisionTable, nil
}

func (d DecisionTableBuilder) BuildWithoutValidation() DecisionTable {
	d.decisionTable = d.covertExpressionsIntoEntries(d.decisionTable, d.decisionTable.ExpressionLanguage)
	return DecisionTable{data: d.decisionTable}
}
