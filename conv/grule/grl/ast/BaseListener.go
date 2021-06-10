package ast

import "github.com/global-soft-ba/decisionTable/ast"

func CreateGRLTreeWalker(listener GrlListenerInterface) ast.TreeWalker {
	return ast.CreateTreeWalker(&BaseListener{listener})
}

type BaseListener struct {
	grl GrlListenerInterface
}

func (l *BaseListener) EnterNode(node ast.Node) {}
func (l *BaseListener) ExitNode(node ast.Node) {
	switch node.(type) {
	case Integer:
		l.grl.ExitInteger(node.(Integer))
	case String:
		l.grl.ExitString(node.(String))
	case MathOperations:
		l.grl.ExitMathOperations(node.(MathOperations))
	case ComparisonOperations:
		l.grl.ExitComparisonOperations(node.(ComparisonOperations))
	case LogicalOperations:
		l.grl.ExitLogicalOperations(node.(LogicalOperations))
	}
}
