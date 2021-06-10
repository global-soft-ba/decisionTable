package visitors

import (
	"github.com/global-soft-ba/decisionTable/converters/grule/grlmodel"
	mapper2 "github.com/global-soft-ba/decisionTable/converters/grule/termconverter/sfeel/mapper"
	gen "github.com/global-soft-ba/decisionTable/parser/sfeel/generated"
)

func CreateBoolVisitor(expr grlmodel.Term, maps mapper2.TermMapper) *BoolVisitor {
	return &BoolVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

type BoolVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel.Term
	maps mapper2.TermMapper
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
