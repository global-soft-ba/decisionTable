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

func (d DTableBuilder) AddRule(description string) DTableRuleBuilderInterface {
	ruleBuilder := DTableRuleBuilder{
		input:       []model.Entry{},
		output:      []model.Entry{},
		description: description,
		builder:     d,
	}

	return ruleBuilder
}

func (d DTableBuilder) Build() (DecisionTable, []error) {
	validtr := validator.CreateDTableValidator(d.dTableData)

	valid, err := validtr.Validate()
	if valid != true {
		return DecisionTable{}, err
	}

	table := d.createDecisionTable()
	table.interferences = validtr.ValidateInterferences()
	return table, nil
}

func (d DTableBuilder) createDecisionTable() DecisionTable {
	dTable := DecisionTable{
		key:              d.dTableData.Key,
		name:             d.dTableData.Name,
		hitPolicy:        d.dTableData.HitPolicy,
		collectOperator:  d.dTableData.CollectOperator,
		notationStandard: d.dTableData.NotationStandard,
		valid:            true,
		interferences:    false,
		inputFields:      d.dTableData.InputFields,
		outputFields:     d.dTableData.OutputFields,
		rules:            d.dTableData.Rules,
	}

	return dTable
}

type DTableRuleBuilder struct {
	input  []model.Entry
	output []model.Entry

	description string
	builder     DTableBuilder
}

func (r DTableRuleBuilder) AddInputEntry(expr string, exprLang model.ExpressionLanguage) DTableRuleBuilderInterface {
	entry := model.CreateEntry(expr, exprLang)
	r.input = append(r.input, entry)
	return r
}

func (r DTableRuleBuilder) AddOutputEntry(expr string, exprLang model.ExpressionLanguage) DTableRuleBuilderInterface {
	entry := model.CreateEntry(expr, exprLang)
	r.output = append(r.output, entry)
	return r
}

func (r DTableRuleBuilder) BuildRule() DTableBuilderInterface {
	rule := model.Rule{Description: r.description, InputEntries: r.input, OutputEntries: r.output}
	r.builder.dTableData.Rules = append(r.builder.dTableData.Rules, rule)
	return r.builder
}
