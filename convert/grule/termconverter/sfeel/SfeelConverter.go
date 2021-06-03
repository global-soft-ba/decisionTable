package sfeel

import (
	grlmodel2 "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/grl/symbols"
	"decisionTable/convert/grule/termconverter/sfeel/visitors"
	"decisionTable/data"
	"decisionTable/parser/sfeel/parser"
)

func CreateSfeelTermToGrlConverter() SFeelTermConverter {
	return SFeelTermConverter{symbols.SettingsGRL}
}

func CreateSfeelTermToJsonConverter() SFeelTermConverter {
	return SFeelTermConverter{symbols.SettingsJSON}
}

// SFeelTermConverter is a converter between the ParseTree and our grl data model
type SFeelTermConverter struct {
	maps symbols.TermMapper
}

func (c SFeelTermConverter) ConvertExpression(expr grlmodel2.Term) grlmodel2.Term {
	prs := parser.CreateSfeelParser(expr.Expression)

	switch expr.Typ {
	case data.Integer:
		tree := prs.Parse().ValidIntegerInput()
		conv := visitors.CreateIntegerVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.Float:
		tree := prs.Parse().ValidNumberInput()
		conv := visitors.CreateNumberVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.String:
		tree := prs.Parse().ValidStringInput()
		conv := visitors.CreateStringVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.Boolean:
		tree := prs.Parse().ValidBoolInput()
		conv := visitors.CreateBoolVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.DateTime:
		tree := prs.Parse().ValidDateTimeInput()
		conv := visitors.CreateDateTimeVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	default:
		return grlmodel2.Term{}
	}
}

func (c SFeelTermConverter) ConvertAssignments(expr grlmodel2.Term) grlmodel2.Term {
	prs := parser.CreateSfeelParser(expr.Expression)

	switch expr.Typ {
	case data.Integer:
		tree := prs.Parse().ValidIntegerOutput()
		conv := visitors.CreateIntegerVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.Float:
		tree := prs.Parse().ValidNumberOutput()
		conv := visitors.CreateNumberVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.String:
		tree := prs.Parse().ValidStringOutput()
		conv := visitors.CreateStringVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.Boolean:
		tree := prs.Parse().ValidBoolOutput()
		conv := visitors.CreateBoolVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case data.DateTime:
		tree := prs.Parse().ValidDateTimeOutput()
		conv := visitors.CreateDateTimeVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	default:
		return grlmodel2.Term{}
	}
}
