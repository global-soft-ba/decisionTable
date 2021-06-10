package ast

import (
	"github.com/global-soft-ba/decisionTable/ast"
)

func CreateSFeelTreeWalker(listener SFeelListenerInterface) ast.TreeWalker {
	return ast.CreateTreeWalker(&BaseListener{listener})
}

type BaseListener struct {
	sfeel SFeelListenerInterface
}

func (l *BaseListener) EnterNode(node ast.Node) {}
func (l *BaseListener) ExitNode(node ast.Node) {
	switch node.(type) {
	case EmptyStatement:
		l.sfeel.ExitEmptyStatement(node.(EmptyStatement))
	case QualifiedName:
		l.sfeel.ExitQualifiedName(node.(QualifiedName))
	case Integer:
		l.sfeel.ExitInteger(node.(Integer))
	case Float:
		l.sfeel.ExitFloat(node.(Float))
	case Boolean:
		l.sfeel.ExitBoolean(node.(Boolean))
	case String:
		l.sfeel.ExitString(node.(String))
	case DateTime:
		l.sfeel.ExitDateTime(node.(DateTime))
	case Interval:
		l.sfeel.ExitInterval(node.(Interval))
	case UnaryTest:
		l.sfeel.ExitUnaryTest(node.(UnaryTest))
	case UnaryTests:
		l.sfeel.ExitUnaryTests(node.(UnaryTests))
	}
}
