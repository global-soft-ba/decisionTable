package sfeel

import (
	"decisionTable/model"
	"decisionTable/validator/expressions/sfeel/errors"
	gen "decisionTable/validator/expressions/sfeel/generated"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//TODO Also Validate if the given entry type an the used operator/type in expression match

func CreateParser() Parser {
	return Parser{}
}

type Parser struct {
}

//TODO Check if, due to the use of pointer this structure is thread safe. Until that, we use a new single parser instance for each validation

func (p Parser) ValidateInputEntry(entry model.Entry) (bool, []error) {

	newP := createSingleParser()
	is := antlr.NewInputStream(entry.Expression())
	newP.lexer.SetInputStream(is)
	newP.parser.InputEntry()

	if len(newP.errorListener.Errors) > 0 {
		return false, newP.errorListener.Errors
	}

	return true, nil
}

func (p Parser) ValidateOutputEntry(entry model.Entry) (bool, []error) {
	newP := createSingleParser()
	is := antlr.NewInputStream(entry.Expression())
	newP.lexer.SetInputStream(is)
	newP.parser.OutputEntry()

	if len(newP.errorListener.Errors) > 0 {
		return false, newP.errorListener.Errors
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

type singleParser struct {
	lexer         *gen.SFeelLexer
	parser        *gen.SFeelParser
	errorListener *errors.ErrorListener
}

func createSingleParser() singleParser {
	// Create ErrorListener
	errorListener := &errors.ErrorListener{}

	lexer := createLexer(errorListener)
	prs := createParser(lexer, errorListener)

	return singleParser{lexer, prs, errorListener}
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
