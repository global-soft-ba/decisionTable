package ast

import (
	"bytes"
	"fmt"
	"github.com/global-soft-ba/decisionTable/ast"
	"reflect"
)

type ArithmeticNegation struct {
	ParserRule Rule
	Value      ast.Node
}

func (l ArithmeticNegation) ParserLiteral() string { return l.ParserRule.Literal }
func (l ArithmeticNegation) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("- %s", l.Value.String()))
	return out.String()
}
func (l ArithmeticNegation) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l.Value)
}
func (l ArithmeticNegation) GetChildren() []ast.Node {
	return []ast.Node{l.Value}
}

type SimpleValue struct {
	ParserRule Rule
	Operator   Token
	Value      ast.Node
}

func (l SimpleValue) ParserLiteral() string { return l.ParserRule.Literal }
func (l SimpleValue) String() string {
	var out bytes.Buffer
	out.WriteString(l.Operator.Literal)
	out.WriteString(l.Value.String())
	return out.String()
}
func (l SimpleValue) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l.Value)
}
func (l SimpleValue) GetChildren() []ast.Node {
	return []ast.Node{l.Value}
}

type Parentheses struct {
	ParserRule Rule
	Value      ast.Node
}

func (l Parentheses) ParserLiteral() string { return l.ParserRule.Literal }
func (l Parentheses) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("(%s)", l.Value.String()))
	return out.String()
}
func (l Parentheses) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l.Value)
}
func (l Parentheses) GetChildren() []ast.Node {
	return []ast.Node{l.Value}
}

type SimpleExpression struct {
	ParserRule Rule
	Operator   Token
	Value      ast.Node
}

func (l SimpleExpression) ParserLiteral() string { return l.ParserRule.Literal }
func (l SimpleExpression) String() string {
	var out bytes.Buffer
	out.WriteString(l.Operator.Literal)
	out.WriteString(l.Value.String())
	return out.String()
}
func (l SimpleExpression) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l.Value)
}
func (l SimpleExpression) GetChildren() []ast.Node {
	return []ast.Node{l.Value}
}

type SimpleExpressions struct {
	ParserRule        Rule
	SimpleExpressions []ast.Node
}

func (l SimpleExpressions) ParserLiteral() string {
	return l.ParserRule.Literal
}
func (l SimpleExpressions) String() string {
	var out bytes.Buffer

	for i, val := range l.SimpleExpressions {
		if i > 0 {
			out.WriteString(",")
		}
		out.WriteString(val.String())
	}

	return out.String()
}
func (l SimpleExpressions) GetOperandDataType() reflect.Type {
	return checkDataTypePrecedences(l.SimpleExpressions...)
}
func (l SimpleExpressions) GetChildren() []ast.Node {
	return l.SimpleExpressions
}

type Expression struct {
	ParserRule       Rule
	SimpleExpression ast.Node
}

func (l Expression) ParserLiteral() string {
	return l.ParserRule.Literal
}
func (l Expression) String() string {
	var out bytes.Buffer
	out.WriteString(l.SimpleExpression.String())
	return out.String()
}
func (l Expression) GetOperandDataType() reflect.Type {
	return checkDataTypePrecedences(l.SimpleExpression)
}
func (l Expression) GetChildren() []ast.Node {
	return []ast.Node{l.SimpleExpression}
}

type ArithmeticExpression struct {
	ParserRule Rule
	Left       ast.Node
	Operator   Token
	Right      ast.Node
}

func (l ArithmeticExpression) ParserLiteral() string {
	return l.ParserRule.Literal
}
func (l ArithmeticExpression) String() string {
	var out bytes.Buffer
	out.WriteString(l.Left.String())
	out.WriteString(l.Operator.Literal)
	out.WriteString(l.Right.String())
	return out.String()
}
func (l ArithmeticExpression) GetOperandDataType() reflect.Type {
	return checkDataTypePrecedences(l.Left, l.Right)
}
func (l ArithmeticExpression) GetChildren() []ast.Node {
	return []ast.Node{l.Left, l.Right}
}

type Comparison struct {
	ParserRule Rule
	Left       ast.Node
	Operator   Token
	Right      ast.Node
}

func (l Comparison) ParserLiteral() string { return l.ParserRule.Literal }
func (l Comparison) String() string {
	var out bytes.Buffer
	out.WriteString(l.Left.String())
	out.WriteString(l.Operator.Literal)
	out.WriteString(l.Right.String())
	return out.String()
}
func (l Comparison) GetOperandDataType() reflect.Type {
	return checkDataTypePrecedences(l.Left, l.Right)
}
func (l Comparison) GetChildren() []ast.Node {
	return []ast.Node{l.Left, l.Right}
}
