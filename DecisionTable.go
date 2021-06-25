package decisionTable

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/converters/interfaces"
	"github.com/global-soft-ba/decisionTable/model"
)

var (
	ErrDTableNotValid = errors.New("decision table must be valid before converting")
)

func CreateDecisionTable() DecisionTableBuilderInterface {
	d := DecisionTableBuilder{}
	return d
}

type DecisionTable struct {
	key              string
	name             string
	hitPolicy        model.HitPolicy
	collectOperator  model.CollectOperator
	notationStandard model.DTableStandard
	interferences    bool
	valid            bool

	inputFields  []model.Field
	outputFields []model.Field
	rules        []model.Rule
}

func (d DecisionTable) Key() string {
	return d.key
}

func (d DecisionTable) Name() string {
	return d.name
}

func (d DecisionTable) HitPolicy() model.HitPolicy {
	return d.hitPolicy
}

func (d DecisionTable) CollectOperator() model.CollectOperator {
	return d.collectOperator
}

func (d DecisionTable) NotationStandard() model.DTableStandard {
	return d.notationStandard
}

func (d DecisionTable) Valid() bool {
	return d.valid
}

func (d DecisionTable) InputFields() []model.Field {
	return d.inputFields
}

func (d DecisionTable) OutputFields() []model.Field {
	return d.outputFields
}

func (d DecisionTable) Rules() []model.Rule {
	return d.rules
}

func (d DecisionTable) Interferences() bool {
	return d.interferences
}

func (d DecisionTable) Convert(converter interfaces.ConverterInterface) (interface{}, error) {

	if !d.valid {
		return []string{}, ErrDTableNotValid
	}

	dTable := model.TableData{
		Key:              d.key,
		Name:             d.name,
		HitPolicy:        d.hitPolicy,
		CollectOperator:  d.collectOperator,
		NotationStandard: d.notationStandard,
		InputFields:      d.inputFields,
		OutputFields:     d.outputFields,
		Interferences:    d.interferences,
		Rules:            d.rules,
	}
	return converter.Convert(dTable)
}
