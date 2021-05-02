package grl

import (
	gen "decisionTable/parser/sfeel/generated"
	parser2 "decisionTable/parser/sfeel/parser"
	"fmt"
)

func CreateSFeelVisitor() *Visitor {
	return &Visitor{}
}

// Visitor Visitor is a converter between the ParseTree and our grl data model
type Visitor struct {
	gen.BaseSFeelVisitor
}

func (v *Visitor) VisitComparisonIntegerInputRule(ctx *gen.ComparisonIntegerInputRuleContext) interface{} {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetTokenType())

	return ctx.ComparisonInteger().Accept(v)
}

func (v *Visitor) VisitComparisonInteger(ctx *gen.ComparisonIntegerContext) interface{} {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetText())
	ops := ctx.ComparisonOps().Accept(v)
	val := ctx.INTEGER()

	fmt.Println("Ops,val=", ops, val)

	return ops
}

func (v *Visitor) VisitComparisonOps(ctx *gen.ComparisonOpsContext) interface{} {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetText())
	return ctx.GetStart().GetTokenType()
}

func (v *Visitor) VisitEmptyIntegerInputRule(ctx *gen.EmptyIntegerInputRuleContext) interface{} {
	fmt.Println("Empty Rule")
	return ctx
}

func (v Visitor) testVisitor() {

	// const ENTRY = `{{define "ENTRY"}}{{.Identifier}}.{{.Name}} {{.Expression}}{{end}}`

	prs := parser2.CreateParser("-")
	tree := prs.Parse().ValidIntegerInput()
	vis := CreateSFeelVisitor()
	x := tree.Accept(vis)

	fmt.Println("OUT", x)

}
