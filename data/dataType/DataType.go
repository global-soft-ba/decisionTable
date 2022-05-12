package dataType

import (
	"fmt"
	"strings"
)

const (
	ErrDecisionTableUnknownDataType = "unknown data type"
)

type DataType string

const (
	Boolean  DataType = "boolean"
	Integer  DataType = "integer"
	Long     DataType = "long"
	Float    DataType = "float"
	Double   DataType = "double"
	String   DataType = "string"
	Date     DataType = "date"
	Time     DataType = "time"
	DateTime DataType = "datetime"
)

func (dt *DataType) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	switch s {
	case "boolean":
		*dt = Boolean
	case "integer":
		*dt = Integer
	case "long":
		*dt = Long
	case "float":
		*dt = Float
	case "double":
		*dt = Double
	case "string":
		*dt = String
	case "date":
		*dt = Date
	case "time":
		*dt = Time
	case "datetime":
		*dt = DateTime

	default:
		return fmt.Errorf("%s \"%s\"", ErrDecisionTableUnknownDataType, s)
	}

	return nil
}
