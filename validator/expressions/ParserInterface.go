package expressions

import "decisionTable/model"

type ParserInterface interface {
	ValidateEntry(entry model.Entry) (bool, error)
}
