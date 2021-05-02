// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSFeelVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSFeelVisitor) VisitEntry(ctx *EntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidIntegerInput(ctx *ValidIntegerInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidNumberInput(ctx *ValidNumberInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidStringInput(ctx *ValidStringInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidBoolInput(ctx *ValidBoolInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidDateTimeInput(ctx *ValidDateTimeInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidIntegerOutput(ctx *ValidIntegerOutputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidNumberOutput(ctx *ValidNumberOutputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidStringOutput(ctx *ValidStringOutputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidBoolOutput(ctx *ValidBoolOutputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitValidDateTimeOutput(ctx *ValidDateTimeOutputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparisonRule(ctx *EqualcomparisonRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparisionsRule(ctx *ComparisionsRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitRangeRule(ctx *RangeRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDisjunctionRule(ctx *DisjunctionRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitNegationRule(ctx *NegationRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEmptyInputRule(ctx *EmptyInputRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitStrings(ctx *StringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitBools(ctx *BoolsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDatetime(ctx *DatetimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparison(ctx *EqualcomparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparisonInteger(ctx *EqualcomparisonIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparisonNumber(ctx *EqualcomparisonNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparisonBool(ctx *EqualcomparisonBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparisonStrings(ctx *EqualcomparisonStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitEqualcomparisonDateTime(ctx *EqualcomparisonDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparisonOps(ctx *ComparisonOpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparisonInteger(ctx *ComparisonIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparisonDateTime(ctx *ComparisonDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitComparisonNumber(ctx *ComparisonNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitRanges(ctx *RangesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitRop(ctx *RopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitRangeInteger(ctx *RangeIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitRangeNumber(ctx *RangeNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitRangeDateTime(ctx *RangeDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDisjunctions(ctx *DisjunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDisjunctionsInteger(ctx *DisjunctionsIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDisjunctionsNumber(ctx *DisjunctionsNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDisjunctionsString(ctx *DisjunctionsStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitDisjunctionsDateTime(ctx *DisjunctionsDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFeelVisitor) VisitNegation(ctx *NegationContext) interface{} {
	return v.VisitChildren(ctx)
}
