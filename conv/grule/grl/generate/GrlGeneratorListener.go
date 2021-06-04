package generate

import (
	"bytes"
	"decisionTable/ast"
	grl "decisionTable/conv/grule/grl/ast"
	templates "decisionTable/conv/grule/grl/generate/grl"
	"text/template"
)

func CreateGrlGeneratorListener() (GrlGeneratorListener, error) {

	tpl, err := templates.GenerateGrlTemplates()
	if err != nil {
		return GrlGeneratorListener{}, err
	}
	return GrlGeneratorListener{
		OperationsTemplates: tpl,
		OperatorTable:       templates.OperatorTable,
		stack:               ast.NewStack()}, nil
}

type operation struct {
	Left  string
	Op    string
	Right string
}

type GrlGeneratorListener struct {
	grl.GrlListener
	stack               *ast.Stack
	OperationsTemplates template.Template
	OperatorTable       map[int]string
}

func (g GrlGeneratorListener) GetCode() string {
	return g.stack.Pop().(string)
}

func (g GrlGeneratorListener) executeTemplate(op int, tmplName string) string {
	var tpl bytes.Buffer

	rightVal := g.stack.Pop()
	leftVal := g.stack.Pop()
	operator := g.OperatorTable[op]

	data := operation{
		Left:  leftVal.(string),
		Op:    operator,
		Right: rightVal.(string),
	}

	g.OperationsTemplates.ExecuteTemplate(&tpl, tmplName, data)

	return tpl.String()
}

func (g GrlGeneratorListener) ExitLogicalOperations(ctx grl.LogicalOperations) {
	r := g.executeTemplate(ctx.Operator, templates.BinaryOperation)
	g.stack.Push(r)
}

func (g GrlGeneratorListener) ExitComparisonOperations(ctx grl.ComparisonOperations) {
	r := g.executeTemplate(ctx.Operator, templates.BinaryOperation)
	g.stack.Push(r)
}

func (g GrlGeneratorListener) ExitInteger(ctx grl.Integer) {
	g.stack.Push(ctx.String())
}

func (g GrlGeneratorListener) ExitString(ctx grl.String) {
	g.stack.Push(ctx.String())
}
