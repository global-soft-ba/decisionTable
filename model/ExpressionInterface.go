package model

type ExpressionInterface interface {
	String() string
	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType DataTyp) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []Field) ([]Field, []error)
}
