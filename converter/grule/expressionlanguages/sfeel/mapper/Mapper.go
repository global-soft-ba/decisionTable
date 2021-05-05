package mapper

import (
	"bytes"
	"decisionTable/converter/grule/grlmodel"
	"regexp"
	"strconv"
	"text/template"
)

const (

	//Logical operator token
	EQUAL = 1
	AND   = 2
	OR    = 3

	//Operation Templates
	NEGATION        = 4
	EQUALCOMPARISON = 5
	COMPARISON      = 6
	RANGES          = 7
	DISJUNCTIONS    = 8
	DISJUNCTIONTERM = 9
	DATEANDTIME     = 10
)

type TemplateData struct {
	Expr grlmodel.Term
	Op   string
	Val  string
}

type Mapper struct {
	TargetToken         map[int]string
	ComparisonOperators map[int]string
	StartRanges         map[int]string
	EndRanges           map[int]string
	Templates           map[int]string
}

func (m Mapper) executeTemplate(tmpl string, data interface{}) string {
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

func (m Mapper) MapEmpty(expr grlmodel.Term) string {
	return ""
}

func (m Mapper) MapComparison(expr grlmodel.Term, tokenTyp int, val string) string {
	tmpl := m.Templates[COMPARISON]
	op := m.ComparisonOperators[tokenTyp]

	data := TemplateData{expr, op, val}
	return m.executeTemplate(tmpl, data)
}

func (m Mapper) MapEqualComparison(expr grlmodel.Term, val string) string {
	op := m.TargetToken[EQUAL]
	tmpl := m.Templates[EQUALCOMPARISON]

	data := TemplateData{expr, op, val}
	return m.executeTemplate(tmpl, data)
}

func (m Mapper) MapRange(expr grlmodel.Term, opStart int, valStart string, opEnd int, valEnd string) string {
	tmpl := m.Templates[RANGES]
	ops := m.StartRanges[opStart]
	ope := m.EndRanges[opEnd]

	start := TemplateData{expr, ops, valStart}
	end := TemplateData{expr, ope, valEnd}

	data := map[string]TemplateData{"start": start, "end": end}
	return m.executeTemplate(tmpl, data)
}

func (m Mapper) MapDisjunctions(terms []interface{}) string {
	tmpl := m.Templates[DISJUNCTIONS]
	data := map[string]string{
		"first":  terms[0].(string),
		"second": terms[1].(string),
	}
	return m.executeTemplate(tmpl, data)
}

func (m Mapper) MapDisjunctionsTerm(term interface{}) string {
	tmpl := m.Templates[DISJUNCTIONTERM]
	return m.executeTemplate(tmpl, term.(string))
}

func (m Mapper) MapNegation(term interface{}) string {
	tmpl := m.Templates[NEGATION]
	return m.executeTemplate(tmpl, term.(string))
}

func (m Mapper) MapDateAndTimeFormat(expr string) string {
	//regex := regexp.MustCompile(`DateAndTime\("|-|T|:|"\)`)
	regex := regexp.MustCompile(`\D+`)
	split := regex.Split(expr, -1)

	var format []int
	for _, val := range split {
		s, _ := strconv.Atoi(val)
		format = append(format, s)
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
