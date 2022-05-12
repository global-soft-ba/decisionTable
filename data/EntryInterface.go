package data

import (
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
)

type EntryInterface interface {
	String() string
	ExpressionLanguage() expressionLanguage.ExpressionLanguage

	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType dataType.DataType) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []field.Field) ([]field.Field, []error)

	Convert(listener ast.SFeelListenerInterface)
}
