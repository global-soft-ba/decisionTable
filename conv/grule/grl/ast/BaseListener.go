package ast

import (
	"github.com/global-soft-ba/decisionTable/ast"
)

func CreateGRLTreeWalker(listener GrlListenerInterface) ast.TreeWalker {
	return ast.CreateTreeWalker(&BaseListener{listener})
}

type BaseListener struct {
	grl GrlListenerInterface
}

func (l *BaseListener) EnterNode(node ast.Node) {}
func (l *BaseListener) ExitNode(node ast.Node) {
	switch node.(type) {
	case EmptyStatement:
		l.grl.ExitEmptyStatement(node.(EmptyStatement))
	case QualifiedName:
		l.grl.ExitQualifiedName(node.(QualifiedName))
	case Boolean:
		l.grl.ExitBoolean(node.(Boolean))
	case Integer:
		l.grl.ExitInteger(node.(Integer))
	case String:
		l.grl.ExitString(node.(String))
	case Float:
		l.grl.ExitFloat(node.(Float))
	case DateTime:
		l.grl.ExitDateTime(node.(DateTime))
	case Parentheses:
		l.grl.ExitParentheses(node.(Parentheses))
	case ArithmeticNegation:
		l.grl.ExitArithmeticNegation(node.(ArithmeticNegation))

	case MathOperations:
		l.grl.ExitMathOperations(node.(MathOperations))
	case ComparisonOperations:
		l.grl.ExitComparisonOperations(node.(ComparisonOperations))
	case LogicalOperations:
		l.grl.ExitLogicalOperations(node.(LogicalOperations))
	case AssignmentOperations:
		l.grl.ExitAssignmentOperations(node.(AssignmentOperations))
	case PowOperation:
		l.grl.ExitPowOperation(node.(PowOperation))
	}
}
