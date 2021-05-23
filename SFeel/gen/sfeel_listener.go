// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SFeelListener is a complete listener for a parse tree produced by SFeelParser.
type SFeelListener interface {
	antlr.ParseTreeListener

	// EnterInputEntry is called when entering the inputEntry production.
	EnterInputEntry(c *InputEntryContext)

	// EnterOutputEntry is called when entering the outputEntry production.
	EnterOutputEntry(c *OutputEntryContext)

	// EnterSimpleUnaryTests is called when entering the SimpleUnaryTests production.
	EnterSimpleUnaryTests(c *SimpleUnaryTestsContext)

	// EnterNegationSimpleUnaryTests is called when entering the NegationSimpleUnaryTests production.
	EnterNegationSimpleUnaryTests(c *NegationSimpleUnaryTestsContext)

	// EnterEmptySimpleUnaryTests is called when entering the EmptySimpleUnaryTests production.
	EnterEmptySimpleUnaryTests(c *EmptySimpleUnaryTestsContext)

	// EnterSimple_positive_unary_tests is called when entering the simple_positive_unary_tests production.
	EnterSimple_positive_unary_tests(c *Simple_positive_unary_testsContext)

	// EnterSimple_positive_unary_test is called when entering the simple_positive_unary_test production.
	EnterSimple_positive_unary_test(c *Simple_positive_unary_testContext)

	// EnterUnaryComparison is called when entering the UnaryComparison production.
	EnterUnaryComparison(c *UnaryComparisonContext)

	// EnterEqualUnaryComparison is called when entering the EqualUnaryComparison production.
	EnterEqualUnaryComparison(c *EqualUnaryComparisonContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterOpen_interval_start is called when entering the open_interval_start production.
	EnterOpen_interval_start(c *Open_interval_startContext)

	// EnterClosed_interval_start is called when entering the closed_interval_start production.
	EnterClosed_interval_start(c *Closed_interval_startContext)

	// EnterOpen_interval_end is called when entering the open_interval_end production.
	EnterOpen_interval_end(c *Open_interval_endContext)

	// EnterClosed_interval_end is called when entering the closed_interval_end production.
	EnterClosed_interval_end(c *Closed_interval_endContext)

	// EnterSimple_expressions is called when entering the simple_expressions production.
	EnterSimple_expressions(c *Simple_expressionsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSimple_expression is called when entering the simple_expression production.
	EnterSimple_expression(c *Simple_expressionContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterAdditionOrSubtraction is called when entering the AdditionOrSubtraction production.
	EnterAdditionOrSubtraction(c *AdditionOrSubtractionContext)

	// EnterValue is called when entering the Value production.
	EnterValue(c *ValueContext)

	// EnterMultiplicationOrDivision is called when entering the MultiplicationOrDivision production.
	EnterMultiplicationOrDivision(c *MultiplicationOrDivisionContext)

	// EnterArithmeticNegation is called when entering the ArithmeticNegation production.
	EnterArithmeticNegation(c *ArithmeticNegationContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterPower is called when entering the Power production.
	EnterPower(c *PowerContext)

	// EnterEndpoint is called when entering the endpoint production.
	EnterEndpoint(c *EndpointContext)

	// EnterSimple_value is called when entering the simple_value production.
	EnterSimple_value(c *Simple_valueContext)

	// EnterQualified_name is called when entering the qualified_name production.
	EnterQualified_name(c *Qualified_nameContext)

	// EnterSimple_literal is called when entering the simple_literal production.
	EnterSimple_literal(c *Simple_literalContext)

	// EnterDate_time_literal is called when entering the date_time_literal production.
	EnterDate_time_literal(c *Date_time_literalContext)

	// EnterNumeric_literal is called when entering the numeric_literal production.
	EnterNumeric_literal(c *Numeric_literalContext)

	// EnterInteger_literal is called when entering the integer_literal production.
	EnterInteger_literal(c *Integer_literalContext)

	// EnterReal_literal is called when entering the real_literal production.
	EnterReal_literal(c *Real_literalContext)

	// EnterBoolean_literal is called when entering the boolean_literal production.
	EnterBoolean_literal(c *Boolean_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// ExitInputEntry is called when exiting the inputEntry production.
	ExitInputEntry(c *InputEntryContext)

	// ExitOutputEntry is called when exiting the outputEntry production.
	ExitOutputEntry(c *OutputEntryContext)

	// ExitSimpleUnaryTests is called when exiting the SimpleUnaryTests production.
	ExitSimpleUnaryTests(c *SimpleUnaryTestsContext)

	// ExitNegationSimpleUnaryTests is called when exiting the NegationSimpleUnaryTests production.
	ExitNegationSimpleUnaryTests(c *NegationSimpleUnaryTestsContext)

	// ExitEmptySimpleUnaryTests is called when exiting the EmptySimpleUnaryTests production.
	ExitEmptySimpleUnaryTests(c *EmptySimpleUnaryTestsContext)

	// ExitSimple_positive_unary_tests is called when exiting the simple_positive_unary_tests production.
	ExitSimple_positive_unary_tests(c *Simple_positive_unary_testsContext)

	// ExitSimple_positive_unary_test is called when exiting the simple_positive_unary_test production.
	ExitSimple_positive_unary_test(c *Simple_positive_unary_testContext)

	// ExitUnaryComparison is called when exiting the UnaryComparison production.
	ExitUnaryComparison(c *UnaryComparisonContext)

	// ExitEqualUnaryComparison is called when exiting the EqualUnaryComparison production.
	ExitEqualUnaryComparison(c *EqualUnaryComparisonContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitOpen_interval_start is called when exiting the open_interval_start production.
	ExitOpen_interval_start(c *Open_interval_startContext)

	// ExitClosed_interval_start is called when exiting the closed_interval_start production.
	ExitClosed_interval_start(c *Closed_interval_startContext)

	// ExitOpen_interval_end is called when exiting the open_interval_end production.
	ExitOpen_interval_end(c *Open_interval_endContext)

	// ExitClosed_interval_end is called when exiting the closed_interval_end production.
	ExitClosed_interval_end(c *Closed_interval_endContext)

	// ExitSimple_expressions is called when exiting the simple_expressions production.
	ExitSimple_expressions(c *Simple_expressionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSimple_expression is called when exiting the simple_expression production.
	ExitSimple_expression(c *Simple_expressionContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitAdditionOrSubtraction is called when exiting the AdditionOrSubtraction production.
	ExitAdditionOrSubtraction(c *AdditionOrSubtractionContext)

	// ExitValue is called when exiting the Value production.
	ExitValue(c *ValueContext)

	// ExitMultiplicationOrDivision is called when exiting the MultiplicationOrDivision production.
	ExitMultiplicationOrDivision(c *MultiplicationOrDivisionContext)

	// ExitArithmeticNegation is called when exiting the ArithmeticNegation production.
	ExitArithmeticNegation(c *ArithmeticNegationContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitPower is called when exiting the Power production.
	ExitPower(c *PowerContext)

	// ExitEndpoint is called when exiting the endpoint production.
	ExitEndpoint(c *EndpointContext)

	// ExitSimple_value is called when exiting the simple_value production.
	ExitSimple_value(c *Simple_valueContext)

	// ExitQualified_name is called when exiting the qualified_name production.
	ExitQualified_name(c *Qualified_nameContext)

	// ExitSimple_literal is called when exiting the simple_literal production.
	ExitSimple_literal(c *Simple_literalContext)

	// ExitDate_time_literal is called when exiting the date_time_literal production.
	ExitDate_time_literal(c *Date_time_literalContext)

	// ExitNumeric_literal is called when exiting the numeric_literal production.
	ExitNumeric_literal(c *Numeric_literalContext)

	// ExitInteger_literal is called when exiting the integer_literal production.
	ExitInteger_literal(c *Integer_literalContext)

	// ExitReal_literal is called when exiting the real_literal production.
	ExitReal_literal(c *Real_literalContext)

	// ExitBoolean_literal is called when exiting the boolean_literal production.
	ExitBoolean_literal(c *Boolean_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)
}
