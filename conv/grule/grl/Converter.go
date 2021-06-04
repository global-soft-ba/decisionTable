package grl

import (
	grule "decisionTable/conv/grule/data"
	grlconv "decisionTable/conv/grule/grl/conv"
	table "decisionTable/data"
	"errors"
)

var (
	ErrGruleExpressionLanguageNotSupported = errors.New("expression language is not supported")
)

type Converter struct{}

func (c Converter) Convert(fieldName string, entry table.EntryInterface, targetFormat grule.OutputFormat) (string, error) {

	lang := entry.ExpressionLanguage()

	switch lang {
	case table.SFEEL:
		if targetFormat == grule.GRL {
			grlConverter := grlconv.CreateSFeelToGrlConverter()
			res, err := grlConverter.Convert(fieldName, entry)
			if err != nil {
				return "", err
			}
			return res, nil

		} else {
			panic("implement me")
		}
	}

	return "", ErrGruleExpressionLanguageNotSupported
}
