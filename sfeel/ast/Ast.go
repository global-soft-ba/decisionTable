package ast

import (
	"reflect"
)

// The base Node interface
type Node interface {
	ParserLiteral() string
	String() string
	GetOperandType() reflect.Type
	GetChildren() []Node
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

func GetAllQualifiedNames(node Node) []QualifiedName {
	if (reflect.TypeOf(node) == reflect.TypeOf(QualifiedName{})) {
		return []QualifiedName{node.(QualifiedName)}
	}

	var out []QualifiedName
	for _, val := range node.GetChildren() {
		qf := GetAllQualifiedNames(val)
		if qf != nil {
			out = append(out, qf...)
		}
	}

	return out
}
