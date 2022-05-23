package sfeel

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/lang/sfeel/antlr"
	sfeel "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
)

type EntryConverter struct {
	ast sfeel.Node
}

func CreateInputEntryConverter(entry string) data.EntryConverterInterface {
	tree, err := antlr.CreateParser(entry).ParseInput()
	if err != nil {
		return EntryConverter{ast: nil}
	}

	return EntryConverter{ast: tree}
}

func CreateOutputEntryConverter(entry string) data.EntryConverterInterface {
	tree, err := antlr.CreateParser(entry).ParseOutput()
	if err != nil {
		return EntryConverter{ast: nil}
	}

	return EntryConverter{ast: tree}
}

func (e EntryConverter) Convert(listener sfeel.SFeelListenerInterface) {
	treeWalker := sfeel.CreateSFeelTreeWalker(listener)
	treeWalker.Walk(e.ast)
}
