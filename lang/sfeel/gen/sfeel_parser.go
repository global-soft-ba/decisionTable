// Code generated from SFeel.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 215,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 69, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 74, 10, 5, 12, 5, 14, 5, 77, 11,
	5, 3, 6, 3, 6, 5, 6, 81, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 86, 10, 7, 3, 8,
	3, 8, 5, 8, 90, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 97, 10, 8, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 7, 14, 112, 10, 14, 12, 14, 14, 14, 115, 11, 14, 3, 14, 5,
	14, 118, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16, 125, 10, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 136,
	10, 17, 3, 17, 3, 17, 3, 17, 7, 17, 141, 10, 17, 12, 17, 14, 17, 144, 11,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 154,
	10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	7, 18, 165, 10, 18, 12, 18, 14, 18, 168, 11, 18, 3, 19, 3, 19, 5, 19, 172,
	10, 19, 3, 20, 3, 20, 5, 20, 176, 10, 20, 3, 21, 3, 21, 3, 21, 7, 21, 181,
	10, 21, 12, 21, 14, 21, 184, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	190, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 5, 24, 198, 10,
	24, 3, 24, 3, 24, 5, 24, 202, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 28, 6, 28, 211, 10, 28, 13, 28, 14, 28, 212, 3, 28, 2, 4, 32,
	34, 29, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 2, 10, 3, 2, 23, 26, 3, 2, 7, 8,
	4, 2, 4, 4, 9, 9, 3, 2, 23, 28, 3, 2, 19, 20, 3, 2, 21, 22, 3, 2, 11, 14,
	3, 2, 15, 16, 2, 214, 2, 56, 3, 2, 2, 2, 4, 59, 3, 2, 2, 2, 6, 68, 3, 2,
	2, 2, 8, 70, 3, 2, 2, 2, 10, 80, 3, 2, 2, 2, 12, 85, 3, 2, 2, 2, 14, 89,
	3, 2, 2, 2, 16, 98, 3, 2, 2, 2, 18, 100, 3, 2, 2, 2, 20, 102, 3, 2, 2,
	2, 22, 104, 3, 2, 2, 2, 24, 106, 3, 2, 2, 2, 26, 117, 3, 2, 2, 2, 28, 119,
	3, 2, 2, 2, 30, 124, 3, 2, 2, 2, 32, 135, 3, 2, 2, 2, 34, 153, 3, 2, 2,
	2, 36, 171, 3, 2, 2, 2, 38, 175, 3, 2, 2, 2, 40, 177, 3, 2, 2, 2, 42, 189,
	3, 2, 2, 2, 44, 191, 3, 2, 2, 2, 46, 197, 3, 2, 2, 2, 48, 203, 3, 2, 2,
	2, 50, 205, 3, 2, 2, 2, 52, 207, 3, 2, 2, 2, 54, 210, 3, 2, 2, 2, 56, 57,
	5, 6, 4, 2, 57, 58, 7, 2, 2, 3, 58, 3, 3, 2, 2, 2, 59, 60, 5, 26, 14, 2,
	60, 61, 7, 2, 2, 3, 61, 5, 3, 2, 2, 2, 62, 69, 5, 8, 5, 2, 63, 64, 7, 3,
	2, 2, 64, 65, 5, 8, 5, 2, 65, 66, 7, 4, 2, 2, 66, 69, 3, 2, 2, 2, 67, 69,
	5, 24, 13, 2, 68, 62, 3, 2, 2, 2, 68, 63, 3, 2, 2, 2, 68, 67, 3, 2, 2,
	2, 69, 7, 3, 2, 2, 2, 70, 75, 5, 10, 6, 2, 71, 72, 7, 5, 2, 2, 72, 74,
	5, 10, 6, 2, 73, 71, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 9, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 81, 5, 12,
	7, 2, 79, 81, 5, 14, 8, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81,
	11, 3, 2, 2, 2, 82, 83, 9, 2, 2, 2, 83, 86, 5, 36, 19, 2, 84, 86, 5, 36,
	19, 2, 85, 82, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 13, 3, 2, 2, 2, 87,
	90, 5, 16, 9, 2, 88, 90, 5, 18, 10, 2, 89, 87, 3, 2, 2, 2, 89, 88, 3, 2,
	2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 5, 36, 19, 2, 92, 93, 7, 6, 2, 2, 93,
	96, 5, 36, 19, 2, 94, 97, 5, 20, 11, 2, 95, 97, 5, 22, 12, 2, 96, 94, 3,
	2, 2, 2, 96, 95, 3, 2, 2, 2, 97, 15, 3, 2, 2, 2, 98, 99, 9, 3, 2, 2, 99,
	17, 3, 2, 2, 2, 100, 101, 7, 9, 2, 2, 101, 19, 3, 2, 2, 2, 102, 103, 9,
	4, 2, 2, 103, 21, 3, 2, 2, 2, 104, 105, 7, 8, 2, 2, 105, 23, 3, 2, 2, 2,
	106, 107, 7, 22, 2, 2, 107, 25, 3, 2, 2, 2, 108, 113, 5, 30, 16, 2, 109,
	110, 7, 5, 2, 2, 110, 112, 5, 30, 16, 2, 111, 109, 3, 2, 2, 2, 112, 115,
	3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 118, 3, 2,
	2, 2, 115, 113, 3, 2, 2, 2, 116, 118, 5, 24, 13, 2, 117, 108, 3, 2, 2,
	2, 117, 116, 3, 2, 2, 2, 118, 27, 3, 2, 2, 2, 119, 120, 5, 30, 16, 2, 120,
	29, 3, 2, 2, 2, 121, 125, 5, 38, 20, 2, 122, 125, 5, 34, 18, 2, 123, 125,
	5, 32, 17, 2, 124, 121, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3,
	2, 2, 2, 125, 31, 3, 2, 2, 2, 126, 127, 8, 17, 1, 2, 127, 128, 5, 38, 20,
	2, 128, 129, 9, 5, 2, 2, 129, 130, 5, 28, 15, 2, 130, 136, 3, 2, 2, 2,
	131, 132, 5, 34, 18, 2, 132, 133, 9, 5, 2, 2, 133, 134, 5, 28, 15, 2, 134,
	136, 3, 2, 2, 2, 135, 126, 3, 2, 2, 2, 135, 131, 3, 2, 2, 2, 136, 142,
	3, 2, 2, 2, 137, 138, 12, 3, 2, 2, 138, 139, 9, 5, 2, 2, 139, 141, 5, 28,
	15, 2, 140, 137, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2,
	142, 143, 3, 2, 2, 2, 143, 33, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146,
	8, 18, 1, 2, 146, 154, 5, 38, 20, 2, 147, 148, 7, 7, 2, 2, 148, 149, 5,
	34, 18, 2, 149, 150, 7, 4, 2, 2, 150, 154, 3, 2, 2, 2, 151, 152, 7, 22,
	2, 2, 152, 154, 5, 34, 18, 6, 153, 145, 3, 2, 2, 2, 153, 147, 3, 2, 2,
	2, 153, 151, 3, 2, 2, 2, 154, 166, 3, 2, 2, 2, 155, 156, 12, 5, 2, 2, 156,
	157, 7, 18, 2, 2, 157, 165, 5, 34, 18, 6, 158, 159, 12, 4, 2, 2, 159, 160,
	9, 6, 2, 2, 160, 165, 5, 34, 18, 5, 161, 162, 12, 3, 2, 2, 162, 163, 9,
	7, 2, 2, 163, 165, 5, 34, 18, 4, 164, 155, 3, 2, 2, 2, 164, 158, 3, 2,
	2, 2, 164, 161, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2,
	166, 167, 3, 2, 2, 2, 167, 35, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 172,
	5, 40, 21, 2, 170, 172, 5, 42, 22, 2, 171, 169, 3, 2, 2, 2, 171, 170, 3,
	2, 2, 2, 172, 37, 3, 2, 2, 2, 173, 176, 5, 40, 21, 2, 174, 176, 5, 42,
	22, 2, 175, 173, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 39, 3, 2, 2, 2,
	177, 182, 7, 17, 2, 2, 178, 179, 7, 10, 2, 2, 179, 181, 7, 17, 2, 2, 180,
	178, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183,
	3, 2, 2, 2, 183, 41, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 190, 5, 46,
	24, 2, 186, 190, 5, 54, 28, 2, 187, 190, 5, 52, 27, 2, 188, 190, 5, 44,
	23, 2, 189, 185, 3, 2, 2, 2, 189, 186, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2,
	189, 188, 3, 2, 2, 2, 190, 43, 3, 2, 2, 2, 191, 192, 9, 8, 2, 2, 192, 193,
	7, 7, 2, 2, 193, 194, 5, 54, 28, 2, 194, 195, 7, 4, 2, 2, 195, 45, 3, 2,
	2, 2, 196, 198, 7, 22, 2, 2, 197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2,
	198, 201, 3, 2, 2, 2, 199, 202, 5, 48, 25, 2, 200, 202, 5, 50, 26, 2, 201,
	199, 3, 2, 2, 2, 201, 200, 3, 2, 2, 2, 202, 47, 3, 2, 2, 2, 203, 204, 7,
	29, 2, 2, 204, 49, 3, 2, 2, 2, 205, 206, 7, 30, 2, 2, 206, 51, 3, 2, 2,
	2, 207, 208, 9, 9, 2, 2, 208, 53, 3, 2, 2, 2, 209, 211, 7, 31, 2, 2, 210,
	209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213,
	3, 2, 2, 2, 213, 55, 3, 2, 2, 2, 23, 68, 75, 80, 85, 89, 96, 113, 117,
	124, 135, 142, 153, 164, 166, 171, 175, 182, 189, 197, 201, 212,
}
var literalNames = []string{
	"", "'not('", "')'", "','", "'..'", "'('", "']'", "'['", "'.'", "'date'",
	"'time'", "'date and time'", "'duration'", "'true'", "'false'", "", "'**'",
	"'*'", "'/'", "'+'", "'-'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Name", "POW",
	"MUL", "DIV", "ADD", "SUB", "LESS", "LESSEQ", "GREATER", "GREATEREQ", "EQUAL",
	"NOTEQUAL", "INTEGER", "REAL", "STRING", "WS",
}

