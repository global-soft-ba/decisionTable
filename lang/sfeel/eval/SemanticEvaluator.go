package eval

import (
	"errors"
	"fmt"
	ast "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	"reflect"
)

func CreateTypeCheckingEvaluator() EvaluatorInterface {
	return SemanticEvaluator{}
}

type SemanticEvaluator struct {
	node ast.Node
}

func (e SemanticEvaluator) Eval(node ast.Node) (bool, []error) {
	e.node = node

	switch e.node.(type) {
	case ast.Integer:
		// No additional Evaluation necessary
		return true, nil
	case ast.Float:
		// No additional evaluation necessary
		return true, nil
	case ast.String:
		// No additional evaluation necessary
		return true, nil
	case ast.DateTime:
		// No additional evaluation necessary
		return true, nil
	case ast.Boolean:
		// No additional evaluation necessary
		return true, nil

	}

	return false, []error{errors.New(fmt.Sprintf("given node type %s is not allowed as input entry", reflect.TypeOf(e.node)))}
}
