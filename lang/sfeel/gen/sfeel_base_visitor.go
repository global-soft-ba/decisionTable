// Code generated from SFeel.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSFeelVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSFeelVisitor) VisitInput(ctx *InputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitOutput(ctx *OutputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimpleUnaryTests(ctx *SimpleUnaryTestsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitNegationSimpleUnaryTests(ctx *NegationSimpleUnaryTestsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEmptySimpleUnaryTests(ctx *EmptySimpleUnaryTestsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimple_positive_unary_tests(ctx *Simple_positive_unary_testsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimple_positive_unary_test(ctx *Simple_positive_unary_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitUnaryComparison(ctx *UnaryComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualUnaryComparison(ctx *EqualUnaryComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitOpen_interval_start(ctx *Open_interval_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitClosed_interval_start(ctx *Closed_interval_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitOpen_interval_end(ctx *Open_interval_endContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitClosed_interval_end(ctx *Closed_interval_endContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEmpty_expression(ctx *Empty_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimpleExpressions(ctx *SimpleExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEmptySimpleExpressions(ctx *EmptySimpleExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimple_expression(ctx *Simple_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitAdditionOrSubtraction(ctx *AdditionOrSubtractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitMultiplicationOrDivision(ctx *MultiplicationOrDivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitArithmeticNegation(ctx *ArithmeticNegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEndpoint(ctx *EndpointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimple_value(ctx *Simple_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitQualified_name(ctx *Qualified_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitSimple_literal(ctx *Simple_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDate_time_literal(ctx *Date_time_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitNumeric_literal(ctx *Numeric_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitInteger_literal(ctx *Integer_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitReal_literal(ctx *Real_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitBoolean_literal(ctx *Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}
