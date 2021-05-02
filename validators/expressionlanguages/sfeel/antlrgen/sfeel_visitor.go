// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SFeelParser.
type SFeelVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SFeelParser#entry.
	VisitEntry(ctx *EntryContext) interface{}

	// Visit a parse tree produced by SFeelParser#validIntegerInput.
	VisitValidIntegerInput(ctx *ValidIntegerInputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validNumberInput.
	VisitValidNumberInput(ctx *ValidNumberInputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validStringInput.
	VisitValidStringInput(ctx *ValidStringInputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validBoolInput.
	VisitValidBoolInput(ctx *ValidBoolInputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validDateTimeInput.
	VisitValidDateTimeInput(ctx *ValidDateTimeInputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validIntegerOutput.
	VisitValidIntegerOutput(ctx *ValidIntegerOutputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validNumberOutput.
	VisitValidNumberOutput(ctx *ValidNumberOutputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validStringOutput.
	VisitValidStringOutput(ctx *ValidStringOutputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validBoolOutput.
	VisitValidBoolOutput(ctx *ValidBoolOutputContext) interface{}

	// Visit a parse tree produced by SFeelParser#validDateTimeOutput.
	VisitValidDateTimeOutput(ctx *ValidDateTimeOutputContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualcomparisonRule.
	VisitEqualcomparisonRule(ctx *EqualcomparisonRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#ComparisionsRule.
	VisitComparisionsRule(ctx *ComparisionsRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#RangeRule.
	VisitRangeRule(ctx *RangeRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#DisjunctionRule.
	VisitDisjunctionRule(ctx *DisjunctionRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#NegationRule.
	VisitNegationRule(ctx *NegationRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyInputRule.
	VisitEmptyInputRule(ctx *EmptyInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by SFeelParser#strings.
	VisitStrings(ctx *StringsContext) interface{}

	// Visit a parse tree produced by SFeelParser#bools.
	VisitBools(ctx *BoolsContext) interface{}

	// Visit a parse tree produced by SFeelParser#datetime.
	VisitDatetime(ctx *DatetimeContext) interface{}

	// Visit a parse tree produced by SFeelParser#equalcomparison.
	VisitEqualcomparison(ctx *EqualcomparisonContext) interface{}

	// Visit a parse tree produced by SFeelParser#equalcomparisonInteger.
	VisitEqualcomparisonInteger(ctx *EqualcomparisonIntegerContext) interface{}

	// Visit a parse tree produced by SFeelParser#equalcomparisonNumber.
	VisitEqualcomparisonNumber(ctx *EqualcomparisonNumberContext) interface{}

	// Visit a parse tree produced by SFeelParser#equalcomparisonBool.
	VisitEqualcomparisonBool(ctx *EqualcomparisonBoolContext) interface{}

	// Visit a parse tree produced by SFeelParser#equalcomparisonStrings.
	VisitEqualcomparisonStrings(ctx *EqualcomparisonStringsContext) interface{}

	// Visit a parse tree produced by SFeelParser#equalcomparisonDateTime.
	VisitEqualcomparisonDateTime(ctx *EqualcomparisonDateTimeContext) interface{}

	// Visit a parse tree produced by SFeelParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by SFeelParser#comparisonOps.
	VisitComparisonOps(ctx *ComparisonOpsContext) interface{}

	// Visit a parse tree produced by SFeelParser#comparisonInteger.
	VisitComparisonInteger(ctx *ComparisonIntegerContext) interface{}

	// Visit a parse tree produced by SFeelParser#comparisonDateTime.
	VisitComparisonDateTime(ctx *ComparisonDateTimeContext) interface{}

	// Visit a parse tree produced by SFeelParser#comparisonNumber.
	VisitComparisonNumber(ctx *ComparisonNumberContext) interface{}

	// Visit a parse tree produced by SFeelParser#ranges.
	VisitRanges(ctx *RangesContext) interface{}

	// Visit a parse tree produced by SFeelParser#rop.
	VisitRop(ctx *RopContext) interface{}

	// Visit a parse tree produced by SFeelParser#rangeInteger.
	VisitRangeInteger(ctx *RangeIntegerContext) interface{}

	// Visit a parse tree produced by SFeelParser#rangeNumber.
	VisitRangeNumber(ctx *RangeNumberContext) interface{}

	// Visit a parse tree produced by SFeelParser#rangeDateTime.
	VisitRangeDateTime(ctx *RangeDateTimeContext) interface{}

	// Visit a parse tree produced by SFeelParser#disjunctions.
	VisitDisjunctions(ctx *DisjunctionsContext) interface{}

	// Visit a parse tree produced by SFeelParser#disjunctionsInteger.
	VisitDisjunctionsInteger(ctx *DisjunctionsIntegerContext) interface{}

	// Visit a parse tree produced by SFeelParser#disjunctionsNumber.
	VisitDisjunctionsNumber(ctx *DisjunctionsNumberContext) interface{}

	// Visit a parse tree produced by SFeelParser#disjunctionsString.
	VisitDisjunctionsString(ctx *DisjunctionsStringContext) interface{}

	// Visit a parse tree produced by SFeelParser#disjunctionsDateTime.
	VisitDisjunctionsDateTime(ctx *DisjunctionsDateTimeContext) interface{}

	// Visit a parse tree produced by SFeelParser#negation.
	VisitNegation(ctx *NegationContext) interface{}
}
