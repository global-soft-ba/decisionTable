package converters

import (
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mappings"
	gen "decisionTable/parser/sfeel/generated"
	"fmt"
)

func CreateIntegerConverter(maps maps.ConverterMapping) *IntegerConverter {
	return &IntegerConverter{gen.BaseSFeelVisitor{}, maps}
}

// IntegerConverter IntegerConverter is a converter between the ParseTree and our grl data model
type IntegerConverter struct {
	gen.BaseSFeelVisitor
	maps maps.ConverterMapping
}

func (v *IntegerConverter) VisitComparisonIntegerInputRule(ctx *gen.ComparisonIntegerInputRuleContext) interface{} {
	return ctx.ComparisonInteger().Accept(v)
}

func (v *IntegerConverter) VisitComparisonInteger(ctx *gen.ComparisonIntegerContext) interface{} {
	// fmt.Println("InInput:", ctx.GetText(), ":", ctx.GetStart().GetTokenType(), ":", ctx.GetStop().GetText())
	//ops := ctx.ComparisonOps().Accept(v)
	//val := ctx.INTEGER().GetText()

	// r := v.tempalte.Execute("ComInteger", ComparisonOperators[ops] , val)
	return ""
}

func (v *IntegerConverter) VisitComparisonOps(ctx *gen.ComparisonOpsContext) interface{} {
	return ctx.GetStart().GetTokenType()
}

func (v *IntegerConverter) VisitEmptyIntegerInputRule(ctx *gen.EmptyIntegerInputRuleContext) interface{} {
	fmt.Println("Empty Rule")
	return ctx
}

/*
func (v IntegerConverter) testVisitor() {

	// const ENTRY = `{{define "ENTRY"}}{{.Identifier}}.{{.Name}} {{.Expression}}{{end}}`

	prs := parser.CreateParser("-")
	tree := prs.Parse().ValidIntegerInput()
	vis := CreateConverter()
	x := tree.Accept(vis)

	fmt.Println("OUT", x)

}
*/
