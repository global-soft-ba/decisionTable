package visitors

import (
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/grlmodel"
	gen "decisionTable/parser/sfeel/generated"
)

func CreateStringVisitor(expr grlmodel.Term, maps maps.Mapper) *StringVisitor {
	return &StringVisitor{gen.BaseSFeelVisitor{}, expr, maps}
}

type StringVisitor struct {
	gen.BaseSFeelVisitor
	expr grlmodel.Term
	maps maps.Mapper
}

func (v *StringVisitor) VisitEqualComparisonStringInputRule(ctx *gen.EqualComparisonStringInputRuleContext) interface{} {
	val := ctx.EqualcomparisonStrings().GetText()
	return v.maps.MapEqualComparison(v.expr, val)
}

func (v *StringVisitor) VisitEmptyStringInputRule(ctx *gen.EmptyStringInputRuleContext) interface{} {
	return v.maps.MapEmpty(v.expr)
}

func (v *StringVisitor) VisitDisjunctionsStringInputRule(ctx *gen.DisjunctionsStringInputRuleContext) interface{} {
	entireTerm := ctx.DisjunctionsString().Accept(v)
	return v.maps.MapDisjunctionsTerm(entireTerm)
}

func (v *StringVisitor) VisitDisjunctionsString(ctx *gen.DisjunctionsStringContext) interface{} {
	var result []interface{}

	for _, val := range ctx.AllDisjunctionsString() {
		resString := val.Accept(v)
		result = append(result, resString)
	}

	if ctx.Strings() != nil {
		term := v.maps.MapEqualComparison(v.expr, ctx.Strings().GetText())
		return v.maps.MapDisjunctionsTerm(term)
	}

	return v.maps.MapDisjunctions(result)
}

func (v *StringVisitor) VisitNegationStringInputRule(ctx *gen.NegationStringInputRuleContext) interface{} {
	var term interface{}

	if ctx.EqualcomparisonStrings() != nil {
		val := ctx.EqualcomparisonStrings().GetText()
		term = v.maps.MapEqualComparison(v.expr, val)
	}

	if ctx.DisjunctionsString() != nil {
		term = ctx.DisjunctionsString().Accept(v)
	}

	return v.maps.MapNegation(term)
}

func (v *StringVisitor) VisitStrings(ctx *gen.StringsContext) interface{} {
	return ctx.STRING()
}
