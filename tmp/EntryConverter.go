package tmp

import (
	"fmt"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/lang/sfeel"
)

const (
	ErrDecisionTableEntryExpressionLanguageIsInvalid = "invalid entry expression language"
)

func CreateInputEntryConverter(el expressionLanguage.ExpressionLanguage, entry string) (data.EntryConverterInterface, error) {
	switch el {

	case expressionLanguage.SFEEL:
		return sfeel.CreateInputEntryConverter(entry), nil

	case expressionLanguage.FEEL:
		panic("implement me")

	default:
		return nil, fmt.Errorf("%s \"%s\"", ErrDecisionTableEntryExpressionLanguageIsInvalid, el)
	}
}

func CreateOutputEntryConverter(el expressionLanguage.ExpressionLanguage, entry string) (data.EntryConverterInterface, error) {
	switch el {

	case expressionLanguage.SFEEL:
		return sfeel.CreateOutputEntryConverter(entry), nil

	case expressionLanguage.FEEL:
		panic("implement me")

	default:
		return nil, fmt.Errorf("%s \"%s\"", ErrDecisionTableEntryExpressionLanguageIsInvalid, el)
	}
}
