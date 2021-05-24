package eval

import (
	"decisionTable/SFeel/ast"
	"errors"
	"fmt"
	"reflect"
)

func CreateEvaluator() EvaluatorInterface {
	return Evaluator{}
}

type Evaluator struct {
	node ast.Node
}

func (e Evaluator) Eval(node ast.Node) (bool, []error) {
	e.node = node

	switch e.node.(type) {
	case ast.UnaryTests:
		// No additional Evaluation necessary
		return true, nil
	case ast.EmptyUnaryTest:
		// No additional evaluation necessary
		return true, nil
	}

	return false, []error{errors.New(fmt.Sprintf("given node type %s cannot be evaluated at this level or is unknown", reflect.TypeOf(e.node)))}
}
