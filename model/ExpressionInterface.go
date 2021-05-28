package model

type ExpressionInterface interface {
	String() string
	ValidateDataTypeOfExpression(varType DataTyp) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []Field) ([]Field, []error)
}
