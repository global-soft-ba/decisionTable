package SFeel

import (
	"decisionTable/SFeel/ast"
	"decisionTable/SFeel/eval"
)

func CreateTerm(exp string) Term {
	return Term{ast: nil, expression: exp}
}

type Term struct {
	ast        ast.Node
	expression string
}

func (a Term) String() string {
	return a.ast.String()
}

func (a Term) Eval(evaluator eval.EvaluatorInterface) (bool, []error) {
	return evaluator.Eval(a.ast)
}
