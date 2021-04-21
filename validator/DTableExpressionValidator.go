package validator

import "decisionTable/model"

func CreateDTableExpressionValidator(dTable model.DTableData) DTableValidatorInterface {
	r := DTableExpressionValidator{dTable: dTable}
	return r
}

type DTableExpressionValidator struct {
	dTable model.DTableData
}

func (d DTableExpressionValidator) Validate() (bool, []error) {
	//Switch per ExpressionLanguage
	//Expression validation => Separate type mit []Entry und Inputfields
	return true, nil
}
