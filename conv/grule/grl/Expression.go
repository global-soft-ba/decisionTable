package grl

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/ast"
	grule "github.com/global-soft-ba/decisionTable/conv/grule/data"
	grl "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	conv "github.com/global-soft-ba/decisionTable/conv/grule/grl/conv"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl/generate"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
)

var (
	ErrGruleExpressionLanguageNotSupported = errors.New("expression language of entry is not supported")
	ErrGruleOutputFormatNotSupported       = errors.New("output format not supported")
)

func CreateExpression(field field.Field, el expressionLanguage.ExpressionLanguage, entryType entryType.EntryType, entry string) (grule.ExpressionInterface, error) {

	switch el {
	case expressionLanguage.SFEEL:
		res, err := conv.CreateSFeelToGrlAstConverter().ConvertToGrlAst(field, entryType, entry)
		if err != nil {
			return nil, err
		}
		return Expression{tree: res}, nil
	}

	return nil, ErrGruleExpressionLanguageNotSupported
}

type Expression struct {
	tree ast.Node
}

func (e Expression) String() string {
	return e.tree.String()
}

func (e Expression) Convert(targetFormat grule.OutputFormat) (string, error) {
	if targetFormat == grule.GRL {
		// Walk the grl ast tree and convert
		gen, err := generate.CreateGrlGeneratorListener()
		if err != nil {
			return "", err
		}

		walker := grl.CreateGRLTreeWalker(&gen)
		walker.Walk(e.tree)
		return gen.GetCode(), nil
	}
	return "", ErrGruleOutputFormatNotSupported
}
