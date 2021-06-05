package data

import "decisionTable/data"

type Term struct {
	Name               string
	Key                string
	Typ                data.DataTyp
	Expression         ExpressionInterface
	ExpressionLanguage data.ExpressionLanguage
}
