package grl

import (
	"bytes"
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/convert/grule/mapper"
	"decisionTable/convert/grule/tableconverter/grl/templates"
	"decisionTable/convert/grule/termconverter"
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

func CreateDTableToGrlConverter() DTableToGrlConverter {
	conv := termconverter.CreateTermConverterFactory()
	output := model.GRL
	return DTableToGrlConverter{conv, output}
}

type DTableToGrlConverter struct {
	expConvFac termconverter.TermConverterFactory
	format     model.OutputFormat
}

func (c DTableToGrlConverter) Convert(data model.TableData) (interface{}, error) {

	grlModel, err := mapper.CreateDTableMapper().MapDTableToRuleSet(data)
	if err != nil {
		return []string{}, err
	}

	result, err := c.convertRuleSetIntoGRL(grlModel)
	if err != nil {
		return []string{}, err
	}

	return result, nil
}

func (c DTableToGrlConverter) buildTemplate(hitPolicy model.HitPolicy, interference bool) (*template.Template, error) {

	var t *template.Template
	// TODo build only GRL Template / What with Json?
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
	_, err = t.New(Interference).Parse(c.buildInterferenceTemplate(interference))
	if err != nil {
		return &template.Template{}, err
	}

	return t, err
}

func (c DTableToGrlConverter) buildPolicyTemplate(hitPolicy model.HitPolicy) string {
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

func (c DTableToGrlConverter) buildInterferenceTemplate(interference bool) string {
	switch interference {
	case true:
		return templates.INTERFERENCE
	default:
		return templates.NONINTERFERENCE
	}
}

func (c DTableToGrlConverter) convertRuleSetIntoGRL(ruleSet grlmodel.RuleSet) ([]string, error) {
	var result []string

	tmpl, err := c.buildTemplate(ruleSet.HitPolicy, ruleSet.Interference)
	if err != nil {
		return []string{}, err
	}
	for _, v := range ruleSet.Rules {
		convRule, errConv := c.convertRuleExpressionsIntoGRL(v)
		if errConv != nil {
			return []string{}, err
		}

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, convRule) // Converts along the templates

		if err != nil {
			return []string{}, err
		}

		result = append(result, tpl.String())
	}

	return result, nil
}

func (c DTableToGrlConverter) convertRuleExpressionsIntoGRL(rule grlmodel.Rule) (grlmodel.Rule, error) {
	result, err := c.convertInputExpressionsIntoGRL(rule)
	if err != nil {
		return result, err
	}

	result, err = c.convertOutputExpressionsIntoGRL(result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c DTableToGrlConverter) convertInputExpressionsIntoGRL(rule grlmodel.Rule) (grlmodel.Rule, error) {
	result := rule
	var expr []grlmodel.Term

	for _, v := range rule.Expressions {
		expConv, err := c.expConvFac.GetExpressionConverter(v.ExpressionLanguage, c.format)
		if err != nil {
			return grlmodel.Rule{}, err
		}

		term := expConv.ConvertExpression(v)
		//Remove empty expressions
		if term.Expression != "" {
			expr = append(expr, term)
		}
	}

	result.Expressions = expr
	return result, nil
}

func (c DTableToGrlConverter) convertOutputExpressionsIntoGRL(rule grlmodel.Rule) (grlmodel.Rule, error) {
	result := rule
	var assgn []grlmodel.Term

	for _, v := range rule.Assignments {
		expConv, err := c.expConvFac.GetExpressionConverter(v.ExpressionLanguage, c.format)
		if err != nil {
			return grlmodel.Rule{}, err
		}
		term := expConv.ConvertAssignments(v)
		//Remove empty expressions
		if term.Expression != "" {
			assgn = append(assgn, term)
		}
	}

	result.Assignments = assgn
	return result, nil
}
