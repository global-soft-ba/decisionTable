package mapper

import (
	converterModel "decisionTable/converter/grule/model"
	"decisionTable/model"
	"errors"
	"strconv"
)

var (
	ErrMapperIndexOutOfBound = errors.New("index of input or output fields is out of bound ")
)

type DTableToGrlMapper struct{}

func (c DTableToGrlMapper) MapToRuleSet(data model.DTableData) (converterModel.RuleSet, error) {
	result := converterModel.RuleSet{
		Key:             data.Key,
		Name:            data.Name,
		HitPolicy:       data.HitPolicy,
		CollectOperator: data.CollectOperator,
		Rules:           []converterModel.Rule{},
	}

	for i, val := range data.Rules {
		rule, err := c.mapRule(i, val, data.InputFields, data.OutputFields)
		if err != nil {
			return converterModel.RuleSet{}, err
		}
		result.Rules = append(result.Rules, rule)
	}

	return result, nil
}

func (c DTableToGrlMapper) mapRule(id int, rule model.Rule, inputFields []model.Field, outputFields []model.Field) (converterModel.Rule, error) {
	r := converterModel.Rule{
		Name:        strconv.Itoa(id),
		Description: rule.Description,
		Salience:    id,
		Expressions: nil,
		Assignments: nil,
	}

	fieldInputCount := len(inputFields)
	for i, val := range rule.InputEntries {
		if i > fieldInputCount {
			return converterModel.Rule{}, ErrMapperIndexOutOfBound
		}

		expr := converterModel.Expression{
			Name:       inputFields[i].Name,
			Identifier: inputFields[i].Label,
			Expression: val.Expression,
		}
		r.Expressions = append(r.Expressions, expr)
	}

	fieldOutputCount := len(outputFields)
	for i, val := range rule.OutputEntries {
		if i > fieldOutputCount {
			return converterModel.Rule{}, ErrMapperIndexOutOfBound
		}
		asst := converterModel.Assignment{
			Name:       outputFields[i].Name,
			Identifier: outputFields[i].Label,
			Expression: val.Expression,
		}
		r.Assignments = append(r.Assignments, asst)
	}

	return r, nil

}
