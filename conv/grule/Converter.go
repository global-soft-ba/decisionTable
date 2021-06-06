package grule

import (
	"decisionTable/conv/grule/conv"
	"decisionTable/conv/grule/generate"
	"decisionTable/conv/interfaces"
	"decisionTable/data"
)

func CreateConverter() interfaces.ConverterInterface {
	return Converter{}
}

type Converter struct{}

func (c Converter) Convert(data data.Table, format string) (interface{}, error) {

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
