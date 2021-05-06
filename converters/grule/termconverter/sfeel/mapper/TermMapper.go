package mapper

import (
	"bytes"
	"decisionTable/converters/grule/grlmodel"
	"errors"
	"regexp"
	"strconv"
	"text/template"
)

var (
	ErrExpressionLanguageDateTimeFormat = errors.New("could not format datetime value correct")
)

const (

	// Operator token
	EQUAL  = 1
	AND    = 2
	OR     = 3
	ASSIGN = 4

	//Operation Templates
	NEGATION        = 4
	EQUALCOMPARISON = 5
	COMPARISON      = 6
	RANGES          = 7
	DISJUNCTIONS    = 8
	DISJUNCTIONTERM = 9
	DATEANDTIME     = 10
	ASSIGNMENT      = 11
)

type TemplateData struct {
	Expr grlmodel.Term
	Op   string
	Val  string
}

type TermMapper struct {
	TargetToken         map[int]string
	ComparisonOperators map[int]string
	StartRanges         map[int]string
	EndRanges           map[int]string
	Templates           map[int]string
}

func (m TermMapper) executeTemplate(tmpl string, data interface{}) string {
	var t *template.Template

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err.Error()
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)

	if err != nil {
		return err.Error()
	}
	return tpl.String()
}

func (m TermMapper) MapEmpty(expr grlmodel.Term) string {
	return ""
}

func (m TermMapper) MapComparison(expr grlmodel.Term, tokenTyp int, val string) string {
	tmpl := m.Templates[COMPARISON]
	op := m.ComparisonOperators[tokenTyp]

	data := TemplateData{expr, op, val}
	return m.executeTemplate(tmpl, data)
}

func (m TermMapper) MapEqualComparison(expr grlmodel.Term, val string) string {
	op := m.TargetToken[EQUAL]
	tmpl := m.Templates[EQUALCOMPARISON]

	data := TemplateData{expr, op, val}
	return m.executeTemplate(tmpl, data)
}

func (m TermMapper) MapRange(expr grlmodel.Term, opStart int, valStart string, opEnd int, valEnd string) string {
	tmpl := m.Templates[RANGES]
	ops := m.StartRanges[opStart]
	ope := m.EndRanges[opEnd]

	start := TemplateData{expr, ops, valStart}
	end := TemplateData{expr, ope, valEnd}

	data := map[string]TemplateData{"start": start, "end": end}
	return m.executeTemplate(tmpl, data)
}

func (m TermMapper) MapDisjunctions(terms []interface{}) string {
	tmpl := m.Templates[DISJUNCTIONS]
	data := map[string]string{
		"first":  terms[0].(string),
		"second": terms[1].(string),
	}
	return m.executeTemplate(tmpl, data)
}

func (m TermMapper) MapDisjunctionsTerm(term interface{}) string {
	tmpl := m.Templates[DISJUNCTIONTERM]
	return m.executeTemplate(tmpl, term.(string))
}

func (m TermMapper) MapNegation(term interface{}) string {
	tmpl := m.Templates[NEGATION]
	return m.executeTemplate(tmpl, term.(string))
}

func (m TermMapper) MapDateAndTimeFormat(expr string) string {
	regex := regexp.MustCompile(`\D+`)
	split := regex.Split(expr, -1)

	var format []int
	for _, val := range split {
		s, _ := strconv.Atoi(val)
		format = append(format, s)
	}

	if len(format) == 0 || len(format) < 6 {
		return ErrExpressionLanguageDateTimeFormat.Error()
	}

	tmpl := m.Templates[DATEANDTIME]
	data := map[string]int{
		"Year":    format[1],
		"Month":   format[2],
		"Day":     format[3],
		"Hour":    format[4],
		"Minutes": format[5],
		"Seconds": format[6],
	}
	return m.executeTemplate(tmpl, data)
}

func (m TermMapper) MapAssignment(expr grlmodel.Term, val string) string {
	tmpl := m.Templates[ASSIGNMENT]
	op := m.TargetToken[ASSIGN]

	data := TemplateData{expr, op, val}
	return m.executeTemplate(tmpl, data)
}
