package ast

import (
	"bytes"
	"github.com/global-soft-ba/decisionTable/ast"
	sfeel "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
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

type QualifiedName struct {
	Val []string
}

func (q QualifiedName) String() string {
	var out bytes.Buffer
	for i, val := range q.Val {
		if i > 0 {
			out.WriteString(sfeel.SFeelSeparatorQualifiedName)
		}
		out.WriteString(val)
	}

	return out.String()
}

func (q QualifiedName) GetQualifiedName() string {
	return q.String()
}

func (q QualifiedName) GetChildren() []ast.Node {
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
