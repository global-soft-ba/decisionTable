package conv

import (
	"decisionTable/ast"
	"decisionTable/data"
)

func CreateSFeelToGrlAstConverter() SFeelToGrlAstConverter {
	listener := CreateSFeelToGrlAstConverterListener()
	return SFeelToGrlAstConverter{listener: listener}
}

type SFeelToGrlAstConverter struct {
	listener SFeelToGrlAstConverterListener
}

func (c SFeelToGrlAstConverter) ConvertToGrlAst(field data.FieldInterface, sfeelEntry data.EntryInterface) (ast.Node, error) {
	c.listener.field = field
	sfeelEntry.Convert(&c.listener)
	grlTree := c.listener.GetAST()

	return grlTree, nil
}
