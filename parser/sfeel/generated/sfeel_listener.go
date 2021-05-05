// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SFeelListener is a complete listener for a parse tree produced by SFeelParser.
type SFeelListener interface {
	antlr.ParseTreeListener

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterEqualComparisonIntegerInputRule is called when entering the EqualComparisonIntegerInputRule production.
	EnterEqualComparisonIntegerInputRule(c *EqualComparisonIntegerInputRuleContext)

	// EnterComparisonIntegerInputRule is called when entering the ComparisonIntegerInputRule production.
	EnterComparisonIntegerInputRule(c *ComparisonIntegerInputRuleContext)

	// EnterRangeComparisonIntegerInputRule is called when entering the RangeComparisonIntegerInputRule production.
	EnterRangeComparisonIntegerInputRule(c *RangeComparisonIntegerInputRuleContext)

	// EnterDisjunctionsIntegerInputRule is called when entering the DisjunctionsIntegerInputRule production.
	EnterDisjunctionsIntegerInputRule(c *DisjunctionsIntegerInputRuleContext)

	// EnterNegationIntegerInputRule is called when entering the NegationIntegerInputRule production.
	EnterNegationIntegerInputRule(c *NegationIntegerInputRuleContext)

	// EnterEmptyIntegerInputRule is called when entering the EmptyIntegerInputRule production.
	EnterEmptyIntegerInputRule(c *EmptyIntegerInputRuleContext)

	// EnterEqualComparisonNumberInputRule is called when entering the EqualComparisonNumberInputRule production.
	EnterEqualComparisonNumberInputRule(c *EqualComparisonNumberInputRuleContext)

	// EnterComparisonNumberInputRule is called when entering the ComparisonNumberInputRule production.
	EnterComparisonNumberInputRule(c *ComparisonNumberInputRuleContext)

	// EnterRangeComparisonNumberInputRule is called when entering the RangeComparisonNumberInputRule production.
	EnterRangeComparisonNumberInputRule(c *RangeComparisonNumberInputRuleContext)

	// EnterDisjunctionsNumberInputRule is called when entering the DisjunctionsNumberInputRule production.
	EnterDisjunctionsNumberInputRule(c *DisjunctionsNumberInputRuleContext)

	// EnterNegationNumberInputRule is called when entering the NegationNumberInputRule production.
	EnterNegationNumberInputRule(c *NegationNumberInputRuleContext)

	// EnterEmptyNumberInputRule is called when entering the EmptyNumberInputRule production.
	EnterEmptyNumberInputRule(c *EmptyNumberInputRuleContext)

	// EnterEqualComparisonStringInputRule is called when entering the EqualComparisonStringInputRule production.
	EnterEqualComparisonStringInputRule(c *EqualComparisonStringInputRuleContext)

	// EnterDisjunctionsStringInputRule is called when entering the DisjunctionsStringInputRule production.
	EnterDisjunctionsStringInputRule(c *DisjunctionsStringInputRuleContext)

	// EnterNegationStringInputRule is called when entering the NegationStringInputRule production.
	EnterNegationStringInputRule(c *NegationStringInputRuleContext)

	// EnterEmptyStringInputRule is called when entering the EmptyStringInputRule production.
	EnterEmptyStringInputRule(c *EmptyStringInputRuleContext)

	// EnterEqualComparisonBoolInputRule is called when entering the EqualComparisonBoolInputRule production.
	EnterEqualComparisonBoolInputRule(c *EqualComparisonBoolInputRuleContext)

	// EnterEmptyBoolInputRule is called when entering the EmptyBoolInputRule production.
	EnterEmptyBoolInputRule(c *EmptyBoolInputRuleContext)

	// EnterEqualComparisonDateTimeInputRule is called when entering the EqualComparisonDateTimeInputRule production.
	EnterEqualComparisonDateTimeInputRule(c *EqualComparisonDateTimeInputRuleContext)

	// EnterComparisonDateTimeInputRule is called when entering the ComparisonDateTimeInputRule production.
	EnterComparisonDateTimeInputRule(c *ComparisonDateTimeInputRuleContext)

	// EnterRangeComparisonDateTimeInputRule is called when entering the RangeComparisonDateTimeInputRule production.
	EnterRangeComparisonDateTimeInputRule(c *RangeComparisonDateTimeInputRuleContext)

	// EnterDisjunctionsDateTimeInputRule is called when entering the DisjunctionsDateTimeInputRule production.
	EnterDisjunctionsDateTimeInputRule(c *DisjunctionsDateTimeInputRuleContext)

	// EnterNegationDateTimeInputRule is called when entering the NegationDateTimeInputRule production.
	EnterNegationDateTimeInputRule(c *NegationDateTimeInputRuleContext)

	// EnterEmptyDateTimeInputRule is called when entering the EmptyDateTimeInputRule production.
	EnterEmptyDateTimeInputRule(c *EmptyDateTimeInputRuleContext)

	// EnterIntegerOutputRule is called when entering the IntegerOutputRule production.
	EnterIntegerOutputRule(c *IntegerOutputRuleContext)

	// EnterEmptyIntegerOutputRule is called when entering the EmptyIntegerOutputRule production.
	EnterEmptyIntegerOutputRule(c *EmptyIntegerOutputRuleContext)

	// EnterNumberOutputRule is called when entering the NumberOutputRule production.
	EnterNumberOutputRule(c *NumberOutputRuleContext)

	// EnterEmptyNumberOutputRule is called when entering the EmptyNumberOutputRule production.
	EnterEmptyNumberOutputRule(c *EmptyNumberOutputRuleContext)

	// EnterStringOutputRule is called when entering the StringOutputRule production.
	EnterStringOutputRule(c *StringOutputRuleContext)

	// EnterEmptyStringOutputRule is called when entering the EmptyStringOutputRule production.
	EnterEmptyStringOutputRule(c *EmptyStringOutputRuleContext)

	// EnterBoolOutputRule is called when entering the BoolOutputRule production.
	EnterBoolOutputRule(c *BoolOutputRuleContext)

	// EnterEmptyBoolOutputRule is called when entering the EmptyBoolOutputRule production.
	EnterEmptyBoolOutputRule(c *EmptyBoolOutputRuleContext)

	// EnterDateTimeOutputRule is called when entering the DateTimeOutputRule production.
	EnterDateTimeOutputRule(c *DateTimeOutputRuleContext)

	// EnterEmptyDateTimeOutputRule is called when entering the EmptyDateTimeOutputRule production.
	EnterEmptyDateTimeOutputRule(c *EmptyDateTimeOutputRuleContext)

	// EnterEqualcomparisonRule is called when entering the EqualcomparisonRule production.
	EnterEqualcomparisonRule(c *EqualcomparisonRuleContext)

	// EnterComparisionsRule is called when entering the ComparisionsRule production.
	EnterComparisionsRule(c *ComparisionsRuleContext)

	// EnterRangeRule is called when entering the RangeRule production.
	EnterRangeRule(c *RangeRuleContext)

	// EnterDisjunctionRule is called when entering the DisjunctionRule production.
	EnterDisjunctionRule(c *DisjunctionRuleContext)

	// EnterNegationRule is called when entering the NegationRule production.
	EnterNegationRule(c *NegationRuleContext)

	// EnterEmptyInputRule is called when entering the EmptyInputRule production.
	EnterEmptyInputRule(c *EmptyInputRuleContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterStrings is called when entering the strings production.
	EnterStrings(c *StringsContext)

	// EnterBools is called when entering the bools production.
	EnterBools(c *BoolsContext)

	// EnterDatetime is called when entering the datetime production.
	EnterDatetime(c *DatetimeContext)

	// EnterEqualcomparison is called when entering the equalcomparison production.
	EnterEqualcomparison(c *EqualcomparisonContext)

	// EnterEqualcomparisonInteger is called when entering the equalcomparisonInteger production.
	EnterEqualcomparisonInteger(c *EqualcomparisonIntegerContext)

	// EnterEqualcomparisonNumber is called when entering the equalcomparisonNumber production.
	EnterEqualcomparisonNumber(c *EqualcomparisonNumberContext)

	// EnterEqualcomparisonBool is called when entering the equalcomparisonBool production.
	EnterEqualcomparisonBool(c *EqualcomparisonBoolContext)

	// EnterEqualcomparisonStrings is called when entering the equalcomparisonStrings production.
	EnterEqualcomparisonStrings(c *EqualcomparisonStringsContext)

	// EnterEqualcomparisonDateTime is called when entering the equalcomparisonDateTime production.
	EnterEqualcomparisonDateTime(c *EqualcomparisonDateTimeContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterComparisonOps is called when entering the comparisonOps production.
	EnterComparisonOps(c *ComparisonOpsContext)

	// EnterComparisonInteger is called when entering the comparisonInteger production.
	EnterComparisonInteger(c *ComparisonIntegerContext)

	// EnterComparisonDateTime is called when entering the comparisonDateTime production.
	EnterComparisonDateTime(c *ComparisonDateTimeContext)

	// EnterComparisonNumber is called when entering the comparisonNumber production.
	EnterComparisonNumber(c *ComparisonNumberContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterRop is called when entering the rop production.
	EnterRop(c *RopContext)

	// EnterRangeInteger is called when entering the rangeInteger production.
	EnterRangeInteger(c *RangeIntegerContext)

	// EnterRangeNumber is called when entering the rangeNumber production.
	EnterRangeNumber(c *RangeNumberContext)

	// EnterRangeDateTime is called when entering the rangeDateTime production.
	EnterRangeDateTime(c *RangeDateTimeContext)

	// EnterDisjunctions is called when entering the disjunctions production.
	EnterDisjunctions(c *DisjunctionsContext)

	// EnterDisjunctionsInteger is called when entering the disjunctionsInteger production.
	EnterDisjunctionsInteger(c *DisjunctionsIntegerContext)

	// EnterDisjunctionsNumber is called when entering the disjunctionsNumber production.
	EnterDisjunctionsNumber(c *DisjunctionsNumberContext)

	// EnterDisjunctionsString is called when entering the disjunctionsString production.
	EnterDisjunctionsString(c *DisjunctionsStringContext)

	// EnterDisjunctionsDateTime is called when entering the disjunctionsDateTime production.
	EnterDisjunctionsDateTime(c *DisjunctionsDateTimeContext)

	// EnterNegation is called when entering the negation production.
	EnterNegation(c *NegationContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitEqualComparisonIntegerInputRule is called when exiting the EqualComparisonIntegerInputRule production.
	ExitEqualComparisonIntegerInputRule(c *EqualComparisonIntegerInputRuleContext)

	// ExitComparisonIntegerInputRule is called when exiting the ComparisonIntegerInputRule production.
	ExitComparisonIntegerInputRule(c *ComparisonIntegerInputRuleContext)

	// ExitRangeComparisonIntegerInputRule is called when exiting the RangeComparisonIntegerInputRule production.
	ExitRangeComparisonIntegerInputRule(c *RangeComparisonIntegerInputRuleContext)

	// ExitDisjunctionsIntegerInputRule is called when exiting the DisjunctionsIntegerInputRule production.
	ExitDisjunctionsIntegerInputRule(c *DisjunctionsIntegerInputRuleContext)

	// ExitNegationIntegerInputRule is called when exiting the NegationIntegerInputRule production.
	ExitNegationIntegerInputRule(c *NegationIntegerInputRuleContext)

	// ExitEmptyIntegerInputRule is called when exiting the EmptyIntegerInputRule production.
	ExitEmptyIntegerInputRule(c *EmptyIntegerInputRuleContext)

	// ExitEqualComparisonNumberInputRule is called when exiting the EqualComparisonNumberInputRule production.
	ExitEqualComparisonNumberInputRule(c *EqualComparisonNumberInputRuleContext)

	// ExitComparisonNumberInputRule is called when exiting the ComparisonNumberInputRule production.
	ExitComparisonNumberInputRule(c *ComparisonNumberInputRuleContext)

	// ExitRangeComparisonNumberInputRule is called when exiting the RangeComparisonNumberInputRule production.
	ExitRangeComparisonNumberInputRule(c *RangeComparisonNumberInputRuleContext)

	// ExitDisjunctionsNumberInputRule is called when exiting the DisjunctionsNumberInputRule production.
	ExitDisjunctionsNumberInputRule(c *DisjunctionsNumberInputRuleContext)

	// ExitNegationNumberInputRule is called when exiting the NegationNumberInputRule production.
	ExitNegationNumberInputRule(c *NegationNumberInputRuleContext)

	// ExitEmptyNumberInputRule is called when exiting the EmptyNumberInputRule production.
	ExitEmptyNumberInputRule(c *EmptyNumberInputRuleContext)

	// ExitEqualComparisonStringInputRule is called when exiting the EqualComparisonStringInputRule production.
	ExitEqualComparisonStringInputRule(c *EqualComparisonStringInputRuleContext)

	// ExitDisjunctionsStringInputRule is called when exiting the DisjunctionsStringInputRule production.
	ExitDisjunctionsStringInputRule(c *DisjunctionsStringInputRuleContext)

	// ExitNegationStringInputRule is called when exiting the NegationStringInputRule production.
	ExitNegationStringInputRule(c *NegationStringInputRuleContext)

	// ExitEmptyStringInputRule is called when exiting the EmptyStringInputRule production.
	ExitEmptyStringInputRule(c *EmptyStringInputRuleContext)

	// ExitEqualComparisonBoolInputRule is called when exiting the EqualComparisonBoolInputRule production.
	ExitEqualComparisonBoolInputRule(c *EqualComparisonBoolInputRuleContext)

	// ExitEmptyBoolInputRule is called when exiting the EmptyBoolInputRule production.
	ExitEmptyBoolInputRule(c *EmptyBoolInputRuleContext)

	// ExitEqualComparisonDateTimeInputRule is called when exiting the EqualComparisonDateTimeInputRule production.
	ExitEqualComparisonDateTimeInputRule(c *EqualComparisonDateTimeInputRuleContext)

	// ExitComparisonDateTimeInputRule is called when exiting the ComparisonDateTimeInputRule production.
	ExitComparisonDateTimeInputRule(c *ComparisonDateTimeInputRuleContext)

	// ExitRangeComparisonDateTimeInputRule is called when exiting the RangeComparisonDateTimeInputRule production.
	ExitRangeComparisonDateTimeInputRule(c *RangeComparisonDateTimeInputRuleContext)

	// ExitDisjunctionsDateTimeInputRule is called when exiting the DisjunctionsDateTimeInputRule production.
	ExitDisjunctionsDateTimeInputRule(c *DisjunctionsDateTimeInputRuleContext)

	// ExitNegationDateTimeInputRule is called when exiting the NegationDateTimeInputRule production.
	ExitNegationDateTimeInputRule(c *NegationDateTimeInputRuleContext)

	// ExitEmptyDateTimeInputRule is called when exiting the EmptyDateTimeInputRule production.
	ExitEmptyDateTimeInputRule(c *EmptyDateTimeInputRuleContext)

	// ExitIntegerOutputRule is called when exiting the IntegerOutputRule production.
	ExitIntegerOutputRule(c *IntegerOutputRuleContext)

	// ExitEmptyIntegerOutputRule is called when exiting the EmptyIntegerOutputRule production.
	ExitEmptyIntegerOutputRule(c *EmptyIntegerOutputRuleContext)

	// ExitNumberOutputRule is called when exiting the NumberOutputRule production.
	ExitNumberOutputRule(c *NumberOutputRuleContext)

	// ExitEmptyNumberOutputRule is called when exiting the EmptyNumberOutputRule production.
	ExitEmptyNumberOutputRule(c *EmptyNumberOutputRuleContext)

	// ExitStringOutputRule is called when exiting the StringOutputRule production.
	ExitStringOutputRule(c *StringOutputRuleContext)

	// ExitEmptyStringOutputRule is called when exiting the EmptyStringOutputRule production.
	ExitEmptyStringOutputRule(c *EmptyStringOutputRuleContext)

	// ExitBoolOutputRule is called when exiting the BoolOutputRule production.
	ExitBoolOutputRule(c *BoolOutputRuleContext)

	// ExitEmptyBoolOutputRule is called when exiting the EmptyBoolOutputRule production.
	ExitEmptyBoolOutputRule(c *EmptyBoolOutputRuleContext)

	// ExitDateTimeOutputRule is called when exiting the DateTimeOutputRule production.
	ExitDateTimeOutputRule(c *DateTimeOutputRuleContext)

	// ExitEmptyDateTimeOutputRule is called when exiting the EmptyDateTimeOutputRule production.
	ExitEmptyDateTimeOutputRule(c *EmptyDateTimeOutputRuleContext)

	// ExitEqualcomparisonRule is called when exiting the EqualcomparisonRule production.
	ExitEqualcomparisonRule(c *EqualcomparisonRuleContext)

	// ExitComparisionsRule is called when exiting the ComparisionsRule production.
	ExitComparisionsRule(c *ComparisionsRuleContext)

	// ExitRangeRule is called when exiting the RangeRule production.
	ExitRangeRule(c *RangeRuleContext)

	// ExitDisjunctionRule is called when exiting the DisjunctionRule production.
	ExitDisjunctionRule(c *DisjunctionRuleContext)

	// ExitNegationRule is called when exiting the NegationRule production.
	ExitNegationRule(c *NegationRuleContext)

	// ExitEmptyInputRule is called when exiting the EmptyInputRule production.
	ExitEmptyInputRule(c *EmptyInputRuleContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitStrings is called when exiting the strings production.
	ExitStrings(c *StringsContext)

	// ExitBools is called when exiting the bools production.
	ExitBools(c *BoolsContext)

	// ExitDatetime is called when exiting the datetime production.
	ExitDatetime(c *DatetimeContext)

	// ExitEqualcomparison is called when exiting the equalcomparison production.
	ExitEqualcomparison(c *EqualcomparisonContext)

	// ExitEqualcomparisonInteger is called when exiting the equalcomparisonInteger production.
	ExitEqualcomparisonInteger(c *EqualcomparisonIntegerContext)

	// ExitEqualcomparisonNumber is called when exiting the equalcomparisonNumber production.
	ExitEqualcomparisonNumber(c *EqualcomparisonNumberContext)

	// ExitEqualcomparisonBool is called when exiting the equalcomparisonBool production.
	ExitEqualcomparisonBool(c *EqualcomparisonBoolContext)

	// ExitEqualcomparisonStrings is called when exiting the equalcomparisonStrings production.
	ExitEqualcomparisonStrings(c *EqualcomparisonStringsContext)

	// ExitEqualcomparisonDateTime is called when exiting the equalcomparisonDateTime production.
	ExitEqualcomparisonDateTime(c *EqualcomparisonDateTimeContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitComparisonOps is called when exiting the comparisonOps production.
	ExitComparisonOps(c *ComparisonOpsContext)

	// ExitComparisonInteger is called when exiting the comparisonInteger production.
	ExitComparisonInteger(c *ComparisonIntegerContext)

	// ExitComparisonDateTime is called when exiting the comparisonDateTime production.
	ExitComparisonDateTime(c *ComparisonDateTimeContext)

	// ExitComparisonNumber is called when exiting the comparisonNumber production.
	ExitComparisonNumber(c *ComparisonNumberContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitRop is called when exiting the rop production.
	ExitRop(c *RopContext)

	// ExitRangeInteger is called when exiting the rangeInteger production.
	ExitRangeInteger(c *RangeIntegerContext)

	// ExitRangeNumber is called when exiting the rangeNumber production.
	ExitRangeNumber(c *RangeNumberContext)

	// ExitRangeDateTime is called when exiting the rangeDateTime production.
	ExitRangeDateTime(c *RangeDateTimeContext)

	// ExitDisjunctions is called when exiting the disjunctions production.
	ExitDisjunctions(c *DisjunctionsContext)

	// ExitDisjunctionsInteger is called when exiting the disjunctionsInteger production.
	ExitDisjunctionsInteger(c *DisjunctionsIntegerContext)

	// ExitDisjunctionsNumber is called when exiting the disjunctionsNumber production.
	ExitDisjunctionsNumber(c *DisjunctionsNumberContext)

	// ExitDisjunctionsString is called when exiting the disjunctionsString production.
	ExitDisjunctionsString(c *DisjunctionsStringContext)

	// ExitDisjunctionsDateTime is called when exiting the disjunctionsDateTime production.
	ExitDisjunctionsDateTime(c *DisjunctionsDateTimeContext)

	// ExitNegation is called when exiting the negation production.
	ExitNegation(c *NegationContext)
}
