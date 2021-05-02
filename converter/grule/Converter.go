package grule

import (
	"bytes"
	"decisionTable/converter/grule/builder"
	"decisionTable/converter/grule/expressionlanguages/sfeel"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/converter/grule/templates/grl"
	"decisionTable/model"
	"text/template"
)

const (
	Rule         = "Rule"
	RuleName     = "RuleName"
	Salience     = "Salience"
	Expressions  = "Expressions"
	Assignments  = "Assignments"
	Entries      = "Entries"
	Interference = "Interference"
)

func CreateConverter() Converter {
	conv := sfeel.CreateGrlExpressionConverter()
	return Converter{conv}
}

type Converter struct {
	expConv sfeel.ExpressionConverter
}

func (c Converter) Convert(data model.TableData) (interface{}, error) {

	grlModel, err := builder.CreateGruleBuilder().MapDTableToRuleSet(data)
	if err != nil {
		return []string{}, err
	}

	result, err := c.convertRuleSetIntoGRL(grlModel)
	if err != nil {
		return []string{}, err
	}

	return result, nil
}

func (c Converter) buildTemplate(hitPolicy model.HitPolicy, interference bool) (*template.Template, error) {

	var t *template.Template
	// TODo build only GRL Tempalte / What with Json?
	t, err := template.New(Rule).Parse(grl.RULE)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(RuleName).Parse(grl.RULENAME)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Expressions).Parse(grl.WHEN)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Assignments).Parse(grl.THEN)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Entries).Parse(grl.ENTRY)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Salience).Parse(c.buildPolicyTemplate(hitPolicy))
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Interference).Parse(c.buildInterferenceTemplate(interference))
	if err != nil {
		return &template.Template{}, err
	}

	return t, err
}

func (c Converter) buildPolicyTemplate(hitPolicy model.HitPolicy) string {
	switch hitPolicy {
	case model.Unique:
		return grl.UNIQUE
	case model.First:
		return grl.FIRST
	case model.Priority:
		return grl.PRIORITY
	default:
		return grl.DEFAULT
	}
}

func (c Converter) buildInterferenceTemplate(interference bool) string {
	switch interference {
	case true:
		return grl.INTERFERENCE
	default:
		return grl.NONINTERFERENCE
	}
}

func (c Converter) convertRuleSetIntoGRL(ruleSet grlmodel.RuleSet) ([]string, error) {
	var result []string

	tmpl, err := c.buildTemplate(ruleSet.HitPolicy, ruleSet.Interference)
	if err != nil {
		return []string{}, err
	}
	for _, v := range ruleSet.Rules {
		convRule := c.convertRuleExpressionsIntoGRL(v)

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, convRule) // Converts along the templates

		if err != nil {
			return []string{}, err
		}

		result = append(result, tpl.String())
	}

	return result, nil
}

func (c Converter) convertRuleExpressionsIntoGRL(rule grlmodel.Rule) grlmodel.Rule {
	result := rule
	for i, v := range rule.Expressions {
		//convExpression := c.expConv.Convert(v)
		convExpression := v
		result.Expressions[i] = convExpression

	}

	return result
}
