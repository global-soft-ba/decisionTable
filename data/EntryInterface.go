package data

import (
	"decisionTable/lang/sfeel/ast"
)

type EntryInterface interface {
	String() string
	ExpressionLanguage() ExpressionLanguage

	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType DataTyp) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []FieldInterface) ([]FieldInterface, []error)

	Convert(listener ast.SFeelListenerInterface)
}
