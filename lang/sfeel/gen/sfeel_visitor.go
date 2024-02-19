// Code generated from SFeel.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SFeelParser.
type SFeelVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SFeelParser#input.
	VisitInput(ctx *InputContext) interface{}

	// Visit a parse tree produced by SFeelParser#output.
	VisitOutput(ctx *OutputContext) interface{}

	// Visit a parse tree produced by SFeelParser#SimpleUnaryTests.
	VisitSimpleUnaryTests(ctx *SimpleUnaryTestsContext) interface{}

	// Visit a parse tree produced by SFeelParser#NegationSimpleUnaryTests.
	VisitNegationSimpleUnaryTests(ctx *NegationSimpleUnaryTestsContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptySimpleUnaryTests.
	VisitEmptySimpleUnaryTests(ctx *EmptySimpleUnaryTestsContext) interface{}

	// Visit a parse tree produced by SFeelParser#simple_positive_unary_tests.
	VisitSimple_positive_unary_tests(ctx *Simple_positive_unary_testsContext) interface{}

	// Visit a parse tree produced by SFeelParser#simple_positive_unary_test.
	VisitSimple_positive_unary_test(ctx *Simple_positive_unary_testContext) interface{}

	// Visit a parse tree produced by SFeelParser#UnaryComparison.
	VisitUnaryComparison(ctx *UnaryComparisonContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualUnaryComparison.
	VisitEqualUnaryComparison(ctx *EqualUnaryComparisonContext) interface{}

	// Visit a parse tree produced by SFeelParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by SFeelParser#open_interval_start.
	VisitOpen_interval_start(ctx *Open_interval_startContext) interface{}

	// Visit a parse tree produced by SFeelParser#closed_interval_start.
	VisitClosed_interval_start(ctx *Closed_interval_startContext) interface{}

	// Visit a parse tree produced by SFeelParser#open_interval_end.
	VisitOpen_interval_end(ctx *Open_interval_endContext) interface{}

	// Visit a parse tree produced by SFeelParser#closed_interval_end.
	VisitClosed_interval_end(ctx *Closed_interval_endContext) interface{}

	// Visit a parse tree produced by SFeelParser#empty_expression.
	VisitEmpty_expression(ctx *Empty_expressionContext) interface{}

	// Visit a parse tree produced by SFeelParser#SimpleExpressions.
	VisitSimpleExpressions(ctx *SimpleExpressionsContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptySimpleExpressions.
	VisitEmptySimpleExpressions(ctx *EmptySimpleExpressionsContext) interface{}

	// Visit a parse tree produced by SFeelParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SFeelParser#simple_expression.
	VisitSimple_expression(ctx *Simple_expressionContext) interface{}

	// Visit a parse tree produced by SFeelParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by SFeelParser#AdditionOrSubtraction.
	VisitAdditionOrSubtraction(ctx *AdditionOrSubtractionContext) interface{}

	// Visit a parse tree produced by SFeelParser#Value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by SFeelParser#MultiplicationOrDivision.
	VisitMultiplicationOrDivision(ctx *MultiplicationOrDivisionContext) interface{}

	// Visit a parse tree produced by SFeelParser#ArithmeticNegation.
	VisitArithmeticNegation(ctx *ArithmeticNegationContext) interface{}

	// Visit a parse tree produced by SFeelParser#Parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by SFeelParser#Power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by SFeelParser#endpoint.
	VisitEndpoint(ctx *EndpointContext) interface{}

	// Visit a parse tree produced by SFeelParser#simple_value.
	VisitSimple_value(ctx *Simple_valueContext) interface{}

	// Visit a parse tree produced by SFeelParser#qualified_name.
	VisitQualified_name(ctx *Qualified_nameContext) interface{}

	// Visit a parse tree produced by SFeelParser#simple_literal.
	VisitSimple_literal(ctx *Simple_literalContext) interface{}

	// Visit a parse tree produced by SFeelParser#date_time_literal.
	VisitDate_time_literal(ctx *Date_time_literalContext) interface{}

	// Visit a parse tree produced by SFeelParser#numeric_literal.
	VisitNumeric_literal(ctx *Numeric_literalContext) interface{}

	// Visit a parse tree produced by SFeelParser#integer_literal.
	VisitInteger_literal(ctx *Integer_literalContext) interface{}

	// Visit a parse tree produced by SFeelParser#real_literal.
	VisitReal_literal(ctx *Real_literalContext) interface{}

	// Visit a parse tree produced by SFeelParser#boolean_literal.
	VisitBoolean_literal(ctx *Boolean_literalContext) interface{}

	// Visit a parse tree produced by SFeelParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}
}
