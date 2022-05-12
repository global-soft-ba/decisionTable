package decisionTable

import (
	"fmt"
	"github.com/global-soft-ba/decisionTable/conv"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"github.com/global-soft-ba/decisionTable/validator"
)

var (
	ErrDecisionTableNotValid = "decision table is not valid"
)

type DecisionTable struct {
	data data.DecisionTable
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

func (d DecisionTable) Rules() []data.Rule {
	return d.data.Rules
}

func (d DecisionTable) Interferences() bool {
	return d.data.Interferences
}

func (d DecisionTable) CheckIfContainsInterferences() bool {
	for _, inputField := range d.data.InputFields {
		if d.checkIfContainsField(inputField, d.data.OutputFields) {
			return true
		}
	}

	return false
}

func (d DecisionTable) checkIfContainsField(field field.Field, fieldSet []field.Field) bool {
	for _, fieldFromSet := range fieldSet {
		if field.Name == fieldFromSet.Name {
			return true
		}
	}

	return false
}

func (d *DecisionTable) Validate(standard standard.Standard) error {
	v := validator.NewDecisionTableValidator()

	if err := v.Validate(d.data, standard); err != nil {
		return fmt.Errorf("%s: %w", ErrDecisionTableNotValid, err)
	}

	d.data.Interferences = d.CheckIfContainsInterferences()

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
