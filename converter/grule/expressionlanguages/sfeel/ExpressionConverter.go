package sfeel

import (
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/expressionlanguages/sfeel/visitors"
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
	maps maps.Mapper
}

// TODO Seperate Input and Output! (Maybe the input works as well as with output)
func (c ExpressionConverter) Convert(expr grlmodel.Term) grlmodel.Term {
	prs := parser.CreateSfeelParser(expr.Expression)

	switch expr.Typ {
	case model.Integer:
		tree := prs.Parse().ValidIntegerInput()
		conv := visitors.CreateIntegerVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case model.Float:
		tree := prs.Parse().ValidNumberInput()
		conv := visitors.CreateNumberVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	default:
		return grlmodel.Term{}
	}
}
