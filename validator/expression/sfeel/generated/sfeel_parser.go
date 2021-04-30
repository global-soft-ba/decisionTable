// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 169,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 65, 10, 8, 3, 9, 3, 9, 5, 9,
	69, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 5, 13, 81, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 5, 17, 101, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 107, 10,
	18, 3, 18, 3, 18, 3, 18, 7, 18, 112, 10, 18, 12, 18, 14, 18, 115, 11, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 123, 10, 19, 12, 19, 14,
	19, 126, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 132, 10, 20, 3, 20,
	3, 20, 3, 20, 7, 20, 137, 10, 20, 12, 20, 14, 20, 140, 11, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 148, 10, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 158, 10, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 5, 23, 164, 10, 23, 3, 23, 5, 23, 167, 10, 23, 3, 23, 2, 5,
	34, 36, 38, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 2, 5, 3, 2, 8, 9, 3, 2, 12, 15, 3, 2, 16, 17,
	2, 172, 2, 46, 3, 2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 52, 3, 2, 2, 2, 8, 54,
	3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 58, 3, 2, 2, 2, 14, 64, 3, 2, 2, 2,
	16, 68, 3, 2, 2, 2, 18, 70, 3, 2, 2, 2, 20, 72, 3, 2, 2, 2, 22, 75, 3,
	2, 2, 2, 24, 80, 3, 2, 2, 2, 26, 82, 3, 2, 2, 2, 28, 84, 3, 2, 2, 2, 30,
	90, 3, 2, 2, 2, 32, 100, 3, 2, 2, 2, 34, 102, 3, 2, 2, 2, 36, 116, 3, 2,
	2, 2, 38, 127, 3, 2, 2, 2, 40, 141, 3, 2, 2, 2, 42, 157, 3, 2, 2, 2, 44,
	166, 3, 2, 2, 2, 46, 47, 5, 42, 22, 2, 47, 48, 7, 2, 2, 3, 48, 3, 3, 2,
	2, 2, 49, 50, 5, 44, 23, 2, 50, 51, 7, 2, 2, 3, 51, 5, 3, 2, 2, 2, 52,
	53, 9, 2, 2, 2, 53, 7, 3, 2, 2, 2, 54, 55, 7, 10, 2, 2, 55, 9, 3, 2, 2,
	2, 56, 57, 7, 11, 2, 2, 57, 11, 3, 2, 2, 2, 58, 59, 7, 20, 2, 2, 59, 13,
	3, 2, 2, 2, 60, 65, 5, 10, 6, 2, 61, 65, 5, 8, 5, 2, 62, 65, 5, 12, 7,
	2, 63, 65, 5, 6, 4, 2, 64, 60, 3, 2, 2, 2, 64, 61, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 15, 3, 2, 2, 2, 66, 69, 5, 20, 11,
	2, 67, 69, 5, 22, 12, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 17,
	3, 2, 2, 2, 70, 71, 9, 3, 2, 2, 71, 19, 3, 2, 2, 2, 72, 73, 5, 18, 10,
	2, 73, 74, 5, 12, 7, 2, 74, 21, 3, 2, 2, 2, 75, 76, 5, 18, 10, 2, 76, 77,
	5, 6, 4, 2, 77, 23, 3, 2, 2, 2, 78, 81, 5, 28, 15, 2, 79, 81, 5, 30, 16,
	2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 25, 3, 2, 2, 2, 82, 83,
	9, 4, 2, 2, 83, 27, 3, 2, 2, 2, 84, 85, 5, 26, 14, 2, 85, 86, 5, 6, 4,
	2, 86, 87, 7, 3, 2, 2, 87, 88, 5, 6, 4, 2, 88, 89, 5, 26, 14, 2, 89, 29,
	3, 2, 2, 2, 90, 91, 5, 26, 14, 2, 91, 92, 5, 12, 7, 2, 92, 93, 7, 3, 2,
	2, 93, 94, 5, 12, 7, 2, 94, 95, 3, 2, 2, 2, 95, 96, 5, 26, 14, 2, 96, 31,
	3, 2, 2, 2, 97, 101, 5, 34, 18, 2, 98, 101, 5, 36, 19, 2, 99, 101, 5, 38,
	20, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101,
	33, 3, 2, 2, 2, 102, 106, 8, 18, 1, 2, 103, 107, 5, 6, 4, 2, 104, 107,
	5, 22, 12, 2, 105, 107, 5, 28, 15, 2, 106, 103, 3, 2, 2, 2, 106, 104, 3,
	2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 113, 3, 2, 2, 2, 108, 109, 12, 4, 2,
	2, 109, 110, 7, 18, 2, 2, 110, 112, 5, 34, 18, 5, 111, 108, 3, 2, 2, 2,
	112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114,
	35, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117, 8, 19, 1, 2, 117, 118,
	5, 8, 5, 2, 118, 124, 3, 2, 2, 2, 119, 120, 12, 4, 2, 2, 120, 121, 7, 18,
	2, 2, 121, 123, 5, 36, 19, 5, 122, 119, 3, 2, 2, 2, 123, 126, 3, 2, 2,
	2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 37, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 131, 8, 20, 1, 2, 128, 132, 5, 12, 7, 2, 129, 132,
	5, 20, 11, 2, 130, 132, 5, 30, 16, 2, 131, 128, 3, 2, 2, 2, 131, 129, 3,
	2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 138, 3, 2, 2, 2, 133, 134, 12, 4, 2,
	2, 134, 135, 7, 18, 2, 2, 135, 137, 5, 38, 20, 5, 136, 133, 3, 2, 2, 2,
	137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139,
	39, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 142, 7, 19, 2, 2, 142, 147,
	7, 4, 2, 2, 143, 148, 5, 14, 8, 2, 144, 148, 5, 16, 9, 2, 145, 148, 5,
	24, 13, 2, 146, 148, 5, 32, 17, 2, 147, 143, 3, 2, 2, 2, 147, 144, 3, 2,
	2, 2, 147, 145, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 150, 7, 5, 2, 2, 150, 41, 3, 2, 2, 2, 151, 158, 5, 14, 8, 2, 152,
	158, 5, 16, 9, 2, 153, 158, 5, 24, 13, 2, 154, 158, 5, 32, 17, 2, 155,
	158, 5, 40, 21, 2, 156, 158, 7, 6, 2, 2, 157, 151, 3, 2, 2, 2, 157, 152,
	3, 2, 2, 2, 157, 153, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 157, 155, 3, 2,
	2, 2, 157, 156, 3, 2, 2, 2, 158, 43, 3, 2, 2, 2, 159, 164, 5, 10, 6, 2,
	160, 164, 5, 8, 5, 2, 161, 164, 5, 12, 7, 2, 162, 164, 5, 6, 4, 2, 163,
	159, 3, 2, 2, 2, 163, 160, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 162,
	3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 167, 7, 6, 2, 2, 166, 163, 3, 2,
	2, 2, 166, 165, 3, 2, 2, 2, 167, 45, 3, 2, 2, 2, 15, 64, 68, 80, 100, 106,
	113, 124, 131, 138, 147, 157, 163, 166,
}
var literalNames = []string{
	"", "'..'", "'('", "')'", "'-'", "", "", "", "", "", "'<'", "'<='", "'>'",
	"'>='", "'['", "']'", "','", "'not'",
}
var symbolicNames = []string{
	"", "", "", "", "", "SIGN", "INTEGER", "FLOAT", "STRING", "BOOL", "LESS",
	"LESSEQ", "GREATER", "GREATEREQ", "RANGEIN", "RANGEOUT", "DISJUNCTION",
	"NEGATION", "DATEANDTIME", "FORMAT",
}

