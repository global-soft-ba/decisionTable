package conv

import (
	grlmodel2 "decisionTable/conv/grule/data"
	"decisionTable/data"
	"errors"
	"strconv"
)

var (
	ErrMapperIndexOutOfBound = errors.New("index of input or output fields is out of bound ")
	ErrEmptyExpression       = errors.New("empty expression detected and will be skiped")
)

func CreateDTableMapper() DTableMapper {
	return DTableMapper{}
}

type DTableMapper struct{}

func (c DTableMapper) MapDTableToRuleSet(data data.Table) (grlmodel2.RuleSet, error) {
	result := grlmodel2.RuleSet{
		Key:             data.Key,
		Name:            data.Name,
		HitPolicy:       data.HitPolicy,
		CollectOperator: data.CollectOperator,
		Interference:    data.Interferences,
		Rules:           []grlmodel2.Rule{},
	}

	maxRules := len(data.Rules)
	for i, val := range data.Rules {
		rule, err := c.mapEntriesAndFieldsToRule(i, maxRules, val, data.InputFields, data.OutputFields)
		if err != nil {
			return grlmodel2.RuleSet{}, err
		}
		result.Rules = append(result.Rules, rule)
	}
	return result, nil
}

func (c DTableMapper) mapEntriesAndFieldsToRule(columId int, maxRules int, rule data.Rule, inputFields []data.Field, outputFields []data.Field) (grlmodel2.Rule, error) {
	r := grlmodel2.Rule{
		Name:        strconv.Itoa(columId),
		Description: rule.Description,
		Salience:    columId,
		InvSalience: maxRules - columId - 1, //Necessary for HitPolicies
		Expressions: nil,
		Assignments: nil,
	}

	expr, err := c.mapEntryToExpressions(rule.InputEntries, inputFields)
	if err != nil {
		return grlmodel2.Rule{}, err
	}

	asst, err := c.mapEntryToExpressions(rule.OutputEntries, outputFields)
	if err != nil {
		return grlmodel2.Rule{}, err
	}

	r.Expressions = expr
	r.Assignments = asst
	return r, nil
}

func (c DTableMapper) mapEntryToExpressions(entries []data.Entry, fields []data.Field) ([]grlmodel2.Term, error) {
	var result []grlmodel2.Term

	fieldCount := len(fields)
	for i, val := range entries {
		if i >= fieldCount {
			return []grlmodel2.Term{}, ErrMapperIndexOutOfBound
		}
		entry, err := c.mapToExpression(val, fields[i])

		if err != nil && err != ErrEmptyExpression {
			return []grlmodel2.Term{}, err
		}

		//EmptyEntries will be droped
		if err != ErrEmptyExpression {
			result = append(result, entry)
		}
	}
	return result, nil
}

func (c DTableMapper) mapToExpression(entry data.Entry, field data.Field) (grlmodel2.Term, error) {
	expr := grlmodel2.Term{
		Name:               field.Name,
		Key:                field.Key,
		Typ:                field.Typ,
		Expression:         entry.Expression(),
		ExpressionLanguage: entry.ExpressionLanguage(),
	}

	return expr, nil
}
