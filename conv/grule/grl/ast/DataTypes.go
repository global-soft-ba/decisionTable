package ast

import (
	"bytes"
	"fmt"
	"github.com/global-soft-ba/decisionTable/ast"
	sfeel "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	"strconv"
	"time"
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
	Sign *string
	Val  int64
}

func (i Integer) String() string {
	if i.Sign != nil {
		return fmt.Sprintf("%s%s", *i.Sign, strconv.FormatInt(i.Val, 10))
	}
	return strconv.FormatInt(i.Val, 10)
}

func (i Integer) GetChildren() []ast.Node {
	return nil
}

type Float struct {
	Val float64
}

func (i Float) String() string {
	return strconv.FormatFloat(i.Val, 'g', -1, 64)
}

func (i Float) GetChildren() []ast.Node {
	return nil
}

type Boolean struct {
	Val bool
}

func (i Boolean) String() string {
	return strconv.FormatBool(i.Val)
}

func (i Boolean) GetChildren() []ast.Node {
	return nil
}

type DateTime struct {
	Val time.Time
}

func (i DateTime) String() string {
	return i.Val.String()
}

func (i DateTime) GetChildren() []ast.Node {
	return nil
}

type EmptyStatement struct{}

func (i EmptyStatement) String() string {
	return ""
}

func (i EmptyStatement) GetChildren() []ast.Node {
	return nil
}

type Parentheses struct {
	Value ast.Node
}

func (i Parentheses) String() string {
	return fmt.Sprintf("(%s)", i.Value.String())
}

func (i Parentheses) GetChildren() []ast.Node {
	return []ast.Node{i.Value}
}

type ArithmeticNegation struct {
	Value ast.Node
}

func (i ArithmeticNegation) String() string {
	return fmt.Sprintf("-%s", i.Value.String())
}

func (i ArithmeticNegation) GetChildren() []ast.Node {
	return []ast.Node{i.Value}
}
