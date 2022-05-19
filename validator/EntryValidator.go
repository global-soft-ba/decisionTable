package validator

import (
	"fmt"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/lang/sfeel"
)

const (
	ErrDecisionTableEntryExpressionLanguageIsInvalid = "invalid entry expression language"
)

func CreateInputEntryValidator(el expressionLanguage.ExpressionLanguage, entry string) (data.EntryValidatorInterface, error) {
	switch el {

	case expressionLanguage.SFEEL:
		return sfeel.CreateInputEntryValidator(entry), nil

	case expressionLanguage.FEEL:
		panic("implement me")

	default:
		return nil, fmt.Errorf("%s \"%s\"", ErrDecisionTableEntryExpressionLanguageIsInvalid, el)
	}
}

func CreateOutputEntryValidator(el expressionLanguage.ExpressionLanguage, entry string) (data.EntryValidatorInterface, error) {
	switch el {

	case expressionLanguage.SFEEL:
		return sfeel.CreateOutputEntryValidator(entry), nil

	case expressionLanguage.FEEL:
		panic("implement me")

	default:
		return nil, fmt.Errorf("%s \"%s\"", ErrDecisionTableEntryExpressionLanguageIsInvalid, el)
	}
}
