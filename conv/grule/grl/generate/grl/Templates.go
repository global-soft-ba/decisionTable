package grl

import (
	grl "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	"text/template"
)

var (
	OperatorTable = map[int]string{
		grl.AssignmentEQUAL:             "=",
		grl.ComparisonOperatorLESS:      "<",
		grl.ComparisonOperatorLESSEQ:    "<=",
		grl.ComparisonOperatorGREATER:   ">",
		grl.ComparisonOperatorGREATEREQ: ">=",
		grl.LogicalAND:                  "&&",
		grl.LogicalOR:                   "||",
		grl.ComparisonOperatorEQUAL:     "==",
		grl.ComparisonOperatorNOTEQUAL:  "!=",
		grl.MathADD:                     "+",
		grl.MathSUB:                     "-",
		grl.MathDIV:                     "/",
		grl.MathMUL:                     "*",
	}
	PowOperation        = `Pow({{.Base}}, {{ .Exp }})`
	AssignmentOperation = `{{.Left}} {{ .Op }} {{ .Right }}`
	BinaryOperation     = `({{.Left}} {{ .Op }} {{ .Right }})`
	Negation            = `!({{.}})`
)

func GenerateGrlTemplates() (template.Template, error) {
	var t *template.Template

	t, err := template.New(BinaryOperation).Parse(BinaryOperation)
	if err != nil {
		return template.Template{}, err
	}

	_, err = t.New(Negation).Parse(Negation)
	if err != nil {
		return template.Template{}, err
	}

	_, err = t.New(AssignmentOperation).Parse(AssignmentOperation)
	if err != nil {
		return template.Template{}, err
	}

	_, err = t.New(PowOperation).Parse(PowOperation)
	if err != nil {
		return template.Template{}, err
	}

	return *t, err
}
