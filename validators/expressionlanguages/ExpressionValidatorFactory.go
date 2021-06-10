package expressionlanguages

import (
	err "errors"
	"github.com/global-soft-ba/decisionTable/model"
	sFeelValidator "github.com/global-soft-ba/decisionTable/validators/expressionlanguages/sfeel"
)

var (
	ErrDTableNoParserFoundForExpressionLanguage = err.New("no parser found for expression language")
)

func CreateParserFactory() ExpressionValidatorFactory {
	return ExpressionValidatorFactory{}
}

type ExpressionValidatorFactory struct {
}

func (p ExpressionValidatorFactory) GetParser(notation model.ExpressionLanguage) (ExpressionValidatorInterface, error) {
	switch notation {
	case model.SFEEL:
		return sFeelValidator.CreateValidator(), nil
	default:
		return nil, ErrDTableNoParserFoundForExpressionLanguage
	}
}
