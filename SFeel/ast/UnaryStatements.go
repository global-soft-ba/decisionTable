package ast

import "bytes"

type UnaryStatements struct {
	ParserRules Rule
	Negation    Rule
	UnaryTests  []Node
}

func (l UnaryStatements) ParserLiteral() string {
	return l.ParserRules.Literal
}

func (l UnaryStatements) String() string {
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
		out2.WriteString("(")
		out2.WriteString(out.String())
		out2.WriteString(")")
	} else {
		out2 = out
	}

	return out2.String()
}

type UnaryStatement struct {
	ParserRule Rule
	Operator   Token
	Value      Node
}

func (l UnaryStatement) ParserLiteral() string { return l.ParserRule.Literal }
func (l UnaryStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.Operator.Literal)
	out.WriteString(l.Value.String())
	return out.String()
}

type Interval struct {
	ParserRule        Rule
	StartIntervalRule Rule
	EndIntervalRule   Rule
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
