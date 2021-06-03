package convert

import (
	"decisionTable/convert/grule"
	"decisionTable/convert/interfaces"
	"decisionTable/data"
	"errors"
)

var (
	ErrDTableNoConverterFoundForTableStandard = errors.New("no parser found for decision table standard")
)

func CreateTableConverterFactory() DTableConverterFactory {
	return DTableConverterFactory{}
}

type DTableConverterFactory struct{}

func (c DTableConverterFactory) GetTableConverter(standard data.DTableStandard, format data.OutputFormat) (interfaces.ConverterInterface, error) {
	switch standard {
	case data.GRULE:
		conv, err := grule.CreateDTableToGruleConverterFactory().GetFormatConverter(format)
		if err != nil {
			return nil, err
		}
		return conv, nil
	case data.DROOLS:
		panic("implement me")
	case data.DMN:
		panic("implement me")
	default:
		return nil, ErrDTableNoConverterFoundForTableStandard
	}
}
