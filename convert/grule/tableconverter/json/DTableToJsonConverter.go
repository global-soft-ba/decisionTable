package json

import (
	"decisionTable/convert/grule/termconverter"
	"decisionTable/model"
)

func CreateDTableToJsonConverter() DTableToJsonConverter {
	conv := termconverter.CreateTermConverterFactory()
	output := model.JSON
	return DTableToJsonConverter{conv, output}
}

type DTableToJsonConverter struct {
	expConvFac termconverter.TermConverterFactory
	format     model.OutputFormat
}

func (c DTableToJsonConverter) Convert(data model.TableData) (interface{}, error) {
	panic("not implemented")
}
