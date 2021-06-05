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

func (c SFeelToGrlAstConverter) ConvertToGrlAst(fieldName string, sfeelEntry data.EntryInterface) (ast.Node, error) {
	c.listener.fieldName = fieldName
	sfeelEntry.Convert(&c.listener)
	grlTree := c.listener.GetAST()

	return grlTree, nil
}
