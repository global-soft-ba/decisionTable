// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SFeelListener is a complete listener for a parse tree produced by SFeelParser.
type SFeelListener interface {
	antlr.ParseTreeListener

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterValidIntegerInput is called when entering the validIntegerInput production.
	EnterValidIntegerInput(c *ValidIntegerInputContext)

	// EnterValidNumberInput is called when entering the validNumberInput production.
	EnterValidNumberInput(c *ValidNumberInputContext)

	// EnterValidStringInput is called when entering the validStringInput production.
	EnterValidStringInput(c *ValidStringInputContext)

	// EnterValidBoolInput is called when entering the validBoolInput production.
	EnterValidBoolInput(c *ValidBoolInputContext)

	// EnterValidDateTimeInput is called when entering the validDateTimeInput production.
	EnterValidDateTimeInput(c *ValidDateTimeInputContext)

	// EnterValidIntegerOutput is called when entering the validIntegerOutput production.
	EnterValidIntegerOutput(c *ValidIntegerOutputContext)

	// EnterValidNumberOutput is called when entering the validNumberOutput production.
	EnterValidNumberOutput(c *ValidNumberOutputContext)

	// EnterValidStringOutput is called when entering the validStringOutput production.
	EnterValidStringOutput(c *ValidStringOutputContext)

	// EnterValidBoolOutput is called when entering the validBoolOutput production.
	EnterValidBoolOutput(c *ValidBoolOutputContext)

	// EnterValidDateTimeOutput is called when entering the validDateTimeOutput production.
	EnterValidDateTimeOutput(c *ValidDateTimeOutputContext)

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

	// ExitValidIntegerInput is called when exiting the validIntegerInput production.
	ExitValidIntegerInput(c *ValidIntegerInputContext)

	// ExitValidNumberInput is called when exiting the validNumberInput production.
	ExitValidNumberInput(c *ValidNumberInputContext)

	// ExitValidStringInput is called when exiting the validStringInput production.
	ExitValidStringInput(c *ValidStringInputContext)

	// ExitValidBoolInput is called when exiting the validBoolInput production.
	ExitValidBoolInput(c *ValidBoolInputContext)

	// ExitValidDateTimeInput is called when exiting the validDateTimeInput production.
	ExitValidDateTimeInput(c *ValidDateTimeInputContext)

	// ExitValidIntegerOutput is called when exiting the validIntegerOutput production.
	ExitValidIntegerOutput(c *ValidIntegerOutputContext)

	// ExitValidNumberOutput is called when exiting the validNumberOutput production.
	ExitValidNumberOutput(c *ValidNumberOutputContext)

	// ExitValidStringOutput is called when exiting the validStringOutput production.
	ExitValidStringOutput(c *ValidStringOutputContext)

	// ExitValidBoolOutput is called when exiting the validBoolOutput production.
	ExitValidBoolOutput(c *ValidBoolOutputContext)

	// ExitValidDateTimeOutput is called when exiting the validDateTimeOutput production.
	ExitValidDateTimeOutput(c *ValidDateTimeOutputContext)

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
