package visitors

import (
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/grlmodel"
	gen "decisionTable/parser/sfeel/generated"
)

func CreateBoolVisitor(expr grlmodel.Term, maps maps.Mapper) *BoolVisitor {
	return &BoolVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

type BoolVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel.Term
	maps maps.Mapper
}

func (v *BoolVisitor) VisitEqualComparisonBoolInputRule(ctx *gen.EqualComparisonBoolInputRuleContext) interface{} {
	val := ctx.EqualcomparisonBool().GetText()
	return v.maps.MapEqualComparison(v.expr, val)
}

func (v *BoolVisitor) VisitEmptyBoolInputRule(ctx *gen.EmptyBoolInputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}
