package grule

import (
	"decisionTable/converter"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/converter/grule/mapper"
	"decisionTable/converter/grule/templates"
	"decisionTable/model"
	"html/template"
)

func CreateDTableToGrlConverter() Converter {
	return Converter{}
}

type Converter struct {
}

func (c Converter) Convert(data model.DTableData) ([]string, error) {
	if data.NotationStandard != model.GRULE {
		return []string{}, converter.ErrDTableNotationStandard
	}

	grlModel, err := mapper.CreateGruleMapper().MapToRuleSet(data)
	if err != nil {
		return []string{}, err
	}

	result, err := c.converting(grlModel)
	if err != nil {
		return []string{}, err
	}

	return result, nil
}

func (c Converter) converting(ruleSet grlmodel.RuleSet) ([]string, error) {
	/*

		var result []byte
		templateRaw,err := c.selectPolicyTemplate(data.HitPolicy)
		if err != nil {
			return []byte{}, err
		}

		t, err := template.New("Rule").Parse(templateRaw)

		//2.) Convertiere jede Regel gemäß dem Template

		for k, v := range data.Rules {

			// Execute Template ()
		}

		return result, nil*/
	return []string{}, nil
}

func (c Converter) selectPolicyTemplate(hitPolicy model.HitPolicy) (string, error) {
	switch hitPolicy {
	case model.Unique:
		return templates.UNIQUE, nil
	case model.First:
		return templates.FIRST, nil
	case model.Priority:
		return templates.PRIORITY, nil

	default:
		return "", converter.ErrDTableHitPolicy
	}
}

func (c Converter) loadTokenTemplates() (template.Template, error) {

	return template.Template{}, nil
}

func (c Converter) checkForInterference(table model.DTableData) bool {
	return false
}

func (c Converter) createRuleEntry() {

}
