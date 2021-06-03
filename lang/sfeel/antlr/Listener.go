package antlr

import (
	"decisionTable/lang/sfeel/antlr/errors"
	ast "decisionTable/lang/sfeel/ast"
	"decisionTable/lang/sfeel/gen"
	"reflect"
	"strconv"
)

func CreateListener() Listener {
	listener := Listener{stack: newStack(), tokenMap: TokenTable}
	return listener
}

type Listener struct {
	parser.BaseSFeelListener
	tokenMap map[int]int
	Errors   []error
	stack    *stack
}

func (s *Listener) GetAST() ast.Node {
	if s.stack.length == 0 {
		return nil
	}
	return s.stack.Pop().(ast.Node)
}

//Unary Tests
func (s *Listener) ExitEmptySimpleUnaryTests(ctx *parser.EmptySimpleUnaryTestsContext) {
	lit := ctx.GetStart().GetText()
	tkn := ast.Token{
		Type:    -1,
		Literal: lit,
	}
	empty := ast.EmptyStatement{ParserToken: tkn}
	s.stack.Push(empty)
}

func (s *Listener) ExitNegationSimpleUnaryTests(ctx *parser.NegationSimpleUnaryTestsContext) {
	const NEGATIONTOKEN = "not("
	neg := ctx.GetStart().GetText()

	if neg == NEGATIONTOKEN {
		val := s.stack.Pop().(ast.UnaryTests)
		val.Negation = ast.Rule{
			Type:    parser.SFeelParserRULE_simple_unary_tests,
			Literal: neg,
		}
		s.stack.Push(val)
	}
}

func (s *Listener) ExitSimple_positive_unary_tests(ctx *parser.Simple_positive_unary_testsContext) {
	prs := ast.Rule{
		Type:    parser.SFeelParserRULE_simple_positive_unary_tests,
		Literal: ctx.GetText(),
	}

	suT := ast.UnaryTests{ParserRules: prs, Negation: ast.Rule{Type: -1}}
	length := s.stack.Len()
	for i := 0; i < length; i++ {
		val := s.stack.Pop().(ast.Node)
		suT.UnaryTests = append([]ast.Node{val}, suT.UnaryTests...)
	}

	s.stack.Push(suT)
}

func (s *Listener) ExitEqualUnaryComparison(ctx *parser.EqualUnaryComparisonContext) {
	rule := ast.Rule{Type: parser.SFeelParserRULE_unary_comparison, Literal: ctx.GetText()}
	val := s.stack.Pop()

	n := ast.UnaryTest{
		ParserRule: rule,
		Operator:   ast.Token{Type: -1},
		Value:      val.(ast.Node),
	}
	s.stack.Push(n)
}

func (s *Listener) ExitUnaryComparison(ctx *parser.UnaryComparisonContext) {
	prs := ast.Rule{Type: parser.SFeelParserRULE_unary_comparison, Literal: ctx.GetText()}
	op := ast.Token{Type: s.tokenMap[ctx.GetStart().GetTokenType()], Literal: ctx.GetStart().GetText()}
	val := s.stack.Pop()

	n := ast.UnaryTest{
		ParserRule: prs,
		Operator:   op,
		Value:      val.(ast.Node),
	}
	s.stack.Push(n)
}

func (s *Listener) ExitInterval(ctx *parser.IntervalContext) {
	prs := ast.Rule{
		Type:    parser.SFeelParserRULE_interval,
		Literal: ctx.GetText(),
	}
	endToken := s.stack.Pop().(ast.Token)
	endVal := s.stack.Pop().(ast.Node)
	startVal := s.stack.Pop().(ast.Node)
	startToken := s.stack.Pop().(ast.Token)

	inVal := ast.Interval{
		ParserRule:        prs,
		StartIntervalRule: startToken,
		EndIntervalRule:   endToken,
		StartValue:        startVal,
		EndValue:          endVal,
	}

	s.stack.Push(inVal)

}

func (s *Listener) ExitOpen_interval_start(ctx *parser.Open_interval_startContext) {
	r := ast.Token{Type: s.tokenMap[parser.SFeelParserRULE_open_interval_start], Literal: ctx.GetText()}
	s.stack.Push(r)
}

