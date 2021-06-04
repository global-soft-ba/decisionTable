package conv

import (
	grl "decisionTable/conv/grule/grl/ast"
	"decisionTable/conv/grule/grl/generate"
	"decisionTable/data"
)

func CreateSFeelToGrlConverter() SFeelToGrlConverter {
	listener := CreateSFeelToGrlAstConverterListener()
	return SFeelToGrlConverter{listener: listener}
}

type SFeelToGrlConverter struct {
	listener SFeelToGrlAstConverterListener
}

func (c SFeelToGrlConverter) Convert(fieldName string, sfeelEntry data.EntryInterface) (string, error) {
	c.listener.fieldName = fieldName

	// Build grl ast data model from sfeel ast model
	sfeelEntry.Convert(&c.listener)
	grlTree := c.listener.GetAST()

	// Walk the grl ast tree and convert
	gen, err := generate.CreateGrlGeneratorListener()

	if err != nil {
		return "", err
	}

	walker := grl.CreateGRLTreeWalker(gen)
	walker.Walk(grlTree)

	return gen.GetCode(), nil
}
