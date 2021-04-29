package sfeel

import (
	"decisionTable/model"
	"decisionTable/validator/expressions/sfeel/errors"
	gen "decisionTable/validator/expressions/sfeel/generated"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func CreateParser() Parser {
	// Create ErrorListener
	errorListener := &errors.ErrorListener{}

	lexer := createLexer(errorListener)
	prs := createParser(lexer, errorListener)

	return Parser{lexer, prs, errorListener}
}

func createLexer(e *errors.ErrorListener) *gen.SFeelLexer {
	// Setup the input - Initial Load
	is := antlr.NewInputStream("-")
	lexer := gen.NewSFeelLexer(is)

	// Add ErrorListener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(e)
	return lexer
}

func createParser(lexer *gen.SFeelLexer, errors *errors.ErrorListener) *gen.SFeelParser {
	//Create TokenStream and Lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	prs := gen.NewSFeelParser(stream)

	// Add ErrorListener
	prs.RemoveErrorListeners()
	prs.AddErrorListener(errors)
	return prs
}

type Parser struct {
	lexer         *gen.SFeelLexer
	parser        *gen.SFeelParser
	errorListener *errors.ErrorListener
}

//TODO Check if, due to the use of pointer this structure is thread safe.

func (p Parser) ValidateEntry(entry model.Entry) (bool, []error) {
	//Reset Parser before new validation
	is := antlr.NewInputStream(entry.Expression())
	p.lexer.SetInputStream(is)
	p.parser.Start()

	if len(p.errorListener.Errors) > 0 {
		return false, p.errorListener.Errors
	}

	return true, nil
}

/*
func (p Parser) ConvertEntry(entry model.Entry,listener antlr.ParseTreeListener){
	//Reset Parser before new converting
	is := antlr.NewInputStream(entry.Expression())
	p.lexer.SetInputStream(is)
	p.parser.Start()

	antlr.NewParseTreeWalker().Walk(listener,p.parser.Start())

	if len(p.errorListener.Errors) > 0 {
		return false, p.errorListener.Errors
	}
}
*/
