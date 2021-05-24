package sfeel

import (
	"decisionTable/sfeel/antlr"
	"decisionTable/sfeel/ast"
	"decisionTable/sfeel/eval"
)

func CreateInputEntry(exp string) (Entry, []error) {
	tree, err := antlr.CreateParser(exp).Parse()
	if err != nil {
		return Entry{}, err
	}

	evl, err := eval.CreateInputEntryEvaluator().Eval(tree)
	if !evl {
		return Entry{}, err
	}

	return Entry{ast: tree, expression: exp}, nil
}

func CreateOutputEntry(exp string) (Entry, []error) {
	tree, err := antlr.CreateParser(exp).Parse()
	if err != nil {
		return Entry{}, err
	}

	evl, err := eval.CreateOutputEntryEvaluator().Eval(tree)
	if !evl {
		return Entry{}, err
	}

	return Entry{ast: tree, expression: exp}, nil
}

type Entry struct {
	ast        ast.Node
	expression string
}

func (e Entry) String() string {
	if e.ast == nil {
		return ""
	}
	return e.ast.String()
}
