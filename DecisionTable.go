package decisionTable

import (
	"encoding/json"
	"fmt"
	"github.com/global-soft-ba/decisionTable/conv"
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/rule"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"github.com/global-soft-ba/decisionTable/validator"
)

var (
	ErrDecisionTableNotValid = "decision table is not valid"
)

type DecisionTable struct {
	data decisionTable.DecisionTable
}

func (d DecisionTable) ID() string {
	return d.data.ID
}

func (d DecisionTable) Name() string {
	return d.data.Name
}

func (d DecisionTable) HitPolicy() hitPolicy.HitPolicy {
	return d.data.HitPolicy
}

func (d DecisionTable) CollectOperator() collectOperator.CollectOperator {
	return d.data.CollectOperator
}

func (d DecisionTable) ExpressionLanguage() expressionLanguage.ExpressionLanguage {
	return d.data.ExpressionLanguage
}

func (d DecisionTable) Standard() standard.Standard {
	return d.data.Standard
}

func (d DecisionTable) InputFields() []field.Field {
	return d.data.InputFields
}

func (d DecisionTable) OutputFields() []field.Field {
	return d.data.OutputFields
}

func (d DecisionTable) Rules() []rule.Rule {
	return d.data.Rules
}

func (d DecisionTable) Validate(standard standard.Standard) error {
	v := validator.NewDecisionTableValidator()

	if err := v.Validate(d.data, standard); err != nil {
		return fmt.Errorf("%s: %w", ErrDecisionTableNotValid, err)
	}

	return nil
}

func (d DecisionTable) Convert(standard standard.Standard) (interface{}, error) {
	if err := d.Validate(standard); err != nil {
		return nil, err
	}

	converter := conv.CreateConverter()

	res, err := converter.Convert(d.data, standard)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d DecisionTable) Serialize() (string, error) {
	s, err := json.MarshalIndent(d.data, "", "\t")
	if err != nil {
		return "{}", fmt.Errorf("%s: %w", ErrDecisionTableSerializationError, err)
	}

	return string(s), nil
}

func Unserialize(s string) (DecisionTable, error) {
	var dt decisionTable.DecisionTable
	if err := json.Unmarshal([]byte(s), &dt); err != nil {
		return DecisionTable{}, fmt.Errorf("%s: %w", ErrDecisionTableUnserializationError, err)
	}

	return DecisionTable{data: dt}, nil
}
