package decisionTable

import (
	"encoding/json"
	"fmt"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"github.com/global-soft-ba/decisionTable/lang/sfeel"
)

var (
	ErrDecisionTableSerializationError   = "decision table serialization error"
	ErrDecisionTableUnserializationError = "decision table unserialization error"
)

type DecisionTableBuilder struct {
	data data.DecisionTable
}

func NewDecisionTableBuilder() DecisionTableBuilderInterface {
	return DecisionTableBuilder{}
}

func (d DecisionTableBuilder) SetID(id string) DecisionTableBuilderInterface {
	d.data.ID = id
	return d
}

func (d DecisionTableBuilder) SetName(name string) DecisionTableBuilderInterface {
	d.data.Name = name
	return d
}

func (d DecisionTableBuilder) SetHitPolicy(hitPolicy hitPolicy.HitPolicy) DecisionTableBuilderInterface {
	d.data.HitPolicy = hitPolicy
	return d
}

func (d DecisionTableBuilder) SetCollectOperator(collectOperator collectOperator.CollectOperator) DecisionTableBuilderInterface {
	d.data.CollectOperator = collectOperator
	return d
}

func (d DecisionTableBuilder) SetExpressionLanguage(expressionLanguage expressionLanguage.ExpressionLanguage) DecisionTableBuilderInterface {
	d.data.ExpressionLanguage = expressionLanguage
	return d
}

func (d DecisionTableBuilder) SetStandard(standard standard.Standard) DecisionTableBuilderInterface {
	d.data.Standard = standard
	return d
}

func (d DecisionTableBuilder) AddInputField(inputField field.Field) DecisionTableBuilderInterface {
	d.data.InputFields = append(d.data.InputFields, inputField)
	return d
}

func (d DecisionTableBuilder) AddOutputField(outputField field.Field) DecisionTableBuilderInterface {
	d.data.OutputFields = append(d.data.OutputFields, outputField)
	return d
}

func (d DecisionTableBuilder) AddRule(rule Rule) DecisionTableBuilderInterface {
	d.data.Rules = append(d.data.Rules, rule.data)
	return d
}

func (d DecisionTableBuilder) convertExpressionsIntoEntries(decisionTable data.DecisionTable, el expressionLanguage.ExpressionLanguage) data.DecisionTable {
	for i, rule := range decisionTable.Rules {
		for _, inputExpression := range rule.InputExpressions {
			switch el {
			case expressionLanguage.SFEEL:
				inputEntry := sfeel.CreateInputEntry(inputExpression)
				decisionTable.Rules[i].InputEntries = append(decisionTable.Rules[i].InputEntries, inputEntry)
			}
		}

		for _, outputExpression := range rule.OutputExpressions {
			switch el {
			case expressionLanguage.SFEEL:
				outputEntry := sfeel.CreateOutputEntry(outputExpression)
				decisionTable.Rules[i].OutputEntries = append(decisionTable.Rules[i].OutputEntries, outputEntry)
			}
		}
	}

	return decisionTable
}

func (d DecisionTableBuilder) Build() (DecisionTable, error) {
	d.data = d.convertExpressionsIntoEntries(d.data, d.data.ExpressionLanguage)
	decisionTable := DecisionTable{data: d.data}

	if err := decisionTable.Validate(decisionTable.Standard()); err != nil {
		return DecisionTable{}, err
	}

	return decisionTable, nil
}

func (d DecisionTableBuilder) BuildWithoutValidation() DecisionTable {
	d.data = d.convertExpressionsIntoEntries(d.data, d.data.ExpressionLanguage)
	return DecisionTable{data: d.data}
}

func (d DecisionTableBuilder) Serialize() (string, error) {
	s, err := json.MarshalIndent(d.data, "", "\t")
	if err != nil {
		return "{}", fmt.Errorf("%s: %w", ErrDecisionTableSerializationError, err)
	}

	return string(s), nil
}

func Unserialize(s string) (DecisionTableBuilderInterface, error) {
	var decisionTable data.DecisionTable
	if err := json.Unmarshal([]byte(s), &decisionTable); err != nil {
		return DecisionTableBuilder{}, fmt.Errorf("%s: %w", ErrDecisionTableUnserializationError, err)
	}

	return DecisionTableBuilder{data: decisionTable}, nil
}