var ruleNames = []string{
	"input", "output", "simple_unary_tests", "simple_positive_unary_tests",
	"simple_positive_unary_test", "unary_comparison", "interval", "open_interval_start",
	"closed_interval_start", "open_interval_end", "closed_interval_end", "empty_expression",
	"simple_expressions", "expression", "simple_expression", "comparison",
	"arithmetic_expression", "endpoint", "simple_value", "qualified_name",
	"simple_literal", "date_time_literal", "numeric_literal", "integer_literal",
	"real_literal", "boolean_literal", "string_literal",
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
	SFeelParserEOF       = antlr.TokenEOF
	SFeelParserT__0      = 1
	SFeelParserT__1      = 2
	SFeelParserT__2      = 3
	SFeelParserT__3      = 4
	SFeelParserT__4      = 5
	SFeelParserT__5      = 6
	SFeelParserT__6      = 7
	SFeelParserT__7      = 8
	SFeelParserT__8      = 9
	SFeelParserT__9      = 10
	SFeelParserT__10     = 11
	SFeelParserT__11     = 12
	SFeelParserT__12     = 13
	SFeelParserT__13     = 14
	SFeelParserName      = 15
	SFeelParserPOW       = 16
	SFeelParserMUL       = 17
	SFeelParserDIV       = 18
	SFeelParserADD       = 19
	SFeelParserSUB       = 20
	SFeelParserLESS      = 21
	SFeelParserLESSEQ    = 22
	SFeelParserGREATER   = 23
	SFeelParserGREATEREQ = 24
	SFeelParserEQUAL     = 25
	SFeelParserNOTEQUAL  = 26
	SFeelParserINTEGER   = 27
	SFeelParserREAL      = 28
	SFeelParserSTRING    = 29
	SFeelParserWS        = 30
)

// SFeelParser rules.
const (
	SFeelParserRULE_input                       = 0
	SFeelParserRULE_output                      = 1
	SFeelParserRULE_simple_unary_tests          = 2
	SFeelParserRULE_simple_positive_unary_tests = 3
	SFeelParserRULE_simple_positive_unary_test  = 4
	SFeelParserRULE_unary_comparison            = 5
	SFeelParserRULE_interval                    = 6
	SFeelParserRULE_open_interval_start         = 7
	SFeelParserRULE_closed_interval_start       = 8
	SFeelParserRULE_open_interval_end           = 9
	SFeelParserRULE_closed_interval_end         = 10
	SFeelParserRULE_empty_expression            = 11
	SFeelParserRULE_simple_expressions          = 12
	SFeelParserRULE_expression                  = 13
	SFeelParserRULE_simple_expression           = 14
	SFeelParserRULE_comparison                  = 15
	SFeelParserRULE_arithmetic_expression       = 16
	SFeelParserRULE_endpoint                    = 17
	SFeelParserRULE_simple_value                = 18
	SFeelParserRULE_qualified_name              = 19
	SFeelParserRULE_simple_literal              = 20
	SFeelParserRULE_date_time_literal           = 21
	SFeelParserRULE_numeric_literal             = 22
	SFeelParserRULE_integer_literal             = 23
	SFeelParserRULE_real_literal                = 24
	SFeelParserRULE_boolean_literal             = 25
	SFeelParserRULE_string_literal              = 26
)

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_input
	return p
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) Simple_unary_tests() ISimple_unary_testsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_unary_testsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_unary_testsContext)
}

func (s *InputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterInput(s)
	}
}

func (s *InputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitInput(s)
	}
}

