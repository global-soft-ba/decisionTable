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

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterComparisondatetime is called when entering the comparisondatetime production.
	EnterComparisondatetime(c *ComparisondatetimeContext)

	// EnterComparisonnumber is called when entering the comparisonnumber production.
	EnterComparisonnumber(c *ComparisonnumberContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterRop is called when entering the rop production.
	EnterRop(c *RopContext)

	// EnterRangenumber is called when entering the rangenumber production.
	EnterRangenumber(c *RangenumberContext)

	// EnterRangedatetime is called when entering the rangedatetime production.
	EnterRangedatetime(c *RangedatetimeContext)

	// EnterDisjunctions is called when entering the disjunctions production.
	EnterDisjunctions(c *DisjunctionsContext)

	// EnterDisjunctionsNumber is called when entering the disjunctionsNumber production.
	EnterDisjunctionsNumber(c *DisjunctionsNumberContext)

	// EnterDisjunctionsString is called when entering the disjunctionsString production.
	EnterDisjunctionsString(c *DisjunctionsStringContext)

	// EnterDisjunctionsDateTime is called when entering the disjunctionsDateTime production.
	EnterDisjunctionsDateTime(c *DisjunctionsDateTimeContext)

	// EnterNegation is called when entering the negation production.
	EnterNegation(c *NegationContext)

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

	// EnterOutputAssignment is called when entering the OutputAssignment production.
	EnterOutputAssignment(c *OutputAssignmentContext)

	// EnterEmptyOutputRule is called when entering the EmptyOutputRule production.
	EnterEmptyOutputRule(c *EmptyOutputRuleContext)

	// ExitInputEntry is called when exiting the inputEntry production.
	ExitInputEntry(c *InputEntryContext)

	// ExitOutputEntry is called when exiting the outputEntry production.
	ExitOutputEntry(c *OutputEntryContext)

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

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitComparisondatetime is called when exiting the comparisondatetime production.
	ExitComparisondatetime(c *ComparisondatetimeContext)

	// ExitComparisonnumber is called when exiting the comparisonnumber production.
	ExitComparisonnumber(c *ComparisonnumberContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitRop is called when exiting the rop production.
	ExitRop(c *RopContext)

	// ExitRangenumber is called when exiting the rangenumber production.
	ExitRangenumber(c *RangenumberContext)

	// ExitRangedatetime is called when exiting the rangedatetime production.
	ExitRangedatetime(c *RangedatetimeContext)

	// ExitDisjunctions is called when exiting the disjunctions production.
	ExitDisjunctions(c *DisjunctionsContext)

	// ExitDisjunctionsNumber is called when exiting the disjunctionsNumber production.
	ExitDisjunctionsNumber(c *DisjunctionsNumberContext)

	// ExitDisjunctionsString is called when exiting the disjunctionsString production.
	ExitDisjunctionsString(c *DisjunctionsStringContext)

	// ExitDisjunctionsDateTime is called when exiting the disjunctionsDateTime production.
	ExitDisjunctionsDateTime(c *DisjunctionsDateTimeContext)

	// ExitNegation is called when exiting the negation production.
	ExitNegation(c *NegationContext)

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

	// ExitOutputAssignment is called when exiting the OutputAssignment production.
	ExitOutputAssignment(c *OutputAssignmentContext)

	// ExitEmptyOutputRule is called when exiting the EmptyOutputRule production.
	ExitEmptyOutputRule(c *EmptyOutputRuleContext)
}
