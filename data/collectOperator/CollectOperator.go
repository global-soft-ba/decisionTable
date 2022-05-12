package collectOperator

import (
	"fmt"
	"strings"
)

const (
	ErrDecisionTableUnknownCollectOperator = "unknown collect operator"
)

type CollectOperator string

const (
	List  CollectOperator = "list"
	Sum   CollectOperator = "sum"
	Min   CollectOperator = "min"
	Max   CollectOperator = "max"
	Count CollectOperator = "count"
)

func (co *CollectOperator) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	switch s {
	case "list":
		*co = List
	case "sum":
		*co = Sum
	case "min":
		*co = Min
	case "max":
		*co = Max
	case "count":
		*co = Count
	case "":
		*co = ""

	default:
		return fmt.Errorf("%s \"%s\"", ErrDecisionTableUnknownCollectOperator, s)
	}

	return nil
}
