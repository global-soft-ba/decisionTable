package data

import (
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
)

type Term struct {
	Field              field.Field
	Expression         ExpressionInterface
	ExpressionLanguage expressionLanguage.ExpressionLanguage
}
