package expressionlanguages

import (
	"github.com/global-soft-ba/decisionTable/model"
)

type ExpressionValidatorInterface interface {
	ValidateInputEntry(field model.Field, entry model.Entry) (bool, []error)
	ValidateOutputEntry(field model.Field, entry model.Entry) (bool, []error)
}
