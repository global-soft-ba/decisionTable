// Code generated from SFeel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SFeel

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SFeelParser struct {
	*antlr.BaseParser
}

var sfeelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sfeelParserInit() {
	staticData := &sfeelParserStaticData
	staticData.literalNames = []string{
		"", "'not('", "')'", "','", "'..'", "'('", "']'", "'['", "'.'", "'date'",
		"'time'", "'date and time'", "'duration'", "'true'", "'false'", "",
		"'**'", "'*'", "'/'", "'+'", "'-'", "'<'", "'<='", "'>'", "'>='", "'='",
		"'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Name",
		"POW", "MUL", "DIV", "ADD", "SUB", "LESS", "LESSEQ", "GREATER", "GREATEREQ",
		"EQUAL", "NOTEQUAL", "INTEGER", "REAL", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"input", "output", "simple_unary_tests", "simple_positive_unary_tests",
		"simple_positive_unary_test", "unary_comparison", "interval", "open_interval_start",
		"closed_interval_start", "open_interval_end", "closed_interval_end",
		"empty_expression", "simple_expressions", "expression", "simple_expression",
		"comparison", "arithmetic_expression", "endpoint", "simple_value", "qualified_name",
		"simple_literal", "date_time_literal", "numeric_literal", "integer_literal",
		"real_literal", "boolean_literal", "string_literal",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 213, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 67, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 72, 8, 3, 10, 3, 12, 3, 75,
		9, 3, 1, 4, 1, 4, 3, 4, 79, 8, 4, 1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1,
		6, 1, 6, 3, 6, 88, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 95, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 5, 12, 110, 8, 12, 10, 12, 12, 12, 113, 9, 12, 1, 12, 3,
		12, 116, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 123, 8, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 134,
		8, 15, 1, 15, 1, 15, 1, 15, 5, 15, 139, 8, 15, 10, 15, 12, 15, 142, 9,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 152,
		8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 163, 8, 16, 10, 16, 12, 16, 166, 9, 16, 1, 17, 1, 17, 3, 17, 170, 8,
		17, 1, 18, 1, 18, 3, 18, 174, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 179, 8,
		19, 10, 19, 12, 19, 182, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 188,
		8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 3, 22, 196, 8, 22, 1,
		22, 1, 22, 3, 22, 200, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 4, 26, 209, 8, 26, 11, 26, 12, 26, 210, 1, 26, 0, 2, 30, 32, 27,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 0, 8, 1, 0, 21, 24, 1, 0, 5, 6, 2, 0, 2,
		2, 7, 7, 1, 0, 21, 26, 1, 0, 17, 18, 1, 0, 19, 20, 1, 0, 9, 12, 1, 0, 13,
		14, 212, 0, 54, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 66, 1, 0, 0, 0, 6, 68,
		1, 0, 0, 0, 8, 78, 1, 0, 0, 0, 10, 83, 1, 0, 0, 0, 12, 87, 1, 0, 0, 0,
		14, 96, 1, 0, 0, 0, 16, 98, 1, 0, 0, 0, 18, 100, 1, 0, 0, 0, 20, 102, 1,
		0, 0, 0, 22, 104, 1, 0, 0, 0, 24, 115, 1, 0, 0, 0, 26, 117, 1, 0, 0, 0,
		28, 122, 1, 0, 0, 0, 30, 133, 1, 0, 0, 0, 32, 151, 1, 0, 0, 0, 34, 169,
		1, 0, 0, 0, 36, 173, 1, 0, 0, 0, 38, 175, 1, 0, 0, 0, 40, 187, 1, 0, 0,
		0, 42, 189, 1, 0, 0, 0, 44, 195, 1, 0, 0, 0, 46, 201, 1, 0, 0, 0, 48, 203,
		1, 0, 0, 0, 50, 205, 1, 0, 0, 0, 52, 208, 1, 0, 0, 0, 54, 55, 3, 4, 2,
		0, 55, 56, 5, 0, 0, 1, 56, 1, 1, 0, 0, 0, 57, 58, 3, 24, 12, 0, 58, 59,
		5, 0, 0, 1, 59, 3, 1, 0, 0, 0, 60, 67, 3, 6, 3, 0, 61, 62, 5, 1, 0, 0,
		62, 63, 3, 6, 3, 0, 63, 64, 5, 2, 0, 0, 64, 67, 1, 0, 0, 0, 65, 67, 3,
		22, 11, 0, 66, 60, 1, 0, 0, 0, 66, 61, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0,
		67, 5, 1, 0, 0, 0, 68, 73, 3, 8, 4, 0, 69, 70, 5, 3, 0, 0, 70, 72, 3, 8,
		4, 0, 71, 69, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74,
		1, 0, 0, 0, 74, 7, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 79, 3, 10, 5, 0,
		77, 79, 3, 12, 6, 0, 78, 76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 9, 1,
		0, 0, 0, 80, 81, 7, 0, 0, 0, 81, 84, 3, 34, 17, 0, 82, 84, 3, 34, 17, 0,
		83, 80, 1, 0, 0, 0, 83, 82, 1, 0, 0, 0, 84, 11, 1, 0, 0, 0, 85, 88, 3,
		14, 7, 0, 86, 88, 3, 16, 8, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0,
		88, 89, 1, 0, 0, 0, 89, 90, 3, 34, 17, 0, 90, 91, 5, 4, 0, 0, 91, 94, 3,
		34, 17, 0, 92, 95, 3, 18, 9, 0, 93, 95, 3, 20, 10, 0, 94, 92, 1, 0, 0,
		0, 94, 93, 1, 0, 0, 0, 95, 13, 1, 0, 0, 0, 96, 97, 7, 1, 0, 0, 97, 15,
		1, 0, 0, 0, 98, 99, 5, 7, 0, 0, 99, 17, 1, 0, 0, 0, 100, 101, 7, 2, 0,
		0, 101, 19, 1, 0, 0, 0, 102, 103, 5, 6, 0, 0, 103, 21, 1, 0, 0, 0, 104,
		105, 5, 20, 0, 0, 105, 23, 1, 0, 0, 0, 106, 111, 3, 28, 14, 0, 107, 108,
		5, 3, 0, 0, 108, 110, 3, 28, 14, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1,
		0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 116, 1, 0, 0,
		0, 113, 111, 1, 0, 0, 0, 114, 116, 3, 22, 11, 0, 115, 106, 1, 0, 0, 0,
		115, 114, 1, 0, 0, 0, 116, 25, 1, 0, 0, 0, 117, 118, 3, 28, 14, 0, 118,
		27, 1, 0, 0, 0, 119, 123, 3, 36, 18, 0, 120, 123, 3, 32, 16, 0, 121, 123,
		3, 30, 15, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 121, 1,
		0, 0, 0, 123, 29, 1, 0, 0, 0, 124, 125, 6, 15, -1, 0, 125, 126, 3, 36,
		18, 0, 126, 127, 7, 3, 0, 0, 127, 128, 3, 26, 13, 0, 128, 134, 1, 0, 0,
		0, 129, 130, 3, 32, 16, 0, 130, 131, 7, 3, 0, 0, 131, 132, 3, 26, 13, 0,
		132, 134, 1, 0, 0, 0, 133, 124, 1, 0, 0, 0, 133, 129, 1, 0, 0, 0, 134,
		140, 1, 0, 0, 0, 135, 136, 10, 1, 0, 0, 136, 137, 7, 3, 0, 0, 137, 139,
		3, 26, 13, 0, 138, 135, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1,
		0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 31, 1, 0, 0, 0, 142, 140, 1, 0, 0,
		0, 143, 144, 6, 16, -1, 0, 144, 152, 3, 36, 18, 0, 145, 146, 5, 5, 0, 0,
		146, 147, 3, 32, 16, 0, 147, 148, 5, 2, 0, 0, 148, 152, 1, 0, 0, 0, 149,
		150, 5, 20, 0, 0, 150, 152, 3, 32, 16, 4, 151, 143, 1, 0, 0, 0, 151, 145,
		1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 164, 1, 0, 0, 0, 153, 154, 10, 3,
		0, 0, 154, 155, 5, 16, 0, 0, 155, 163, 3, 32, 16, 4, 156, 157, 10, 2, 0,
		0, 157, 158, 7, 4, 0, 0, 158, 163, 3, 32, 16, 3, 159, 160, 10, 1, 0, 0,
		160, 161, 7, 5, 0, 0, 161, 163, 3, 32, 16, 2, 162, 153, 1, 0, 0, 0, 162,
		156, 1, 0, 0, 0, 162, 159, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 33, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 167, 170, 3, 38, 19, 0, 168, 170, 3, 40, 20, 0, 169, 167, 1, 0, 0,
		0, 169, 168, 1, 0, 0, 0, 170, 35, 1, 0, 0, 0, 171, 174, 3, 38, 19, 0, 172,
		174, 3, 40, 20, 0, 173, 171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 37,
		1, 0, 0, 0, 175, 180, 5, 15, 0, 0, 176, 177, 5, 8, 0, 0, 177, 179, 5, 15,
		0, 0, 178, 176, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0,
		180, 181, 1, 0, 0, 0, 181, 39, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 188,
		3, 44, 22, 0, 184, 188, 3, 52, 26, 0, 185, 188, 3, 50, 25, 0, 186, 188,
		3, 42, 21, 0, 187, 183, 1, 0, 0, 0, 187, 184, 1, 0, 0, 0, 187, 185, 1,
		0, 0, 0, 187, 186, 1, 0, 0, 0, 188, 41, 1, 0, 0, 0, 189, 190, 7, 6, 0,
		0, 190, 191, 5, 5, 0, 0, 191, 192, 3, 52, 26, 0, 192, 193, 5, 2, 0, 0,
		193, 43, 1, 0, 0, 0, 194, 196, 5, 20, 0, 0, 195, 194, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 200, 3, 46, 23, 0, 198, 200,
		3, 48, 24, 0, 199, 197, 1, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 45, 1, 0,
		0, 0, 201, 202, 5, 27, 0, 0, 202, 47, 1, 0, 0, 0, 203, 204, 5, 28, 0, 0,
		204, 49, 1, 0, 0, 0, 205, 206, 7, 7, 0, 0, 206, 51, 1, 0, 0, 0, 207, 209,
		5, 29, 0, 0, 208, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0,
		0, 0, 210, 211, 1, 0, 0, 0, 211, 53, 1, 0, 0, 0, 21, 66, 73, 78, 83, 87,
		94, 111, 115, 122, 133, 140, 151, 162, 164, 169, 173, 180, 187, 195, 199,
		210,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SFeelParserInit initializes any static state used to implement SFeelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSFeelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SFeelParserInit() {
	staticData := &sfeelParserStaticData
	staticData.once.Do(sfeelParserInit)
}

// NewSFeelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSFeelParser(input antlr.TokenStream) *SFeelParser {
	SFeelParserInit()
	this := new(SFeelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sfeelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_unary_testsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_expressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmpty_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_positive_unary_testsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_positive_unary_testsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_positive_unary_testContext); ok {
			len++
		}
	}

	tst := make([]ISimple_positive_unary_testContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_positive_unary_testContext); ok {
			tst[i] = t.(ISimple_positive_unary_testContext)
			i++
		}
	}

	return tst
}

