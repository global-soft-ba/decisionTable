// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SFeelParser.
type SFeelVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SFeelParser#entry.
	VisitEntry(ctx *EntryContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualComparisonIntegerInputRule.
	VisitEqualComparisonIntegerInputRule(ctx *EqualComparisonIntegerInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#ComparisonIntegerInputRule.
	VisitComparisonIntegerInputRule(ctx *ComparisonIntegerInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#RangeComparisonIntegerInputRule.
	VisitRangeComparisonIntegerInputRule(ctx *RangeComparisonIntegerInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#DisjunctionsIntegerInputRule.
	VisitDisjunctionsIntegerInputRule(ctx *DisjunctionsIntegerInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#NegationIntegerInputRule.
	VisitNegationIntegerInputRule(ctx *NegationIntegerInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyIntegerInputRule.
	VisitEmptyIntegerInputRule(ctx *EmptyIntegerInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualComparisonNumberInputRule.
	VisitEqualComparisonNumberInputRule(ctx *EqualComparisonNumberInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#ComparisonNumberInputRule.
	VisitComparisonNumberInputRule(ctx *ComparisonNumberInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#RangeComparisonNumberInputRule.
	VisitRangeComparisonNumberInputRule(ctx *RangeComparisonNumberInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#DisjunctionsNumberInputRule.
	VisitDisjunctionsNumberInputRule(ctx *DisjunctionsNumberInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#NegationNumberInputRule.
	VisitNegationNumberInputRule(ctx *NegationNumberInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyNumberInputRule.
	VisitEmptyNumberInputRule(ctx *EmptyNumberInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualComparisonStringInputRule.
	VisitEqualComparisonStringInputRule(ctx *EqualComparisonStringInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#ComparisonStringInputRule.
	VisitComparisonStringInputRule(ctx *ComparisonStringInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#NegationStringInputRule.
	VisitNegationStringInputRule(ctx *NegationStringInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyStringInputRule.
	VisitEmptyStringInputRule(ctx *EmptyStringInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualComparisonBoolInputRule.
	VisitEqualComparisonBoolInputRule(ctx *EqualComparisonBoolInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyBoolInputRule.
	VisitEmptyBoolInputRule(ctx *EmptyBoolInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EqualComparisonDateTimeInputRule.
	VisitEqualComparisonDateTimeInputRule(ctx *EqualComparisonDateTimeInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#ComparisonDateTimeInputRule.
	VisitComparisonDateTimeInputRule(ctx *ComparisonDateTimeInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#RangeComparisonDateTimeInputRule.
	VisitRangeComparisonDateTimeInputRule(ctx *RangeComparisonDateTimeInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#DisjunctionsDateTimeInputRule.
	VisitDisjunctionsDateTimeInputRule(ctx *DisjunctionsDateTimeInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#NegationDateTimeInputRule.
	VisitNegationDateTimeInputRule(ctx *NegationDateTimeInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyDateTimeInputRule.
	VisitEmptyDateTimeInputRule(ctx *EmptyDateTimeInputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#IntegerOutputRule.
	VisitIntegerOutputRule(ctx *IntegerOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyIntegerOutputRule.
	VisitEmptyIntegerOutputRule(ctx *EmptyIntegerOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#NumberOutputRule.
	VisitNumberOutputRule(ctx *NumberOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyNumberOutputRule.
	VisitEmptyNumberOutputRule(ctx *EmptyNumberOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#StringOutputRule.
	VisitStringOutputRule(ctx *StringOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyStringOutputRule.
	VisitEmptyStringOutputRule(ctx *EmptyStringOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#BoolOutputRule.
	VisitBoolOutputRule(ctx *BoolOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyBoolOutputRule.
	VisitEmptyBoolOutputRule(ctx *EmptyBoolOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#DateTimeOutputRule.
	VisitDateTimeOutputRule(ctx *DateTimeOutputRuleContext) interface{}

	// Visit a parse tree produced by SFeelParser#EmptyDateTimeOutputRule.
	VisitEmptyDateTimeOutputRule(ctx *EmptyDateTimeOutputRuleContext) interface{}

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
