package eval

import (
	"github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
)

type EvaluatorInterface interface {
	Eval(node ast.Node) (bool, []error)
}
