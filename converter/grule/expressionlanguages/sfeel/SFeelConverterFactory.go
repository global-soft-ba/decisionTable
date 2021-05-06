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

func (s ConverterFactory) GetConverterForFormat(format model.OutputFormat) (SFeelConverter, error) {
	switch format {
	case model.GRL:
		return CreateSfeelToGrlConverter(), nil
	case model.JSON:
		return CreateSfeelToJsonConverter(), nil
	default:
		return SFeelConverter{}, ErrExpressionLanguageConverterForOutputFormatNotFound
	}
}