func (s *InputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitInput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Input() (localctx IInputContext) {
	this := p
	_ = this

	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SFeelParserRULE_input)

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
		p.Simple_unary_tests()
	}
	{
		p.SetState(55)
		p.Match(SFeelParserEOF)
	}

	return localctx
}

// IOutputContext is an interface to support dynamic dispatch.
type IOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutputContext differentiates from other interfaces.
	IsOutputContext()
}

type OutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputContext() *OutputContext {
	var p = new(OutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_output
	return p
}

func (*OutputContext) IsOutputContext() {}

func NewOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputContext {
	var p = new(OutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_output

	return p
}

func (s *OutputContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputContext) Simple_expressions() ISimple_expressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_expressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionsContext)
}

func (s *OutputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *OutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterOutput(s)
	}
}

func (s *OutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitOutput(s)
	}
}

func (s *OutputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitOutput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Output() (localctx IOutputContext) {
	this := p
	_ = this

	localctx = NewOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SFeelParserRULE_output)

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
		p.SetState(57)
		p.Simple_expressions()
	}
	{
		p.SetState(58)
		p.Match(SFeelParserEOF)
	}

	return localctx
}

// ISimple_unary_testsContext is an interface to support dynamic dispatch.
type ISimple_unary_testsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_unary_testsContext differentiates from other interfaces.
	IsSimple_unary_testsContext()
}

type Simple_unary_testsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_unary_testsContext() *Simple_unary_testsContext {
	var p = new(Simple_unary_testsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_unary_tests
	return p
}

func (*Simple_unary_testsContext) IsSimple_unary_testsContext() {}

func NewSimple_unary_testsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_unary_testsContext {
	var p = new(Simple_unary_testsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_unary_tests

	return p
}

func (s *Simple_unary_testsContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_unary_testsContext) CopyFrom(ctx *Simple_unary_testsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Simple_unary_testsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_unary_testsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EmptySimpleUnaryTestsContext struct {
	*Simple_unary_testsContext
}

func NewEmptySimpleUnaryTestsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptySimpleUnaryTestsContext {
	var p = new(EmptySimpleUnaryTestsContext)

	p.Simple_unary_testsContext = NewEmptySimple_unary_testsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Simple_unary_testsContext))

	return p
}

func (s *EmptySimpleUnaryTestsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptySimpleUnaryTestsContext) Empty_expression() IEmpty_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmpty_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmpty_expressionContext)
}

func (s *EmptySimpleUnaryTestsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEmptySimpleUnaryTests(s)
	}
}

func (s *EmptySimpleUnaryTestsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEmptySimpleUnaryTests(s)
	}
}

func (s *EmptySimpleUnaryTestsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitEmptySimpleUnaryTests(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleUnaryTestsContext struct {
	*Simple_unary_testsContext
}

func NewSimpleUnaryTestsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleUnaryTestsContext {
	var p = new(SimpleUnaryTestsContext)

	p.Simple_unary_testsContext = NewEmptySimple_unary_testsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Simple_unary_testsContext))

	return p
}

func (s *SimpleUnaryTestsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleUnaryTestsContext) Simple_positive_unary_tests() ISimple_positive_unary_testsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_positive_unary_testsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_positive_unary_testsContext)
}

func (s *SimpleUnaryTestsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimpleUnaryTests(s)
	}
}

func (s *SimpleUnaryTestsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimpleUnaryTests(s)
	}
}

func (s *SimpleUnaryTestsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimpleUnaryTests(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegationSimpleUnaryTestsContext struct {
	*Simple_unary_testsContext
}

func NewNegationSimpleUnaryTestsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationSimpleUnaryTestsContext {
	var p = new(NegationSimpleUnaryTestsContext)

	p.Simple_unary_testsContext = NewEmptySimple_unary_testsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Simple_unary_testsContext))

	return p
}

func (s *NegationSimpleUnaryTestsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationSimpleUnaryTestsContext) Simple_positive_unary_tests() ISimple_positive_unary_testsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_positive_unary_testsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_positive_unary_testsContext)
}

func (s *NegationSimpleUnaryTestsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNegationSimpleUnaryTests(s)
	}
}

func (s *NegationSimpleUnaryTestsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNegationSimpleUnaryTests(s)
	}
}

func (s *NegationSimpleUnaryTestsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitNegationSimpleUnaryTests(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_unary_tests() (localctx ISimple_unary_testsContext) {
	this := p
	_ = this

	localctx = NewSimple_unary_testsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SFeelParserRULE_simple_unary_tests)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleUnaryTestsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Simple_positive_unary_tests()
		}

	case 2:
		localctx = NewNegationSimpleUnaryTestsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(SFeelParserT__0)
		}
		{
			p.SetState(62)
			p.Simple_positive_unary_tests()
		}
		{
			p.SetState(63)
			p.Match(SFeelParserT__1)
		}

	case 3:
		localctx = NewEmptySimpleUnaryTestsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.Empty_expression()
		}

	}

	return localctx
}

// ISimple_positive_unary_testsContext is an interface to support dynamic dispatch.
type ISimple_positive_unary_testsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_positive_unary_testsContext differentiates from other interfaces.
	IsSimple_positive_unary_testsContext()
}

type Simple_positive_unary_testsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_positive_unary_testsContext() *Simple_positive_unary_testsContext {
	var p = new(Simple_positive_unary_testsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_positive_unary_tests
	return p
}

func (*Simple_positive_unary_testsContext) IsSimple_positive_unary_testsContext() {}

func NewSimple_positive_unary_testsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_positive_unary_testsContext {
	var p = new(Simple_positive_unary_testsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_positive_unary_tests

	return p
}

func (s *Simple_positive_unary_testsContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_positive_unary_testsContext) AllSimple_positive_unary_test() []ISimple_positive_unary_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimple_positive_unary_testContext)(nil)).Elem())
	var tst = make([]ISimple_positive_unary_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimple_positive_unary_testContext)
		}
	}

	return tst
}

func (s *Simple_positive_unary_testsContext) Simple_positive_unary_test(i int) ISimple_positive_unary_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_positive_unary_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimple_positive_unary_testContext)
}

func (s *Simple_positive_unary_testsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_positive_unary_testsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_positive_unary_testsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimple_positive_unary_tests(s)
	}
}

func (s *Simple_positive_unary_testsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimple_positive_unary_tests(s)
	}
}

