package expressionlanguages

import (
	"decisionTable/model"
	sFeelValidator "decisionTable/validators/expressionlanguages/sfeel"
	err "errors"
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
