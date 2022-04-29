package decisionTable

import (
	"fmt"
	"github.com/global-soft-ba/decisionTable/conv"
	"github.com/global-soft-ba/decisionTable/data"
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

func (d DecisionTable) HitPolicy() data.HitPolicy {
	return d.data.HitPolicy
}

func (d DecisionTable) CollectOperator() data.CollectOperator {
	return d.data.CollectOperator
}

func (d DecisionTable) ExpressionLanguage() data.ExpressionLanguage {
	return d.data.ExpressionLanguage
}

func (d DecisionTable) Standard() data.Standard {
	return d.data.Standard
}

func (d DecisionTable) InputFields() []data.FieldInterface {
	return d.data.InputFields
}

func (d DecisionTable) OutputFields() []data.FieldInterface {
	return d.data.OutputFields
}

func (d DecisionTable) Rules() []data.Rule {
	return d.data.Rules
}

func (d DecisionTable) Interferences() bool {
	return d.data.Interferences
}

func (d DecisionTable) CheckIfContainsInterferences() bool {
	outputFields := d.data.OutputFields

	for _, inputField := range d.data.InputFields {
		if d.checkIfContainsField(inputField, outputFields) {
			return true
		}
	}

	return false
}

func (d DecisionTable) checkIfContainsField(field data.FieldInterface, fieldSet []data.FieldInterface) bool {
	for _, fieldFromSet := range fieldSet {
		if field.ID() == fieldFromSet.ID() {
			return true
		}
	}

	return false
}

func (d *DecisionTable) Validate(standard data.Standard) error {
	v := validator.NewDecisionTableValidator()

	if err := v.Validate(d.data, standard); err != nil {
		return fmt.Errorf(ErrDecisionTableNotValid+": %w", err)
	}

	d.data.Interferences = d.CheckIfContainsInterferences()

	return nil
}

func (d DecisionTable) Convert(standard data.Standard) (interface{}, error) { // TODO: Change return type from interface{} to []string{}?
	if err := d.Validate(standard); err != nil {
		return []string{}, err
	}

	tableConverter := conv.CreateConverter()

	res, err := tableConverter.Convert(d.data, string(standard)) // TODO: Get rid of casting
	if err != nil {
		return nil, err
	}

	return res, nil
}