func (s *Simple_positive_unary_testsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimple_positive_unary_tests(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_positive_unary_tests() (localctx ISimple_positive_unary_testsContext) {
	this := p
	_ = this

	localctx = NewSimple_positive_unary_testsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SFeelParserRULE_simple_positive_unary_tests)
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
		p.Simple_positive_unary_test()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SFeelParserT__2 {
		{
			p.SetState(69)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(70)
			p.Simple_positive_unary_test()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISimple_positive_unary_testContext is an interface to support dynamic dispatch.
type ISimple_positive_unary_testContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_positive_unary_testContext differentiates from other interfaces.
	IsSimple_positive_unary_testContext()
}

type Simple_positive_unary_testContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_positive_unary_testContext() *Simple_positive_unary_testContext {
	var p = new(Simple_positive_unary_testContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_positive_unary_test
	return p
}

func (*Simple_positive_unary_testContext) IsSimple_positive_unary_testContext() {}

func NewSimple_positive_unary_testContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_positive_unary_testContext {
	var p = new(Simple_positive_unary_testContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_positive_unary_test

	return p
}

func (s *Simple_positive_unary_testContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_positive_unary_testContext) Unary_comparison() IUnary_comparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_comparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_comparisonContext)
}

func (s *Simple_positive_unary_testContext) Interval() IIntervalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalContext)
}

func (s *Simple_positive_unary_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_positive_unary_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_positive_unary_testContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimple_positive_unary_test(s)
	}
}

func (s *Simple_positive_unary_testContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimple_positive_unary_test(s)
	}
}

func (s *Simple_positive_unary_testContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimple_positive_unary_test(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_positive_unary_test() (localctx ISimple_positive_unary_testContext) {
	this := p
	_ = this

	localctx = NewSimple_positive_unary_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SFeelParserRULE_simple_positive_unary_test)

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

	switch p.GetTokenStream().LA(1) {
	case SFeelParserT__8, SFeelParserT__9, SFeelParserT__10, SFeelParserT__11, SFeelParserT__12, SFeelParserT__13, SFeelParserName, SFeelParserSUB, SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ, SFeelParserINTEGER, SFeelParserREAL, SFeelParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Unary_comparison()
		}

	case SFeelParserT__4, SFeelParserT__5, SFeelParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Interval()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnary_comparisonContext is an interface to support dynamic dispatch.
type IUnary_comparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_comparisonContext differentiates from other interfaces.
	IsUnary_comparisonContext()
}

type Unary_comparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_comparisonContext() *Unary_comparisonContext {
	var p = new(Unary_comparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_unary_comparison
	return p
}

func (*Unary_comparisonContext) IsUnary_comparisonContext() {}

func NewUnary_comparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_comparisonContext {
	var p = new(Unary_comparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_unary_comparison

	return p
}

func (s *Unary_comparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_comparisonContext) CopyFrom(ctx *Unary_comparisonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Unary_comparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_comparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualUnaryComparisonContext struct {
	*Unary_comparisonContext
}

func NewEqualUnaryComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualUnaryComparisonContext {
	var p = new(EqualUnaryComparisonContext)

	p.Unary_comparisonContext = NewEmptyUnary_comparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Unary_comparisonContext))

	return p
}

func (s *EqualUnaryComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualUnaryComparisonContext) Endpoint() IEndpointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndpointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndpointContext)
}

func (s *EqualUnaryComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualUnaryComparison(s)
	}
}

func (s *EqualUnaryComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualUnaryComparison(s)
	}
}

func (s *EqualUnaryComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitEqualUnaryComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryComparisonContext struct {
	*Unary_comparisonContext
}

func NewUnaryComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryComparisonContext {
	var p = new(UnaryComparisonContext)

	p.Unary_comparisonContext = NewEmptyUnary_comparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Unary_comparisonContext))

	return p
}

func (s *UnaryComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryComparisonContext) Endpoint() IEndpointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndpointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndpointContext)
}

func (s *UnaryComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESS, 0)
}

func (s *UnaryComparisonContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESSEQ, 0)
}

func (s *UnaryComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATER, 0)
}

func (s *UnaryComparisonContext) GREATEREQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATEREQ, 0)
}

func (s *UnaryComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterUnaryComparison(s)
	}
}

func (s *UnaryComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitUnaryComparison(s)
	}
}

func (s *UnaryComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitUnaryComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Unary_comparison() (localctx IUnary_comparisonContext) {
	this := p
	_ = this

	localctx = NewUnary_comparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SFeelParserRULE_unary_comparison)
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ:
		localctx = NewUnaryComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserLESS)|(1<<SFeelParserLESSEQ)|(1<<SFeelParserGREATER)|(1<<SFeelParserGREATEREQ))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(81)
			p.Endpoint()
		}

	case SFeelParserT__8, SFeelParserT__9, SFeelParserT__10, SFeelParserT__11, SFeelParserT__12, SFeelParserT__13, SFeelParserName, SFeelParserSUB, SFeelParserINTEGER, SFeelParserREAL, SFeelParserSTRING:
		localctx = NewEqualUnaryComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Endpoint()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntervalContext is an interface to support dynamic dispatch.
type IIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalContext differentiates from other interfaces.
	IsIntervalContext()
}

type IntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalContext() *IntervalContext {
	var p = new(IntervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_interval
	return p
}

func (*IntervalContext) IsIntervalContext() {}

func NewIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalContext {
	var p = new(IntervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_interval

	return p
}

func (s *IntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalContext) AllEndpoint() []IEndpointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEndpointContext)(nil)).Elem())
	var tst = make([]IEndpointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEndpointContext)
		}
	}

	return tst
}

func (s *IntervalContext) Endpoint(i int) IEndpointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndpointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEndpointContext)
}

func (s *IntervalContext) Open_interval_start() IOpen_interval_startContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpen_interval_startContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpen_interval_startContext)
}

func (s *IntervalContext) Closed_interval_start() IClosed_interval_startContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClosed_interval_startContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClosed_interval_startContext)
}

func (s *IntervalContext) Open_interval_end() IOpen_interval_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpen_interval_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpen_interval_endContext)
}

func (s *IntervalContext) Closed_interval_end() IClosed_interval_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClosed_interval_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClosed_interval_endContext)
}

func (s *IntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterInterval(s)
	}
}

func (s *IntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitInterval(s)
	}
}

