package grule

import (
	"bytes"
	"decisionTable/converter"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/converter/grule/mapper"
	"decisionTable/converter/grule/templates"
	"decisionTable/model"
	"text/template"
)

const (
	Rule        = "Rule"
	RuleName    = "RuleName"
	Salience    = "Salience"
	Expressions = "Expressions"
	Assignments = "Assignments"
	Entries     = "Entries"
)

func CreateDTableToGrlConverter() Converter {
	return Converter{}
}

type Converter struct {
}

func (c Converter) Convert(data model.DTableData) ([]string, error) {
	if data.NotationStandard != model.GRULE {
		return []string{}, converter.ErrDTableNotationStandard
	}

	grlModel, err := mapper.CreateGruleMapper().MapToRuleSet(data)
	if err != nil {
		return []string{}, err
	}

	result, err := c.createRuleSet(grlModel)
	if err != nil {
		return []string{}, err
	}

	return result, nil
}

func (c Converter) createRuleSet(ruleSet grlmodel.RuleSet) ([]string, error) {
	var result []string

	tmpl, err := c.buildTemplate(ruleSet.HitPolicy)
	if err != nil {
		return []string{}, err
	}
	for _, v := range ruleSet.Rules {
		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, v)
		if err != nil {
			return []string{}, err
		}

		result = append(result, tpl.String())
	}

	return result, nil
}

func (c Converter) buildTemplate(hitPolicy model.HitPolicy) (*template.Template, error) {

	var t *template.Template

	t, err := template.New(Rule).Parse(templates.RULE)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(RuleName).Parse(templates.RULENAME)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Expressions).Parse(templates.WHEN)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Assignments).Parse(templates.THEN)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Entries).Parse(templates.ENTRY)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Salience).Parse(c.buildPolicyTemplate(hitPolicy))
	if err != nil {
		return &template.Template{}, err
	}

	return t, err
}

func (c Converter) buildPolicyTemplate(hitPolicy model.HitPolicy) string {
	switch hitPolicy {
	case model.Unique:
		return templates.UNIQUE
	case model.First:
		return templates.FIRST
	case model.Priority:
		return templates.PRIORITY
	default:
		return templates.DEFAULT
	}
}

func (c Converter) checkForInterference(table model.DTableData) bool {
	return false
}
