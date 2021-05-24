package eval

import (
	"decisionTable/sfeel/ast"
	"errors"
	"fmt"
	"reflect"
)

//Todo implement output entry rules

func CreateOutputEntryEvaluator() EvaluatorInterface {
	return OutputEntryEvaluator{}
}

type OutputEntryEvaluator struct {
	node ast.Node
}

func (e OutputEntryEvaluator) Eval(node ast.Node) (bool, []error) {
	e.node = node
	/*
		switch e.node.(type) {
			default: return true, nil
		}
	*/
	return false, []error{errors.New(fmt.Sprintf("given node type %s is not allowed as output entry", reflect.TypeOf(e.node)))}
}
