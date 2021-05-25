package eval

import (
	"decisionTable/sfeel/ast"
	"errors"
	"fmt"
	"reflect"
)

func CreateInputEntryEvaluator() EvaluatorInterface {
	return InputEntryEvaluator{}
}

type InputEntryEvaluator struct {
	node ast.Node
}

func (e InputEntryEvaluator) Eval(node ast.Node) (bool, []error) {
	e.node = node

	switch e.node.(type) {
	case ast.UnaryTests:
		// No additional Evaluation necessary
		return true, nil
	case ast.EmptyUnaryTest:
		// No additional evaluation necessary
		return true, nil
	}

	return false, []error{errors.New(fmt.Sprintf("given node type %s is not allowed as input entry", reflect.TypeOf(e.node)))}
}