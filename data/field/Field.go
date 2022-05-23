package field

import "github.com/global-soft-ba/decisionTable/data/dataType"

type Field struct {
	Name string            `json:"name"`
	Type dataType.DataType `json:"type"`
}
