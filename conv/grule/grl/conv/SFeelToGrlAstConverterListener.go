package conv

import (
	"decisionTable/ast"
	grl "decisionTable/conv/grule/grl/ast"
	"decisionTable/data"
	sfeel "decisionTable/lang/sfeel/ast"
)

var OperatorMappingTable = map[int]int{
	sfeel.SFeelOperatorLESS:        grl.ComparisonOperatorLESS,
	sfeel.SFeelOperatorLESSEQ:      grl.ComparisonOperatorLESSEQ,
	sfeel.SFeelOperatorGREATER:     grl.ComparisonOperatorGREATER,
	sfeel.SFeelOperatorGREATEREQ:   grl.ComparisonOperatorGREATEREQ,
	sfeel.SFeelClosedIntervalSTART: grl.ComparisonOperatorGREATEREQ, // [1..   [e1..e2] if and only if o ≥ e1 and o ≤ e2
	sfeel.SFeelClosedIntervalEND:   grl.ComparisonOperatorLESSEQ,    // ..1]
	sfeel.SFeelOpenIntervalSTART:   grl.ComparisonOperatorGREATER,   // ]1..
	sfeel.SFeelOpenIntervalEND:     grl.ComparisonOperatorLESS,      // ..1[
}

func CreateSFeelToGrlAstConverterListener() SFeelToGrlAstConverterListener {
	return SFeelToGrlAstConverterListener{symbolMapping: OperatorMappingTable, stack: ast.NewStack()}
}

type SFeelToGrlAstConverterListener struct {
	sfeel.SFeelListener
	field         data.FieldInterface
	stack         *ast.Stack
	symbolMapping map[int]int
	Errors        []error
}

func (l *SFeelToGrlAstConverterListener) GetAST() ast.Node {
	return l.stack.Pop().(ast.Node)
}

func (l *SFeelToGrlAstConverterListener) ExitInteger(ctx sfeel.Integer) {
	i := grl.Integer{Val: ctx.Value}
	l.stack.Push(i)
}

func (l *SFeelToGrlAstConverterListener) ExitQualifiedName(ctx sfeel.QualifiedName) {
	q := grl.QualifiedName{Val: ctx.Value}
	l.stack.Push(q)
}

func (l *SFeelToGrlAstConverterListener) ExitInterval(ctx sfeel.Interval) {

	rightVal := l.stack.Pop()
	leftVal := l.stack.Pop()

	left := grl.ComparisonOperations{
		Left:     grl.String{Val: l.field.GetQualifiedName()},
		Operator: l.symbolMapping[ctx.StartIntervalRule.Type],
		Right:    leftVal.(ast.Node),
	}

	right := grl.ComparisonOperations{
		Left:     grl.String{Val: l.field.GetQualifiedName()},
		Operator: l.symbolMapping[ctx.EndIntervalRule.Type],
		Right:    rightVal.(ast.Node),
	}

	logicOp := grl.LogicalOperations{
		Left:     left,
		Operator: grl.LogicalAND,
		Right:    right,
	}

	l.stack.Push(logicOp)
}
