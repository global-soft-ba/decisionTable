package ast

import (
	"bytes"
	"strconv"
	"time"
)

type QualifiedName struct {
	ParserRule Rule
	Value      []string
}

func (l QualifiedName) ParserLiteral() string { return l.ParserRule.Literal }
func (l QualifiedName) String() string {
	var out bytes.Buffer
	for i, val := range l.Value {
		if i > 0 {
			out.WriteString(".")
		}
		out.WriteString(val)
	}

	return out.String()
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

type Boolean struct {
	ParserRule Rule
	Value      bool
}

func (l Boolean) ParserLiteral() string { return l.ParserRule.Literal }
func (l Boolean) String() string        { return strconv.FormatBool(l.Value) }

type String struct {
	ParserRule Rule
	Value      string
}

func (l String) ParserLiteral() string { return l.ParserRule.Literal }
func (l String) String() string        { return l.Value }

type DateTime struct {
	ParserRule Rule
	Value      time.Time
}

func (l DateTime) ParserLiteral() string { return l.ParserRule.Literal }
func (l DateTime) String() string        { return l.Value.String() }
