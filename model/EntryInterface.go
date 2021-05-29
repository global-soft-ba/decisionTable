package model

type EntryInterface interface {
	String() string
	ExpressionLanguage() ExpressionLanguage
	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType DataTyp) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []Field) ([]Field, []error)
}
