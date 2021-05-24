package eval

import "decisionTable/SFeel/ast"

type EvaluatorInterface interface {
	Eval(node ast.Node) (bool, []error)
}
