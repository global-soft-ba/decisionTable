package grlmodel

import "decisionTable/data"

type Term struct {
	Name               string
	Key                string
	Typ                data.DataTyp
	Expression         string
	ExpressionLanguage data.ExpressionLanguage
}
