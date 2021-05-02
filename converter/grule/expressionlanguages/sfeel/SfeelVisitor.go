package sfeel

import (
	antlrgen "decisionTable/validators/expressionlanguages/sfeel/antlrgen"
	"decisionTable/validators/expressionlanguages/sfeel/parser"
	"fmt"
)

func CreateSFeelVisitor() *Visitor {
	return &Visitor{}
}

type Visitor struct {
	antlrgen.BaseSFeelVisitor
}

func (v *Visitor) VisitValidIntegerInput(ctx *antlrgen.ValidIntegerInputContext) interface{} {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetTokenType())
	return ctx.ComparisonInteger().Accept(v)
}

func (v *Visitor) VisitComparisonInteger(ctx *antlrgen.ComparisonIntegerContext) interface{} {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetText())
	ops := ctx.ComparisonOps().Accept(v)
	val := ctx.INTEGER()

	fmt.Println("Ops,val=", ops, val)

	return ops
}

func (v *Visitor) VisitComparisonOps(ctx *antlrgen.ComparisonOpsContext) interface{} {
	fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetText())
	return ctx.GetStart().GetText()
}

func (v Visitor) testVisitor() {
	prs := parser.CreateParser("<2")
	tree := prs.Parse().ValidIntegerInput()
	vis := CreateSFeelVisitor()
	x := tree.Accept(vis)
	fmt.Println("OUT", x)

}
