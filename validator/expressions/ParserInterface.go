package expressions

import "decisionTable/model"

type ParserInterface interface {
	ValidateInputEntry(entry model.Entry) (bool, error)
	ValidateOutputEntry(entry model.Entry) (bool, error)
}
