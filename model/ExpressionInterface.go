package model

type ExpressionInterface interface {
	String() string
	ValidateDataTypeOfExpression(varType DataTyp) (bool, error)
	ValidateExistenceOfFieldReferences(fields []Field) ([]Field, []error)
}
