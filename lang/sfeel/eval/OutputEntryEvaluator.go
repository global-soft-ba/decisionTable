package eval

import (
	"errors"
	"fmt"
	ast "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
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
	switch e.node.(type) {
	case ast.EmptyStatement:
		// No additional evaluation necessary
		return true, nil

	}

	return false, []error{errors.New(fmt.Sprintf("given node type %s is not allowed as output entry", reflect.TypeOf(e.node)))}
}
