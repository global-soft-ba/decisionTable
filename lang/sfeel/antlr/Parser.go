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

func (p Parser) ParseInput() (ast.Node, []error) {
	parseTree := p.Parser().Input()
	if okay := p.inputSyntaxCheck(parseTree); !okay {
		return nil, p.errorListener.Errors
	}

	return p.buildInputAst(parseTree)
}

func (p Parser) inputSyntaxCheck(tree parser2.IInputContext) bool {
	base := parser2.BaseSFeelListener{}
	antlr.ParseTreeWalkerDefault.Walk(&base, tree)
	if len(p.errorListener.Errors) != 0 {
		return false
	}

	return true
}

func (p Parser) buildInputAst(tree parser2.IInputContext) (ast.Node, []error) {
	lis := CreateListener()
	antlr.ParseTreeWalkerDefault.Walk(&lis, tree)
	if len(lis.Errors) != 0 {
		return nil, lis.Errors
	}
	return lis.GetAST(), nil
}

func (p Parser) ParseOutput() (ast.Node, []error) {
	parseTree := p.Parser().Output()
	if okay := p.outputSyntaxCheck(parseTree); !okay {
		return nil, p.errorListener.Errors
	}

	return p.buildOutputAst(parseTree)
}

func (p Parser) outputSyntaxCheck(tree parser2.IOutputContext) bool {
	base := parser2.BaseSFeelListener{}
	antlr.ParseTreeWalkerDefault.Walk(&base, tree)
	if len(p.errorListener.Errors) != 0 {
		return false
	}

	return true
}

func (p Parser) buildOutputAst(tree parser2.IOutputContext) (ast.Node, []error) {
	lis := CreateListener()
	antlr.ParseTreeWalkerDefault.Walk(&lis, tree)
	if len(lis.Errors) != 0 {
		return nil, lis.Errors
	}
	return lis.GetAST(), nil
}
