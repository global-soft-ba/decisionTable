package generate

import "decisionTable/conv/grule/data"

func CreateGruleGenerator() GruleGenerator {
	return GruleGenerator{}
}

type GruleGenerator struct{}

func (g GruleGenerator) Generate(rules data.RuleSet) (interface{}, error) {
	return nil, nil
}
