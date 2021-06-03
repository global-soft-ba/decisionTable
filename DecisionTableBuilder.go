package main

import (
	"decisionTable/data"
	sfeel2 "decisionTable/lang/sfeel"
	"decisionTable/valid"
)

type DecisionTableBuilder struct {
	data data.Table
}

func (d DecisionTableBuilder) SetDefinitionKey(key string) DecisionTableBuilderInterface {
	d.data.Key = key
	return d
}

func (d DecisionTableBuilder) SetName(name string) DecisionTableBuilderInterface {
	d.data.Name = name
	return d
}

func (d DecisionTableBuilder) SetHitPolicy(policy data.HitPolicy) DecisionTableBuilderInterface {
	d.data.HitPolicy = policy
	return d
}

func (d DecisionTableBuilder) SetCollectOperator(collector data.CollectOperator) DecisionTableBuilderInterface {
	d.data.CollectOperator = collector
	return d
}

func (d DecisionTableBuilder) SetNotationStandard(lang data.DTableStandard) DecisionTableBuilderInterface {
	d.data.NotationStandard = lang
	return d
}

func (d DecisionTableBuilder) AddInputField(name string, label string, typ data.DataTyp) DecisionTableBuilderInterface {
	field := data.Field{Name: name, Key: label, Typ: typ}
	d.data.InputFields = append(d.data.InputFields, field)
	return d
}

func (d DecisionTableBuilder) AddOutputField(name string, label string, typ data.DataTyp) DecisionTableBuilderInterface {
	field := data.Field{Name: name, Key: label, Typ: typ}
	d.data.OutputFields = append(d.data.OutputFields, field)
	return d
}

func (d DecisionTableBuilder) AddRule(description string) DecisionTableRuleBuilderInterface {
	ruleBuilder := DTableRuleBuilder{
		input:       []data.EntryInterface{},
		output:      []data.EntryInterface{},
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
	input  []data.EntryInterface
	output []data.EntryInterface

	description string
	builder     DecisionTableBuilder
}

func (r DTableRuleBuilder) AddInputEntry(expr string, exprLang data.ExpressionLanguage) DecisionTableRuleBuilderInterface {
	switch exprLang {
	case data.SFEEL:
		entry := sfeel2.CreateInputEntry(expr)
		r.input = append(r.input, entry)
		return r
	}

	return nil
}

func (r DTableRuleBuilder) AddOutputEntry(expr string, exprLang data.ExpressionLanguage) DecisionTableRuleBuilderInterface {
	switch exprLang {
	case data.SFEEL:
		entry := sfeel2.CreateOutputEntry(expr)
		r.input = append(r.input, entry)
		return r
	}

	return nil
}

func (r DTableRuleBuilder) BuildRule() DecisionTableBuilderInterface {
	rule := data.Rule{Description: r.description, InputEntries: r.input, OutputEntries: r.output}
	r.builder.data.Rules = append(r.builder.data.Rules, rule)
	return r.builder
}
