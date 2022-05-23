// Code generated from SFeel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SFeelLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var sfeellexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sfeellexerLexerInit() {
	staticData := &sfeellexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "Name", "Name_start_char",
		"Name_part_char", "POW", "MUL", "DIV", "ADD", "SUB", "LESS", "LESSEQ",
		"GREATER", "GREATEREQ", "EQUAL", "NOTEQUAL", "INTEGER", "REAL", "STRING",
		"NONCONTROL_CHAR", "LETTER", "LOWER", "UPPER", "DIGIT", "SPACE", "SYMBOL",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 239, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 146, 8, 14, 10, 14, 12, 14, 149,
		9, 14, 1, 15, 1, 15, 3, 15, 153, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 158,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 5, 28, 188, 8,
		28, 10, 28, 12, 28, 191, 9, 28, 1, 29, 4, 29, 194, 8, 29, 11, 29, 12, 29,
		195, 1, 29, 1, 29, 5, 29, 200, 8, 29, 10, 29, 12, 29, 203, 9, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 5, 30, 209, 8, 30, 10, 30, 12, 30, 212, 9, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 220, 8, 31, 1, 32, 1, 32,
		3, 32, 224, 8, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 0, 0, 39, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 0, 33, 0, 35, 16, 37, 17, 39, 18, 41, 19, 43, 20,
		45, 21, 47, 22, 49, 23, 51, 24, 53, 25, 55, 26, 57, 27, 59, 28, 61, 29,
		63, 0, 65, 0, 67, 0, 69, 0, 71, 0, 73, 0, 75, 0, 77, 30, 1, 0, 5, 4, 0,
		97, 122, 228, 228, 246, 246, 252, 252, 4, 0, 65, 90, 196, 196, 214, 214,
		220, 220, 2, 0, 9, 9, 32, 32, 5, 0, 33, 33, 35, 47, 58, 64, 91, 96, 123,
		126, 3, 0, 9, 10, 13, 13, 32, 32, 241, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 1, 79, 1, 0, 0, 0, 3, 84, 1, 0, 0, 0, 5, 86, 1, 0,
		0, 0, 7, 88, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 93, 1, 0, 0, 0, 13, 95,
		1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0, 19, 104, 1, 0, 0, 0,
		21, 109, 1, 0, 0, 0, 23, 123, 1, 0, 0, 0, 25, 132, 1, 0, 0, 0, 27, 137,
		1, 0, 0, 0, 29, 143, 1, 0, 0, 0, 31, 152, 1, 0, 0, 0, 33, 157, 1, 0, 0,
		0, 35, 159, 1, 0, 0, 0, 37, 162, 1, 0, 0, 0, 39, 164, 1, 0, 0, 0, 41, 166,
		1, 0, 0, 0, 43, 168, 1, 0, 0, 0, 45, 170, 1, 0, 0, 0, 47, 172, 1, 0, 0,
		0, 49, 175, 1, 0, 0, 0, 51, 177, 1, 0, 0, 0, 53, 180, 1, 0, 0, 0, 55, 182,
		1, 0, 0, 0, 57, 185, 1, 0, 0, 0, 59, 193, 1, 0, 0, 0, 61, 206, 1, 0, 0,
		0, 63, 219, 1, 0, 0, 0, 65, 223, 1, 0, 0, 0, 67, 225, 1, 0, 0, 0, 69, 227,
		1, 0, 0, 0, 71, 229, 1, 0, 0, 0, 73, 231, 1, 0, 0, 0, 75, 233, 1, 0, 0,
		0, 77, 235, 1, 0, 0, 0, 79, 80, 5, 110, 0, 0, 80, 81, 5, 111, 0, 0, 81,
		82, 5, 116, 0, 0, 82, 83, 5, 40, 0, 0, 83, 2, 1, 0, 0, 0, 84, 85, 5, 41,
		0, 0, 85, 4, 1, 0, 0, 0, 86, 87, 5, 44, 0, 0, 87, 6, 1, 0, 0, 0, 88, 89,
		5, 46, 0, 0, 89, 90, 5, 46, 0, 0, 90, 8, 1, 0, 0, 0, 91, 92, 5, 40, 0,
		0, 92, 10, 1, 0, 0, 0, 93, 94, 5, 93, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96,
		5, 91, 0, 0, 96, 14, 1, 0, 0, 0, 97, 98, 5, 46, 0, 0, 98, 16, 1, 0, 0,
		0, 99, 100, 5, 100, 0, 0, 100, 101, 5, 97, 0, 0, 101, 102, 5, 116, 0, 0,
		102, 103, 5, 101, 0, 0, 103, 18, 1, 0, 0, 0, 104, 105, 5, 116, 0, 0, 105,
		106, 5, 105, 0, 0, 106, 107, 5, 109, 0, 0, 107, 108, 5, 101, 0, 0, 108,
		20, 1, 0, 0, 0, 109, 110, 5, 100, 0, 0, 110, 111, 5, 97, 0, 0, 111, 112,
		5, 116, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114, 5, 32, 0, 0, 114, 115,
		5, 97, 0, 0, 115, 116, 5, 110, 0, 0, 116, 117, 5, 100, 0, 0, 117, 118,
		5, 32, 0, 0, 118, 119, 5, 116, 0, 0, 119, 120, 5, 105, 0, 0, 120, 121,
		5, 109, 0, 0, 121, 122, 5, 101, 0, 0, 122, 22, 1, 0, 0, 0, 123, 124, 5,
		100, 0, 0, 124, 125, 5, 117, 0, 0, 125, 126, 5, 114, 0, 0, 126, 127, 5,
		97, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129, 5, 105, 0, 0, 129, 130, 5,
		111, 0, 0, 130, 131, 5, 110, 0, 0, 131, 24, 1, 0, 0, 0, 132, 133, 5, 116,
		0, 0, 133, 134, 5, 114, 0, 0, 134, 135, 5, 117, 0, 0, 135, 136, 5, 101,
		0, 0, 136, 26, 1, 0, 0, 0, 137, 138, 5, 102, 0, 0, 138, 139, 5, 97, 0,
		0, 139, 140, 5, 108, 0, 0, 140, 141, 5, 115, 0, 0, 141, 142, 5, 101, 0,
		0, 142, 28, 1, 0, 0, 0, 143, 147, 3, 31, 15, 0, 144, 146, 3, 33, 16, 0,
		145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 30, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 153, 3,
		67, 33, 0, 151, 153, 3, 69, 34, 0, 152, 150, 1, 0, 0, 0, 152, 151, 1, 0,
		0, 0, 153, 32, 1, 0, 0, 0, 154, 158, 3, 31, 15, 0, 155, 158, 5, 95, 0,
		0, 156, 158, 3, 71, 35, 0, 157, 154, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 156, 1, 0, 0, 0, 158, 34, 1, 0, 0, 0, 159, 160, 5, 42, 0, 0, 160,
		161, 5, 42, 0, 0, 161, 36, 1, 0, 0, 0, 162, 163, 5, 42, 0, 0, 163, 38,
		1, 0, 0, 0, 164, 165, 5, 47, 0, 0, 165, 40, 1, 0, 0, 0, 166, 167, 5, 43,
		0, 0, 167, 42, 1, 0, 0, 0, 168, 169, 5, 45, 0, 0, 169, 44, 1, 0, 0, 0,
		170, 171, 5, 60, 0, 0, 171, 46, 1, 0, 0, 0, 172, 173, 5, 60, 0, 0, 173,
		174, 5, 61, 0, 0, 174, 48, 1, 0, 0, 0, 175, 176, 5, 62, 0, 0, 176, 50,
		1, 0, 0, 0, 177, 178, 5, 62, 0, 0, 178, 179, 5, 61, 0, 0, 179, 52, 1, 0,
		0, 0, 180, 181, 5, 61, 0, 0, 181, 54, 1, 0, 0, 0, 182, 183, 5, 33, 0, 0,
		183, 184, 5, 61, 0, 0, 184, 56, 1, 0, 0, 0, 185, 189, 2, 48, 57, 0, 186,
		188, 3, 71, 35, 0, 187, 186, 1, 0, 0, 0, 188, 191, 1, 0, 0, 0, 189, 187,
		1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 58, 1, 0, 0, 0, 191, 189, 1, 0,
		0, 0, 192, 194, 3, 71, 35, 0, 193, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0,
		0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		201, 5, 46, 0, 0, 198, 200, 3, 71, 35, 0, 199, 198, 1, 0, 0, 0, 200, 203,
		1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 204, 1, 0,
		0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 2, 48, 57, 0, 205, 60, 1, 0, 0, 0,
		206, 210, 5, 34, 0, 0, 207, 209, 3, 63, 31, 0, 208, 207, 1, 0, 0, 0, 209,
		212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213,
		1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214, 5, 34, 0, 0, 214, 62, 1, 0,
		0, 0, 215, 220, 3, 65, 32, 0, 216, 220, 3, 71, 35, 0, 217, 220, 3, 75,
		37, 0, 218, 220, 3, 73, 36, 0, 219, 215, 1, 0, 0, 0, 219, 216, 1, 0, 0,
		0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 64, 1, 0, 0, 0, 221,
		224, 3, 67, 33, 0, 222, 224, 3, 69, 34, 0, 223, 221, 1, 0, 0, 0, 223, 222,
		1, 0, 0, 0, 224, 66, 1, 0, 0, 0, 225, 226, 7, 0, 0, 0, 226, 68, 1, 0, 0,
		0, 227, 228, 7, 1, 0, 0, 228, 70, 1, 0, 0, 0, 229, 230, 2, 48, 57, 0, 230,
		72, 1, 0, 0, 0, 231, 232, 7, 2, 0, 0, 232, 74, 1, 0, 0, 0, 233, 234, 7,
		3, 0, 0, 234, 76, 1, 0, 0, 0, 235, 236, 7, 4, 0, 0, 236, 237, 1, 0, 0,
		0, 237, 238, 6, 38, 0, 0, 238, 78, 1, 0, 0, 0, 10, 0, 147, 152, 157, 189,
		195, 201, 210, 219, 223, 1, 6, 0, 0,
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

