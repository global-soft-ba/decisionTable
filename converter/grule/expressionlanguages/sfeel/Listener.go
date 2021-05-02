package sfeel

import (
	antlrg "decisionTable/validators/expressionlanguages/sfeel/antlrgen"
	"decisionTable/validators/expressionlanguages/sfeel/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func CreateSfeelListener() *Listener {
	return &Listener{}
}

type Listener struct {
	*antlrg.BaseSFeelListener
}

func (l *Listener) EnterValidIntegerInput(ctx *antlrg.ValidIntegerInputContext) {
	fmt.Println("InInput:", ctx.GetText())
}

func (l *Listener) EnterComparisonInteger(ctx *antlrg.ComparisonIntegerContext) {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetText())
}

func (l *Listener) ExitComparisonInteger(ctx *antlrg.ComparisonIntegerContext) {
	fmt.Println("InInput:", ctx.GetText(), ctx.GetStart().GetText())
}

func testListener() {
	prs := parser.CreateParser(">2")
	tree := prs.Parse().ValidIntegerInput()
	lis := CreateSfeelListener()

	antlr.NewParseTreeWalker().Walk(lis, tree)
}

/*
func (p Validator) ConvertEntry(entry model.Entry,listener antlr.ParseTreeListener){
	//Reset Validator before new converting
	is := antlr.NewInputStream(entry.Expression())
	p.lexer.SetInputStream(is)
	p.parser.Start()

	antlr.NewParseTreeWalker().Walk(Listener{},p.parser.Start())

	if len(p.errorListener.Errors) > 0 {
		return false, p.errorListener.Errors
	}
}
*/
