package grule

import (
	err "errors"
	"github.com/global-soft-ba/decisionTable/converters/grule/tableconverter/grl"
	"github.com/global-soft-ba/decisionTable/converters/grule/tableconverter/json"
	intf "github.com/global-soft-ba/decisionTable/converters/interfaces"
	"github.com/global-soft-ba/decisionTable/model"
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
