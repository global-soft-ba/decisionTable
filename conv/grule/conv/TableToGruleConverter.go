package conv

import (
	"errors"
	grule "github.com/global-soft-ba/decisionTable/conv/grule/data"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl/conv"
	dtable "github.com/global-soft-ba/decisionTable/data"
	"strconv"
)

var (
	ErrMapperIndexOutOfBound = errors.New("index of input or output fields is out of bound ")
)

func CreateTableToGruleConverter() TableToGruleConverter {
	return TableToGruleConverter{}
}

type TableToGruleConverter struct{}

func (c TableToGruleConverter) Convert(table dtable.Table) (grule.RuleSet, error) {
	return c.convertTableToRuleSet(table)
}

func (c TableToGruleConverter) convertTableToRuleSet(table dtable.Table) (grule.RuleSet, error) {
	result := grule.RuleSet{
		Key:             table.Key,
		Name:            table.Name,
		HitPolicy:       table.HitPolicy,
		CollectOperator: table.CollectOperator,
		Interference:    table.Interferences,
		Rules:           []grule.Rule{},
	}

	res, err := c.convertIntoGruleRules(table)
	if err != nil {
		return grule.RuleSet{}, err
	}

	result.Rules = res
	return result, nil
}

func (c TableToGruleConverter) convertIntoGruleRules(table dtable.Table) ([]grule.Rule, error) {
	var res []grule.Rule

	maxRules := len(table.Rules)
	for i, val := range table.Rules {
		rule, err := c.convertEntriesToRule(i, maxRules, val, table.InputFields, table.OutputFields)
		if err != nil {
			return []grule.Rule{}, err
		}
		res = append(res, rule)
	}

	return res, nil
}

func (c TableToGruleConverter) convertEntriesToRule(columId int, maxRules int, rule dtable.Rule, inputFields []dtable.FieldInterface, outputFields []dtable.FieldInterface) (grule.Rule, error) {
	r := grule.Rule{
		Name:        strconv.Itoa(columId),
		Description: rule.Description,
		Salience:    columId,
		InvSalience: maxRules - columId - 1, //Necessary for HitPolicies
		Expressions: nil,
		Assignments: nil,
	}

	expr, err := c.convertEntriesToTerms(rule.InputEntries, inputFields)
	if err != nil {
		return grule.Rule{}, err
	}

	asst, err := c.convertEntriesToTerms(rule.OutputEntries, outputFields)
	if err != nil {
		return grule.Rule{}, err
	}

	r.Expressions = expr
	r.Assignments = asst
	return r, nil
}

func (c TableToGruleConverter) convertEntriesToTerms(entries []dtable.EntryInterface, fields []dtable.FieldInterface) ([]grule.Term, error) {
	var result []grule.Term

	fieldCount := len(fields)
	for i, val := range entries {
		if i >= fieldCount {
			return []grule.Term{}, ErrMapperIndexOutOfBound
		}

		entry, err := c.convertEntryToExpression(val, fields[i])

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

func (c TableToGruleConverter) convertEntryToExpression(entry dtable.EntryInterface, field dtable.FieldInterface) (grule.Term, error) {
	term := grule.Term{
		Field:              field,
		Expression:         nil,
		ExpressionLanguage: entry.ExpressionLanguage(),
	}

	res, err := grl.CreateExpression(field, entry)
	if err != nil {
		return grule.Term{}, err
	}
	term.Expression = res

	return term, nil
}
