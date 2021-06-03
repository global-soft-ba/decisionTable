package data

import "decisionTable/lang/sfeel/conv"

type EntryInterface interface {
	String() string
	ExpressionLanguage() ExpressionLanguage

	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType DataTyp) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []Field) ([]Field, []error)

	Convert(listener conv.SFeelBaseListenerInterface) conv.SFeelBaseListenerInterface
}
