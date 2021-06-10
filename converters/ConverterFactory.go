package converters

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/converters/grule"
	"github.com/global-soft-ba/decisionTable/converters/interfaces"
	"github.com/global-soft-ba/decisionTable/model"
)

var (
	ErrDTableNoConverterFoundForTableStandard = errors.New("no parser found for decision table standard")
)

func CreateTableConverterFactory() DTableConverterFactory {
	return DTableConverterFactory{}
}

type DTableConverterFactory struct{}

func (c DTableConverterFactory) GetTableConverter(standard model.DTableStandard, format model.OutputFormat) (interfaces.ConverterInterface, error) {
	switch standard {
	case model.GRULE:
		conv, err := grule.CreateDTableToGruleConverterFactory().GetFormatConverter(format)
		if err != nil {
			return nil, err
		}
		return conv, nil
	case model.DROOLS:
		panic("implement me")
	case model.DMN:
		panic("implement me")
	default:
		return nil, ErrDTableNoConverterFoundForTableStandard
	}
}
