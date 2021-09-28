package ast

import (
	"github.com/global-soft-ba/decisionTable/ast"
	"reflect"
)

// Extend interface for SFeel and Parser Topics
type Node interface {
	ast.Node
	ParserLiteral() string
	GetOperandDataType() reflect.Type
}

func GetAllQualifiedNames(node ast.Node) []QualifiedName {
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

func GetAllTreeNodeTypes(node ast.Node) []string {
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
