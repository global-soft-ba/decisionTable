package parser

import (
	"decisionTable/validators/expressionlanguages/sfeel/errors"
	gen "decisionTable/validators/expressionlanguages/sfeel/generated"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func CreateParser(expr string) Parser {
	// Create ErrorListener
	errorListener := &errors.ErrorListener{}

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
	errorListener *errors.ErrorListener
}

func (p Parser) Parse() *gen.SFeelParser {
	return p.parser
}

func (p Parser) Errors() []error {
	return p.errorListener.Errors
}
