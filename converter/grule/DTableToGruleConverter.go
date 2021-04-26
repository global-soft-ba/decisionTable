package grule

import (
	"decisionTable/converter"
	"decisionTable/converter/grule/mapper"
	"decisionTable/converter/grule/templates"
	"decisionTable/model"
	"html/template"
)

type DTableToGruleConverter struct {
}

func (c DTableToGruleConverter) Convert(data model.DTableData) ([]byte, error) {
	if data.NotationStandard != model.GRULE {
		return []byte{}, converter.ErrDTableNotationStandard
	}

	_, err := mapper.DTableToGrlMapper{}.MapToRuleSet(data)
	if err != nil {
		return []byte{}, err
	}

	return []byte{}, nil
}

func (c DTableToGruleConverter) converting(data model.DTableData) ([]byte, error) {
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
	return []byte{}, nil
}

func (c DTableToGruleConverter) selectPolicyTemplate(hitPolicy model.HitPolicy) (string, error) {
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

func (c DTableToGruleConverter) loadTokenTemplates() (template.Template, error) {

	return template.Template{}, nil
}

func (c DTableToGruleConverter) checkForInterference(table model.DTableData) bool {
	return false
}

func (c DTableToGruleConverter) createRuleEntry() {

}