var ruleNames = []string{
	"inputEntry", "outputEntry", "number", "strings", "bools", "datetime",
	"equalcomparison", "comparison", "op", "comparisondatetime", "comparisonnumber",
	"ranges", "rop", "rangenumber", "rangedatetime", "disjunctions", "disjunctionsNumber",
	"disjunctionsString", "disjunctionsDateTime", "negation", "inputexpression",
	"outputexpression",
}

type SFeelParser struct {
	*antlr.BaseParser
}

// NewSFeelParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *SFeelParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSFeelParser(input antlr.TokenStream) *SFeelParser {
	this := new(SFeelParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SFeel.g4"

	return this
}

// SFeelParser tokens.
const (
	SFeelParserEOF         = antlr.TokenEOF
	SFeelParserT__0        = 1
	SFeelParserT__1        = 2
	SFeelParserT__2        = 3
	SFeelParserT__3        = 4
	SFeelParserSIGN        = 5
	SFeelParserINTEGER     = 6
	SFeelParserFLOAT       = 7
	SFeelParserSTRING      = 8
	SFeelParserBOOL        = 9
	SFeelParserLESS        = 10
	SFeelParserLESSEQ      = 11
	SFeelParserGREATER     = 12
	SFeelParserGREATEREQ   = 13
	SFeelParserRANGEIN     = 14
	SFeelParserRANGEOUT    = 15
	SFeelParserDISJUNCTION = 16
	SFeelParserNEGATION    = 17
	SFeelParserDATEANDTIME = 18
	SFeelParserFORMAT      = 19
)

