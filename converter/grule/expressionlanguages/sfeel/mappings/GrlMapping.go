package mappings

import (
	gen "decisionTable/parser/sfeel/generated"
)

var GrlMapping = ConverterMapping{
	TargetToken: map[int]string{
		EQUAL:    "=",
		AND:      "&&",
		OR:       "||",
		NEGATION: `!( {{ expr }} )`,
	},
	ComparisonOperators: map[int]string{
		gen.SFeelLexerLESS:      "<",
		gen.SFeelLexerLESSEQ:    "<=",
		gen.SFeelLexerGREATER:   ">",
		gen.SFeelLexerGREATEREQ: ">=",
	},
	Operations: map[int]string{
		gen.SFeelParserRULE_equalcomparisonInteger: ` input.field = val `,
		gen.SFeelParserRULE_comparisonInteger:      `{{define template Exp.InputField}} < {{Exp.val}}`,
		gen.SFeelParserRULE_rangeInteger:           `{{define template ( {{template singleRange .StartVal} && {{template singleRange .EndVal}  ) }}`,
		gen.SFeelParserRULE_disjunctionsInteger:    ` input =1 {{templates.OR}} input = 2 || input = 3 || ....`,
	},
	StartRanges: map[int]string{
		gen.SFeelParserRANGEIN:  `{{<=}}`, // [1..
		gen.SFeelParserRANGEOUT: `{{<}}`,  // ]1..
	},
	EndRanges: map[int]string{
		gen.SFeelParserRANGEIN:  `{{<}}`,  // ..1[
		gen.SFeelParserRANGEOUT: `{{<=}}`, // ..1]
	},
}
