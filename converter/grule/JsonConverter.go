package grule

import (
	"decisionTable/converter/grule/expressionlanguages/sfeel"
	"decisionTable/model"
)

func CreateJsonConverter() JsonConverter {
	conv := sfeel.CreateJsonExpressionConverter()
	return JsonConverter{conv}
}

type JsonConverter struct {
	expConv sfeel.ExpressionConverter
}

func (j JsonConverter) Convert(data model.TableData) (interface{}, error) {
	panic("not implemented")
}
