package grule

import (
	"decisionTable/convert/grule/tableconverter/grl"
	"decisionTable/convert/grule/tableconverter/json"
	intf "decisionTable/convert/interfaces"
	"decisionTable/model"
	err "errors"
)

var (
	ErrDTableNoConverterFoundForOutputFormat = err.New("no converter found for decision table output formats")
)

func CreateDTableToGruleConverterFactory() DTableToGruleConverterFactory {
	return DTableToGruleConverterFactory{}
}

type DTableToGruleConverterFactory struct{}

func (g DTableToGruleConverterFactory) GetFormatConverter(format model.OutputFormat) (intf.ConverterInterface, error) {
	switch format {
	case model.GRL:
		return grl.CreateDTableToGrlConverter(), nil
	case model.JSON:
		return json.CreateDTableToJsonConverter(), nil
	default:
		return nil, ErrDTableNoConverterFoundForOutputFormat
	}
}
