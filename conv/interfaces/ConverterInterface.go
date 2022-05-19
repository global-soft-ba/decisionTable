package interfaces

import (
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type ConverterInterface interface {
	Convert(decisionTable decisionTable.DecisionTable, standard standard.Standard) (interface{}, error)
}
