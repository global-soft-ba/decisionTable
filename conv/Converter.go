package conv

import (
	"decisionTable/conv/grule"
	"decisionTable/data"
	"errors"
)

var (
	ErrDTableNoConverterFoundForTableStandard = errors.New("no parser found for decision table standard")
)

func CreateConverter() Converter {
	return Converter{}
}

type Converter struct{}

func (c Converter) Convert(table data.Table) (interface{}, error) {

	switch table.NotationStandard {
	case data.GRULE:
		conv := grule.CreateGruleConverter()
		res, err := conv.Convert(table)
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
