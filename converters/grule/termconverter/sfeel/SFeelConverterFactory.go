package sfeel

import (
	"decisionTable/model"
	"errors"
)

var (
	ErrExpressionLanguageConverterForOutputFormatNotFound = errors.New("no converter for output language found")
)

func CreateSFeelConverterFactory() ConverterFactory {
	return ConverterFactory{}
}

type ConverterFactory struct{}

func (s ConverterFactory) GetConverterForFormat(format model.OutputFormat) (SFeelTermConverter, error) {
	switch format {
	case model.GRL:
		return CreateSfeelTermToGrlConverter(), nil
	case model.JSON:
		return CreateSfeelTermToJsonConverter(), nil
	default:
		return SFeelTermConverter{}, ErrExpressionLanguageConverterForOutputFormatNotFound
	}
}
