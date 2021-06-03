package conv

import (
	"decisionTable/lang/sfeel/ast"
)

func CreateTreeWalker(listener SFeelBaseListenerInterface) TreeWalker {
	return TreeWalker{listener: listener}
}

type TreeWalker struct {
	listener SFeelBaseListenerInterface
}

func (w TreeWalker) Walk(node ast.Node) SFeelBaseListenerInterface {
	//Children of Node
	for _, val := range node.GetChildren() {
		w.Walk(val)
	}

	//Node
	w.ExitNode(node)
	return w.listener
}

func (w TreeWalker) ExitNode(node ast.Node) {
	switch node.(type) {
	case ast.EmptyStatement:
		w.listener.ExitEmptyStatement(node.(ast.EmptyStatement))
	case ast.QualifiedName:
		w.listener.ExitQualifiedName(node.(ast.QualifiedName))
	case ast.Integer:
		w.listener.ExitInteger(node.(ast.Integer))
	case ast.Float:
		w.listener.ExitFloat(node.(ast.Float))
	case ast.Boolean:
		w.listener.ExitBoolean(node.(ast.Boolean))
	case ast.String:
		w.listener.ExitString(node.(ast.String))
	case ast.DateTime:
		w.listener.ExitDateTime(node.(ast.DateTime))
	case ast.Interval:
		w.listener.ExitInterval(node.(ast.Interval))
	case ast.UnaryTest:
		w.listener.ExitUnaryTest(node.(ast.UnaryTest))
	case ast.UnaryTests:
		w.listener.ExitUnaryTests(node.(ast.UnaryTests))
	}

}