func (s *Listener) ExitClosed_interval_start(ctx *parser.Closed_interval_startContext) {
	r := ast.Token{Type: s.tokenMap[parser.SFeelParserRULE_closed_interval_start], Literal: ctx.GetText()}
	s.stack.Push(r)
}

func (s *Listener) ExitOpen_interval_end(ctx *parser.Open_interval_endContext) {
	r := ast.Token{Type: s.tokenMap[parser.SFeelParserRULE_open_interval_end], Literal: ctx.GetText()}
	s.stack.Push(r)
}

func (s *Listener) ExitClosed_interval_end(ctx *parser.Closed_interval_endContext) {
	r := ast.Token{Type: s.tokenMap[parser.SFeelParserRULE_closed_interval_end], Literal: ctx.GetText()}
	s.stack.Push(r)
}

// Datatypes and Primitives
func (s *Listener) ExitNumeric_literal(ctx *parser.Numeric_literalContext) {
	sign := ctx.GetStart().GetTokenType()

	if sign == parser.SFeelParserSUB {
		nType := s.stack.Peek()
		switch nType.(type) {
		case ast.Integer:
			n := s.stack.Pop()
			negInt := n.(ast.Integer)
			negInt.SignRule = ast.Rule{Type: parser.SFeelParserRULE_numeric_literal, Literal: ctx.GetStart().GetText()}
			s.stack.Push(negInt)
		case ast.Float:
			n := s.stack.Pop()
			negReal := n.(ast.Float)
			negReal.SignRule = ast.Rule{Type: parser.SFeelParserRULE_numeric_literal, Literal: ctx.GetStart().GetText()}
			s.stack.Push(negReal)
		default:
			s.Errors = append(s.Errors, errors.NewError("unknown numeric data type: %s ", reflect.TypeOf(nType)))
		}
	}
}

func (s *Listener) ExitQualified_name(ctx *parser.Qualified_nameContext) {
	rule := ast.Rule{Type: parser.SFeelParserRULE_qualified_name, Literal: ctx.GetText()}
	var names []string
	for _, val := range ctx.AllName() {
		names = append(names, val.GetText())
	}

	q := ast.QualifiedName{ParserRule: rule, Value: names}
	s.stack.Push(q)
}

func (s *Listener) ExitInteger_literal(ctx *parser.Integer_literalContext) {
	rule := ast.Rule{Type: parser.SFeelParserRULE_integer_literal, Literal: ctx.GetText()}
	val, err := strconv.ParseInt(rule.Literal, 10, 64)

	if err != nil {
		s.Errors = append(s.Errors, err)
	} else {
		n := ast.Integer{ParserRule: rule, SignRule: ast.Rule{Type: -1}, Value: val}
		s.stack.Push(n)
	}
}

func (s *Listener) ExitReal_literal(ctx *parser.Real_literalContext) {
	rule := ast.Rule{Type: parser.SFeelParserRULE_real_literal, Literal: ctx.GetText()}
	val, err := strconv.ParseFloat(rule.Literal, 64)

	if err != nil {
		s.Errors = append(s.Errors, err)
	} else {
		n := ast.Float{ParserRule: rule, SignRule: ast.Rule{Type: -1}, Value: val}
		s.stack.Push(n)
	}
}

func (s *Listener) ExitString_literal(ctx *parser.String_literalContext) {
	rule := ast.Rule{Type: parser.SFeelParserRULE_string_literal, Literal: ctx.GetText()}
	val := rule.Literal
	n := ast.String{ParserRule: rule, Value: val}
	s.stack.Push(n)
}

func (s *Listener) ExitBoolean_literal(ctx *parser.Boolean_literalContext) {
	rule := ast.Rule{Type: parser.SFeelParserRULE_boolean_literal, Literal: ctx.GetText()}
	val, err := strconv.ParseBool(rule.Literal)

	if err != nil {
		s.Errors = append(s.Errors, err)
	} else {
		n := ast.Boolean{ParserRule: rule, Value: val}
		s.stack.Push(n)
	}
}

func (s *Listener) ExitDate_time_literal(ctx *parser.Date_time_literalContext) {
	//ToDo Date-Time Datatype
	panic("implement me")
}
