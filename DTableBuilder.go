package main

import (
	"decisionTable/model"
	"decisionTable/validator"
)

type DTableBuilder struct {
	dTableData model.DTableData
}

func (d DTableBuilder) SetDefinitionKey(key string) DTableBuilderInterface {
	d.dTableData.Key = key
	return d
}

func (d DTableBuilder) SetName(name string) DTableBuilderInterface {
	d.dTableData.Name = name
	return d
}

func (d DTableBuilder) SetHitPolicy(policy model.HitPolicy) DTableBuilderInterface {
	d.dTableData.HitPolicy = policy
	return d
}

func (d DTableBuilder) SetCollectOperator(collector model.CollectOperator) DTableBuilderInterface {
	d.dTableData.CollectOperator = collector
	return d
}

func (d DTableBuilder) SetNotationStandard(lang model.DTableStandard) DTableBuilderInterface {
	d.dTableData.NotationStandard = lang
	return d
}

func (d DTableBuilder) AddInputField(name string, label string, typ model.VariableTyp) DTableBuilderInterface {
	field := model.Field{Name: name, Label: label, Typ: typ}
	d.dTableData.InputFields = append(d.dTableData.InputFields, field)
	return d
}

func (d DTableBuilder) AddOutputField(name string, label string, typ model.VariableTyp) DTableBuilderInterface {
	field := model.Field{Name: name, Label: label, Typ: typ}
	d.dTableData.OutputFields = append(d.dTableData.OutputFields, field)
	return d
}

func (d DTableBuilder) AddRule(input []model.Entry, output []model.Entry, description string) DTableBuilderInterface {
	r := model.Rule{Description: description, InputEntries: input, OutputEntries: output}
	d.dTableData.Rules = append(d.dTableData.Rules, r)
	return d
}

func (d DTableBuilder) Build() (DecisionTable, []error) {
	table := d.createDecisionTable()
	if valid, errs := d.validate(); valid != true {
		return table, errs
	}

	table.valid = true
	return table, nil
}

func (d DTableBuilder) createDecisionTable() DecisionTable {
	dTable := DecisionTable{
		key:              d.dTableData.Key,
		name:             d.dTableData.Name,
		hitPolicy:        d.dTableData.HitPolicy,
		collectOperator:  d.dTableData.CollectOperator,
		notationStandard: d.dTableData.NotationStandard,
		valid:            false,
		inputFields:      d.dTableData.InputFields,
		outputFields:     d.dTableData.OutputFields,
		rules:            d.dTableData.Rules,
	}

	return dTable
}

func (d DTableBuilder) validate() (bool, []error) {
	//General Validation
	if valid, err := validator.CreateDTableValidator(d.dTableData).Validate(); valid != true {
		return valid, err
	}

	return true, []error{}
}
