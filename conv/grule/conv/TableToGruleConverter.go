package conv

import (
	grule "decisionTable/conv/grule/data"
	"decisionTable/data"
)

func CreateTableToGruleConverter() TableToGruleConverter {
	return TableToGruleConverter{}
}

type TableToGruleConverter struct{}

func (c TableToGruleConverter) Convert(table data.Table) (grule.RuleSet, error) {
	return grule.RuleSet{}, nil
}

/*

func (c TableToGruleConverter) convertRuleSetIntoGRL(ruleSet grule.RuleSet) ([]string, error) {
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
		err = tmpl.Execute(&tpl, convRule) // Converts along the grl

		if err != nil {
			return []string{}, err
		}

		result = append(result, tpl.String())
	}

	return result, nil
}

func (c TableToGruleConverter) convertRuleExpressionsIntoGRL(rule grule.Rule) (grule.Rule, error) {
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

func (c TableToGruleConverter) convertInputExpressionsIntoGRL(rule grule.Rule) (grule.Rule, error) {
	result := rule
	var expr []grule.Term

	for _, v := range rule.Expressions {
		expConv, err := c.expConvFac.GetExpressionConverter(v.ExpressionLanguage, c.format)
		if err != nil {
			return grule.Rule{}, err
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

func (c TableToGruleConverter) convertOutputExpressionsIntoGRL(rule grule.Rule) (grule.Rule, error) {
	result := rule
	var assgn []grule.Term

	for _, v := range rule.Assignments {
		expConv, err := c.expConvFac.GetExpressionConverter(v.ExpressionLanguage, c.format)
		if err != nil {
			return grule.Rule{}, err
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
*/
