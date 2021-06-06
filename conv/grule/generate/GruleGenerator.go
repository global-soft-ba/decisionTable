package generate

import (
	"bytes"
	grule "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/generate/grl"
	"decisionTable/conv/grule/generate/json"
	"errors"
	"text/template"
)

var (
	ErrGruleOutputFormatNotSupported = errors.New("output format not supported")
)

func CreateGruleGenerator() GruleGenerator {
	return GruleGenerator{}
}

type GruleGenerator struct {
	templates *template.Template
	format    grule.OutputFormat
}

func (g *GruleGenerator) Generate(rules grule.RuleSet, targetFormat string) (interface{}, error) {
	switch targetFormat {
	case string(grule.GRL):
		tmpl, err := grl.GenerateTemplates(rules.HitPolicy, rules.Interference)
		if err != nil {
			return nil, err
		}
		g.templates = tmpl
		g.format = grule.GRL
		return g.generate(rules)

	case string(grule.JSON):
		tmpl, err := json.GenerateTemplates(rules.HitPolicy, rules.Interference)
		if err != nil {
			return nil, err
		}
		g.templates = tmpl
		g.format = grule.JSON
		return g.generate(rules)

	}
	return nil, ErrGruleOutputFormatNotSupported
}

func (g *GruleGenerator) generate(ruleSet grule.RuleSet) ([]string, error) {
	var result []string
	for _, v := range ruleSet.Rules {
		var tpl bytes.Buffer
		err := g.templates.Execute(&tpl, v)
		if err != nil {
			return []string{}, err
		}
		result = append(result, tpl.String())
	}
	return result, nil
}
