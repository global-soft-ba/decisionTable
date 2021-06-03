package json

import (
	"decisionTable/convert/grule/termconverter"
	"decisionTable/data"
)

func CreateDTableToJsonConverter() DTableToJsonConverter {
	conv := termconverter.CreateTermConverterFactory()
	output := data.JSON
	return DTableToJsonConverter{conv, output}
}

type DTableToJsonConverter struct {
	expConvFac termconverter.TermConverterFactory
	format     data.OutputFormat
}

func (c DTableToJsonConverter) Convert(data data.Table) (interface{}, error) {
	panic("not implemented")
}