// SFeelParser rules.
const (
	SFeelParserRULE_inputEntry           = 0
	SFeelParserRULE_outputEntry          = 1
	SFeelParserRULE_number               = 2
	SFeelParserRULE_strings              = 3
	SFeelParserRULE_bools                = 4
	SFeelParserRULE_datetime             = 5
	SFeelParserRULE_equalcomparison      = 6
	SFeelParserRULE_comparison           = 7
	SFeelParserRULE_op                   = 8
	SFeelParserRULE_comparisondatetime   = 9
	SFeelParserRULE_comparisonnumber     = 10
	SFeelParserRULE_ranges               = 11
	SFeelParserRULE_rop                  = 12
	SFeelParserRULE_rangenumber          = 13
	SFeelParserRULE_rangedatetime        = 14
	SFeelParserRULE_disjunctions         = 15
	SFeelParserRULE_disjunctionsNumber   = 16
	SFeelParserRULE_disjunctionsString   = 17
	SFeelParserRULE_disjunctionsDateTime = 18
	SFeelParserRULE_negation             = 19
	SFeelParserRULE_inputexpression      = 20
	SFeelParserRULE_outputexpression     = 21
)

// IInputEntryContext is an interface to support dynamic dispatch.
type IInputEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputEntryContext differentiates from other interfaces.
	IsInputEntryContext()
}

type InputEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputEntryContext() *InputEntryContext {
	var p = new(InputEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_inputEntry
	return p
}

func (*InputEntryContext) IsInputEntryContext() {}

func NewInputEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputEntryContext {
	var p = new(InputEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_inputEntry

	return p
}

func (s *InputEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *InputEntryContext) Inputexpression() IInputexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputexpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputexpressionContext)
}

func (s *InputEntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *InputEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterInputEntry(s)
	}
}

func (s *InputEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitInputEntry(s)
	}
}

func (p *SFeelParser) InputEntry() (localctx IInputEntryContext) {
	localctx = NewInputEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SFeelParserRULE_inputEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Inputexpression()
	}
	{
		p.SetState(45)
		p.Match(SFeelParserEOF)
	}

	return localctx
}

// IOutputEntryContext is an interface to support dynamic dispatch.
type IOutputEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutputEntryContext differentiates from other interfaces.
	IsOutputEntryContext()
}

type OutputEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputEntryContext() *OutputEntryContext {
	var p = new(OutputEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_outputEntry
	return p
}

func (*OutputEntryContext) IsOutputEntryContext() {}

func NewOutputEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputEntryContext {
	var p = new(OutputEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_outputEntry

	return p
}

func (s *OutputEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputEntryContext) Outputexpression() IOutputexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutputexpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutputexpressionContext)
}

func (s *OutputEntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *OutputEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterOutputEntry(s)
	}
}

func (s *OutputEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitOutputEntry(s)
	}
}

func (p *SFeelParser) OutputEntry() (localctx IOutputEntryContext) {
	localctx = NewOutputEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SFeelParserRULE_outputEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Outputexpression()
	}
	{
		p.SetState(48)
		p.Match(SFeelParserEOF)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SFeelParserFLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SFeelParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SFeelParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserINTEGER || _la == SFeelParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringsContext is an interface to support dynamic dispatch.
type IStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringsContext differentiates from other interfaces.
	IsStringsContext()
}

type StringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringsContext() *StringsContext {
	var p = new(StringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_strings
	return p
}

func (*StringsContext) IsStringsContext() {}

func NewStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringsContext {
	var p = new(StringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_strings

	return p
}

func (s *StringsContext) GetParser() antlr.Parser { return s.parser }

func (s *StringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(SFeelParserSTRING, 0)
}

func (s *StringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterStrings(s)
	}
}

func (s *StringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitStrings(s)
	}
}

func (p *SFeelParser) Strings() (localctx IStringsContext) {
	localctx = NewStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SFeelParserRULE_strings)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(SFeelParserSTRING)
	}

	return localctx
}

// IBoolsContext is an interface to support dynamic dispatch.
type IBoolsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolsContext differentiates from other interfaces.
	IsBoolsContext()
}

type BoolsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolsContext() *BoolsContext {
	var p = new(BoolsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_bools
	return p
}

func (*BoolsContext) IsBoolsContext() {}

func NewBoolsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolsContext {
	var p = new(BoolsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_bools

	return p
}

func (s *BoolsContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SFeelParserBOOL, 0)
}

func (s *BoolsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterBools(s)
	}
}

func (s *BoolsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitBools(s)
	}
}

func (p *SFeelParser) Bools() (localctx IBoolsContext) {
	localctx = NewBoolsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SFeelParserRULE_bools)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(SFeelParserBOOL)
	}

	return localctx
}

// IDatetimeContext is an interface to support dynamic dispatch.
type IDatetimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimeContext differentiates from other interfaces.
	IsDatetimeContext()
}

type DatetimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimeContext() *DatetimeContext {
	var p = new(DatetimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_datetime
	return p
}

func (*DatetimeContext) IsDatetimeContext() {}

func NewDatetimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimeContext {
	var p = new(DatetimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_datetime

	return p
}

func (s *DatetimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimeContext) DATEANDTIME() antlr.TerminalNode {
	return s.GetToken(SFeelParserDATEANDTIME, 0)
}

func (s *DatetimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDatetime(s)
	}
}

func (s *DatetimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDatetime(s)
	}
}

