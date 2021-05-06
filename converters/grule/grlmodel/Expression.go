package grlmodel

import "decisionTable/model"

type Term struct {
	Name               string
	Identifier         string
	Typ                model.VariableTyp
	Expression         string
	ExpressionLanguage model.ExpressionLanguage
}
