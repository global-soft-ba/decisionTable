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

// EnterEntry is called when production entry is entered.
func (s *BaseSFeelListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseSFeelListener) ExitEntry(ctx *EntryContext) {}

// EnterValidIntegerInput is called when production validIntegerInput is entered.
func (s *BaseSFeelListener) EnterValidIntegerInput(ctx *ValidIntegerInputContext) {}

// ExitValidIntegerInput is called when production validIntegerInput is exited.
func (s *BaseSFeelListener) ExitValidIntegerInput(ctx *ValidIntegerInputContext) {}

// EnterValidNumberInput is called when production validNumberInput is entered.
func (s *BaseSFeelListener) EnterValidNumberInput(ctx *ValidNumberInputContext) {}

// ExitValidNumberInput is called when production validNumberInput is exited.
func (s *BaseSFeelListener) ExitValidNumberInput(ctx *ValidNumberInputContext) {}

// EnterValidStringInput is called when production validStringInput is entered.
func (s *BaseSFeelListener) EnterValidStringInput(ctx *ValidStringInputContext) {}

// ExitValidStringInput is called when production validStringInput is exited.
func (s *BaseSFeelListener) ExitValidStringInput(ctx *ValidStringInputContext) {}

// EnterValidBoolInput is called when production validBoolInput is entered.
func (s *BaseSFeelListener) EnterValidBoolInput(ctx *ValidBoolInputContext) {}

// ExitValidBoolInput is called when production validBoolInput is exited.
func (s *BaseSFeelListener) ExitValidBoolInput(ctx *ValidBoolInputContext) {}

// EnterValidDateTimeInput is called when production validDateTimeInput is entered.
func (s *BaseSFeelListener) EnterValidDateTimeInput(ctx *ValidDateTimeInputContext) {}

// ExitValidDateTimeInput is called when production validDateTimeInput is exited.
func (s *BaseSFeelListener) ExitValidDateTimeInput(ctx *ValidDateTimeInputContext) {}

// EnterValidIntegerOutput is called when production validIntegerOutput is entered.
func (s *BaseSFeelListener) EnterValidIntegerOutput(ctx *ValidIntegerOutputContext) {}

// ExitValidIntegerOutput is called when production validIntegerOutput is exited.
func (s *BaseSFeelListener) ExitValidIntegerOutput(ctx *ValidIntegerOutputContext) {}

// EnterValidNumberOutput is called when production validNumberOutput is entered.
func (s *BaseSFeelListener) EnterValidNumberOutput(ctx *ValidNumberOutputContext) {}

// ExitValidNumberOutput is called when production validNumberOutput is exited.
func (s *BaseSFeelListener) ExitValidNumberOutput(ctx *ValidNumberOutputContext) {}

// EnterValidStringOutput is called when production validStringOutput is entered.
func (s *BaseSFeelListener) EnterValidStringOutput(ctx *ValidStringOutputContext) {}

// ExitValidStringOutput is called when production validStringOutput is exited.
func (s *BaseSFeelListener) ExitValidStringOutput(ctx *ValidStringOutputContext) {}

// EnterValidBoolOutput is called when production validBoolOutput is entered.
func (s *BaseSFeelListener) EnterValidBoolOutput(ctx *ValidBoolOutputContext) {}

// ExitValidBoolOutput is called when production validBoolOutput is exited.
func (s *BaseSFeelListener) ExitValidBoolOutput(ctx *ValidBoolOutputContext) {}

// EnterValidDateTimeOutput is called when production validDateTimeOutput is entered.
func (s *BaseSFeelListener) EnterValidDateTimeOutput(ctx *ValidDateTimeOutputContext) {}

// ExitValidDateTimeOutput is called when production validDateTimeOutput is exited.
func (s *BaseSFeelListener) ExitValidDateTimeOutput(ctx *ValidDateTimeOutputContext) {}

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

// EnterEmptyInputRule is called when production EmptyInputRule is entered.
func (s *BaseSFeelListener) EnterEmptyInputRule(ctx *EmptyInputRuleContext) {}

// ExitEmptyInputRule is called when production EmptyInputRule is exited.
func (s *BaseSFeelListener) ExitEmptyInputRule(ctx *EmptyInputRuleContext) {}

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

// EnterEqualcomparisonInteger is called when production equalcomparisonInteger is entered.
func (s *BaseSFeelListener) EnterEqualcomparisonInteger(ctx *EqualcomparisonIntegerContext) {}

// ExitEqualcomparisonInteger is called when production equalcomparisonInteger is exited.
func (s *BaseSFeelListener) ExitEqualcomparisonInteger(ctx *EqualcomparisonIntegerContext) {}

// EnterEqualcomparisonNumber is called when production equalcomparisonNumber is entered.
func (s *BaseSFeelListener) EnterEqualcomparisonNumber(ctx *EqualcomparisonNumberContext) {}

