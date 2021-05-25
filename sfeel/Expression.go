package sfeel

import (
	"decisionTable/sfeel/antlr"
	"decisionTable/sfeel/ast"
	"decisionTable/sfeel/eval"
)

func CreateInputExpression(exp string) (Expression, []error) {
	tree, err := antlr.CreateParser(exp).Parse()
	if err != nil {
		return Expression{}, err
	}

	evl, err := eval.CreateInputEntryEvaluator().Eval(tree)
	if !evl {
		return Expression{}, err
	}

	return Expression{ast: tree, expression: exp}, nil
}

func CreateOutputExpression(exp string) (Expression, []error) {
	tree, err := antlr.CreateParser(exp).Parse()
	if err != nil {
		return Expression{}, err
	}

	evl, err := eval.CreateOutputEntryEvaluator().Eval(tree)
	if !evl {
		return Expression{}, err
	}

	return Expression{ast: tree, expression: exp}, nil
}

type Expression struct {
	ast        ast.Node
	expression string
}

func (e Expression) String() string {
	if e.ast == nil {
		return ""
	}
	return e.ast.String()
}
