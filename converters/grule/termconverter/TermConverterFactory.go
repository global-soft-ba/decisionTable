package termconverter

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/converters/grule/termconverter/sfeel"
	"github.com/global-soft-ba/decisionTable/model"
)

var (
	ErrDTableNoConverterFoundForExpressionLanguage = errors.New("no expression language found for source expression language found")
)

func CreateTermConverterFactory() TermConverterFactory {
	return TermConverterFactory{}
}

type TermConverterFactory struct{}

func (e TermConverterFactory) GetExpressionConverter(sourceLang model.ExpressionLanguage, outputFormat model.OutputFormat) (TermConverterInterface, error) {

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
