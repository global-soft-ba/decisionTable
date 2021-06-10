package interfaces

import (
	"github.com/global-soft-ba/decisionTable/data"
)

type ConverterInterface interface {
	Convert(data data.Table, outputFormat string) (interface{}, error)
}
