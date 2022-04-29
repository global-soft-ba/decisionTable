package grule

import (
	"github.com/global-soft-ba/decisionTable/conv/grule/conv"
	"github.com/global-soft-ba/decisionTable/conv/grule/generate"
	"github.com/global-soft-ba/decisionTable/conv/interfaces"
	"github.com/global-soft-ba/decisionTable/data"
)

func CreateConverter() interfaces.ConverterInterface {
	return Converter{}
}

type Converter struct{}

func (c Converter) Convert(data data.DecisionTable, format string) (interface{}, error) {

	converter := conv.CreateTableToGruleConverter()
	grule, err := converter.Convert(data)
	if err != nil {
		return []string{}, err
	}

	generator := generate.CreateGruleGenerator()
	res, err := generator.Generate(grule, format)
	if err != nil {
		return []string{}, err
	}

	return res, nil
}
