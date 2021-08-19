package ast

type GrlListenerInterface interface {
	ExitInteger(ctx Integer)
	ExitString(ctx String)
	ExitFloat(ctx Float)
	ExitBoolean(ctx Boolean)
	ExitDateTime(ctx DateTime)

	ExitParentheses(ctx Parentheses)
	ExitArithmeticNegation(ctx ArithmeticNegation)
	ExitMathOperations(ctx MathOperations)
	ExitLogicalOperations(ctx LogicalOperations)
	ExitComparisonOperations(ctx ComparisonOperations)
	ExitEmptyStatement(ctx EmptyStatement)
	ExitAssignmentOperations(ctx AssignmentOperations)
	ExitPowOperation(ctx PowOperation)
}
