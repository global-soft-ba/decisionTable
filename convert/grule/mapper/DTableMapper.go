package mapper

import (
	"decisionTable/convert/grule/grlmodel"
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

func (c DTableMapper) MapDTableToRuleSet(data data.Table) (grlmodel.RuleSet, error) {
	result := grlmodel.RuleSet{
		Key:             data.Key,
		Name:            data.Name,
		HitPolicy:       data.HitPolicy,
		CollectOperator: data.CollectOperator,
		Interference:    data.Interferences,
		Rules:           []grlmodel.Rule{},
	}

	maxRules := len(data.Rules)
	for i, val := range data.Rules {
		rule, err := c.mapEntriesAndFieldsToRule(i, maxRules, val, data.InputFields, data.OutputFields)
		if err != nil {
			return grlmodel.RuleSet{}, err
		}
		result.Rules = append(result.Rules, rule)
	}
	return result, nil
}

func (c DTableMapper) mapEntriesAndFieldsToRule(columId int, maxRules int, rule data.Rule, inputFields []data.Field, outputFields []data.Field) (grlmodel.Rule, error) {
	r := grlmodel.Rule{
		Name:        strconv.Itoa(columId),
		Description: rule.Description,
		Salience:    columId,
		InvSalience: maxRules - columId - 1, //Necessary for HitPolicies
		Expressions: nil,
		Assignments: nil,
	}

	expr, err := c.mapEntryToExpressions(rule.InputEntries, inputFields)
	if err != nil {
		return grlmodel.Rule{}, err
	}

	asst, err := c.mapEntryToExpressions(rule.OutputEntries, outputFields)
	if err != nil {
		return grlmodel.Rule{}, err
	}

	r.Expressions = expr
	r.Assignments = asst
	return r, nil
}

func (c DTableMapper) mapEntryToExpressions(entries []data.Entry, fields []data.Field) ([]grlmodel.Term, error) {
	var result []grlmodel.Term

	fieldCount := len(fields)
	for i, val := range entries {
		if i >= fieldCount {
			return []grlmodel.Term{}, ErrMapperIndexOutOfBound
		}
		entry, err := c.mapToExpression(val, fields[i])

		if err != nil && err != ErrEmptyExpression {
			return []grlmodel.Term{}, err
		}

		//EmptyEntries will be droped
		if err != ErrEmptyExpression {
			result = append(result, entry)
		}
	}
	return result, nil
}

func (c DTableMapper) mapToExpression(entry data.Entry, field data.Field) (grlmodel.Term, error) {
	expr := grlmodel.Term{
		Name:               field.Name,
		Key:                field.Key,
		Typ:                field.Typ,
		Expression:         entry.Expression(),
		ExpressionLanguage: entry.ExpressionLanguage(),
	}

	return expr, nil
}
