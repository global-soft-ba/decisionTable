package ast

import (
	"decisionTable/ast"
	"decisionTable/data"
	"strconv"
)

type String struct {
	Val string
}

func (s String) String() string {
	return s.Val
}

func (s String) GetChildren() []ast.Node {
	return nil
}

//TODo Qualified Name and FieldInterface
type QualifiedName struct {
	val data.FieldInterface
}

type Integer struct {
	Val int64
}

func (i Integer) String() string {
	return strconv.FormatInt(i.Val, 10)
}

func (i Integer) GetChildren() []ast.Node {
	return nil
}
