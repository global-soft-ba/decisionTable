package eval

import (
	"decisionTable/sfeel/ast"
	"errors"
	"fmt"
	"reflect"
)

func CreateTypeCheckingEvaluator() EvaluatorInterface {
	return TypeCheckingEvaluator{}
}

type TypeCheckingEvaluator struct {
	node ast.Node
}

func (e TypeCheckingEvaluator) Eval(node ast.Node) (bool, []error) {
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
