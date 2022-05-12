package grule

import (
	"github.com/global-soft-ba/decisionTable/conv/grule/conv"
	"github.com/global-soft-ba/decisionTable/conv/grule/generate"
	"github.com/global-soft-ba/decisionTable/conv/interfaces"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

func CreateConverter() interfaces.ConverterInterface {
	return Converter{}
}

type Converter struct{}

func (c Converter) Convert(decisionTable data.DecisionTable, standard standard.Standard) (interface{}, error) {

	converter := conv.CreateTableToGruleConverter()
	grule, err := converter.Convert(decisionTable)
	if err != nil {
		return []string{}, err
	}

	generator := generate.CreateGruleGenerator()
	res, err := generator.Generate(grule, string(standard))
	if err != nil {
		return []string{}, err
	}

	return res, nil
}
