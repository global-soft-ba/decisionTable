package interfaces

import (
	"github.com/global-soft-ba/decisionTable/model"
)

type ConverterInterface interface {
	Convert(data model.TableData) (interface{}, error)
}
