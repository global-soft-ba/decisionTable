package antlr

import (
	antlrErrors "decisionTable/sfeel/antlr/errors"
	"decisionTable/sfeel/ast"
	gen "decisionTable/sfeel/gen"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ToDo Parser cannot be reused. So we assume that for each evaluation of an expression a new parser must be created

func CreateParser(expr string) Parser {
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

	return Parser{lexer, prs, errorListener}
}

type Parser struct {
	lexer         *gen.SFeelLexer
	parser        *gen.SFeelParser
	errorListener *antlrErrors.ErrorListener
}

func (p Parser) Lexer() *gen.SFeelLexer {
	return p.lexer
}

func (p Parser) Parser() *gen.SFeelParser {
	return p.parser
}

func (p Parser) Errors() []error {
	return p.errorListener.Errors
}

func (p Parser) Parse() (ast.Node, []error) {
	parseTree := p.Parser().Start()
	if okay := p.syntaxCheck(parseTree); !okay {
		return nil, p.errorListener.Errors
	}

	return p.buildAst(parseTree)
}

func (p Parser) syntaxCheck(tree gen.IStartContext) bool {
	base := gen.BaseSFeelListener{}
	antlr.ParseTreeWalkerDefault.Walk(&base, tree)
	if len(p.errorListener.Errors) != 0 {
		return false
	}

	return true
}

func (p Parser) buildAst(tree gen.IStartContext) (ast.Node, []error) {
	lis := CreateListener()
	antlr.ParseTreeWalkerDefault.Walk(&lis, tree)
	if len(lis.Errors) != 0 {
		return nil, lis.Errors
	}
	return lis.GetAST(), nil
}
