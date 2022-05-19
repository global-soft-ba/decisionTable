package data

import (
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/field"
)

type EntryValidatorInterface interface {
	Validate() (bool, []error)
	ValidateDataTypeOfExpression(varType dataType.DataType) (bool, error)
	ValidateExistenceOfFieldReferencesInExpression(fields []field.Field) ([]field.Field, []error)
}
