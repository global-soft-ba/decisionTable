package visitors

import (
	"decisionTable/convert/grule/grlmodel"
	mapper2 "decisionTable/convert/grule/termconverter/sfeel/mapper"
	gen "decisionTable/parser/sfeel/generated"
)

func CreateIntegerVisitor(expr grlmodel.Term, maps mapper2.TermMapper) *IntegerVisitor {
	return &IntegerVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

// IntegerVisitor IntegerVisitor is a converter between the ParseTree and our grl data model
type IntegerVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel.Term
	maps mapper2.TermMapper
}

func (v *IntegerVisitor) VisitEqualComparisonIntegerInputRule(ctx *gen.EqualComparisonIntegerInputRuleContext) interface{} {
	val := ctx.EqualcomparisonInteger().GetText()
	return v.maps.MapEqualComparison(v.expr, val)
}

func (v *IntegerVisitor) VisitComparisonIntegerInputRule(ctx *gen.ComparisonIntegerInputRuleContext) interface{} {
	return ctx.ComparisonInteger().Accept(v)
}

func (v *IntegerVisitor) VisitComparisonInteger(ctx *gen.ComparisonIntegerContext) interface{} {
	tokenType := ctx.ComparisonOps().Accept(v).(int)
	val := ctx.INTEGER().GetText()

	return v.maps.MapComparison(v.expr, tokenType, val)
}

func (v *IntegerVisitor) VisitComparisonOps(ctx *gen.ComparisonOpsContext) interface{} {
	return ctx.GetStart().GetTokenType()
}

func (v *IntegerVisitor) VisitEmptyIntegerInputRule(ctx *gen.EmptyIntegerInputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}

func (v *IntegerVisitor) VisitRangeComparisonIntegerInputRule(ctx *gen.RangeComparisonIntegerInputRuleContext) interface{} {
	return ctx.RangeInteger().Accept(v)
}

func (v *IntegerVisitor) VisitRangeInteger(ctx *gen.RangeIntegerContext) interface{} {
	opStart := ctx.GetStart().GetTokenType()
	opEnd := ctx.GetStop().GetTokenType()
	valStart := ctx.INTEGER(0).GetText()
	valEnd := ctx.INTEGER(1).GetText()
	return v.maps.MapRange(v.expr, opStart, valStart, opEnd, valEnd)
}

func (v *IntegerVisitor) VisitDisjunctionsIntegerInputRule(ctx *gen.DisjunctionsIntegerInputRuleContext) interface{} {
	entireTerm := ctx.DisjunctionsInteger().Accept(v)
	return v.maps.MapDisjunctionsTerm(entireTerm)
}

func (v *IntegerVisitor) VisitDisjunctionsInteger(ctx *gen.DisjunctionsIntegerContext) interface{} {
	var result []interface{}

	for _, val := range ctx.AllDisjunctionsInteger() {
		resString := val.Accept(v)
		result = append(result, resString)
	}

	if ctx.INTEGER() != nil {
		term := v.maps.MapEqualComparison(v.expr, ctx.INTEGER().GetText())
		return v.maps.MapDisjunctionsTerm(term)
	}

	if ctx.ComparisonInteger() != nil {
		comp := ctx.ComparisonInteger().Accept(v)
		return v.maps.MapDisjunctionsTerm(comp)
	}

	if ctx.RangeInteger() != nil {
		rng := ctx.RangeInteger().Accept(v)
		return v.maps.MapDisjunctionsTerm(rng)
	}

	return v.maps.MapDisjunctions(result)
}

func (v *IntegerVisitor) VisitNegationIntegerInputRule(ctx *gen.NegationIntegerInputRuleContext) interface{} {
	var term interface{}

	if ctx.EqualcomparisonInteger() != nil {
		val := ctx.EqualcomparisonInteger().GetText()
		term = v.maps.MapEqualComparison(v.expr, val)
	}

	if ctx.ComparisonInteger() != nil {
		term = ctx.ComparisonInteger().Accept(v)
	}

	if ctx.RangeInteger() != nil {
		term = ctx.RangeInteger().Accept(v)
	}

	if ctx.DisjunctionsInteger() != nil {
		term = ctx.DisjunctionsInteger().Accept(v)
	}

	return v.maps.MapNegation(term)
}

// Assignment Rules
func (v *IntegerVisitor) VisitIntegerAssignmentOutputRule(ctx *gen.IntegerAssignmentOutputRuleContext) interface{} {
	val := ctx.INTEGER().GetText()
	return v.maps.MapAssignment(v.expr, val)
}

func (v *IntegerVisitor) VisitEmptyIntegerOutputRule(ctx *gen.EmptyIntegerOutputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}
