package expressionLanguage

import (
	"fmt"
	"strings"
)

type ExpressionLanguage string

const (
	ErrDecisionTableUnknownExpressionLanguage = "unknown expression language"
)

const (
	SFEEL ExpressionLanguage = "sfeel"
	FEEL  ExpressionLanguage = "feel"
)

func (el *ExpressionLanguage) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	switch s {
	case "sfeel":
		*el = SFEEL
	case "feel":
		*el = FEEL

	default:
		return fmt.Errorf("%s \"%s\"", ErrDecisionTableUnknownExpressionLanguage, s)
	}

	return nil
}
