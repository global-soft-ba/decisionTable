package ast

import (
	"bytes"
	"decisionTable/ast"
	"reflect"
)

type UnaryTests struct {
	ParserRules Rule
	Negation    Rule
	UnaryTests  []ast.Node
}

func (l UnaryTests) ParserLiteral() string {
	return l.ParserRules.Literal
}
func (l UnaryTests) String() string {
	var out bytes.Buffer

	for i, val := range l.UnaryTests {
		if i > 0 {
			out.WriteString(",")
		}
		out.WriteString(val.String())
	}

	var out2 bytes.Buffer
	if l.Negation.Type != -1 {
		out2.WriteString(l.Negation.Literal)
		out2.WriteString(out.String())
		out2.WriteString(")")
	} else {
		out2 = out
	}

	return out2.String()
}
func (l UnaryTests) GetOperandDataType() reflect.Type {
	var result []ast.Node
	for _, val := range l.UnaryTests {
		ut := val.(UnaryTest)
		result = append(result, ut.Value)
	}
	return checkDataTypePrecedences(result...)
}
func (l UnaryTests) GetChildren() []ast.Node {
	return l.UnaryTests
}

type UnaryTest struct {
	ParserRule Rule
	Operator   Token
	Value      Node
}

func (l UnaryTest) ParserLiteral() string { return l.ParserRule.Literal }
func (l UnaryTest) String() string {
	var out bytes.Buffer
	out.WriteString(l.Operator.Literal)
	out.WriteString(l.Value.String())
	return out.String()
}
func (l UnaryTest) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l.Value)
}
func (l UnaryTest) GetChildren() []ast.Node {
	return []ast.Node{l.Value}
}

type Interval struct {
	ParserRule        Rule
	StartIntervalRule Token
	EndIntervalRule   Token
	StartValue        Node
	EndValue          Node
}

func (l Interval) ParserLiteral() string { return l.ParserRule.Literal }
func (l Interval) String() string {
	var out bytes.Buffer

	out.WriteString(l.StartIntervalRule.Literal)
	out.WriteString(l.StartValue.String())
	out.WriteString("..")
	out.WriteString(l.EndValue.String())
	out.WriteString(l.EndIntervalRule.Literal)
	return out.String()
}
func (l Interval) GetOperandDataType() reflect.Type {
	return checkDataTypePrecedence(l.StartValue, l.EndValue)
}
func (l Interval) GetChildren() []ast.Node {
	return []ast.Node{l.StartValue, l.EndValue}
}