func (s *IntervalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitInterval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Interval() (localctx IIntervalContext) {
	this := p
	_ = this

	localctx = NewIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SFeelParserRULE_interval)

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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserT__4, SFeelParserT__5:
		{
			p.SetState(85)
			p.Open_interval_start()
		}

	case SFeelParserT__6:
		{
			p.SetState(86)
			p.Closed_interval_start()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(89)
		p.Endpoint()
	}
	{
		p.SetState(90)
		p.Match(SFeelParserT__3)
	}
	{
		p.SetState(91)
		p.Endpoint()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserT__1, SFeelParserT__6:
		{
			p.SetState(92)
			p.Open_interval_end()
		}

	case SFeelParserT__5:
		{
			p.SetState(93)
			p.Closed_interval_end()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOpen_interval_startContext is an interface to support dynamic dispatch.
type IOpen_interval_startContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpen_interval_startContext differentiates from other interfaces.
	IsOpen_interval_startContext()
}

type Open_interval_startContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpen_interval_startContext() *Open_interval_startContext {
	var p = new(Open_interval_startContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_open_interval_start
	return p
}

func (*Open_interval_startContext) IsOpen_interval_startContext() {}

func NewOpen_interval_startContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Open_interval_startContext {
	var p = new(Open_interval_startContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_open_interval_start

	return p
}

func (s *Open_interval_startContext) GetParser() antlr.Parser { return s.parser }
func (s *Open_interval_startContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Open_interval_startContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Open_interval_startContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterOpen_interval_start(s)
	}
}

func (s *Open_interval_startContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitOpen_interval_start(s)
	}
}

func (s *Open_interval_startContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitOpen_interval_start(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Open_interval_start() (localctx IOpen_interval_startContext) {
	this := p
	_ = this

	localctx = NewOpen_interval_startContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SFeelParserRULE_open_interval_start)
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
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserT__4 || _la == SFeelParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IClosed_interval_startContext is an interface to support dynamic dispatch.
type IClosed_interval_startContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClosed_interval_startContext differentiates from other interfaces.
	IsClosed_interval_startContext()
}

type Closed_interval_startContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosed_interval_startContext() *Closed_interval_startContext {
	var p = new(Closed_interval_startContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_closed_interval_start
	return p
}

func (*Closed_interval_startContext) IsClosed_interval_startContext() {}

func NewClosed_interval_startContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Closed_interval_startContext {
	var p = new(Closed_interval_startContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_closed_interval_start

	return p
}

func (s *Closed_interval_startContext) GetParser() antlr.Parser { return s.parser }
func (s *Closed_interval_startContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Closed_interval_startContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Closed_interval_startContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterClosed_interval_start(s)
	}
}

func (s *Closed_interval_startContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitClosed_interval_start(s)
	}
}

func (s *Closed_interval_startContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitClosed_interval_start(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Closed_interval_start() (localctx IClosed_interval_startContext) {
	this := p
	_ = this

	localctx = NewClosed_interval_startContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SFeelParserRULE_closed_interval_start)

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
		p.SetState(98)
		p.Match(SFeelParserT__6)
	}

	return localctx
}

// IOpen_interval_endContext is an interface to support dynamic dispatch.
type IOpen_interval_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpen_interval_endContext differentiates from other interfaces.
	IsOpen_interval_endContext()
}

type Open_interval_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpen_interval_endContext() *Open_interval_endContext {
	var p = new(Open_interval_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_open_interval_end
	return p
}

func (*Open_interval_endContext) IsOpen_interval_endContext() {}

func NewOpen_interval_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Open_interval_endContext {
	var p = new(Open_interval_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_open_interval_end

	return p
}

func (s *Open_interval_endContext) GetParser() antlr.Parser { return s.parser }
func (s *Open_interval_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Open_interval_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Open_interval_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterOpen_interval_end(s)
	}
}

func (s *Open_interval_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitOpen_interval_end(s)
	}
}

func (s *Open_interval_endContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitOpen_interval_end(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Open_interval_end() (localctx IOpen_interval_endContext) {
	this := p
	_ = this

	localctx = NewOpen_interval_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SFeelParserRULE_open_interval_end)
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
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserT__1 || _la == SFeelParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IClosed_interval_endContext is an interface to support dynamic dispatch.
type IClosed_interval_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClosed_interval_endContext differentiates from other interfaces.
	IsClosed_interval_endContext()
}

type Closed_interval_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosed_interval_endContext() *Closed_interval_endContext {
	var p = new(Closed_interval_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_closed_interval_end
	return p
}

func (*Closed_interval_endContext) IsClosed_interval_endContext() {}

func NewClosed_interval_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Closed_interval_endContext {
	var p = new(Closed_interval_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_closed_interval_end

	return p
}

func (s *Closed_interval_endContext) GetParser() antlr.Parser { return s.parser }
func (s *Closed_interval_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Closed_interval_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Closed_interval_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterClosed_interval_end(s)
	}
}

func (s *Closed_interval_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitClosed_interval_end(s)
	}
}

func (s *Closed_interval_endContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitClosed_interval_end(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Closed_interval_end() (localctx IClosed_interval_endContext) {
	this := p
	_ = this

	localctx = NewClosed_interval_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SFeelParserRULE_closed_interval_end)

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
		p.SetState(102)
		p.Match(SFeelParserT__5)
	}

	return localctx
}

// IEmpty_expressionContext is an interface to support dynamic dispatch.
type IEmpty_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmpty_expressionContext differentiates from other interfaces.
	IsEmpty_expressionContext()
}

type Empty_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmpty_expressionContext() *Empty_expressionContext {
	var p = new(Empty_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_empty_expression
	return p
}

func (*Empty_expressionContext) IsEmpty_expressionContext() {}

func NewEmpty_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Empty_expressionContext {
	var p = new(Empty_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_empty_expression

	return p
}

func (s *Empty_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Empty_expressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(SFeelParserSUB, 0)
}

func (s *Empty_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Empty_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Empty_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEmpty_expression(s)
	}
}

func (s *Empty_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEmpty_expression(s)
	}
}

func (s *Empty_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitEmpty_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Empty_expression() (localctx IEmpty_expressionContext) {
	this := p
	_ = this

	localctx = NewEmpty_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SFeelParserRULE_empty_expression)

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
		p.SetState(104)
		p.Match(SFeelParserSUB)
	}

	return localctx
}

// ISimple_expressionsContext is an interface to support dynamic dispatch.
type ISimple_expressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_expressionsContext differentiates from other interfaces.
	IsSimple_expressionsContext()
}

type Simple_expressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_expressionsContext() *Simple_expressionsContext {
	var p = new(Simple_expressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_expressions
	return p
}

func (*Simple_expressionsContext) IsSimple_expressionsContext() {}

func NewSimple_expressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_expressionsContext {
	var p = new(Simple_expressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_expressions

	return p
}

func (s *Simple_expressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_expressionsContext) CopyFrom(ctx *Simple_expressionsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Simple_expressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_expressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleExpressionsContext struct {
	*Simple_expressionsContext
}

func NewSimpleExpressionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleExpressionsContext {
	var p = new(SimpleExpressionsContext)

	p.Simple_expressionsContext = NewEmptySimple_expressionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Simple_expressionsContext))

	return p
}

func (s *SimpleExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExpressionsContext) AllSimple_expression() []ISimple_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem())
	var tst = make([]ISimple_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimple_expressionContext)
		}
	}

	return tst
}

func (s *SimpleExpressionsContext) Simple_expression(i int) ISimple_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionContext)
}

func (s *SimpleExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimpleExpressions(s)
	}
}

func (s *SimpleExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimpleExpressions(s)
	}
}

func (s *SimpleExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimpleExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptySimpleExpressionsContext struct {
	*Simple_expressionsContext
}

func NewEmptySimpleExpressionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptySimpleExpressionsContext {
	var p = new(EmptySimpleExpressionsContext)

	p.Simple_expressionsContext = NewEmptySimple_expressionsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Simple_expressionsContext))

	return p
}

