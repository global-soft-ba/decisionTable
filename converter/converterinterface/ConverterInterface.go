package converterinterface

import (
	"decisionTable/model"
)

type ConverterInterface interface {
	Convert(data model.TableData) (interface{}, error)
}
