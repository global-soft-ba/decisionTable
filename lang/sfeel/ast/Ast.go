package ast

import (
	"reflect"
)

// The base Node interface which represents abstract syntac trees
type Node interface {
	ParserLiteral() string
	String() string
	GetOperandDataType() reflect.Type
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

func GetAllTreeNodeTypes(node Node) []string {
	if len(node.GetChildren()) == 0 {
		return []string{reflect.TypeOf(node).String()}
	}
	var out []string
	for _, val := range node.GetChildren() {
		qf := GetAllTreeNodeTypes(val)
		if qf != nil {
			out = append(out, qf...)
		}
	}
	out = append(out, reflect.TypeOf(node).String())
	return out
}
