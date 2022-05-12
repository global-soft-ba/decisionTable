package conv

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/ast"
	ast2 "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/field"
)

var ErrEmptyStatement = errors.New("empty statement")

func CreateSFeelToGrlAstConverter() SFeelToGrlAstConverter {
	listener := CreateSFeelToGrlAstConverterListener()
	return SFeelToGrlAstConverter{listener: listener}
}

type SFeelToGrlAstConverter struct {
	listener SFeelToGrlAstConverterListener
}

func (c SFeelToGrlAstConverter) ConvertToGrlAst(field field.Field, sfeelEntry data.EntryInterface) (ast.Node, error) {
	c.listener.field = field
	sfeelEntry.Convert(&c.listener)
	grlTree := c.listener.GetAST()
	switch grlTree.(type) {
	case ast2.EmptyStatement:
		return grlTree, ErrEmptyStatement
	}

	return grlTree, nil
}