func (s *EmptySimpleExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptySimpleExpressionsContext) Empty_expression() IEmpty_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmpty_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmpty_expressionContext)
}

func (s *EmptySimpleExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEmptySimpleExpressions(s)
	}
}

func (s *EmptySimpleExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEmptySimpleExpressions(s)
	}
}

func (s *EmptySimpleExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitEmptySimpleExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_expressions() (localctx ISimple_expressionsContext) {
	this := p
	_ = this

	localctx = NewSimple_expressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SFeelParserRULE_simple_expressions)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleExpressionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Simple_expression()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SFeelParserT__2 {
			{
				p.SetState(107)
				p.Match(SFeelParserT__2)
			}
			{
				p.SetState(108)
				p.Simple_expression()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewEmptySimpleExpressionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Empty_expression()
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Simple_expression() ISimple_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SFeelParserRULE_expression)

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
		p.SetState(117)
		p.Simple_expression()
	}

	return localctx
}

// ISimple_expressionContext is an interface to support dynamic dispatch.
type ISimple_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_expressionContext differentiates from other interfaces.
	IsSimple_expressionContext()
}

type Simple_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_expressionContext() *Simple_expressionContext {
	var p = new(Simple_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_expression
	return p
}

func (*Simple_expressionContext) IsSimple_expressionContext() {}

func NewSimple_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_expressionContext {
	var p = new(Simple_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_expression

	return p
}

func (s *Simple_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_expressionContext) Simple_value() ISimple_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_valueContext)
}

func (s *Simple_expressionContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *Simple_expressionContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *Simple_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimple_expression(s)
	}
}

func (s *Simple_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimple_expression(s)
	}
}

func (s *Simple_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimple_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_expression() (localctx ISimple_expressionContext) {
	this := p
	_ = this

	localctx = NewSimple_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SFeelParserRULE_simple_expression)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Simple_value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.arithmetic_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.comparison(0)
		}

	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
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

func (s *ComparisonContext) GetOperator() antlr.Token { return s.operator }

func (s *ComparisonContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *ComparisonContext) Simple_value() ISimple_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_valueContext)
}

func (s *ComparisonContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESS, 0)
}

func (s *ComparisonContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESSEQ, 0)
}

func (s *ComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATER, 0)
}

func (s *ComparisonContext) GREATEREQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATEREQ, 0)
}

func (s *ComparisonContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SFeelParserEQUAL, 0)
}

func (s *ComparisonContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(SFeelParserNOTEQUAL, 0)
}

func (s *ComparisonContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *ComparisonContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
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

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Comparison() (localctx IComparisonContext) {
	return p.comparison(0)
}

func (p *SFeelParser) comparison(_p int) (localctx IComparisonContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparisonContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, SFeelParserRULE_comparison, _p)
	var _la int

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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(125)
			p.Simple_value()
		}
		{
			p.SetState(126)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ComparisonContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserLESS)|(1<<SFeelParserLESSEQ)|(1<<SFeelParserGREATER)|(1<<SFeelParserGREATEREQ)|(1<<SFeelParserEQUAL)|(1<<SFeelParserNOTEQUAL))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ComparisonContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(127)
			p.Expression()
		}

	case 2:
		{
			p.SetState(129)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(130)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ComparisonContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserLESS)|(1<<SFeelParserLESSEQ)|(1<<SFeelParserGREATER)|(1<<SFeelParserGREATEREQ)|(1<<SFeelParserEQUAL)|(1<<SFeelParserNOTEQUAL))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ComparisonContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(131)
			p.Expression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparisonContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_comparison)
			p.SetState(135)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(136)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ComparisonContext).operator = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserLESS)|(1<<SFeelParserLESSEQ)|(1<<SFeelParserGREATER)|(1<<SFeelParserGREATEREQ)|(1<<SFeelParserEQUAL)|(1<<SFeelParserNOTEQUAL))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ComparisonContext).operator = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(137)
				p.Expression()
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IArithmetic_expressionContext is an interface to support dynamic dispatch.
type IArithmetic_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithmetic_expressionContext differentiates from other interfaces.
	IsArithmetic_expressionContext()
}

type Arithmetic_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmetic_expressionContext() *Arithmetic_expressionContext {
	var p = new(Arithmetic_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_arithmetic_expression
	return p
}

func (*Arithmetic_expressionContext) IsArithmetic_expressionContext() {}

func NewArithmetic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithmetic_expressionContext {
	var p = new(Arithmetic_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_arithmetic_expression

	return p
}

func (s *Arithmetic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithmetic_expressionContext) CopyFrom(ctx *Arithmetic_expressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Arithmetic_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AdditionOrSubtractionContext struct {
	*Arithmetic_expressionContext
	operator antlr.Token
}

func NewAdditionOrSubtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionOrSubtractionContext {
	var p = new(AdditionOrSubtractionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *AdditionOrSubtractionContext) GetOperator() antlr.Token { return s.operator }

func (s *AdditionOrSubtractionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *AdditionOrSubtractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionOrSubtractionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem())
	var tst = make([]IArithmetic_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArithmetic_expressionContext)
		}
	}

	return tst
}

func (s *AdditionOrSubtractionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *AdditionOrSubtractionContext) ADD() antlr.TerminalNode {
	return s.GetToken(SFeelParserADD, 0)
}

func (s *AdditionOrSubtractionContext) SUB() antlr.TerminalNode {
	return s.GetToken(SFeelParserSUB, 0)
}

func (s *AdditionOrSubtractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterAdditionOrSubtraction(s)
	}
}

func (s *AdditionOrSubtractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitAdditionOrSubtraction(s)
	}
}

func (s *AdditionOrSubtractionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitAdditionOrSubtraction(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueContext struct {
	*Arithmetic_expressionContext
}

func NewValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContext {
	var p = new(ValueContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) Simple_value() ISimple_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_valueContext)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationOrDivisionContext struct {
	*Arithmetic_expressionContext
	operator antlr.Token
}

func NewMultiplicationOrDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationOrDivisionContext {
	var p = new(MultiplicationOrDivisionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *MultiplicationOrDivisionContext) GetOperator() antlr.Token { return s.operator }

func (s *MultiplicationOrDivisionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *MultiplicationOrDivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationOrDivisionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem())
	var tst = make([]IArithmetic_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArithmetic_expressionContext)
		}
	}

	return tst
}

func (s *MultiplicationOrDivisionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *MultiplicationOrDivisionContext) MUL() antlr.TerminalNode {
	return s.GetToken(SFeelParserMUL, 0)
}

func (s *MultiplicationOrDivisionContext) DIV() antlr.TerminalNode {
	return s.GetToken(SFeelParserDIV, 0)
}

func (s *MultiplicationOrDivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterMultiplicationOrDivision(s)
	}
}

