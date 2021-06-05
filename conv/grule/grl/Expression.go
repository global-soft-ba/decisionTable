package grl

import (
	"decisionTable/ast"
	grule "decisionTable/conv/grule/data"
	grl "decisionTable/conv/grule/grl/ast"
	conv "decisionTable/conv/grule/grl/conv"
	"decisionTable/conv/grule/grl/generate"
	dTable "decisionTable/data"
	"errors"
)

var (
	ErrGruleExpressionLanguageNotSupported = errors.New("expression language of entry is not supported")
	ErrGruleOutputFormatNotSupported       = errors.New("output format not supported")
)

func CreateExpression(fieldName string, entry dTable.EntryInterface) (grule.ExpressionInterface, error) {

	lang := entry.ExpressionLanguage()

	switch lang {
	case dTable.SFEEL:
		res, err := conv.CreateSFeelToGrlAstConverter().ConvertToGrlAst(fieldName, entry)
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

		walker := grl.CreateGRLTreeWalker(gen)
		walker.Walk(e.tree)
		return gen.GetCode(), nil
	}
	return "", ErrGruleOutputFormatNotSupported
}
