package antlr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/global-soft-ba/decisionTable/lang/sfeel/antlr/errors"
	"github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	parser2 "github.com/global-soft-ba/decisionTable/lang/sfeel/gen"
)

// ToDo Parser cannot be reused (it is not thread safe). So we assume that for each evaluation of an expression a new parser must be created

func CreateParser(expr string) Parser {
	// Create ErrorListener
	errorListener := &errors.ErrorListener{}

	// Create Lexer
	is := antlr.NewInputStream(expr)
	lexer := parser2.NewSFeelLexer(is)

	// Add ErrorListener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	//Create TokenStream with Lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parse
	prs := parser2.NewSFeelParser(stream)

	// Add ErrorListener
	prs.RemoveErrorListeners()
	prs.AddErrorListener(errorListener)

	return Parser{lexer, prs, errorListener}
}

type Parser struct {
	lexer         *parser2.SFeelLexer
	parser        *parser2.SFeelParser
	errorListener *errors.ErrorListener
}

func (p Parser) Lexer() *parser2.SFeelLexer {
	return p.lexer
}

func (p Parser) Parser() *parser2.SFeelParser {
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

func (p Parser) syntaxCheck(tree parser2.IStartContext) bool {
	base := parser2.BaseSFeelListener{}
	antlr.ParseTreeWalkerDefault.Walk(&base, tree)
	if len(p.errorListener.Errors) != 0 {
		return false
	}

	return true
}

func (p Parser) buildAst(tree parser2.IStartContext) (ast.Node, []error) {
	lis := CreateListener()
	antlr.ParseTreeWalkerDefault.Walk(&lis, tree)
	if len(lis.Errors) != 0 {
		return nil, lis.Errors
	}
	return lis.GetAST(), nil
}