func (s *MultiplicationOrDivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitMultiplicationOrDivision(s)
	}
}

func (s *MultiplicationOrDivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitMultiplicationOrDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeticNegationContext struct {
	*Arithmetic_expressionContext
}

func NewArithmeticNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticNegationContext {
	var p = new(ArithmeticNegationContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *ArithmeticNegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticNegationContext) SUB() antlr.TerminalNode {
	return s.GetToken(SFeelParserSUB, 0)
}

func (s *ArithmeticNegationContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *ArithmeticNegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterArithmeticNegation(s)
	}
}

func (s *ArithmeticNegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitArithmeticNegation(s)
	}
}

func (s *ArithmeticNegationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitArithmeticNegation(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesesContext struct {
	*Arithmetic_expressionContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerContext struct {
	*Arithmetic_expressionContext
	operator antlr.Token
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *PowerContext) GetOperator() antlr.Token { return s.operator }

func (s *PowerContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem())
	var tst = make([]IArithmetic_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArithmetic_expressionContext)
		}
	}

	return tst
}

func (s *PowerContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *PowerContext) POW() antlr.TerminalNode {
	return s.GetToken(SFeelParserPOW, 0)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Arithmetic_expression() (localctx IArithmetic_expressionContext) {
	return p.arithmetic_expression(0)
}

func (p *SFeelParser) arithmetic_expression(_p int) (localctx IArithmetic_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArithmetic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithmetic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, SFeelParserRULE_arithmetic_expression, _p)
	var _la int

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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(144)
			p.Simple_value()
		}

	case 2:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(145)
			p.Match(SFeelParserT__4)
		}
		{
			p.SetState(146)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(147)
			p.Match(SFeelParserT__1)
		}

	case 3:
		localctx = NewArithmeticNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(149)
			p.Match(SFeelParserSUB)
		}
		{
			p.SetState(150)
			p.arithmetic_expression(4)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_arithmetic_expression)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(154)

					var _m = p.Match(SFeelParserPOW)

					localctx.(*PowerContext).operator = _m
				}
				{
					p.SetState(155)
					p.arithmetic_expression(4)
				}

			case 2:
				localctx = NewMultiplicationOrDivisionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_arithmetic_expression)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(157)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicationOrDivisionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SFeelParserMUL || _la == SFeelParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicationOrDivisionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(158)
					p.arithmetic_expression(3)
				}

			case 3:
				localctx = NewAdditionOrSubtractionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_arithmetic_expression)
				p.SetState(159)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(160)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditionOrSubtractionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SFeelParserADD || _la == SFeelParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditionOrSubtractionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(161)
					p.arithmetic_expression(2)
				}

			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IEndpointContext is an interface to support dynamic dispatch.
type IEndpointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndpointContext differentiates from other interfaces.
	IsEndpointContext()
}

type EndpointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointContext() *EndpointContext {
	var p = new(EndpointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_endpoint
	return p
}

func (*EndpointContext) IsEndpointContext() {}

func NewEndpointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointContext {
	var p = new(EndpointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_endpoint

	return p
}

func (s *EndpointContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointContext) Qualified_name() IQualified_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_nameContext)
}

func (s *EndpointContext) Simple_literal() ISimple_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_literalContext)
}

func (s *EndpointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEndpoint(s)
	}
}

func (s *EndpointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEndpoint(s)
	}
}

func (s *EndpointContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitEndpoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Endpoint() (localctx IEndpointContext) {
	this := p
	_ = this

	localctx = NewEndpointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SFeelParserRULE_endpoint)

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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Qualified_name()
		}

	case SFeelParserT__8, SFeelParserT__9, SFeelParserT__10, SFeelParserT__11, SFeelParserT__12, SFeelParserT__13, SFeelParserSUB, SFeelParserINTEGER, SFeelParserREAL, SFeelParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Simple_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_valueContext is an interface to support dynamic dispatch.
type ISimple_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_valueContext differentiates from other interfaces.
	IsSimple_valueContext()
}

type Simple_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_valueContext() *Simple_valueContext {
	var p = new(Simple_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_value
	return p
}

func (*Simple_valueContext) IsSimple_valueContext() {}

func NewSimple_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_valueContext {
	var p = new(Simple_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_value

	return p
}

func (s *Simple_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_valueContext) Qualified_name() IQualified_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_nameContext)
}

func (s *Simple_valueContext) Simple_literal() ISimple_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_literalContext)
}

func (s *Simple_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimple_value(s)
	}
}

func (s *Simple_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimple_value(s)
	}
}

func (s *Simple_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimple_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_value() (localctx ISimple_valueContext) {
	this := p
	_ = this

	localctx = NewSimple_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SFeelParserRULE_simple_value)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Qualified_name()
		}

	case SFeelParserT__8, SFeelParserT__9, SFeelParserT__10, SFeelParserT__11, SFeelParserT__12, SFeelParserT__13, SFeelParserSUB, SFeelParserINTEGER, SFeelParserREAL, SFeelParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Simple_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualified_nameContext is an interface to support dynamic dispatch.
type IQualified_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualified_nameContext differentiates from other interfaces.
	IsQualified_nameContext()
}

type Qualified_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualified_nameContext() *Qualified_nameContext {
	var p = new(Qualified_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_qualified_name
	return p
}

func (*Qualified_nameContext) IsQualified_nameContext() {}

func NewQualified_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Qualified_nameContext {
	var p = new(Qualified_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_qualified_name

	return p
}

func (s *Qualified_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Qualified_nameContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(SFeelParserName)
}

func (s *Qualified_nameContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(SFeelParserName, i)
}

func (s *Qualified_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qualified_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Qualified_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterQualified_name(s)
	}
}

func (s *Qualified_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitQualified_name(s)
	}
}

func (s *Qualified_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitQualified_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Qualified_name() (localctx IQualified_nameContext) {
	this := p
	_ = this

	localctx = NewQualified_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SFeelParserRULE_qualified_name)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(SFeelParserName)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(176)
				p.Match(SFeelParserT__7)
			}
			{
				p.SetState(177)
				p.Match(SFeelParserName)
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// ISimple_literalContext is an interface to support dynamic dispatch.
type ISimple_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_literalContext differentiates from other interfaces.
	IsSimple_literalContext()
}

type Simple_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_literalContext() *Simple_literalContext {
	var p = new(Simple_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_simple_literal
	return p
}

func (*Simple_literalContext) IsSimple_literalContext() {}

func NewSimple_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_literalContext {
	var p = new(Simple_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_simple_literal

	return p
}

func (s *Simple_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_literalContext) Numeric_literal() INumeric_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_literalContext)
}

func (s *Simple_literalContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Simple_literalContext) Boolean_literal() IBoolean_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_literalContext)
}

func (s *Simple_literalContext) Date_time_literal() IDate_time_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_time_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_time_literalContext)
}

