package main

import (
	"decisionTable/model"
	"decisionTable/validators"
	"decisionTable/validators/expressionlanguages"
)

type DecisionTableBuilder struct {
	dTableData model.TableData
}

func (d DecisionTableBuilder) SetDefinitionKey(key string) DecisionTableBuilderInterface {
	d.dTableData.Key = key
	return d
}

func (d DecisionTableBuilder) SetName(name string) DecisionTableBuilderInterface {
	d.dTableData.Name = name
	return d
}

func (d DecisionTableBuilder) SetHitPolicy(policy model.HitPolicy) DecisionTableBuilderInterface {
	d.dTableData.HitPolicy = policy
	return d
}

func (d DecisionTableBuilder) SetCollectOperator(collector model.CollectOperator) DecisionTableBuilderInterface {
	d.dTableData.CollectOperator = collector
	return d
}

func (d DecisionTableBuilder) SetNotationStandard(lang model.DTableStandard) DecisionTableBuilderInterface {
	d.dTableData.NotationStandard = lang
	return d
}

func (d DecisionTableBuilder) AddInputField(name string, label string, typ model.VariableTyp) DecisionTableBuilderInterface {
	field := model.Field{Name: name, Key: label, Typ: typ}
	d.dTableData.InputFields = append(d.dTableData.InputFields, field)
	return d
}

func (d DecisionTableBuilder) AddOutputField(name string, label string, typ model.VariableTyp) DecisionTableBuilderInterface {
	field := model.Field{Name: name, Key: label, Typ: typ}
	d.dTableData.OutputFields = append(d.dTableData.OutputFields, field)
	return d
}

func (d DecisionTableBuilder) AddRule(description string) DecisionTableRuleBuilderInterface {
	ruleBuilder := DTableRuleBuilder{
		input:       []model.Entry{},
		output:      []model.Entry{},
		description: description,
		builder:     d,
	}

	return ruleBuilder
}

func (d DecisionTableBuilder) Build() (DecisionTable, []error) {
	parsers := expressionlanguages.CreateParserFactory()
	validtr := validators.CreateDecisionTableValidator(d.dTableData, parsers)

	valid, err := validtr.Validate()
	if valid != true {
		return DecisionTable{}, err
	}

	table := d.createDecisionTable()
	table.interferences = validtr.ValidateContainsInterferences()
	return table, nil
}

func (d DecisionTableBuilder) createDecisionTable() DecisionTable {
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
	builder     DecisionTableBuilder
}

func (r DTableRuleBuilder) AddInputEntry(expr string, exprLang model.ExpressionLanguage) DecisionTableRuleBuilderInterface {
	entry := model.CreateEntry(expr, exprLang)
	r.input = append(r.input, entry)
	return r
}

func (r DTableRuleBuilder) AddOutputEntry(expr string, exprLang model.ExpressionLanguage) DecisionTableRuleBuilderInterface {
	entry := model.CreateEntry(expr, exprLang)
	r.output = append(r.output, entry)
	return r
}

func (r DTableRuleBuilder) BuildRule() DecisionTableBuilderInterface {
	rule := model.Rule{Description: r.description, InputEntries: r.input, OutputEntries: r.output}
	r.builder.dTableData.Rules = append(r.builder.dTableData.Rules, rule)
	return r.builder
}
