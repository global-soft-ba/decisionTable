package main

import (
	"decisionTable/model"
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

func (d DTableBuilder) SetHitPolicy(policy string) DTableBuilderInterface {
	d.dTableData.HitPolicy = policy
	return d
}

func (d DTableBuilder) SetCollectOperator(collector string) DTableBuilderInterface {
	d.dTableData.CollectOperator = collector
	return d
}

func (d DTableBuilder) SetExpressionLanguage(lang string) DTableBuilderInterface {
	d.dTableData.ExpressionLanguage = lang
	return d
}

func (d DTableBuilder) AddInputField(name string, label string, typ string) DTableBuilderInterface {
	field := model.Field{Name: name, Label: label, Typ: typ}
	d.dTableData.InputFields = append(d.dTableData.InputFields, field)
	return d
}

func (d DTableBuilder) AddOutputField(name string, label string, typ string) DTableBuilderInterface {
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
	//Validate()
	//DecisionTable{d.dTableData.Key,...}
	table := d.createDecisionTable()
	return table, nil
}

func (d DTableBuilder) createDecisionTable() DecisionTable {
	dTable := DecisionTable{
		key:                d.dTableData.Key,
		name:               d.dTableData.Name,
		hitPolicy:          d.dTableData.HitPolicy,
		collectOperator:    d.dTableData.CollectOperator,
		expressionLanguage: d.dTableData.ExpressionLanguage,
		inputFields:        d.dTableData.InputFields,
		outputFields:       d.dTableData.OutputFields,
		rules:              d.dTableData.Rules,
	}

	return dTable
}
