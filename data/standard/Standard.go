package standard

import (
	"fmt"
	"strings"
)

const (
	ErrDecisionTableUnknownStandard = "unknown standard"
)

type Standard string

const (
	GRULE  Standard = "grule"
	DMN    Standard = "dmn"
	DROOLS Standard = "drools"
)

func (st *Standard) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	switch s {
	case "grule":
		*st = GRULE
	case "dmn":
		*st = DMN
	case "drools":
		*st = DROOLS

	default:
		return fmt.Errorf("%s \"%s\"", ErrDecisionTableUnknownStandard, s)
	}

	return nil
}
