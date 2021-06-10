package data

import "github.com/global-soft-ba/decisionTable/data"

type Term struct {
	Field              data.FieldInterface
	Expression         ExpressionInterface
	ExpressionLanguage data.ExpressionLanguage
}