// ExitEqualcomparisonNumber is called when production equalcomparisonNumber is exited.
func (s *BaseSFeelListener) ExitEqualcomparisonNumber(ctx *EqualcomparisonNumberContext) {}

// EnterEqualcomparisonBool is called when production equalcomparisonBool is entered.
func (s *BaseSFeelListener) EnterEqualcomparisonBool(ctx *EqualcomparisonBoolContext) {}

// ExitEqualcomparisonBool is called when production equalcomparisonBool is exited.
func (s *BaseSFeelListener) ExitEqualcomparisonBool(ctx *EqualcomparisonBoolContext) {}

// EnterEqualcomparisonStrings is called when production equalcomparisonStrings is entered.
func (s *BaseSFeelListener) EnterEqualcomparisonStrings(ctx *EqualcomparisonStringsContext) {}

// ExitEqualcomparisonStrings is called when production equalcomparisonStrings is exited.
func (s *BaseSFeelListener) ExitEqualcomparisonStrings(ctx *EqualcomparisonStringsContext) {}

// EnterEqualcomparisonDateTime is called when production equalcomparisonDateTime is entered.
func (s *BaseSFeelListener) EnterEqualcomparisonDateTime(ctx *EqualcomparisonDateTimeContext) {}

// ExitEqualcomparisonDateTime is called when production equalcomparisonDateTime is exited.
func (s *BaseSFeelListener) ExitEqualcomparisonDateTime(ctx *EqualcomparisonDateTimeContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSFeelListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSFeelListener) ExitComparison(ctx *ComparisonContext) {}

// EnterComparisonOps is called when production comparisonOps is entered.
func (s *BaseSFeelListener) EnterComparisonOps(ctx *ComparisonOpsContext) {}

// ExitComparisonOps is called when production comparisonOps is exited.
func (s *BaseSFeelListener) ExitComparisonOps(ctx *ComparisonOpsContext) {}

// EnterComparisonInteger is called when production comparisonInteger is entered.
func (s *BaseSFeelListener) EnterComparisonInteger(ctx *ComparisonIntegerContext) {}

// ExitComparisonInteger is called when production comparisonInteger is exited.
func (s *BaseSFeelListener) ExitComparisonInteger(ctx *ComparisonIntegerContext) {}

// EnterComparisonDateTime is called when production comparisonDateTime is entered.
func (s *BaseSFeelListener) EnterComparisonDateTime(ctx *ComparisonDateTimeContext) {}

// ExitComparisonDateTime is called when production comparisonDateTime is exited.
func (s *BaseSFeelListener) ExitComparisonDateTime(ctx *ComparisonDateTimeContext) {}

// EnterComparisonNumber is called when production comparisonNumber is entered.
func (s *BaseSFeelListener) EnterComparisonNumber(ctx *ComparisonNumberContext) {}

// ExitComparisonNumber is called when production comparisonNumber is exited.
func (s *BaseSFeelListener) ExitComparisonNumber(ctx *ComparisonNumberContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseSFeelListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseSFeelListener) ExitRanges(ctx *RangesContext) {}

// EnterRop is called when production rop is entered.
func (s *BaseSFeelListener) EnterRop(ctx *RopContext) {}

// ExitRop is called when production rop is exited.
func (s *BaseSFeelListener) ExitRop(ctx *RopContext) {}

// EnterRangeInteger is called when production rangeInteger is entered.
func (s *BaseSFeelListener) EnterRangeInteger(ctx *RangeIntegerContext) {}

// ExitRangeInteger is called when production rangeInteger is exited.
func (s *BaseSFeelListener) ExitRangeInteger(ctx *RangeIntegerContext) {}

// EnterRangeNumber is called when production rangeNumber is entered.
func (s *BaseSFeelListener) EnterRangeNumber(ctx *RangeNumberContext) {}

// ExitRangeNumber is called when production rangeNumber is exited.
func (s *BaseSFeelListener) ExitRangeNumber(ctx *RangeNumberContext) {}

// EnterRangeDateTime is called when production rangeDateTime is entered.
func (s *BaseSFeelListener) EnterRangeDateTime(ctx *RangeDateTimeContext) {}

// ExitRangeDateTime is called when production rangeDateTime is exited.
func (s *BaseSFeelListener) ExitRangeDateTime(ctx *RangeDateTimeContext) {}

// EnterDisjunctions is called when production disjunctions is entered.
func (s *BaseSFeelListener) EnterDisjunctions(ctx *DisjunctionsContext) {}

// ExitDisjunctions is called when production disjunctions is exited.
func (s *BaseSFeelListener) ExitDisjunctions(ctx *DisjunctionsContext) {}

// EnterDisjunctionsInteger is called when production disjunctionsInteger is entered.
func (s *BaseSFeelListener) EnterDisjunctionsInteger(ctx *DisjunctionsIntegerContext) {}

// ExitDisjunctionsInteger is called when production disjunctionsInteger is exited.
func (s *BaseSFeelListener) ExitDisjunctionsInteger(ctx *DisjunctionsIntegerContext) {}

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
