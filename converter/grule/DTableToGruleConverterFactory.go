package grule

import (
	intf "decisionTable/converter/interfaces"
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
		return CreateDTableToGrlConverter(), nil
	case model.JSON:
		return CreateDTableToJsonConverter(), nil
	default:
		return nil, ErrDTableNoConverterFoundForOutputFormat
	}
}
