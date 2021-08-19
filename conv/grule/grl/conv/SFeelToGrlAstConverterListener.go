package conv

import (
	"github.com/global-soft-ba/decisionTable/ast"
	grl "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	"github.com/global-soft-ba/decisionTable/data"
	sfeel "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
)

var OperatorMappingTable = map[int]int{
	sfeel.SFeelAssignmentEqual:     grl.AssignmentEQUAL, //-1
	sfeel.SFeelOperatorEQUAL:       grl.ComparisonOperatorEQUAL,
	sfeel.SFeelOperatorNOTEQUAL:    grl.ComparisonOperatorNOTEQUAL,
	sfeel.SFeelOperatorLESS:        grl.ComparisonOperatorLESS,
	sfeel.SFeelOperatorLESSEQ:      grl.ComparisonOperatorLESSEQ,
	sfeel.SFeelOperatorGREATER:     grl.ComparisonOperatorGREATER,
	sfeel.SFeelOperatorGREATEREQ:   grl.ComparisonOperatorGREATEREQ,
	sfeel.SFeelClosedIntervalSTART: grl.ComparisonOperatorGREATEREQ, // [1..   [e1..e2] if and only if o ≥ e1 and o ≤ e2
	sfeel.SFeelClosedIntervalEND:   grl.ComparisonOperatorLESSEQ,    // ..1]
	sfeel.SFeelOpenIntervalSTART:   grl.ComparisonOperatorGREATER,   // ]1..
	sfeel.SFeelOpenIntervalEND:     grl.ComparisonOperatorLESS,      // ..1[
	sfeel.SFeelOperatorADD:         grl.MathADD,
	sfeel.SFeelOperatorSUB:         grl.MathSUB,
	sfeel.SFeelOperatorMUL:         grl.MathMUL,
	sfeel.SFeelOperatorDIV:         grl.MathDIV,
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

func (l *SFeelToGrlAstConverterListener) ExitUnaryTest(ctx sfeel.UnaryTest) {
	rightVal := l.stack.Pop()
	comp := grl.ComparisonOperations{
		Left:     grl.String{Val: l.field.GetQualifiedName()},
		Operator: l.symbolMapping[ctx.Operator.Type],
		Right:    rightVal.(ast.Node),
	}

	l.stack.Push(comp)
}

func (l *SFeelToGrlAstConverterListener) ExitSimpleExpression(ctx sfeel.SimpleExpression) {
	rightVal := l.stack.Pop()
	comp := grl.AssignmentOperations{
		Left:     grl.String{Val: l.field.GetQualifiedName()},
		Operator: l.symbolMapping[ctx.Operator.Type],
		Right:    rightVal.(ast.Node),
	}

	l.stack.Push(comp)
}

func (l *SFeelToGrlAstConverterListener) ExitParentheses(ctx sfeel.Parentheses) {
	val := l.stack.Pop()
	comp := grl.Parentheses{
		Value: val.(ast.Node),
	}

	l.stack.Push(comp)
}

func (l *SFeelToGrlAstConverterListener) ExitArithmeticNegation(ctx sfeel.ArithmeticNegation) {
	val := l.stack.Pop()
	an := grl.ArithmeticNegation{
		Value: val.(ast.Node),
	}
	l.stack.Push(an)
}

func (l *SFeelToGrlAstConverterListener) ExitArithmeticExpression(ctx sfeel.ArithmeticExpression) {
	right := l.stack.Pop()
	left := l.stack.Pop()
	if ctx.Operator.Type == sfeel.SFeelOperatorPOW {
		pow := grl.PowOperation{
			Base:     left.(ast.Node),
			Exponent: right.(ast.Node),
		}
		l.stack.Push(pow)
		return
	}
	ae := grl.MathOperations{
		Left:     left.(ast.Node),
		Operator: l.symbolMapping[ctx.Operator.Type],
		Right:    right.(ast.Node),
	}
	l.stack.Push(ae)
}

func (l *SFeelToGrlAstConverterListener) ExitComparison(ctx sfeel.Comparison) {
	//TODO: Fix Comparison
	right := l.stack.Pop()
	left := l.stack.Pop()

	mo := grl.ComparisonOperations{
		Left:     left.(ast.Node),
		Operator: l.symbolMapping[ctx.Operator.Type],
		Right:    right.(ast.Node),
	}
	l.stack.Push(mo)
}

func (l *SFeelToGrlAstConverterListener) ExitInteger(ctx sfeel.Integer) {
	i := grl.Integer{Val: ctx.Value}
	sign := ctx.SignRule.Literal
	if sign != "" {
		i.Sign = &sign
	}
	l.stack.Push(i)
}

func (l *SFeelToGrlAstConverterListener) ExitQualifiedName(ctx sfeel.QualifiedName) {
	q := grl.QualifiedName{Val: ctx.Value}
	l.stack.Push(q)
}

func (l *SFeelToGrlAstConverterListener) ExitString(ctx sfeel.String) {
	q := grl.String{Val: ctx.Value}
	l.stack.Push(q)
}

func (l *SFeelToGrlAstConverterListener) ExitBoolean(ctx sfeel.Boolean) {
	q := grl.Boolean{Val: ctx.Value}
	l.stack.Push(q)
}

func (l *SFeelToGrlAstConverterListener) ExitEmptyStatement(ctx sfeel.EmptyStatement) {
	q := grl.EmptyStatement{}
	l.stack.Push(q)
}

func (l *SFeelToGrlAstConverterListener) ExitFloat(ctx sfeel.Float) {
	q := grl.Float{Val: ctx.Value}
	l.stack.Push(q)
}

func (l *SFeelToGrlAstConverterListener) ExitDateTime(ctx sfeel.DateTime) {
	//TODO: Finish DateTime
	q := grl.DateTime{Val: ctx.Value}
	l.stack.Push(q)
}
