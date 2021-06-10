package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	errors2 "github.com/global-soft-ba/decisionTable/parser/sfeel/errors"
	gen "github.com/global-soft-ba/decisionTable/parser/sfeel/generated"
)

func CreateSfeelParser(expr string) SfeelParser {
	// Create ErrorListener
	errorListener := &errors2.ErrorListener{}

	// Create Lexer
	is := antlr.NewInputStream(expr)
	lexer := gen.NewSFeelLexer(is)

	// Add ErrorListener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	//Create TokenStream with Lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parse
	prs := gen.NewSFeelParser(stream)

	// Add ErrorListener
	prs.RemoveErrorListeners()
	prs.AddErrorListener(errorListener)

	return SfeelParser{lexer, prs, errorListener}
}

type SfeelParser struct {
	lexer         *gen.SFeelLexer
	parser        *gen.SFeelParser
	errorListener *errors2.ErrorListener
}

func (p SfeelParser) Lexer() *gen.SFeelLexer {
	return p.lexer
}

func (p SfeelParser) Parse() *gen.SFeelParser {
	return p.parser
}

func (p SfeelParser) Errors() []error {
	return p.errorListener.Errors
}
