package grule

import (
	"decisionTable/conv/grule"
	"decisionTable/convert/grule/tableconverter/json"
	intf "decisionTable/convert/interfaces"
	"decisionTable/data"
	err "errors"
)

var (
	ErrDTableNoConverterFoundForOutputFormat = err.New("no converter found for decision table output formats")
)

func CreateDTableToGruleConverterFactory() DTableToGruleConverterFactory {
	return DTableToGruleConverterFactory{}
}

type DTableToGruleConverterFactory struct{}

func (g DTableToGruleConverterFactory) GetFormatConverter(format data.OutputFormat) (intf.ConverterInterface, error) {
	switch format {
	case data.GRL:
		return grule.CreateDTableToGrlConverter(), nil
	case data.JSON:
		return json.CreateDTableToJsonConverter(), nil
	default:
		return nil, ErrDTableNoConverterFoundForOutputFormat
	}
}
