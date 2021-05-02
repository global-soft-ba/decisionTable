package converter

import (
	"decisionTable/converter/grule"
	"decisionTable/model"
	err "errors"
)

var (
	ErrDTableNoConverterFoundForTableStandard = err.New("no parser found for decision table standard")
)

func CreateConverterFactory() ConverterFactory {
	return ConverterFactory{}
}

type ConverterFactory struct{}

func (c ConverterFactory) GetConverter(standard model.DTableStandard) (ConverterInterface, error) {
	switch standard {
	case model.GRULE:
		return grule.CreateConverter(), nil
	default:
		return nil, ErrDTableNoConverterFoundForTableStandard
	}
}
