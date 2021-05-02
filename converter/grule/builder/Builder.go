package builder

import (
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/model"
	"errors"
	"strconv"
)

var (
	ErrMapperIndexOutOfBound = errors.New("index of input or output fields is out of bound ")
	ErrEmptyExpression       = errors.New("empty expression detected and will be skiped")
)

func CreateGruleBuilder() GruleBuilder {
	return GruleBuilder{}
}

type GruleBuilder struct{}

func (c GruleBuilder) MapDTableToRuleSet(data model.TableData) (grlmodel.RuleSet, error) {
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

func (c GruleBuilder) mapEntriesAndFieldsToRule(columId int, maxRules int, rule model.Rule, inputFields []model.Field, outputFields []model.Field) (grlmodel.Rule, error) {
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

func (c GruleBuilder) mapEntryToExpressions(entries []model.Entry, fields []model.Field) ([]grlmodel.Expression, error) {
	var result []grlmodel.Expression

	fieldCount := len(fields)
	for i, val := range entries {
		if i >= fieldCount {
			return []grlmodel.Expression{}, ErrMapperIndexOutOfBound
		}
		entry, err := c.mapToExpression(val, fields[i])

		if err != nil && err != ErrEmptyExpression {
			return []grlmodel.Expression{}, err
		}

		//EmptyEntries will be droped
		if err != ErrEmptyExpression {
			result = append(result, entry)
		}
	}
	return result, nil
}

func (c GruleBuilder) mapToExpression(entry model.Entry, field model.Field) (grlmodel.Expression, error) {
	expr := grlmodel.Expression{
		Name:       field.Name,
		Identifier: field.Label,
		Typ:        field.Typ,
		Expression: entry.Expression(),
	}

	return expr, nil
}
