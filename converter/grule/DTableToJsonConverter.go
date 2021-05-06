package grule

import (
	"decisionTable/converter/grule/expressionlanguages"
	"decisionTable/model"
)

func CreateDTableToJsonConverter() DTableToGrlConverter {
	conv := expressionlanguages.CreateExpresionConverterFactory()
	output := model.JSON
	return DTableToGrlConverter{conv, output}
}

type DTableToJsonConverter struct {
	expConvFac expressionlanguages.ExpressionConverterFactory
	format     model.OutputFormat
}

func (c DTableToJsonConverter) Convert(data model.TableData) (interface{}, error) {
	panic("not implemented")
}
