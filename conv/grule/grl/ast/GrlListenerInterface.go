package ast

type GrlListenerInterface interface {
	ExitInteger(ctx Integer)
	ExitString(ctx String)
	//ExitFloat(ctx Float)
	//ExitBoolean(ctx Boolean)
	//ExitDateTime(ctx DateTime)

	ExitMathOperations(ctx MathOperations)
	ExitLogicalOperations(ctx LogicalOperations)
	ExitComparisonOperations(ctx ComparisonOperations)
}
