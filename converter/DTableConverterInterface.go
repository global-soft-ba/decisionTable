package converter

import (
	"decisionTable/model"
	"errors"
)

var (
	ErrDTableHitPolicy        = errors.New("hit policy is not supported by decision table converter")
	ErrDTableNotationStandard = errors.New("notation standard is not supported by decision table converter")
)

type DTableConverterInterface interface {
	Convert(data model.DTableData) ([]byte, error)
}
