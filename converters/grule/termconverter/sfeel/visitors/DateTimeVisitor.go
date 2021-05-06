package visitors

import (
	"decisionTable/converters/grule/grlmodel"
	mapper2 "decisionTable/converters/grule/termconverter/sfeel/mapper"
	gen "decisionTable/parser/sfeel/generated"
)

func CreateDateTimeVisitor(expr grlmodel.Term, maps mapper2.TermMapper) *DateTimeVisitor {
	return &DateTimeVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

type DateTimeVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel.Term
	maps mapper2.TermMapper
}

func (v *DateTimeVisitor) VisitEqualComparisonDateTimeInputRule(ctx *gen.EqualComparisonDateTimeInputRuleContext) interface{} {
	val := ctx.EqualcomparisonDateTime().GetText()
	format := v.maps.MapDateAndTimeFormat(val)
	return v.maps.MapEqualComparison(v.expr, format)
}

func (v *DateTimeVisitor) VisitComparisonDateTimeInputRule(ctx *gen.ComparisonDateTimeInputRuleContext) interface{} {
	return ctx.ComparisonDateTime().Accept(v)
}

func (v *DateTimeVisitor) VisitComparisonDateTime(ctx *gen.ComparisonDateTimeContext) interface{} {
	tokenType := ctx.ComparisonOps().Accept(v).(int)
	val := ctx.Datetime().GetText()
	format := v.maps.MapDateAndTimeFormat(val)
	return v.maps.MapComparison(v.expr, tokenType, format)
}

func (v *DateTimeVisitor) VisitComparisonOps(ctx *gen.ComparisonOpsContext) interface{} {
	return ctx.GetStart().GetTokenType()
}

func (v *DateTimeVisitor) VisitEmptyDateTimeInputRule(ctx *gen.EmptyDateTimeInputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}

func (v *DateTimeVisitor) VisitRangeComparisonDateTimeInputRule(ctx *gen.RangeComparisonDateTimeInputRuleContext) interface{} {
	return ctx.RangeDateTime().Accept(v)
}

func (v *DateTimeVisitor) VisitRangeDateTime(ctx *gen.RangeDateTimeContext) interface{} {
	opStart := ctx.GetStart().GetTokenType()
	opEnd := ctx.GetStop().GetTokenType()
	valStart := v.maps.MapDateAndTimeFormat(ctx.Datetime(0).GetText())
	valEnd := v.maps.MapDateAndTimeFormat(ctx.Datetime(1).GetText())

	return v.maps.MapRange(v.expr, opStart, valStart, opEnd, valEnd)
}

func (v *DateTimeVisitor) VisitDisjunctionsDateTimeInputRule(ctx *gen.DisjunctionsDateTimeInputRuleContext) interface{} {
	entireTerm := ctx.DisjunctionsDateTime().Accept(v)
	return v.maps.MapDisjunctionsTerm(entireTerm)
}

func (v *DateTimeVisitor) VisitDisjunctionsDateTime(ctx *gen.DisjunctionsDateTimeContext) interface{} {
	var result []interface{}

	for _, val := range ctx.AllDisjunctionsDateTime() {
		resString := val.Accept(v)
		result = append(result, resString)
	}

	if ctx.Datetime() != nil {
		format := v.maps.MapDateAndTimeFormat(ctx.Datetime().GetText())
		term := v.maps.MapEqualComparison(v.expr, format)
		return v.maps.MapDisjunctionsTerm(term)
	}

	if ctx.ComparisonDateTime() != nil {
		comp := ctx.ComparisonDateTime().Accept(v)
		return v.maps.MapDisjunctionsTerm(comp)
	}

	if ctx.RangeDateTime() != nil {
		rng := ctx.RangeDateTime().Accept(v)
		return v.maps.MapDisjunctionsTerm(rng)
	}

	return v.maps.MapDisjunctions(result)
}

func (v *DateTimeVisitor) VisitNegationDateTimeInputRule(ctx *gen.NegationDateTimeInputRuleContext) interface{} {
	var term interface{}

	if ctx.EqualcomparisonDateTime() != nil {
		val := ctx.EqualcomparisonDateTime().GetText()
		format := v.maps.MapDateAndTimeFormat(val)
		term = v.maps.MapEqualComparison(v.expr, format)
	}

	if ctx.ComparisonDateTime() != nil {
		term = ctx.ComparisonDateTime().Accept(v)
	}

	if ctx.RangeDateTime() != nil {
		term = ctx.RangeDateTime().Accept(v)
	}

	if ctx.DisjunctionsDateTime() != nil {
		term = ctx.DisjunctionsDateTime().Accept(v)
	}

	return v.maps.MapNegation(term)
}

// Assignment Rules
func (v *DateTimeVisitor) VisitDateTimeAssignmentOutputRule(ctx *gen.DateTimeAssignmentOutputRuleContext) interface{} {
	val := ctx.Datetime().GetText()
	format := v.maps.MapDateAndTimeFormat(val)
	return v.maps.MapAssignment(v.expr, format)
}

func (v *DateTimeVisitor) VisitEmptyDateTimeOutputRule(ctx *gen.EmptyDateTimeOutputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}
