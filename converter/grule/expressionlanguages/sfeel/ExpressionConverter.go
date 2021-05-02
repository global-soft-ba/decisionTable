package sfeel

import (
	"decisionTable/converter/grule/expressionlanguages/sfeel/converters"
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mappings"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
)

func CreateGrlExpressionConverter() ExpressionConverter {
	return ExpressionConverter{maps.GrlMapping}
}

func CreateJsonExpressionConverter() ExpressionConverter {
	panic("not implemented")
}

// ExpressionConverter ExpressionConverter is a converter between the ParseTree and our grl data model
type ExpressionConverter struct {
	maps maps.ConverterMapping
}

func (c ExpressionConverter) Convert(expr grlmodel.Expression) grlmodel.Expression {
	prs := parser.CreateSfeelParser(expr.Expression)

	switch expr.Typ {
	case model.Integer:
		tree := prs.Parse().ValidIntegerInput()
		conv := converters.CreateIntegerConverter(c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return grlmodel.Expression{}
	default:
		return grlmodel.Expression{}
	}
}
