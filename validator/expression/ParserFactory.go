package expression

import (
	"decisionTable/model"
	sfeelParser "decisionTable/validator/expression/sfeel"
	err "errors"
)

var (
	ErrDTableNoParserFoundForExpressionLanguage = err.New("no parser found for expression language")
)

func CreateParserFactory() ParserFactory {
	return ParserFactory{}
}

type ParserFactory struct {
}

func (p ParserFactory) GetParser(notation model.ExpressionLanguage) (ParserInterface, error) {
	switch notation {
	case model.SFEEL:
		return sfeelParser.CreateParser(), nil
	default:
		return nil, ErrDTableNoParserFoundForExpressionLanguage
	}
}
