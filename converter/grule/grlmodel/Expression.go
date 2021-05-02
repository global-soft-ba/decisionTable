package grlmodel

import "decisionTable/model"

type Expression struct {
	Name       string
	Identifier string
	Typ        model.VariableTyp
	Expression string
}
