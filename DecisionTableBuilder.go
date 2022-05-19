package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

var (
	ErrDecisionTableSerializationError   = "decision table serialization error"
	ErrDecisionTableUnserializationError = "decision table unserialization error"
)

type DecisionTableBuilder struct {
	data decisionTable.DecisionTable
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

func (d DecisionTableBuilder) Build() (DecisionTable, error) {
	dt := DecisionTable{data: d.data}

	if err := dt.Validate(dt.Standard()); err != nil {
		return DecisionTable{}, err
	}

	return dt, nil
}

func (d DecisionTableBuilder) BuildWithoutValidation() DecisionTable {
	return DecisionTable{data: d.data}
}
