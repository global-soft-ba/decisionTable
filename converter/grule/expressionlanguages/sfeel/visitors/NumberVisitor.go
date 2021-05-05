package visitors

import (
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/grlmodel"
	gen "decisionTable/parser/sfeel/generated"
)

func CreateNumberVisitor(expr grlmodel.Term, maps maps.Mapper) *NumberVisitor {
	return &NumberVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

type NumberVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel.Term
	maps maps.Mapper
}

func (v *NumberVisitor) VisitEqualComparisonNumberInputRule(ctx *gen.EqualComparisonNumberInputRuleContext) interface{} {
	val := ctx.EqualcomparisonNumber().GetText()
	return v.maps.MapEqualComparison(v.expr, val)
}

func (v *NumberVisitor) VisitComparisonNumberInputRule(ctx *gen.ComparisonNumberInputRuleContext) interface{} {
	return ctx.ComparisonNumber().Accept(v)
}

func (v *NumberVisitor) VisitComparisonNumber(ctx *gen.ComparisonNumberContext) interface{} {
	tokenType := ctx.ComparisonOps().Accept(v).(int)
	val := ctx.Number().GetText()

	return v.maps.MapComparison(v.expr, tokenType, val)
}

func (v *NumberVisitor) VisitComparisonOps(ctx *gen.ComparisonOpsContext) interface{} {
	return ctx.GetStart().GetTokenType()
}

func (v *NumberVisitor) VisitEmptyNumberInputRule(ctx *gen.EmptyNumberInputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}

func (v *NumberVisitor) VisitRangeComparisonNumberInputRule(ctx *gen.RangeComparisonNumberInputRuleContext) interface{} {
	return ctx.RangeNumber().Accept(v)
}

func (v *NumberVisitor) VisitRangeNumber(ctx *gen.RangeNumberContext) interface{} {
	opStart := ctx.GetStart().GetTokenType()
	opEnd := ctx.GetStop().GetTokenType()
	valStart := ctx.Number(0).GetText()
	valEnd := ctx.Number(1).GetText()
	return v.maps.MapRange(v.expr, opStart, valStart, opEnd, valEnd)
}

func (v *NumberVisitor) VisitDisjunctionsNumberInputRule(ctx *gen.DisjunctionsNumberInputRuleContext) interface{} {
	entireTerm := ctx.DisjunctionsNumber().Accept(v)
	return v.maps.MapDisjunctionsTerm(entireTerm)
}

func (v *NumberVisitor) VisitDisjunctionsNumber(ctx *gen.DisjunctionsNumberContext) interface{} {
	var result []interface{}

	for _, val := range ctx.AllDisjunctionsNumber() {
		resString := val.Accept(v)
		result = append(result, resString)
	}

	if ctx.Number() != nil {
		term := v.maps.MapEqualComparison(v.expr, ctx.Number().GetText())
		return v.maps.MapDisjunctionsTerm(term)
	}

	if ctx.ComparisonNumber() != nil {
		comp := ctx.ComparisonNumber().Accept(v)
		return v.maps.MapDisjunctionsTerm(comp)
	}

	if ctx.RangeNumber() != nil {
		rng := ctx.RangeNumber().Accept(v)
		return v.maps.MapDisjunctionsTerm(rng)
	}

	return v.maps.MapDisjunctions(result)
}

func (v *NumberVisitor) VisitNegationNumberInputRule(ctx *gen.NegationNumberInputRuleContext) interface{} {
	var term interface{}

	if ctx.EqualcomparisonNumber() != nil {
		val := ctx.EqualcomparisonNumber().GetText()
		term = v.maps.MapEqualComparison(v.expr, val)
	}

	if ctx.ComparisonNumber() != nil {
		term = ctx.ComparisonNumber().Accept(v)
	}

	if ctx.RangeNumber() != nil {
		term = ctx.RangeNumber().Accept(v)
	}

	if ctx.DisjunctionsNumber() != nil {
		term = ctx.DisjunctionsNumber().Accept(v)
	}

	return v.maps.MapNegation(term)
}
