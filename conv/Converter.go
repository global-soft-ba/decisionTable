package conv

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/conv/grule"
	"github.com/global-soft-ba/decisionTable/data"
)

var (
	ErrDTableNoConverterFoundForTableStandard = errors.New("no parser found for decision table standard")
)

func CreateConverter() Converter {
	return Converter{}
}

type Converter struct{}

func (c Converter) Convert(table data.Table, format string) (interface{}, error) {

	switch table.NotationStandard {
	case data.GRULE:
		conv := grule.CreateConverter()
		res, err := conv.Convert(table, format)
		if err != nil {
			return nil, err
		}
		return res, nil

	case data.DROOLS:
		panic("implement me")
	case data.DMN:
		panic("implement me")
	default:
		return nil, ErrDTableNoConverterFoundForTableStandard
	}
}
