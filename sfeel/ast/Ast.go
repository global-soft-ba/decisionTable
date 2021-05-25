package ast

// The base Node interface
type Node interface {
	ParserLiteral() string
	String() string
}

// SFeelParser tokens
type Token struct {
	Type    int
	Literal string
}

// SFeelParser rules.
type Rule struct {
	Type    int
	Literal string
}