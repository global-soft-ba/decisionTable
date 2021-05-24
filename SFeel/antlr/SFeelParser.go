package antlr

import (
	antlrErrors "decisionTable/SFeel/antlr/errors"
	gen "decisionTable/SFeel/gen"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func CreateSFeelParser(expr string) SFeelParser {
	// Create ErrorListener
	errorListener := &antlrErrors.ErrorListener{}

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

	return SFeelParser{lexer, prs, errorListener}
}

type SFeelParser struct {
	lexer         *gen.SFeelLexer
	parser        *gen.SFeelParser
	errorListener *antlrErrors.ErrorListener
}

func (O SFeelParser) Lexer() *gen.SFeelLexer {
	return O.lexer
}

func (O SFeelParser) Parser() *gen.SFeelParser {
	return O.parser
}

func (O SFeelParser) Errors() []error {
	return O.errorListener.Errors
}