func (s *Simple_positive_unary_testsContext) Simple_positive_unary_test(i int) ISimple_positive_unary_testContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_positive_unary_testContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_comparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_comparisonContext)
}

func (s *Simple_positive_unary_testContext) Interval() IIntervalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndpointContext); ok {
			len++
		}
	}

	tst := make([]IEndpointContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndpointContext); ok {
			tst[i] = t.(IEndpointContext)
			i++
		}
	}

	return tst
}

func (s *IntervalContext) Endpoint(i int) IEndpointContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndpointContext)
}

func (s *IntervalContext) Open_interval_start() IOpen_interval_startContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpen_interval_startContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpen_interval_startContext)
}

func (s *IntervalContext) Closed_interval_start() IClosed_interval_startContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosed_interval_startContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosed_interval_startContext)
}

func (s *IntervalContext) Open_interval_end() IOpen_interval_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpen_interval_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpen_interval_endContext)
}

func (s *IntervalContext) Closed_interval_end() IClosed_interval_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosed_interval_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_expressionContext); ok {
			len++
		}
	}

	tst := make([]ISimple_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_expressionContext); ok {
			tst[i] = t.(ISimple_expressionContext)
			i++
		}
	}

	return tst
}

func (s *SimpleExpressionsContext) Simple_expression(i int) ISimple_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmpty_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_valueContext)
}

func (s *Simple_expressionContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *Simple_expressionContext) Comparison() IComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_valueContext)
}

func (s *ComparisonContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *ComparisonContext) Comparison() IComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditionOrSubtractionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicationOrDivisionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualified_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualified_nameContext)
}

func (s *EndpointContext) Simple_literal() ISimple_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualified_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualified_nameContext)
}

func (s *Simple_valueContext) Simple_literal() ISimple_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumeric_literalContext)
}

func (s *Simple_literalContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Simple_literalContext) Boolean_literal() IBoolean_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_literalContext)
}

func (s *Simple_literalContext) Date_time_literal() IDate_time_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDate_time_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *Numeric_literalContext) Real_literal() IReal_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
