package conv

import (
	"errors"
	grule "github.com/global-soft-ba/decisionTable/conv/grule/data"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl/conv"
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/rule"
	"strconv"
)

var (
	ErrMapperIndexOutOfBound = errors.New("index of input or output fields is out of bound ")
)

func CreateTableToGruleConverter() TableToGruleConverter {
	return TableToGruleConverter{}
}

type TableToGruleConverter struct{}

func (c TableToGruleConverter) Convert(table decisionTable.DecisionTable) (grule.RuleSet, error) {
	return c.convertTableToRuleSet(table)
}

func (c TableToGruleConverter) convertTableToRuleSet(table decisionTable.DecisionTable) (grule.RuleSet, error) {
	result := grule.RuleSet{
		Key:             table.ID,
		Name:            table.Name,
		HitPolicy:       table.HitPolicy,
		CollectOperator: table.CollectOperator,
		Rules:           []grule.Rule{},
	}

	result.Interference = c.checkIfContainsInterferences(table)

	res, err := c.convertIntoGruleRules(table)
	if err != nil {
		return grule.RuleSet{}, err
	}

	result.Rules = res
	return result, nil
}

func (c TableToGruleConverter) convertIntoGruleRules(table decisionTable.DecisionTable) ([]grule.Rule, error) {
	var res []grule.Rule

	el := table.ExpressionLanguage

	maxRules := len(table.Rules)
	for i, val := range table.Rules {
		r, err := c.convertEntriesToRule(el, i, maxRules, val, table.InputFields, table.OutputFields)
		if err != nil {
			return []grule.Rule{}, err
		}
		res = append(res, r)
	}

	return res, nil
}

func (c TableToGruleConverter) convertEntriesToRule(expressionLanguage expressionLanguage.ExpressionLanguage, columId int, maxRules int, rule rule.Rule, inputFields []field.Field, outputFields []field.Field) (grule.Rule, error) {
	r := grule.Rule{
		Name:        strconv.Itoa(columId),
		Annotation:  rule.Annotation,
		Salience:    columId,
		InvSalience: maxRules - columId - 1, //Necessary for HitPolicies
		Expressions: nil,
		Assignments: nil,
	}

	expr, err := c.convertEntriesToTerms(expressionLanguage, entryType.Input, rule.InputEntries, inputFields)
	if err != nil {
		return grule.Rule{}, err
	}

	asst, err := c.convertEntriesToTerms(expressionLanguage, entryType.Output, rule.OutputEntries, outputFields)
	if err != nil {
		return grule.Rule{}, err
	}

	r.Expressions = expr
	r.Assignments = asst
	return r, nil
}

func (c TableToGruleConverter) convertEntriesToTerms(expressionLanguage expressionLanguage.ExpressionLanguage, entryType entryType.EntryType, entries []string, fields []field.Field) ([]grule.Term, error) {
	var result []grule.Term

	fieldCount := len(fields)
	for i, val := range entries {
		if i >= fieldCount {
			return []grule.Term{}, ErrMapperIndexOutOfBound
		}

		entry, err := c.convertEntryToExpression(expressionLanguage, entryType, val, fields[i])

		if err != nil {
			if errors.Is(err, conv.ErrEmptyStatement) {
				continue
			}
			return []grule.Term{}, err
		}
		result = append(result, entry)
	}

	return result, nil
}

func (c TableToGruleConverter) convertEntryToExpression(expressionLanguage expressionLanguage.ExpressionLanguage, entryType entryType.EntryType, entry string, field field.Field) (grule.Term, error) {
	term := grule.Term{
		Field:              field,
		Expression:         nil,
		ExpressionLanguage: expressionLanguage,
	}

	res, err := grl.CreateExpression(field, expressionLanguage, entryType, entry)
	if err != nil {
		return grule.Term{}, err
	}
	term.Expression = res

	return term, nil
}

func (c TableToGruleConverter) checkIfContainsInterferences(decisionTable decisionTable.DecisionTable) bool {
	for _, inputField := range decisionTable.InputFields {
		if c.checkIfContainsField(inputField, decisionTable.OutputFields) {
			return true
		}
	}

	return false
}

func (c TableToGruleConverter) checkIfContainsField(field field.Field, fieldSet []field.Field) bool {
	for _, fieldFromSet := range fieldSet {
		if field.Name == fieldFromSet.Name {
			return true
		}
	}

	return false
}
