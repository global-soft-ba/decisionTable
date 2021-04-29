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

// EnterNumber is called when production number is entered.
func (s *BaseSFeelListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSFeelListener) ExitNumber(ctx *NumberContext) {}

// EnterStrings is called when production strings is entered.
func (s *BaseSFeelListener) EnterStrings(ctx *StringsContext) {}

// ExitStrings is called when production strings is exited.
func (s *BaseSFeelListener) ExitStrings(ctx *StringsContext) {}

// EnterBools is called when production bools is entered.
func (s *BaseSFeelListener) EnterBools(ctx *BoolsContext) {}

// ExitBools is called when production bools is exited.
func (s *BaseSFeelListener) ExitBools(ctx *BoolsContext) {}

// EnterDatetime is called when production datetime is entered.
func (s *BaseSFeelListener) EnterDatetime(ctx *DatetimeContext) {}

// ExitDatetime is called when production datetime is exited.
func (s *BaseSFeelListener) ExitDatetime(ctx *DatetimeContext) {}

// EnterEqualcomparison is called when production equalcomparison is entered.
func (s *BaseSFeelListener) EnterEqualcomparison(ctx *EqualcomparisonContext) {}

// ExitEqualcomparison is called when production equalcomparison is exited.
func (s *BaseSFeelListener) ExitEqualcomparison(ctx *EqualcomparisonContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSFeelListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSFeelListener) ExitComparison(ctx *ComparisonContext) {}

// EnterOp is called when production op is entered.
func (s *BaseSFeelListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseSFeelListener) ExitOp(ctx *OpContext) {}

// EnterComparisonnumber is called when production comparisonnumber is entered.
func (s *BaseSFeelListener) EnterComparisonnumber(ctx *ComparisonnumberContext) {}

// ExitComparisonnumber is called when production comparisonnumber is exited.
func (s *BaseSFeelListener) ExitComparisonnumber(ctx *ComparisonnumberContext) {}

// EnterComparisondatetime is called when production comparisondatetime is entered.
func (s *BaseSFeelListener) EnterComparisondatetime(ctx *ComparisondatetimeContext) {}

// ExitComparisondatetime is called when production comparisondatetime is exited.
func (s *BaseSFeelListener) ExitComparisondatetime(ctx *ComparisondatetimeContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseSFeelListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseSFeelListener) ExitRanges(ctx *RangesContext) {}

// EnterRop is called when production rop is entered.
func (s *BaseSFeelListener) EnterRop(ctx *RopContext) {}

// ExitRop is called when production rop is exited.
func (s *BaseSFeelListener) ExitRop(ctx *RopContext) {}

// EnterRangenumber is called when production rangenumber is entered.
func (s *BaseSFeelListener) EnterRangenumber(ctx *RangenumberContext) {}

// ExitRangenumber is called when production rangenumber is exited.
func (s *BaseSFeelListener) ExitRangenumber(ctx *RangenumberContext) {}

// EnterRangedatetime is called when production rangedatetime is entered.
func (s *BaseSFeelListener) EnterRangedatetime(ctx *RangedatetimeContext) {}

// ExitRangedatetime is called when production rangedatetime is exited.
func (s *BaseSFeelListener) ExitRangedatetime(ctx *RangedatetimeContext) {}

// EnterDisjunctions is called when production disjunctions is entered.
func (s *BaseSFeelListener) EnterDisjunctions(ctx *DisjunctionsContext) {}

// ExitDisjunctions is called when production disjunctions is exited.
func (s *BaseSFeelListener) ExitDisjunctions(ctx *DisjunctionsContext) {}

// EnterDisjunctionsNumber is called when production disjunctionsNumber is entered.
func (s *BaseSFeelListener) EnterDisjunctionsNumber(ctx *DisjunctionsNumberContext) {}

// ExitDisjunctionsNumber is called when production disjunctionsNumber is exited.
func (s *BaseSFeelListener) ExitDisjunctionsNumber(ctx *DisjunctionsNumberContext) {}

// EnterDisjunctionsString is called when production disjunctionsString is entered.
func (s *BaseSFeelListener) EnterDisjunctionsString(ctx *DisjunctionsStringContext) {}

// ExitDisjunctionsString is called when production disjunctionsString is exited.
func (s *BaseSFeelListener) ExitDisjunctionsString(ctx *DisjunctionsStringContext) {}

// EnterDisjunctionsDateTime is called when production disjunctionsDateTime is entered.
func (s *BaseSFeelListener) EnterDisjunctionsDateTime(ctx *DisjunctionsDateTimeContext) {}

// ExitDisjunctionsDateTime is called when production disjunctionsDateTime is exited.
func (s *BaseSFeelListener) ExitDisjunctionsDateTime(ctx *DisjunctionsDateTimeContext) {}

// EnterNegation is called when production negation is entered.
func (s *BaseSFeelListener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production negation is exited.
func (s *BaseSFeelListener) ExitNegation(ctx *NegationContext) {}

// EnterEqualcomparisonRule is called when production EqualcomparisonRule is entered.
func (s *BaseSFeelListener) EnterEqualcomparisonRule(ctx *EqualcomparisonRuleContext) {}

// ExitEqualcomparisonRule is called when production EqualcomparisonRule is exited.
func (s *BaseSFeelListener) ExitEqualcomparisonRule(ctx *EqualcomparisonRuleContext) {}

// EnterComparisionsRule is called when production ComparisionsRule is entered.
func (s *BaseSFeelListener) EnterComparisionsRule(ctx *ComparisionsRuleContext) {}

// ExitComparisionsRule is called when production ComparisionsRule is exited.
func (s *BaseSFeelListener) ExitComparisionsRule(ctx *ComparisionsRuleContext) {}

// EnterRangeRule is called when production RangeRule is entered.
func (s *BaseSFeelListener) EnterRangeRule(ctx *RangeRuleContext) {}

// ExitRangeRule is called when production RangeRule is exited.
func (s *BaseSFeelListener) ExitRangeRule(ctx *RangeRuleContext) {}

// EnterDisjunctionRule is called when production DisjunctionRule is entered.
func (s *BaseSFeelListener) EnterDisjunctionRule(ctx *DisjunctionRuleContext) {}

// ExitDisjunctionRule is called when production DisjunctionRule is exited.
func (s *BaseSFeelListener) ExitDisjunctionRule(ctx *DisjunctionRuleContext) {}

// EnterNegationRule is called when production NegationRule is entered.
func (s *BaseSFeelListener) EnterNegationRule(ctx *NegationRuleContext) {}

// ExitNegationRule is called when production NegationRule is exited.
func (s *BaseSFeelListener) ExitNegationRule(ctx *NegationRuleContext) {}

// EnterEmptyRule is called when production EmptyRule is entered.
func (s *BaseSFeelListener) EnterEmptyRule(ctx *EmptyRuleContext) {}

// ExitEmptyRule is called when production EmptyRule is exited.
func (s *BaseSFeelListener) ExitEmptyRule(ctx *EmptyRuleContext) {}
