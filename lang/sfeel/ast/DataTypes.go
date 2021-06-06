package ast

import (
	"bytes"
	"decisionTable/ast"
	"reflect"
	"strconv"
	"time"
)

func checkDataTypePrecedence(typ1 ast.Node, typ2 ast.Node) reflect.Type {
	if reflect.TypeOf(typ1) == reflect.TypeOf(typ2) {
		return reflect.TypeOf(typ1)
	}
	switch typ1.(type) {
	case Integer:
		if reflect.TypeOf(typ2) == reflect.TypeOf(Float{}) {
			return reflect.TypeOf(Float{})
		}
	case Float:
		if reflect.TypeOf(typ2) == reflect.TypeOf(Integer{}) {
			return reflect.TypeOf(Float{})
		}
	}
	return nil
}
func checkDataTypePrecedences(types ...ast.Node) reflect.Type {
	length := len(types)

	switch length {
	case 0:
		return nil
	case 1:
		return reflect.TypeOf(types[0])
	default:
		init := types[0]
		for i := 1; i < length; i++ {
			result := checkDataTypePrecedence(init, types[i])
			if result == nil {
				return nil
			} else if reflect.TypeOf(init) != result {
				init = types[i]
			}
		}
		return reflect.TypeOf(init)
	}
}

// SFeelParser rules. - just for verifiction to see if the parser works correct
type Rule struct {
	Type    int
	Literal string
}

type EmptyStatement struct {
	ParserToken Token
}

func (l EmptyStatement) ParserLiteral() string {
	return l.ParserToken.Literal
}
func (l EmptyStatement) String() string {
	return ""
}
func (l EmptyStatement) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l EmptyStatement) GetChildren() []ast.Node {
	return nil
}

type QualifiedName struct {
	ParserRule Rule
	Value      []string
}

func (l QualifiedName) ParserLiteral() string { return l.ParserRule.Literal }
func (l QualifiedName) String() string {
	var out bytes.Buffer
	for i, val := range l.Value {
		if i > 0 {
			out.WriteString(SFeelSeparatorQualifiedName)
		}
		out.WriteString(val)
	}

	return out.String()
}
func (l QualifiedName) GetQualifiedName() string {
	return l.String()
}
func (l QualifiedName) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l QualifiedName) GetChildren() []ast.Node {
	return nil
}

type Integer struct {
	ParserRule Rule
	SignRule   Rule
	Value      int64
}

func (l Integer) ParserLiteral() string {
	return l.SignRule.Literal + l.ParserRule.Literal
}
func (l Integer) String() string {
	var out bytes.Buffer
	if l.SignRule.Type != -1 {
		out.WriteString(l.SignRule.Literal)
	}
	out.WriteString(strconv.FormatInt(l.Value, 10))
	return out.String()
}
func (l Integer) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l Integer) GetChildren() []ast.Node {
	return nil
}

type Float struct {
	ParserRule Rule
	SignRule   Rule
	Value      float64
}

func (l Float) ParserLiteral() string {
	return l.SignRule.Literal + l.ParserRule.Literal
}
func (l Float) String() string {
	var out bytes.Buffer
	if l.SignRule.Type != -1 {
		out.WriteString(l.SignRule.Literal)
	}
	out.WriteString(strconv.FormatFloat(l.Value, 'E', -1, 64))
	return out.String()
}
func (l Float) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l Float) GetChildren() []ast.Node {
	return nil
}

type Boolean struct {
	ParserRule Rule
	Value      bool
}

func (l Boolean) ParserLiteral() string { return l.ParserRule.Literal }
func (l Boolean) String() string        { return strconv.FormatBool(l.Value) }
func (l Boolean) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l Boolean) GetChildren() []ast.Node {
	return nil
}

type String struct {
	ParserRule Rule
	Value      string
}

func (l String) ParserLiteral() string { return l.ParserRule.Literal }
func (l String) String() string        { return l.Value }
func (l String) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l String) GetChildren() []ast.Node {
	return nil
}

type DateTime struct {
	ParserRule Rule
	Value      time.Time
}

func (l DateTime) ParserLiteral() string { return l.ParserRule.Literal }
func (l DateTime) String() string        { return l.Value.String() }
func (l DateTime) GetOperandDataType() reflect.Type {
	return reflect.TypeOf(l)
}
func (l DateTime) GetChildren() []ast.Node {
	return nil
}
