package conv

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/conv/grule"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

var (
	ErrDTableNoConverterFoundForTableStandard = errors.New("no parser found for decision table standard")
)

func CreateConverter() Converter {
	return Converter{}
}

type Converter struct{}

func (c Converter) Convert(decisionTable data.DecisionTable, s standard.Standard) (interface{}, error) {

	switch decisionTable.Standard {
	case standard.GRULE:
		conv := grule.CreateConverter()
		res, err := conv.Convert(decisionTable, s)
		if err != nil {
			return nil, err
		}
		return res, nil

	case standard.DROOLS:
		panic("implement me")

	case standard.DMN:
		panic("implement me")

	default:
		return nil, ErrDTableNoConverterFoundForTableStandard
	}
}
