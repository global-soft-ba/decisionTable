package grl

import (
	"bytes"
	gen "decisionTable/parser/sfeel/generated"
	"fmt"
	"text/template"
)

const (
	// GRULE Specific Tokens
	// Logical operator token
	EQUAL    = '='
	AND      = "&&"
	OR       = "||"
	NEGATION = `!( {{ expr }} )`
)

var (
	/*
		Use the lexer config von sFeel.g4
		OneToOne mappings
	*/

	ComparisionOperators = map[int]string{
		gen.SFeelLexerLESS:      "<",
		gen.SFeelLexerLESSEQ:    "<=",
		gen.SFeelLexerGREATER:   ">",
		gen.SFeelLexerGREATEREQ: ">=",
	}
)

var (
	Operations = map[int]string{
		gen.SFeelParserRULE_equalcomparisonInteger: ` input.field = val `,
		gen.SFeelParserRULE_comparisonInteger:      `{{define template Exp.InputField}} < {{Exp.val}}`,
		gen.SFeelParserRULE_rangeInteger:           `{{define template ( {{template singleRange .StartVal} && {{template singleRange .EndVal}  ) }}`,
		gen.SFeelParserRULE_disjunctionsInteger:    ` input =1 {{templates.OR}} input = 2 || input = 3 || ....`,
	}

	startRanges = map[int]string{
		gen.SFeelParserRANGEIN:  `{{<=}}`, // [1..
		gen.SFeelParserRANGEOUT: `{{<}}`,  // ]1..
	}

	endRanges = map[int]string{
		gen.SFeelParserRANGEIN:  `{{<}}`,  // ..1[
		gen.SFeelParserRANGEOUT: `{{<=}}`, // ..1]
	}
)

func test() {

	var t *template.Template

	t, _ = template.New("Expressions").Parse(`{}`)
	var tpl bytes.Buffer
	t.Execute(&tpl, "")

	fmt.Println("Template=", tpl.String())

}
