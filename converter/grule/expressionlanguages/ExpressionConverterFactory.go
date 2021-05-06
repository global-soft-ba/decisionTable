package expressionlanguages

import (
	"decisionTable/converter/grule/expressionlanguages/sfeel"
	"decisionTable/model"
	"errors"
)

var (
	ErrDTableNoConverterFoundForExpressionLanguage = errors.New("no expression language found for source expression language found")
)

func CreateExpresionConverterFactory() ExpressionConverterFactory {
	return ExpressionConverterFactory{}
}

type ExpressionConverterFactory struct{}

func (e ExpressionConverterFactory) GetExpressionConverter(sourceLang model.ExpressionLanguage, outputFormat model.OutputFormat) (ExpressionConverterInterface, error) {

	switch sourceLang {
	case model.SFEEL:
		conv, err := sfeel.CreateSFeelConverterFactory().GetConverterForFormat(outputFormat)
		if err != nil {
			return nil, err
		}
		return conv, nil
	default:
		return nil, ErrDTableNoConverterFoundForExpressionLanguage
	}

}
