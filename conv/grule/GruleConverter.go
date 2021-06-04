package grule

import (
	"decisionTable/conv/grule/conv"
	"decisionTable/conv/grule/generate"
	"decisionTable/data"
)

const (
	Rule         = "Rule"
	RuleName     = "RuleName"
	Salience     = "Salience"
	Expressions  = "Expressions"
	Assignments  = "Assignments"
	Entries      = "Entries"
	Interference = "Interference"
)

func CreateGruleConverter() GruleConverter {
	return GruleConverter{}
}

type GruleConverter struct{}

func (c GruleConverter) Convert(data data.Table) (interface{}, error) {

	converter := conv.CreateTableToGruleConverter()
	grule, err := converter.Convert(data)
	if err != nil {
		return []string{}, err
	}

	generator := generate.CreateGruleGenerator()
	res, err := generator.Generate(grule)
	if err != nil {
		return []string{}, err
	}

	return res, nil
}
