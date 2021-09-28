package generate

import (
	"bytes"
	"errors"
	"github.com/global-soft-ba/decisionTable/ast"
	grl "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	templates "github.com/global-soft-ba/decisionTable/conv/grule/grl/generate/grl"
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
		stack:               ast.NewStack(),
	}, nil
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

	if err := g.OperationsTemplates.ExecuteTemplate(&tpl, tmplName, data); err != nil {
		g.Errors = append(g.Errors, err)
	}

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

func (g *GrlGeneratorListener) ExitMathOperations(ctx grl.MathOperations) {
	r := g.executeTemplate(ctx.Operator, templates.BinaryOperation)
	g.stack.Push(r)
}

func (g *GrlGeneratorListener) ExitAssignmentOperations(ctx grl.AssignmentOperations) {
	r := g.executeTemplate(ctx.Operator, templates.AssignmentOperation)
	g.stack.Push(r)
}

func (g *GrlGeneratorListener) ExitPowOperation(ctx grl.PowOperation) {
	var tpl bytes.Buffer

	rightVal := g.stack.Pop()
	leftVal := g.stack.Pop()

	if err := g.OperationsTemplates.ExecuteTemplate(&tpl, templates.PowOperation, struct {
		Base string
		Exp  string
	}{
		Base: leftVal.(string),
		Exp:  rightVal.(string),
	}); err != nil {
		g.Errors = append(g.Errors, err)
	}

	g.stack.Push(tpl.String())
}

func (g *GrlGeneratorListener) ExitParentheses(ctx grl.Parentheses) {
	parentheses := g.stack.Pop().(grl.Parentheses)
	g.stack.Push(parentheses.String())
}

func (g *GrlGeneratorListener) ExitArithmeticNegation(ctx grl.ArithmeticNegation) {
	negation := g.stack.Pop().(grl.ArithmeticNegation)
	g.stack.Push(negation.String())
}

func (g *GrlGeneratorListener) ExitInteger(ctx grl.Integer) {
	g.stack.Push(ctx.String())
}

func (g *GrlGeneratorListener) ExitString(ctx grl.String) {
	g.stack.Push(ctx.String())
}

func (g *GrlGeneratorListener) ExitBoolean(ctx grl.Boolean) {
	g.stack.Push(ctx.String())
}

func (g *GrlGeneratorListener) ExitDateTime(ctx grl.DateTime) {
	g.stack.Push(ctx.String())
}

func (g *GrlGeneratorListener) ExitEmptyStatement(ctx grl.EmptyStatement) {

}

func (g *GrlGeneratorListener) ExitFloat(ctx grl.Float) {
	g.stack.Push(ctx.String())
}
