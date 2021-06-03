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
	if len(node.GetChildren()) == 0 {
		w.ExitNode(node)
	}

	for _, val := range node.GetChildren() {
		w.Walk(val)
	}

	return w.listener
}

func (w TreeWalker) ExitNode(node ast.Node) {
	switch node.(type) {
	case ast.EmptyStatement:
		w.listener.ExitEmptyStatement("")
	case ast.QualifiedName:
		w.listener.ExitQualifiedName()
	case ast.Integer:
		w.listener.ExitInteger()
	case ast.Float:
		w.listener.ExitFloat()
	}

}
