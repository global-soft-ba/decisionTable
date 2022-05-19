package conv

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/ast"
	ast2 "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/tmp"
)

var ErrEmptyStatement = errors.New("empty statement")

func CreateSFeelToGrlAstConverter() SFeelToGrlAstConverter {
	listener := CreateSFeelToGrlAstConverterListener()
	return SFeelToGrlAstConverter{listener: listener}
}

type SFeelToGrlAstConverter struct {
	listener SFeelToGrlAstConverterListener
}

func (c SFeelToGrlAstConverter) ConvertToGrlAst(field field.Field, et entryType.EntryType, e string) (ast.Node, error) {
	c.listener.field = field

	var converter data.EntryConverterInterface
	var err error
	if et == entryType.Input {
		converter, err = tmp.CreateInputEntryConverter(expressionLanguage.SFEEL, e)
	} else if et == entryType.Output {
		converter, err = tmp.CreateOutputEntryConverter(expressionLanguage.SFEEL, e)
	}

	if err != nil {
		return nil, err
	}

	converter.Convert(&c.listener)
	grlTree := c.listener.GetAST()
	switch grlTree.(type) {
	case ast2.EmptyStatement:
		return grlTree, ErrEmptyStatement
	}

	return grlTree, nil
}
