package sfeel

import (
	"decisionTable/data"
	"errors"
)

var (
	ErrExpressionLanguageConverterForOutputFormatNotFound = errors.New("no converter for output language found")
)

func CreateSFeelConverterFactory() ConverterFactory {
	return ConverterFactory{}
}

type ConverterFactory struct{}

func (s ConverterFactory) GetConverterForFormat(format data.OutputFormat) (SFeelTermConverter, error) {
	switch format {
	case data.GRL:
		return CreateSfeelTermToGrlConverter(), nil
	case data.JSON:
		return CreateSfeelTermToJsonConverter(), nil
	default:
		return SFeelTermConverter{}, ErrExpressionLanguageConverterForOutputFormatNotFound
	}
}
