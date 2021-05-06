package sfeel

import (
	maps "decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/expressionlanguages/sfeel/visitors"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
)

func CreateSfeelToGrlConverter() SFeelConverter {
	return SFeelConverter{maps.SettingsGRL}
}

func CreateSfeelToJsonConverter() SFeelConverter {
	return SFeelConverter{maps.SettingsJSON}
}

// SFeelConverter SFeelConverter is a converter between the ParseTree and our grl data model
type SFeelConverter struct {
	maps maps.Mapper
}

func (c SFeelConverter) ConvertExpression(expr grlmodel.Term) grlmodel.Term {
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

func (c SFeelConverter) ConvertAssignments(expr grlmodel.Term) grlmodel.Term {
	return expr
}