// SFeelLexerInit initializes any static state used to implement SFeelLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSFeelLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SFeelLexerInit() {
	staticData := &sfeellexerLexerStaticData
	staticData.once.Do(sfeellexerLexerInit)
}

// NewSFeelLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSFeelLexer(input antlr.CharStream) *SFeelLexer {
	SFeelLexerInit()
	l := new(SFeelLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &sfeellexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SFeel.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SFeelLexer tokens.
const (
	SFeelLexerT__0      = 1
	SFeelLexerT__1      = 2
	SFeelLexerT__2      = 3
	SFeelLexerT__3      = 4
	SFeelLexerT__4      = 5
	SFeelLexerT__5      = 6
	SFeelLexerT__6      = 7
	SFeelLexerT__7      = 8
	SFeelLexerT__8      = 9
	SFeelLexerT__9      = 10
	SFeelLexerT__10     = 11
	SFeelLexerT__11     = 12
	SFeelLexerT__12     = 13
	SFeelLexerT__13     = 14
	SFeelLexerName      = 15
	SFeelLexerPOW       = 16
	SFeelLexerMUL       = 17
	SFeelLexerDIV       = 18
	SFeelLexerADD       = 19
	SFeelLexerSUB       = 20
	SFeelLexerLESS      = 21
	SFeelLexerLESSEQ    = 22
	SFeelLexerGREATER   = 23
	SFeelLexerGREATEREQ = 24
	SFeelLexerEQUAL     = 25
	SFeelLexerNOTEQUAL  = 26
	SFeelLexerINTEGER   = 27
	SFeelLexerREAL      = 28
	SFeelLexerSTRING    = 29
	SFeelLexerWS        = 30
)
