package termconverter

import (
	"decisionTable/convert/grule/termconverter/sfeel"
	"decisionTable/data"
	"errors"
)

var (
	ErrDTableNoConverterFoundForExpressionLanguage = errors.New("no expression language found for source expression language found")
)

func CreateTermConverterFactory() TermConverterFactory {
	return TermConverterFactory{}
}

type TermConverterFactory struct{}

func (e TermConverterFactory) GetExpressionConverter(sourceLang data.ExpressionLanguage, outputFormat data.OutputFormat) (TermConverterInterface, error) {

	switch sourceLang {
	case data.SFEEL:
		conv, err := sfeel.CreateSFeelConverterFactory().GetConverterForFormat(outputFormat)
		if err != nil {
			return nil, err
		}
		return conv, nil
	default:
		return nil, ErrDTableNoConverterFoundForExpressionLanguage
	}

}
