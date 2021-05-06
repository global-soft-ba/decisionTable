package mapper

import (
	gen "decisionTable/parser/sfeel/generated"
)

var SettingsGRL = TermMapper{
	TargetToken: map[int]string{
		AND:    "&&",
		OR:     "||",
		EQUAL:  "==",
		ASSIGN: "=",
	},
	ComparisonOperators: map[int]string{
		gen.SFeelLexerLESS:      "<",
		gen.SFeelLexerLESSEQ:    "<=",
		gen.SFeelLexerGREATER:   ">",
		gen.SFeelLexerGREATEREQ: ">=",
	},
	StartRanges: map[int]string{
		gen.SFeelParserRANGEIN:  `>=`, // [1..
		gen.SFeelParserRANGEOUT: `>`,  // ]1..
	},
	EndRanges: map[int]string{
		gen.SFeelParserRANGEIN:  `<`,  // ..1[
		gen.SFeelParserRANGEOUT: `<=`, // ..1]
	},
	Templates: map[int]string{
		EQUALCOMPARISON: `{{.Expr.Identifier}}.{{.Expr.Name}} {{ .Op }} {{ .Val }}`,
		COMPARISON:      `{{.Expr.Identifier}}.{{.Expr.Name}} {{ .Op }} {{ .Val }}`,
		RANGES:          `(({{.start.Expr.Identifier}}.{{.start.Expr.Name}} {{.start.Op}} {{.start.Val}}) && ({{.end.Expr.Identifier}}.{{.end.Expr.Name}} {{.end.Op}} {{.end.Val}}))`,
		DISJUNCTIONS:    `{{.first}} || {{.second}}`,
		DISJUNCTIONTERM: `({{.}})`,
		NEGATION:        `!({{.}})`,
		DATEANDTIME:     `MakeTime({{.Year}},{{.Month}},{{.Day}},{{.Hour}},{{.Minutes}},{{.Seconds}})`,
		ASSIGNMENT:      `{{.Expr.Identifier}}.{{.Expr.Name}} {{ .Op }} {{ .Val }}`,
	},
}
