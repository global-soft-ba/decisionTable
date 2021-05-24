package eval

import "decisionTable/sfeel/ast"

type EvaluatorInterface interface {
	Eval(node ast.Node) (bool, []error)
}
