package data

import "decisionTable/data"

type Term struct {
	Field              data.FieldInterface
	Expression         ExpressionInterface
	ExpressionLanguage data.ExpressionLanguage
}
