package data

import (
	"github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
)

type EntryConverterInterface interface {
	Convert(listener ast.SFeelListenerInterface)
}
