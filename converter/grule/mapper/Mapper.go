package mapper

import (
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/model"
	"errors"
	"strconv"
)

var (
	ErrMapperIndexOutOfBound = errors.New("index of input or output fields is out of bound ")
)

func CreateGruleMapper() Mapper {
	return Mapper{}
}

type Mapper struct{}

func (c Mapper) MapToRuleSet(data model.DTableData) (grlmodel.RuleSet, error) {
	result := grlmodel.RuleSet{
		Key:             data.Key,
		Name:            data.Name,
		HitPolicy:       data.HitPolicy,
		CollectOperator: data.CollectOperator,
		Rules:           []grlmodel.Rule{},
	}

	for i, val := range data.Rules {
		rule, err := c.mapToRule(i, val, data.InputFields, data.OutputFields)
		if err != nil {
			return grlmodel.RuleSet{}, err
		}
		result.Rules = append(result.Rules, rule)
	}

	return result, nil
}

func (c Mapper) mapToRule(id int, rule model.Rule, inputFields []model.Field, outputFields []model.Field) (grlmodel.Rule, error) {
	r := grlmodel.Rule{
		Name:        strconv.Itoa(id),
		Description: rule.Description,
		Salience:    id,
		Expressions: nil,
		Assignments: nil,
	}

	expr, err := c.mapToExpressions(rule.InputEntries, inputFields)
	if err != nil {
		return grlmodel.Rule{}, err
	}

	asst, err := c.mapToExpressions(rule.OutputEntries, outputFields)
	if err != nil {
		return grlmodel.Rule{}, err
	}

	r.Expressions = expr
	r.Assignments = asst
	return r, nil
}

func (c Mapper) mapToExpressions(entries []model.Entry, fields []model.Field) ([]grlmodel.Expression, error) {
	var result []grlmodel.Expression

	fieldCount := len(fields)
	for i, val := range entries {
		if i >= fieldCount {
			return []grlmodel.Expression{}, ErrMapperIndexOutOfBound
		}
		entry, err := c.mapToExpression(val, fields[i])
		if err != nil {
			return []grlmodel.Expression{}, err
		}
		result = append(result, entry)
	}
	return result, nil
}

func (c Mapper) mapToExpression(entry model.Entry, field model.Field) (grlmodel.Expression, error) {
	expr := grlmodel.Expression{
		Name:       field.Name,
		Identifier: field.Label,
		Expression: entry.Expression,
	}

	return expr, nil
}
