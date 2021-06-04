package ast

import (
	"bytes"
	"decisionTable/ast"
	"strconv"
)

type MathOperations struct {
	Left     ast.Node
	Operator int
	Right    ast.Node
}

func (m MathOperations) String() string {
	var out bytes.Buffer
	out.WriteString(m.Left.String())
	out.WriteString(strconv.Itoa(m.Operator))
	out.WriteString(m.Right.String())
	return out.String()
}

func (m MathOperations) GetChildren() []ast.Node {
	return []ast.Node{m.Left, m.Right}
}

type LogicalOperations struct {
	Left     ast.Node
	Operator int
	Right    ast.Node
}

func (l LogicalOperations) String() string {
	var out bytes.Buffer
	out.WriteString(l.Left.String())
	out.WriteString(" ")

	out.WriteString("opId:")
	op := strconv.Itoa(l.Operator)
	out.WriteString(op)

	out.WriteString(" ")
	out.WriteString(l.Right.String())
	return out.String()
}

func (l LogicalOperations) GetChildren() []ast.Node {
	return []ast.Node{l.Left, l.Right}
}

type ComparisonOperations struct {
	Left     ast.Node
	Operator int
	Right    ast.Node
}

func (c ComparisonOperations) String() string {
	var out bytes.Buffer
	out.WriteString(c.Left.String())
	out.WriteString(" ")
	out.WriteString("opId:")
	op := strconv.Itoa(c.Operator)
	out.WriteString(op)
	out.WriteString(" ")
	out.WriteString(c.Right.String())
	return out.String()
}

func (c ComparisonOperations) GetChildren() []ast.Node {
	return []ast.Node{c.Left, c.Right}
}
