package eval

import (
	"decisionTable/lang/sfeel/ast"
)

type EvaluatorInterface interface {
	Eval(node ast.Node) (bool, []error)
}
