package grule

import (
	"bytes"
	grlmodel2 "decisionTable/conv/grule/data"
	mapper2 "decisionTable/conv/grule/mapper"
	templates2 "decisionTable/conv/grule/templates"
	"decisionTable/convert/grule/termconverter"
	"decisionTable/data"
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
	output := data.GRL
	return DTableToGrlConverter{conv, output}
}

type DTableToGrlConverter struct {
	expConvFac termconverter.TermConverterFactory
	format     data.OutputFormat
}

func (c DTableToGrlConverter) Convert(data data.Table) (interface{}, error) {

	grlModel, err := mapper2.CreateDTableMapper().MapDTableToRuleSet(data)
	if err != nil {
		return []string{}, err
	}

	result, err := c.convertRuleSetIntoGRL(grlModel)
	if err != nil {
		return []string{}, err
	}

	return result, nil
}

func (c DTableToGrlConverter) buildTemplate(hitPolicy data.HitPolicy, interference bool) (*template.Template, error) {

	var t *template.Template
	// TODo build only GRL Template / What with Json?
	t, err := template.New(Rule).Parse(templates2.RULE)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(RuleName).Parse(templates2.RULENAME)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Expressions).Parse(templates2.WHEN)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Assignments).Parse(templates2.THEN)
	if err != nil {
		return &template.Template{}, err
	}
	_, err = t.New(Entries).Parse(templates2.ENTRY)
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

func (c DTableToGrlConverter) buildPolicyTemplate(hitPolicy data.HitPolicy) string {
	switch hitPolicy {
	case data.Unique:
		return templates2.UNIQUE
	case data.First:
		return templates2.FIRST
	case data.Priority:
		return templates2.PRIORITY
	default:
		return templates2.DEFAULT
	}
}

func (c DTableToGrlConverter) buildInterferenceTemplate(interference bool) string {
	switch interference {
	case true:
		return templates2.INTERFERENCE
	default:
		return templates2.NONINTERFERENCE
	}
}

func (c DTableToGrlConverter) convertRuleSetIntoGRL(ruleSet grlmodel2.RuleSet) ([]string, error) {
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

func (c DTableToGrlConverter) convertRuleExpressionsIntoGRL(rule grlmodel2.Rule) (grlmodel2.Rule, error) {
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

func (c DTableToGrlConverter) convertInputExpressionsIntoGRL(rule grlmodel2.Rule) (grlmodel2.Rule, error) {
	result := rule
	var expr []grlmodel2.Term

	for _, v := range rule.Expressions {
		expConv, err := c.expConvFac.GetExpressionConverter(v.ExpressionLanguage, c.format)
		if err != nil {
			return grlmodel2.Rule{}, err
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

func (c DTableToGrlConverter) convertOutputExpressionsIntoGRL(rule grlmodel2.Rule) (grlmodel2.Rule, error) {
	result := rule
	var assgn []grlmodel2.Term

	for _, v := range rule.Assignments {
		expConv, err := c.expConvFac.GetExpressionConverter(v.ExpressionLanguage, c.format)
		if err != nil {
			return grlmodel2.Rule{}, err
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
