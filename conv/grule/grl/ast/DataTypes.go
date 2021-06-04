package ast

import (
	"decisionTable/ast"
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

type Integer struct {
	Val int64
}

func (i Integer) String() string {
	return strconv.FormatInt(i.Val, 10)
}

func (i Integer) GetChildren() []ast.Node {
	return nil
}
