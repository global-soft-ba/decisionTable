package expressionlanguages

import (
	"decisionTable/model"
)

type ExpressionValidatorInterface interface {
	ValidateInputEntry(field model.Field, entry model.Entry) (bool, []error)
	ValidateOutputEntry(field model.Field, entry model.Entry) (bool, []error)
}
