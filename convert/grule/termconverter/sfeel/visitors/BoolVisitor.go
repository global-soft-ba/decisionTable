package visitors

import (
	grlmodel2 "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/grl/symbols"
	gen "decisionTable/parser/sfeel/generated"
)

func CreateBoolVisitor(expr grlmodel2.Term, maps symbols.TermMapper) *BoolVisitor {
	return &BoolVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

type BoolVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel2.Term
	maps symbols.TermMapper
}

func (v *BoolVisitor) VisitEqualComparisonBoolInputRule(ctx *gen.EqualComparisonBoolInputRuleContext) interface{} {
	val := ctx.EqualcomparisonBool().GetText()
	return v.maps.MapEqualComparison(v.expr, val)
}

func (v *BoolVisitor) VisitEmptyBoolInputRule(ctx *gen.EmptyBoolInputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}

// Assignment Rules
func (v *BoolVisitor) VisitBoolAssignmentOutputRule(ctx *gen.BoolAssignmentOutputRuleContext) interface{} {
	val := ctx.Bools().GetText()
	return v.maps.MapAssignment(v.expr, val)
}

func (v *BoolVisitor) VisitEmptyBoolOutputRule(ctx *gen.EmptyBoolOutputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}
