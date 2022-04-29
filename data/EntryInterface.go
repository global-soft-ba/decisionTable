package data

import (
	"github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
)

type EntryInterface interface {
	String() string
	ExpressionLanguage() ExpressionLanguage

	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType DataType) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []FieldInterface) ([]FieldInterface, []error)

	Convert(listener ast.SFeelListenerInterface)
}
