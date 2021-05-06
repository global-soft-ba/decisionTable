package sfeel

import (
	"decisionTable/converters/grule/grlmodel"
	mapper2 "decisionTable/converters/grule/termconverter/sfeel/mapper"
	"decisionTable/converters/grule/termconverter/sfeel/visitors"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
)

func CreateSfeelTermToGrlConverter() SFeelTermConverter {
	return SFeelTermConverter{mapper2.SettingsGRL}
}

func CreateSfeelTermToJsonConverter() SFeelTermConverter {
	return SFeelTermConverter{mapper2.SettingsJSON}
}

// SFeelTermConverter is a converter between the ParseTree and our grl data model
type SFeelTermConverter struct {
	maps mapper2.TermMapper
}

func (c SFeelTermConverter) ConvertExpression(expr grlmodel.Term) grlmodel.Term {
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
	case model.String:
		tree := prs.Parse().ValidStringInput()
		conv := visitors.CreateStringVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	case model.DateTime:
		tree := prs.Parse().ValidDateTimeInput()
		conv := visitors.CreateDateTimeVisitor(expr, c.maps)
		expr.Expression = tree.Accept(conv).(string)
		return expr
	default:
		return grlmodel.Term{}
	}
}

func (c SFeelTermConverter) ConvertAssignments(expr grlmodel.Term) grlmodel.Term {
	return expr
}
