package generate

import (
	"bytes"
	"decisionTable/ast"
	grl "decisionTable/conv/grule/grl/ast"
	templates "decisionTable/conv/grule/grl/generate/grl"
	"errors"
	"text/template"
)

var (
	ErrConverterGRLSymbolNotFound = errors.New("Operator not found in operator symbol table")
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

//ToDo error handling with stack, type casts and templates
type GrlGeneratorListener struct {
	grl.GrlListener
	stack               *ast.Stack
	OperationsTemplates template.Template
	OperatorTable       map[int]string
	Errors              []error
}

func (g *GrlGeneratorListener) GetCode() string {
	return g.stack.Pop().(string)
}

func (g *GrlGeneratorListener) executeTemplate(op int, tmplName string) string {
	var tpl bytes.Buffer

	rightVal := g.stack.Pop()
	leftVal := g.stack.Pop()
	operator := g.executeOperatorMapping(op)

	data := operation{
		Left:  leftVal.(string),
		Op:    operator,
		Right: rightVal.(string),
	}

	g.OperationsTemplates.ExecuteTemplate(&tpl, tmplName, data)

	return tpl.String()
}

func (g *GrlGeneratorListener) executeOperatorMapping(op int) string {
	val, ok := g.OperatorTable[op]
	if !ok {
		g.Errors = append(g.Errors, ErrConverterGRLSymbolNotFound)
	}

	return val
}

func (g *GrlGeneratorListener) ExitLogicalOperations(ctx grl.LogicalOperations) {
	r := g.executeTemplate(ctx.Operator, templates.BinaryOperation)
	g.stack.Push(r)
}

func (g *GrlGeneratorListener) ExitComparisonOperations(ctx grl.ComparisonOperations) {
	r := g.executeTemplate(ctx.Operator, templates.BinaryOperation)
	g.stack.Push(r)
}

func (g *GrlGeneratorListener) ExitInteger(ctx grl.Integer) {
	g.stack.Push(ctx.String())
}

func (g *GrlGeneratorListener) ExitString(ctx grl.String) {
	g.stack.Push(ctx.String())
}
