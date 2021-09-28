package ast

type SFeelListenerInterface interface {
	ExitEmptyStatement(ctx EmptyStatement)
	ExitQualifiedName(ctx QualifiedName)
	ExitInteger(ctx Integer)
	ExitFloat(ctx Float)
	ExitBoolean(ctx Boolean)
	ExitString(ctx String)
	ExitDateTime(ctx DateTime)
	ExitInterval(ctx Interval)

	ExitUnaryTest(ctx UnaryTest)
	ExitUnaryTests(ctx UnaryTests)
	ExitSimpleValue(ctx SimpleValue)
	ExitSimpleExpression(ctx SimpleExpression)
	ExitParentheses(ctx Parentheses)
	ExitArithmeticNegation(ctx ArithmeticNegation)
	ExitArithmeticExpression(ctx ArithmeticExpression)
	ExitComparison(ctx Comparison)
}
