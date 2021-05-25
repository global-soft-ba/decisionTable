package grlmodel

import "github.com/global-soft-ba/decisionTable/model"

type Term struct {
	Name               string
	Key                string
	Typ                model.VariableTyp
	Expression         string
	ExpressionLanguage model.ExpressionLanguage
}