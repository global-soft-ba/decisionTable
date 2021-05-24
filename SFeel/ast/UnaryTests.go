package ast

import "bytes"

type EmptyUnaryTest struct {
	ParserToken Token
}

func (l EmptyUnaryTest) ParserLiteral() string {
	return l.ParserToken.Literal
}

func (l EmptyUnaryTest) String() string {
	return ""
}

type UnaryTests struct {
	ParserRules Rule
	Negation    Rule
	UnaryTests  []Node
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