func (p *SFeelParser) Datetime() (localctx IDatetimeContext) {
	localctx = NewDatetimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SFeelParserRULE_datetime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(SFeelParserDATEANDTIME)
	}

	return localctx
}

// IEqualcomparisonContext is an interface to support dynamic dispatch.
type IEqualcomparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonContext differentiates from other interfaces.
	IsEqualcomparisonContext()
}

type EqualcomparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonContext() *EqualcomparisonContext {
	var p = new(EqualcomparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparison
	return p
}

func (*EqualcomparisonContext) IsEqualcomparisonContext() {}

func NewEqualcomparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonContext {
	var p = new(EqualcomparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparison

	return p
}

func (s *EqualcomparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonContext) Bools() IBoolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolsContext)
}

func (s *EqualcomparisonContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *EqualcomparisonContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *EqualcomparisonContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *EqualcomparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparison(s)
	}
}

func (s *EqualcomparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparison(s)
	}
}

func (p *SFeelParser) Equalcomparison() (localctx IEqualcomparisonContext) {
	localctx = NewEqualcomparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SFeelParserRULE_equalcomparison)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserBOOL:
		{
			p.SetState(58)
			p.Bools()
		}

	case SFeelParserSTRING:
		{
			p.SetState(59)
			p.Strings()
		}

	case SFeelParserDATEANDTIME:
		{
			p.SetState(60)
			p.Datetime()
		}

	case SFeelParserINTEGER, SFeelParserFLOAT:
		{
			p.SetState(61)
			p.Number()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) Comparisondatetime() IComparisondatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisondatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisondatetimeContext)
}

func (s *ComparisonContext) Comparisonnumber() IComparisonnumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonnumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonnumberContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (p *SFeelParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SFeelParserRULE_comparison)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Comparisondatetime()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Comparisonnumber()
		}

	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) LESS() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESS, 0)
}

func (s *OpContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESSEQ, 0)
}

func (s *OpContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATER, 0)
}

func (s *OpContext) GREATEREQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATEREQ, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *SFeelParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SFeelParserRULE_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserLESS)|(1<<SFeelParserLESSEQ)|(1<<SFeelParserGREATER)|(1<<SFeelParserGREATEREQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisondatetimeContext is an interface to support dynamic dispatch.
type IComparisondatetimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisondatetimeContext differentiates from other interfaces.
	IsComparisondatetimeContext()
}

type ComparisondatetimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisondatetimeContext() *ComparisondatetimeContext {
	var p = new(ComparisondatetimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparisondatetime
	return p
}

func (*ComparisondatetimeContext) IsComparisondatetimeContext() {}

func NewComparisondatetimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisondatetimeContext {
	var p = new(ComparisondatetimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparisondatetime

	return p
}

func (s *ComparisondatetimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisondatetimeContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *ComparisondatetimeContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *ComparisondatetimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisondatetimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisondatetimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisondatetime(s)
	}
}

func (s *ComparisondatetimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisondatetime(s)
	}
}

func (p *SFeelParser) Comparisondatetime() (localctx IComparisondatetimeContext) {
	localctx = NewComparisondatetimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SFeelParserRULE_comparisondatetime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Op()
	}
	{
		p.SetState(71)
		p.Datetime()
	}

	return localctx
}

// IComparisonnumberContext is an interface to support dynamic dispatch.
type IComparisonnumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonnumberContext differentiates from other interfaces.
	IsComparisonnumberContext()
}

type ComparisonnumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonnumberContext() *ComparisonnumberContext {
	var p = new(ComparisonnumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparisonnumber
	return p
}

func (*ComparisonnumberContext) IsComparisonnumberContext() {}

func NewComparisonnumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonnumberContext {
	var p = new(ComparisonnumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparisonnumber

	return p
}

func (s *ComparisonnumberContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonnumberContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *ComparisonnumberContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ComparisonnumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonnumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonnumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisonnumber(s)
	}
}

func (s *ComparisonnumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisonnumber(s)
	}
}

func (p *SFeelParser) Comparisonnumber() (localctx IComparisonnumberContext) {
	localctx = NewComparisonnumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SFeelParserRULE_comparisonnumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Op()
	}
	{
		p.SetState(74)
		p.Number()
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) Rangenumber() IRangenumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangenumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangenumberContext)
}

func (s *RangesContext) Rangedatetime() IRangedatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangedatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangedatetimeContext)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *SFeelParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SFeelParserRULE_ranges)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Rangenumber()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Rangedatetime()
		}

	}

	return localctx
}

// IRopContext is an interface to support dynamic dispatch.
type IRopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRopContext differentiates from other interfaces.
	IsRopContext()
}

type RopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRopContext() *RopContext {
	var p = new(RopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rop
	return p
}

func (*RopContext) IsRopContext() {}

func NewRopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RopContext {
	var p = new(RopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rop

	return p
}

func (s *RopContext) GetParser() antlr.Parser { return s.parser }

func (s *RopContext) RANGEIN() antlr.TerminalNode {
	return s.GetToken(SFeelParserRANGEIN, 0)
}

func (s *RopContext) RANGEOUT() antlr.TerminalNode {
	return s.GetToken(SFeelParserRANGEOUT, 0)
}

func (s *RopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRop(s)
	}
}

func (s *RopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRop(s)
	}
}

func (p *SFeelParser) Rop() (localctx IRopContext) {
	localctx = NewRopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SFeelParserRULE_rop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserRANGEIN || _la == SFeelParserRANGEOUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRangenumberContext is an interface to support dynamic dispatch.
type IRangenumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangenumberContext differentiates from other interfaces.
	IsRangenumberContext()
}

type RangenumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangenumberContext() *RangenumberContext {
	var p = new(RangenumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rangenumber
	return p
}

func (*RangenumberContext) IsRangenumberContext() {}

func NewRangenumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangenumberContext {
	var p = new(RangenumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rangenumber

	return p
}

func (s *RangenumberContext) GetParser() antlr.Parser { return s.parser }

func (s *RangenumberContext) AllRop() []IRopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRopContext)(nil)).Elem())
	var tst = make([]IRopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRopContext)
		}
	}

	return tst
}

func (s *RangenumberContext) Rop(i int) IRopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRopContext)
}

func (s *RangenumberContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *RangenumberContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *RangenumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangenumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangenumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangenumber(s)
	}
}

func (s *RangenumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangenumber(s)
	}
}

func (p *SFeelParser) Rangenumber() (localctx IRangenumberContext) {
	localctx = NewRangenumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SFeelParserRULE_rangenumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Rop()
	}
	{
		p.SetState(83)
		p.Number()
	}
	{
		p.SetState(84)
		p.Match(SFeelParserT__0)
	}
	{
		p.SetState(85)
		p.Number()
	}
	{
		p.SetState(86)
		p.Rop()
	}

	return localctx
}

// IRangedatetimeContext is an interface to support dynamic dispatch.
type IRangedatetimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangedatetimeContext differentiates from other interfaces.
	IsRangedatetimeContext()
}

type RangedatetimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangedatetimeContext() *RangedatetimeContext {
	var p = new(RangedatetimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rangedatetime
	return p
}

func (*RangedatetimeContext) IsRangedatetimeContext() {}

func NewRangedatetimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangedatetimeContext {
	var p = new(RangedatetimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rangedatetime

	return p
}

func (s *RangedatetimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangedatetimeContext) AllRop() []IRopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRopContext)(nil)).Elem())
	var tst = make([]IRopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRopContext)
		}
	}

	return tst
}

func (s *RangedatetimeContext) Rop(i int) IRopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRopContext)
}

func (s *RangedatetimeContext) AllDatetime() []IDatetimeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatetimeContext)(nil)).Elem())
	var tst = make([]IDatetimeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatetimeContext)
		}
	}

	return tst
}

func (s *RangedatetimeContext) Datetime(i int) IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *RangedatetimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangedatetimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangedatetimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangedatetime(s)
	}
}

func (s *RangedatetimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangedatetime(s)
	}
}

func (p *SFeelParser) Rangedatetime() (localctx IRangedatetimeContext) {
	localctx = NewRangedatetimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SFeelParserRULE_rangedatetime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Rop()
	}

	{
		p.SetState(89)
		p.Datetime()
	}
	{
		p.SetState(90)
		p.Match(SFeelParserT__0)
	}
	{
		p.SetState(91)
		p.Datetime()
	}

	{
		p.SetState(93)
		p.Rop()
	}

	return localctx
}

// IDisjunctionsContext is an interface to support dynamic dispatch.
type IDisjunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsContext differentiates from other interfaces.
	IsDisjunctionsContext()
}

type DisjunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsContext() *DisjunctionsContext {
	var p = new(DisjunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctions
	return p
}

func (*DisjunctionsContext) IsDisjunctionsContext() {}

func NewDisjunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsContext {
	var p = new(DisjunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctions

	return p
}

func (s *DisjunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsContext) DisjunctionsNumber() IDisjunctionsNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsNumberContext)
}

func (s *DisjunctionsContext) DisjunctionsString() IDisjunctionsStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsStringContext)
}

func (s *DisjunctionsContext) DisjunctionsDateTime() IDisjunctionsDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsDateTimeContext)
}

func (s *DisjunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctions(s)
	}
}

func (s *DisjunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctions(s)
	}
}

