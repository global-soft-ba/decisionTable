package interfaces

import (
	"decisionTable/data"
)

type ConverterInterface interface {
	Convert(data data.Table, outputFormat string) (interface{}, error)
}
