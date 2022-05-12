package interfaces

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type ConverterInterface interface {
	Convert(decisionTable data.DecisionTable, standard standard.Standard) (interface{}, error)
}