func (p *SFeelParser) Disjunctions() (localctx IDisjunctionsContext) {
	localctx = NewDisjunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SFeelParserRULE_disjunctions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(95)
			p.disjunctionsNumber(0)
		}

	case 2:
		{
			p.SetState(96)
			p.disjunctionsString(0)
		}

	case 3:
		{
			p.SetState(97)
			p.disjunctionsDateTime(0)
		}

	}

	return localctx
}

// IDisjunctionsNumberContext is an interface to support dynamic dispatch.
type IDisjunctionsNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsNumberContext differentiates from other interfaces.
	IsDisjunctionsNumberContext()
}

type DisjunctionsNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsNumberContext() *DisjunctionsNumberContext {
	var p = new(DisjunctionsNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsNumber
	return p
}

func (*DisjunctionsNumberContext) IsDisjunctionsNumberContext() {}

func NewDisjunctionsNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsNumberContext {
	var p = new(DisjunctionsNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsNumber

	return p
}

func (s *DisjunctionsNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsNumberContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *DisjunctionsNumberContext) Comparisonnumber() IComparisonnumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonnumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonnumberContext)
}

func (s *DisjunctionsNumberContext) Rangenumber() IRangenumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangenumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangenumberContext)
}

func (s *DisjunctionsNumberContext) AllDisjunctionsNumber() []IDisjunctionsNumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem())
	var tst = make([]IDisjunctionsNumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsNumberContext)
		}
	}

	return tst
}

func (s *DisjunctionsNumberContext) DisjunctionsNumber(i int) IDisjunctionsNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsNumberContext)
}

func (s *DisjunctionsNumberContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsNumber(s)
	}
}

func (s *DisjunctionsNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsNumber(s)
	}
}

func (p *SFeelParser) DisjunctionsNumber() (localctx IDisjunctionsNumberContext) {
	return p.disjunctionsNumber(0)
}

func (p *SFeelParser) disjunctionsNumber(_p int) (localctx IDisjunctionsNumberContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsNumberContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsNumberContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, SFeelParserRULE_disjunctionsNumber, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER, SFeelParserFLOAT:
		{
			p.SetState(101)
			p.Number()
		}

	case SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ:
		{
			p.SetState(102)
			p.Comparisonnumber()
		}

	case SFeelParserRANGEIN, SFeelParserRANGEOUT:
		{
			p.SetState(103)
			p.Rangenumber()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsNumberContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsNumber)
			p.SetState(106)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(107)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(108)
				p.disjunctionsNumber(3)
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IDisjunctionsStringContext is an interface to support dynamic dispatch.
type IDisjunctionsStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsStringContext differentiates from other interfaces.
	IsDisjunctionsStringContext()
}

type DisjunctionsStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsStringContext() *DisjunctionsStringContext {
	var p = new(DisjunctionsStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsString
	return p
}

func (*DisjunctionsStringContext) IsDisjunctionsStringContext() {}

func NewDisjunctionsStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsStringContext {
	var p = new(DisjunctionsStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsString

	return p
}

func (s *DisjunctionsStringContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsStringContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *DisjunctionsStringContext) AllDisjunctionsString() []IDisjunctionsStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem())
	var tst = make([]IDisjunctionsStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsStringContext)
		}
	}

	return tst
}

func (s *DisjunctionsStringContext) DisjunctionsString(i int) IDisjunctionsStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsStringContext)
}

func (s *DisjunctionsStringContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsString(s)
	}
}

func (s *DisjunctionsStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsString(s)
	}
}

func (p *SFeelParser) DisjunctionsString() (localctx IDisjunctionsStringContext) {
	return p.disjunctionsString(0)
}

func (p *SFeelParser) disjunctionsString(_p int) (localctx IDisjunctionsStringContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsStringContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsStringContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, SFeelParserRULE_disjunctionsString, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Strings()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsStringContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsString)
			p.SetState(117)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(118)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(119)
				p.disjunctionsString(3)
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IDisjunctionsDateTimeContext is an interface to support dynamic dispatch.
type IDisjunctionsDateTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsDateTimeContext differentiates from other interfaces.
	IsDisjunctionsDateTimeContext()
}

type DisjunctionsDateTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsDateTimeContext() *DisjunctionsDateTimeContext {
	var p = new(DisjunctionsDateTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsDateTime
	return p
}

func (*DisjunctionsDateTimeContext) IsDisjunctionsDateTimeContext() {}

func NewDisjunctionsDateTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsDateTimeContext {
	var p = new(DisjunctionsDateTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsDateTime

	return p
}

func (s *DisjunctionsDateTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsDateTimeContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *DisjunctionsDateTimeContext) Comparisondatetime() IComparisondatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisondatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisondatetimeContext)
}

func (s *DisjunctionsDateTimeContext) Rangedatetime() IRangedatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangedatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangedatetimeContext)
}

