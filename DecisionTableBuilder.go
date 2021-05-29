package main

import (
	"decisionTable/model"
	"decisionTable/sfeel"
	"decisionTable/valid"
)

type DecisionTableBuilder struct {
	data model.TableData
}

func (d DecisionTableBuilder) SetDefinitionKey(key string) DecisionTableBuilderInterface {
	d.data.Key = key
	return d
}

func (d DecisionTableBuilder) SetName(name string) DecisionTableBuilderInterface {
	d.data.Name = name
	return d
}

func (d DecisionTableBuilder) SetHitPolicy(policy model.HitPolicy) DecisionTableBuilderInterface {
	d.data.HitPolicy = policy
	return d
}

func (d DecisionTableBuilder) SetCollectOperator(collector model.CollectOperator) DecisionTableBuilderInterface {
	d.data.CollectOperator = collector
	return d
}

func (d DecisionTableBuilder) SetNotationStandard(lang model.DTableStandard) DecisionTableBuilderInterface {
	d.data.NotationStandard = lang
	return d
}

func (d DecisionTableBuilder) AddInputField(name string, label string, typ model.DataTyp) DecisionTableBuilderInterface {
	field := model.Field{Name: name, Key: label, Typ: typ}
	d.data.InputFields = append(d.data.InputFields, field)
	return d
}

func (d DecisionTableBuilder) AddOutputField(name string, label string, typ model.DataTyp) DecisionTableBuilderInterface {
	field := model.Field{Name: name, Key: label, Typ: typ}
	d.data.OutputFields = append(d.data.OutputFields, field)
	return d
}

func (d DecisionTableBuilder) AddRule(description string) DecisionTableRuleBuilderInterface {
	ruleBuilder := DTableRuleBuilder{
		input:       []model.EntryInterface{},
		output:      []model.EntryInterface{},
		description: description,
		builder:     d,
	}

	return ruleBuilder
}

func (d DecisionTableBuilder) Build() (DecisionTable, []error) {

	validtr := valid.CreateDecisionTableValidator(d.data)

	val, err := validtr.Validate()
	if val != true {
		return DecisionTable{}, err
	}

	table := DecisionTable{data: d.data, valid: val}
	table.data.Interferences = validtr.ValidateContainsInterferences()
	return table, nil
}

type DTableRuleBuilder struct {
	input  []model.EntryInterface
	output []model.EntryInterface

	description string
	builder     DecisionTableBuilder
}

func (r DTableRuleBuilder) AddInputEntry(expr string, exprLang model.ExpressionLanguage) DecisionTableRuleBuilderInterface {
	switch exprLang {
	case model.SFEEL:
		entry := sfeel.CreateInputEntry(expr)
		r.input = append(r.input, entry)
		return r
	}

	return nil
}

func (r DTableRuleBuilder) AddOutputEntry(expr string, exprLang model.ExpressionLanguage) DecisionTableRuleBuilderInterface {
	switch exprLang {
	case model.SFEEL:
		entry := sfeel.CreateOutputEntry(expr)
		r.input = append(r.input, entry)
		return r
	}

	return nil
}

func (r DTableRuleBuilder) BuildRule() DecisionTableBuilderInterface {
	rule := model.Rule{Description: r.description, InputEntries: r.input, OutputEntries: r.output}
	r.builder.data.Rules = append(r.builder.data.Rules, rule)
	return r.builder
}