func (s *Simple_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterSimple_literal(s)
	}
}

func (s *Simple_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitSimple_literal(s)
	}
}

func (s *Simple_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitSimple_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Simple_literal() (localctx ISimple_literalContext) {
	this := p
	_ = this

	localctx = NewSimple_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SFeelParserRULE_simple_literal)

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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserSUB, SFeelParserINTEGER, SFeelParserREAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Numeric_literal()
		}

	case SFeelParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.String_literal()
		}

	case SFeelParserT__12, SFeelParserT__13:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.Boolean_literal()
		}

	case SFeelParserT__8, SFeelParserT__9, SFeelParserT__10, SFeelParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(186)
			p.Date_time_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDate_time_literalContext is an interface to support dynamic dispatch.
type IDate_time_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_time_literalContext differentiates from other interfaces.
	IsDate_time_literalContext()
}

type Date_time_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_time_literalContext() *Date_time_literalContext {
	var p = new(Date_time_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_date_time_literal
	return p
}

func (*Date_time_literalContext) IsDate_time_literalContext() {}

func NewDate_time_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_time_literalContext {
	var p = new(Date_time_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_date_time_literal

	return p
}

func (s *Date_time_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_time_literalContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Date_time_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_time_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_time_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDate_time_literal(s)
	}
}

func (s *Date_time_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDate_time_literal(s)
	}
}

func (s *Date_time_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitDate_time_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Date_time_literal() (localctx IDate_time_literalContext) {
	this := p
	_ = this

	localctx = NewDate_time_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SFeelParserRULE_date_time_literal)
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
		p.SetState(189)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserT__8)|(1<<SFeelParserT__9)|(1<<SFeelParserT__10)|(1<<SFeelParserT__11))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(190)
		p.Match(SFeelParserT__4)
	}
	{
		p.SetState(191)
		p.String_literal()
	}
	{
		p.SetState(192)
		p.Match(SFeelParserT__1)
	}

	return localctx
}

// INumeric_literalContext is an interface to support dynamic dispatch.
type INumeric_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumeric_literalContext differentiates from other interfaces.
	IsNumeric_literalContext()
}

type Numeric_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_literalContext() *Numeric_literalContext {
	var p = new(Numeric_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_numeric_literal
	return p
}

func (*Numeric_literalContext) IsNumeric_literalContext() {}

func NewNumeric_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_literalContext {
	var p = new(Numeric_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_numeric_literal

	return p
}

func (s *Numeric_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_literalContext) Integer_literal() IInteger_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *Numeric_literalContext) Real_literal() IReal_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReal_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReal_literalContext)
}

func (s *Numeric_literalContext) SUB() antlr.TerminalNode {
	return s.GetToken(SFeelParserSUB, 0)
}

func (s *Numeric_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNumeric_literal(s)
	}
}

func (s *Numeric_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNumeric_literal(s)
	}
}

func (s *Numeric_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitNumeric_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Numeric_literal() (localctx INumeric_literalContext) {
	this := p
	_ = this

	localctx = NewNumeric_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SFeelParserRULE_numeric_literal)
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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SFeelParserSUB {
		{
			p.SetState(194)
			p.Match(SFeelParserSUB)
		}

	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER:
		{
			p.SetState(197)
			p.Integer_literal()
		}

	case SFeelParserREAL:
		{
			p.SetState(198)
			p.Real_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInteger_literalContext is an interface to support dynamic dispatch.
type IInteger_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_literalContext differentiates from other interfaces.
	IsInteger_literalContext()
}

type Integer_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_literalContext() *Integer_literalContext {
	var p = new(Integer_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_integer_literal
	return p
}

func (*Integer_literalContext) IsInteger_literalContext() {}

func NewInteger_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_literalContext {
	var p = new(Integer_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_integer_literal

	return p
}

func (s *Integer_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_literalContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *Integer_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterInteger_literal(s)
	}
}

func (s *Integer_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitInteger_literal(s)
	}
}

func (s *Integer_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitInteger_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Integer_literal() (localctx IInteger_literalContext) {
	this := p
	_ = this

	localctx = NewInteger_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SFeelParserRULE_integer_literal)

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
		p.SetState(201)
		p.Match(SFeelParserINTEGER)
	}

	return localctx
}

// IReal_literalContext is an interface to support dynamic dispatch.
type IReal_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReal_literalContext differentiates from other interfaces.
	IsReal_literalContext()
}

type Real_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_literalContext() *Real_literalContext {
	var p = new(Real_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_real_literal
	return p
}

func (*Real_literalContext) IsReal_literalContext() {}

func NewReal_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_literalContext {
	var p = new(Real_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_real_literal

	return p
}

func (s *Real_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_literalContext) REAL() antlr.TerminalNode {
	return s.GetToken(SFeelParserREAL, 0)
}

func (s *Real_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterReal_literal(s)
	}
}

func (s *Real_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitReal_literal(s)
	}
}

func (s *Real_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitReal_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Real_literal() (localctx IReal_literalContext) {
	this := p
	_ = this

	localctx = NewReal_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SFeelParserRULE_real_literal)

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
		p.SetState(203)
		p.Match(SFeelParserREAL)
	}

	return localctx
}

// IBoolean_literalContext is an interface to support dynamic dispatch.
type IBoolean_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolean_literalContext differentiates from other interfaces.
	IsBoolean_literalContext()
}

type Boolean_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_literalContext() *Boolean_literalContext {
	var p = new(Boolean_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_boolean_literal
	return p
}

func (*Boolean_literalContext) IsBoolean_literalContext() {}

func NewBoolean_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_literalContext {
	var p = new(Boolean_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_boolean_literal

	return p
}

func (s *Boolean_literalContext) GetParser() antlr.Parser { return s.parser }
func (s *Boolean_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterBoolean_literal(s)
	}
}

func (s *Boolean_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitBoolean_literal(s)
	}
}

func (s *Boolean_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitBoolean_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) Boolean_literal() (localctx IBoolean_literalContext) {
	this := p
	_ = this

	localctx = NewBoolean_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SFeelParserRULE_boolean_literal)
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
		p.SetState(205)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserT__12 || _la == SFeelParserT__13) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_string_literal
	return p
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(SFeelParserSTRING)
}

func (s *String_literalContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(SFeelParserSTRING, i)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (s *String_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SFeelVisitor:
		return t.VisitString_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SFeelParser) String_literal() (localctx IString_literalContext) {
	this := p
	_ = this

	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SFeelParserRULE_string_literal)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(207)
				p.Match(SFeelParserSTRING)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

func (p *SFeelParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ComparisonContext = nil
		if localctx != nil {
			t = localctx.(*ComparisonContext)
		}
		return p.Comparison_Sempred(t, predIndex)

	case 16:
		var t *Arithmetic_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Arithmetic_expressionContext)
		}
		return p.Arithmetic_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SFeelParser) Comparison_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SFeelParser) Arithmetic_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