func (s *DisjunctionsDateTimeContext) AllDisjunctionsDateTime() []IDisjunctionsDateTimeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem())
	var tst = make([]IDisjunctionsDateTimeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsDateTimeContext)
		}
	}

	return tst
}

func (s *DisjunctionsDateTimeContext) DisjunctionsDateTime(i int) IDisjunctionsDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsDateTimeContext)
}

func (s *DisjunctionsDateTimeContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsDateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsDateTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsDateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsDateTime(s)
	}
}

func (s *DisjunctionsDateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsDateTime(s)
	}
}

func (p *SFeelParser) DisjunctionsDateTime() (localctx IDisjunctionsDateTimeContext) {
	return p.disjunctionsDateTime(0)
}

func (p *SFeelParser) disjunctionsDateTime(_p int) (localctx IDisjunctionsDateTimeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsDateTimeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsDateTimeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, SFeelParserRULE_disjunctionsDateTime, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserDATEANDTIME:
		{
			p.SetState(126)
			p.Datetime()
		}

	case SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ:
		{
			p.SetState(127)
			p.Comparisondatetime()
		}

	case SFeelParserRANGEIN, SFeelParserRANGEOUT:
		{
			p.SetState(128)
			p.Rangedatetime()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsDateTimeContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsDateTime)
			p.SetState(131)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(132)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(133)
				p.disjunctionsDateTime(3)
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// INegationContext is an interface to support dynamic dispatch.
type INegationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegationContext differentiates from other interfaces.
	IsNegationContext()
}

type NegationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegationContext() *NegationContext {
	var p = new(NegationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_negation
	return p
}

func (*NegationContext) IsNegationContext() {}

func NewNegationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegationContext {
	var p = new(NegationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_negation

	return p
}

func (s *NegationContext) GetParser() antlr.Parser { return s.parser }

func (s *NegationContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(SFeelParserNEGATION, 0)
}

func (s *NegationContext) Equalcomparison() IEqualcomparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonContext)
}

func (s *NegationContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *NegationContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *NegationContext) Disjunctions() IDisjunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsContext)
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNegation(s)
	}
}

func (p *SFeelParser) Negation() (localctx INegationContext) {
	localctx = NewNegationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SFeelParserRULE_negation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(SFeelParserNEGATION)
	}
	{
		p.SetState(140)
		p.Match(SFeelParserT__1)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(141)
			p.Equalcomparison()
		}

	case 2:
		{
			p.SetState(142)
			p.Comparison()
		}

	case 3:
		{
			p.SetState(143)
			p.Ranges()
		}

	case 4:
		{
			p.SetState(144)
			p.Disjunctions()
		}

	}
	{
		p.SetState(147)
		p.Match(SFeelParserT__2)
	}

	return localctx
}

// IInputexpressionContext is an interface to support dynamic dispatch.
type IInputexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputexpressionContext differentiates from other interfaces.
	IsInputexpressionContext()
}

type InputexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputexpressionContext() *InputexpressionContext {
	var p = new(InputexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_inputexpression
	return p
}

func (*InputexpressionContext) IsInputexpressionContext() {}

func NewInputexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputexpressionContext {
	var p = new(InputexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_inputexpression

	return p
}

func (s *InputexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputexpressionContext) CopyFrom(ctx *InputexpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InputexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DisjunctionRuleContext struct {
	*InputexpressionContext
}

func NewDisjunctionRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DisjunctionRuleContext {
	var p = new(DisjunctionRuleContext)

	p.InputexpressionContext = NewEmptyInputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputexpressionContext))

	return p
}

func (s *DisjunctionRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionRuleContext) Disjunctions() IDisjunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsContext)
}

func (s *DisjunctionRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionRule(s)
	}
}

func (s *DisjunctionRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionRule(s)
	}
}

type ComparisionsRuleContext struct {
	*InputexpressionContext
}

func NewComparisionsRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisionsRuleContext {
	var p = new(ComparisionsRuleContext)

	p.InputexpressionContext = NewEmptyInputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputexpressionContext))

	return p
}

func (s *ComparisionsRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisionsRuleContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ComparisionsRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisionsRule(s)
	}
}

func (s *ComparisionsRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisionsRule(s)
	}
}

type EmptyInputRuleContext struct {
	*InputexpressionContext
}

func NewEmptyInputRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyInputRuleContext {
	var p = new(EmptyInputRuleContext)

	p.InputexpressionContext = NewEmptyInputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputexpressionContext))

	return p
}

func (s *EmptyInputRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyInputRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEmptyInputRule(s)
	}
}

func (s *EmptyInputRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEmptyInputRule(s)
	}
}

type NegationRuleContext struct {
	*InputexpressionContext
}

func NewNegationRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationRuleContext {
	var p = new(NegationRuleContext)

	p.InputexpressionContext = NewEmptyInputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputexpressionContext))

	return p
}

func (s *NegationRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationRuleContext) Negation() INegationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INegationContext)
}

func (s *NegationRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNegationRule(s)
	}
}

func (s *NegationRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNegationRule(s)
	}
}

type EqualcomparisonRuleContext struct {
	*InputexpressionContext
}

func NewEqualcomparisonRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualcomparisonRuleContext {
	var p = new(EqualcomparisonRuleContext)

	p.InputexpressionContext = NewEmptyInputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputexpressionContext))

	return p
}

func (s *EqualcomparisonRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonRuleContext) Equalcomparison() IEqualcomparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonContext)
}

func (s *EqualcomparisonRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonRule(s)
	}
}

func (s *EqualcomparisonRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonRule(s)
	}
}

type RangeRuleContext struct {
	*InputexpressionContext
}

func NewRangeRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeRuleContext {
	var p = new(RangeRuleContext)

	p.InputexpressionContext = NewEmptyInputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputexpressionContext))

	return p
}

func (s *RangeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeRuleContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *RangeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangeRule(s)
	}
}

func (s *RangeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangeRule(s)
	}
}

func (p *SFeelParser) Inputexpression() (localctx IInputexpressionContext) {
	localctx = NewInputexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SFeelParserRULE_inputexpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualcomparisonRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Equalcomparison()
		}

	case 2:
		localctx = NewComparisionsRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Comparison()
		}

	case 3:
		localctx = NewRangeRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Ranges()
		}

	case 4:
		localctx = NewDisjunctionRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Disjunctions()
		}

	case 5:
		localctx = NewNegationRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Negation()
		}

	case 6:
		localctx = NewEmptyInputRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(154)
			p.Match(SFeelParserT__3)
		}

	}

	return localctx
}

// IOutputexpressionContext is an interface to support dynamic dispatch.
type IOutputexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutputexpressionContext differentiates from other interfaces.
	IsOutputexpressionContext()
}

type OutputexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputexpressionContext() *OutputexpressionContext {
	var p = new(OutputexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_outputexpression
	return p
}

func (*OutputexpressionContext) IsOutputexpressionContext() {}

func NewOutputexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputexpressionContext {
	var p = new(OutputexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_outputexpression

	return p
}

func (s *OutputexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputexpressionContext) CopyFrom(ctx *OutputexpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OutputexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OutputAssignmentContext struct {
	*OutputexpressionContext
}

func NewOutputAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OutputAssignmentContext {
	var p = new(OutputAssignmentContext)

	p.OutputexpressionContext = NewEmptyOutputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OutputexpressionContext))

	return p
}

func (s *OutputAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputAssignmentContext) Bools() IBoolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolsContext)
}

func (s *OutputAssignmentContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *OutputAssignmentContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *OutputAssignmentContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *OutputAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterOutputAssignment(s)
	}
}

func (s *OutputAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitOutputAssignment(s)
	}
}

type EmptyOutputRuleContext struct {
	*OutputexpressionContext
}

func NewEmptyOutputRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyOutputRuleContext {
	var p = new(EmptyOutputRuleContext)

	p.OutputexpressionContext = NewEmptyOutputexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OutputexpressionContext))

	return p
}

func (s *EmptyOutputRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyOutputRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEmptyOutputRule(s)
	}
}

func (s *EmptyOutputRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEmptyOutputRule(s)
	}
}

func (p *SFeelParser) Outputexpression() (localctx IOutputexpressionContext) {
	localctx = NewOutputexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SFeelParserRULE_outputexpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER, SFeelParserFLOAT, SFeelParserSTRING, SFeelParserBOOL, SFeelParserDATEANDTIME:
		localctx = NewOutputAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(161)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SFeelParserBOOL:
			{
				p.SetState(157)
				p.Bools()
			}

		case SFeelParserSTRING:
			{
				p.SetState(158)
				p.Strings()
			}

		case SFeelParserDATEANDTIME:
			{
				p.SetState(159)
				p.Datetime()
			}

		case SFeelParserINTEGER, SFeelParserFLOAT:
			{
				p.SetState(160)
				p.Number()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case SFeelParserT__3:
		localctx = NewEmptyOutputRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(SFeelParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *SFeelParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *DisjunctionsNumberContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsNumberContext)
		}
		return p.DisjunctionsNumber_Sempred(t, predIndex)

	case 17:
		var t *DisjunctionsStringContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsStringContext)
		}
		return p.DisjunctionsString_Sempred(t, predIndex)

	case 18:
		var t *DisjunctionsDateTimeContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsDateTimeContext)
		}
		return p.DisjunctionsDateTime_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SFeelParser) DisjunctionsNumber_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SFeelParser) DisjunctionsString_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SFeelParser) DisjunctionsDateTime_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
