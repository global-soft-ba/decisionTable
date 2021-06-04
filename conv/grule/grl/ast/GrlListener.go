package ast

type GrlListener struct{}

func (g GrlListener) ExitInteger(ctx Integer) {}
func (g GrlListener) ExitString(ctx String)   {}

//func (g GrlListener) ExitFloat(ctx Float){}
//func (g GrlListener) ExitBoolean(ctx Boolean){}
//func (g GrlListener) ExitDateTime(ctx DateTime){}

func (g GrlListener) ExitMathOperations(ctx MathOperations)             {}
func (g GrlListener) ExitLogicalOperations(ctx LogicalOperations)       {}
func (g GrlListener) ExitComparisonOperations(ctx ComparisonOperations) {}
