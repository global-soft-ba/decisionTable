package antlr

import (
	"decisionTable/lang/sfeel/ast"
	gen "decisionTable/lang/sfeel/gen"
)

var (
	TokenTable = map[int]int{
		gen.SFeelParserPOW:                        ast.SFeelOperatorPOW,
		gen.SFeelParserMUL:                        ast.SFeelOperatorMUL,
		gen.SFeelParserDIV:                        ast.SFeelOperatorDIV,
		gen.SFeelParserADD:                        ast.SFeelOperatorADD,
		gen.SFeelParserSUB:                        ast.SFeelOperatorSUB,
		gen.SFeelParserLESS:                       ast.SFeelOperatorLESS,
		gen.SFeelParserLESSEQ:                     ast.SFeelOperatorLESSEQ,
		gen.SFeelParserGREATER:                    ast.SFeelOperatorGREATER,
		gen.SFeelParserGREATEREQ:                  ast.SFeelOperatorGREATEREQ,
		gen.SFeelParserEQUAL:                      ast.SFeelOperatorEQUAL,
		gen.SFeelParserNOTEQUAL:                   ast.SFeelOperatorNOTEQUAL,
		gen.SFeelParserRULE_closed_interval_start: ast.SFeelClosedIntervalSTART,
		gen.SFeelParserRULE_closed_interval_end:   ast.SFeelClosedIntervalEND,
		gen.SFeelParserRULE_open_interval_start:   ast.SFeelOpenIntervalSTART,
		gen.SFeelParserRULE_open_interval_end:     ast.SFeelOpenIntervalEND,
	}
)
