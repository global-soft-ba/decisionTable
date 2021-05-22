// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSFeelListener is a complete listener for a parse tree produced by SFeelParser.
type BaseSFeelListener struct{}

var _ SFeelListener = &BaseSFeelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSFeelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSFeelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSFeelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSFeelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSFeelListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSFeelListener) ExitStart(ctx *StartContext) {}

// EnterSimple_expressions is called when production simple_expressions is entered.
func (s *BaseSFeelListener) EnterSimple_expressions(ctx *Simple_expressionsContext) {}

// ExitSimple_expressions is called when production simple_expressions is exited.
func (s *BaseSFeelListener) ExitSimple_expressions(ctx *Simple_expressionsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSFeelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSFeelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSimple_expression is called when production simple_expression is entered.
func (s *BaseSFeelListener) EnterSimple_expression(ctx *Simple_expressionContext) {}

// ExitSimple_expression is called when production simple_expression is exited.
func (s *BaseSFeelListener) ExitSimple_expression(ctx *Simple_expressionContext) {}

// EnterSimpleUnaryTests is called when production SimpleUnaryTests is entered.
func (s *BaseSFeelListener) EnterSimpleUnaryTests(ctx *SimpleUnaryTestsContext) {}

// ExitSimpleUnaryTests is called when production SimpleUnaryTests is exited.
func (s *BaseSFeelListener) ExitSimpleUnaryTests(ctx *SimpleUnaryTestsContext) {}

// EnterNegationSimpleUnaryTests is called when production NegationSimpleUnaryTests is entered.
func (s *BaseSFeelListener) EnterNegationSimpleUnaryTests(ctx *NegationSimpleUnaryTestsContext) {}

// ExitNegationSimpleUnaryTests is called when production NegationSimpleUnaryTests is exited.
func (s *BaseSFeelListener) ExitNegationSimpleUnaryTests(ctx *NegationSimpleUnaryTestsContext) {}

// EnterEmptySimpleUnaryTests is called when production EmptySimpleUnaryTests is entered.
func (s *BaseSFeelListener) EnterEmptySimpleUnaryTests(ctx *EmptySimpleUnaryTestsContext) {}

// ExitEmptySimpleUnaryTests is called when production EmptySimpleUnaryTests is exited.
func (s *BaseSFeelListener) ExitEmptySimpleUnaryTests(ctx *EmptySimpleUnaryTestsContext) {}

// EnterSimple_positive_unary_tests is called when production simple_positive_unary_tests is entered.
func (s *BaseSFeelListener) EnterSimple_positive_unary_tests(ctx *Simple_positive_unary_testsContext) {
}

// ExitSimple_positive_unary_tests is called when production simple_positive_unary_tests is exited.
func (s *BaseSFeelListener) ExitSimple_positive_unary_tests(ctx *Simple_positive_unary_testsContext) {
}

// EnterSimple_positive_unary_test is called when production simple_positive_unary_test is entered.
func (s *BaseSFeelListener) EnterSimple_positive_unary_test(ctx *Simple_positive_unary_testContext) {}

// ExitSimple_positive_unary_test is called when production simple_positive_unary_test is exited.
func (s *BaseSFeelListener) ExitSimple_positive_unary_test(ctx *Simple_positive_unary_testContext) {}

// EnterUnaryComparison is called when production UnaryComparison is entered.
func (s *BaseSFeelListener) EnterUnaryComparison(ctx *UnaryComparisonContext) {}

// ExitUnaryComparison is called when production UnaryComparison is exited.
func (s *BaseSFeelListener) ExitUnaryComparison(ctx *UnaryComparisonContext) {}

// EnterEqualUnaryComparison is called when production EqualUnaryComparison is entered.
func (s *BaseSFeelListener) EnterEqualUnaryComparison(ctx *EqualUnaryComparisonContext) {}

// ExitEqualUnaryComparison is called when production EqualUnaryComparison is exited.
func (s *BaseSFeelListener) ExitEqualUnaryComparison(ctx *EqualUnaryComparisonContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSFeelListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSFeelListener) ExitComparison(ctx *ComparisonContext) {}

// EnterAdditionOrSubtraction is called when production AdditionOrSubtraction is entered.
func (s *BaseSFeelListener) EnterAdditionOrSubtraction(ctx *AdditionOrSubtractionContext) {}

// ExitAdditionOrSubtraction is called when production AdditionOrSubtraction is exited.
func (s *BaseSFeelListener) ExitAdditionOrSubtraction(ctx *AdditionOrSubtractionContext) {}

// EnterValue is called when production Value is entered.
func (s *BaseSFeelListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production Value is exited.
func (s *BaseSFeelListener) ExitValue(ctx *ValueContext) {}

// EnterMultiplicationOrDivision is called when production MultiplicationOrDivision is entered.
func (s *BaseSFeelListener) EnterMultiplicationOrDivision(ctx *MultiplicationOrDivisionContext) {}

// ExitMultiplicationOrDivision is called when production MultiplicationOrDivision is exited.
func (s *BaseSFeelListener) ExitMultiplicationOrDivision(ctx *MultiplicationOrDivisionContext) {}

// EnterArithmeticNegation is called when production ArithmeticNegation is entered.
func (s *BaseSFeelListener) EnterArithmeticNegation(ctx *ArithmeticNegationContext) {}

// ExitArithmeticNegation is called when production ArithmeticNegation is exited.
func (s *BaseSFeelListener) ExitArithmeticNegation(ctx *ArithmeticNegationContext) {}

// EnterParentheses is called when production Parentheses is entered.
func (s *BaseSFeelListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production Parentheses is exited.
func (s *BaseSFeelListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterPower is called when production Power is entered.
func (s *BaseSFeelListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production Power is exited.
func (s *BaseSFeelListener) ExitPower(ctx *PowerContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseSFeelListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseSFeelListener) ExitInterval(ctx *IntervalContext) {}

// EnterOpen_interval_start is called when production open_interval_start is entered.
func (s *BaseSFeelListener) EnterOpen_interval_start(ctx *Open_interval_startContext) {}

// ExitOpen_interval_start is called when production open_interval_start is exited.
func (s *BaseSFeelListener) ExitOpen_interval_start(ctx *Open_interval_startContext) {}

// EnterClosed_interval_start is called when production closed_interval_start is entered.
func (s *BaseSFeelListener) EnterClosed_interval_start(ctx *Closed_interval_startContext) {}

// ExitClosed_interval_start is called when production closed_interval_start is exited.
func (s *BaseSFeelListener) ExitClosed_interval_start(ctx *Closed_interval_startContext) {}

// EnterOpen_interval_end is called when production open_interval_end is entered.
func (s *BaseSFeelListener) EnterOpen_interval_end(ctx *Open_interval_endContext) {}

// ExitOpen_interval_end is called when production open_interval_end is exited.
func (s *BaseSFeelListener) ExitOpen_interval_end(ctx *Open_interval_endContext) {}

// EnterClosed_interval_end is called when production closed_interval_end is entered.
func (s *BaseSFeelListener) EnterClosed_interval_end(ctx *Closed_interval_endContext) {}

// ExitClosed_interval_end is called when production closed_interval_end is exited.
func (s *BaseSFeelListener) ExitClosed_interval_end(ctx *Closed_interval_endContext) {}

// EnterEndpoint is called when production endpoint is entered.
func (s *BaseSFeelListener) EnterEndpoint(ctx *EndpointContext) {}

// ExitEndpoint is called when production endpoint is exited.
func (s *BaseSFeelListener) ExitEndpoint(ctx *EndpointContext) {}

// EnterSimple_value is called when production simple_value is entered.
func (s *BaseSFeelListener) EnterSimple_value(ctx *Simple_valueContext) {}

// ExitSimple_value is called when production simple_value is exited.
func (s *BaseSFeelListener) ExitSimple_value(ctx *Simple_valueContext) {}

// EnterQualified_name is called when production qualified_name is entered.
func (s *BaseSFeelListener) EnterQualified_name(ctx *Qualified_nameContext) {}

// ExitQualified_name is called when production qualified_name is exited.
func (s *BaseSFeelListener) ExitQualified_name(ctx *Qualified_nameContext) {}

// EnterSimple_literal is called when production simple_literal is entered.
func (s *BaseSFeelListener) EnterSimple_literal(ctx *Simple_literalContext) {}

// ExitSimple_literal is called when production simple_literal is exited.
func (s *BaseSFeelListener) ExitSimple_literal(ctx *Simple_literalContext) {}

// EnterDate_time_literal is called when production date_time_literal is entered.
func (s *BaseSFeelListener) EnterDate_time_literal(ctx *Date_time_literalContext) {}

// ExitDate_time_literal is called when production date_time_literal is exited.
func (s *BaseSFeelListener) ExitDate_time_literal(ctx *Date_time_literalContext) {}

// EnterNumeric_literal is called when production numeric_literal is entered.
func (s *BaseSFeelListener) EnterNumeric_literal(ctx *Numeric_literalContext) {}

// ExitNumeric_literal is called when production numeric_literal is exited.
func (s *BaseSFeelListener) ExitNumeric_literal(ctx *Numeric_literalContext) {}

// EnterInteger_literal is called when production integer_literal is entered.
func (s *BaseSFeelListener) EnterInteger_literal(ctx *Integer_literalContext) {}

// ExitInteger_literal is called when production integer_literal is exited.
func (s *BaseSFeelListener) ExitInteger_literal(ctx *Integer_literalContext) {}

// EnterReal_literal is called when production real_literal is entered.
func (s *BaseSFeelListener) EnterReal_literal(ctx *Real_literalContext) {}

// ExitReal_literal is called when production real_literal is exited.
func (s *BaseSFeelListener) ExitReal_literal(ctx *Real_literalContext) {}

// EnterBoolean_literal is called when production boolean_literal is entered.
func (s *BaseSFeelListener) EnterBoolean_literal(ctx *Boolean_literalContext) {}

// ExitBoolean_literal is called when production boolean_literal is exited.
func (s *BaseSFeelListener) ExitBoolean_literal(ctx *Boolean_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseSFeelListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseSFeelListener) ExitString_literal(ctx *String_literalContext) {}
