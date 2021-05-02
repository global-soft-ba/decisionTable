package parser

import (
	errors2 "decisionTable/parser/sfeel/errors"
	gen "decisionTable/parser/sfeel/generated"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func CreateParser(expr string) Parser {
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

	return Parser{lexer, prs, errorListener}
}

type Parser struct {
	lexer         *gen.SFeelLexer
	parser        *gen.SFeelParser
	errorListener *errors2.ErrorListener
}

func (p Parser) Lexer() *gen.SFeelLexer {
	return p.lexer
}

func (p Parser) Parse() *gen.SFeelParser {
	return p.parser
}

func (p Parser) Errors() []error {
	return p.errorListener.Errors
}
