package converter

import (
	"decisionTable/converter/grule"
	"decisionTable/model"
	err "errors"
)

var (
	ErrDTableNoConverterFoundForTableStandard = err.New("no parser found for decision table standard")
	ErrDTableNoConverterFoundForOutputFormat  = err.New("no converter found for decision table output formats")
)

type OutputFormat string

const (
	GRL  OutputFormat = "GRL"
	JSON OutputFormat = "JSON"
)

func CreateConverterFactory(format OutputFormat) ConverterFactory {
	return ConverterFactory{format}
}

type ConverterFactory struct {
	format OutputFormat
}

func (c ConverterFactory) GetConverter(standard model.DTableStandard) (ConverterInterface, error) {
	switch standard {
	case model.GRULE:
		switch c.format {
		case GRL:
			return grule.CreateGrlConverter(), nil
		case JSON:
			return grule.CreateJsonConverter(), nil
		default:
			return nil, ErrDTableNoConverterFoundForOutputFormat
		}

	default:
		return nil, ErrDTableNoConverterFoundForTableStandard
	}
}
