package antlr

import (
	ast "decisionTable/SFeel/ast"
	"decisionTable/SFeel/errors"
	gen "decisionTable/SFeel/gen"
	"reflect"
	"strconv"
)

func CreateSFeelListener() SFeelListener {
	listener := SFeelListener{stack: newStack()}
	return listener
}

type SFeelListener struct {
	gen.BaseSFeelListener
	Errors []error
	stack  *stack
}

func (s *SFeelListener) GetAST() ast.Node {
	if s.stack.length == 0 {
		return nil
	}
	return s.stack.Pop().(ast.Node)
}

func (s *SFeelListener) ExitEmptySimpleUnaryTests(ctx *gen.EmptySimpleUnaryTestsContext) {
	lit := ctx.GetStart().GetText()
	tkn := ast.Token{
		Type:    gen.SFeelParserRULE_simple_unary_tests,
		Literal: lit,
	}
	empty := ast.EmptyUnaryTest{ParserToken: tkn}
	s.stack.Push(empty)
}

func (s *SFeelListener) ExitNegationSimpleUnaryTests(ctx *gen.NegationSimpleUnaryTestsContext) {
	const NEGATIONTOKEN = "not("
	neg := ctx.GetStart().GetText()

	if neg == NEGATIONTOKEN {
		val := s.stack.Pop().(ast.UnaryTests)
		val.Negation = ast.Rule{
			Type:    gen.SFeelParserRULE_simple_unary_tests,
			Literal: neg,
		}
		s.stack.Push(val)
	}
}

func (s *SFeelListener) ExitSimple_positive_unary_tests(ctx *gen.Simple_positive_unary_testsContext) {
	prs := ast.Rule{
		Type:    gen.SFeelParserRULE_simple_positive_unary_tests,
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

func (s *SFeelListener) ExitEqualUnaryComparison(ctx *gen.EqualUnaryComparisonContext) {
	rule := ast.Rule{Type: gen.SFeelParserRULE_unary_comparison, Literal: ctx.GetText()}
	val := s.stack.Pop()

	n := ast.UnaryTest{
		ParserRule: rule,
		Operator:   ast.Token{Type: -1},
		Value:      val.(ast.Node),
	}
	s.stack.Push(n)
}

func (s *SFeelListener) ExitUnaryComparison(ctx *gen.UnaryComparisonContext) {
	prs := ast.Rule{Type: gen.SFeelParserRULE_unary_comparison, Literal: ctx.GetText()}
	op := ast.Token{Type: ctx.GetStart().GetTokenType(), Literal: ctx.GetStart().GetText()}
	val := s.stack.Pop()

	n := ast.UnaryTest{
		ParserRule: prs,
		Operator:   op,
		Value:      val.(ast.Node),
	}
	s.stack.Push(n)
}

func (s *SFeelListener) ExitInterval(ctx *gen.IntervalContext) {
	prs := ast.Rule{
		Type:    gen.SFeelParserRULE_interval,
		Literal: ctx.GetText(),
	}
	endRule := s.stack.Pop().(ast.Rule)
	endVal := s.stack.Pop().(ast.Node)
	startVal := s.stack.Pop().(ast.Node)
	startRule := s.stack.Pop().(ast.Rule)

	inVal := ast.Interval{
		ParserRule:        prs,
		StartIntervalRule: startRule,
		EndIntervalRule:   endRule,
		StartValue:        startVal,
		EndValue:          endVal,
	}

	s.stack.Push(inVal)

}

func (s *SFeelListener) ExitOpen_interval_start(ctx *gen.Open_interval_startContext) {
	r := ast.Rule{Type: gen.SFeelParserRULE_open_interval_start, Literal: ctx.GetText()}
	s.stack.Push(r)
}

func (s *SFeelListener) ExitClosed_interval_start(ctx *gen.Closed_interval_startContext) {
	r := ast.Rule{Type: gen.SFeelParserRULE_closed_interval_start, Literal: ctx.GetText()}
	s.stack.Push(r)
}

func (s *SFeelListener) ExitOpen_interval_end(ctx *gen.Open_interval_endContext) {
	r := ast.Rule{Type: gen.SFeelParserRULE_open_interval_end, Literal: ctx.GetText()}
	s.stack.Push(r)
}

func (s *SFeelListener) ExitClosed_interval_end(ctx *gen.Closed_interval_endContext) {
	r := ast.Rule{Type: gen.SFeelParserRULE_closed_interval_end, Literal: ctx.GetText()}
	s.stack.Push(r)
}

// Datatypes and Primitives
func (s *SFeelListener) ExitNumeric_literal(ctx *gen.Numeric_literalContext) {
	sign := ctx.GetStart().GetTokenType()

	if sign == gen.SFeelParserSUB {
		nType := s.stack.Peek()
		switch nType.(type) {
		case ast.Integer:
			n := s.stack.Pop()
			negInt := n.(ast.Integer)
			negInt.SignRule = ast.Rule{Type: gen.SFeelParserRULE_numeric_literal, Literal: ctx.GetStart().GetText()}
			s.stack.Push(negInt)
		case ast.Float:
			n := s.stack.Pop()
			negReal := n.(ast.Float)
			negReal.SignRule = ast.Rule{Type: gen.SFeelParserRULE_numeric_literal, Literal: ctx.GetStart().GetText()}
			s.stack.Push(negReal)
		default:
			s.Errors = append(s.Errors, errors.NewError("unknown numeric data type: %s ", reflect.TypeOf(nType)))
		}
	}
}

func (s *SFeelListener) ExitQualified_name(ctx *gen.Qualified_nameContext) {
	rule := ast.Rule{Type: gen.SFeelParserRULE_qualified_name, Literal: ctx.GetText()}
	var names []string
	for _, val := range ctx.AllName() {
		names = append(names, val.GetText())
	}

	q := ast.QualifiedName{ParserRule: rule, Value: names}
	s.stack.Push(q)
}

func (s *SFeelListener) ExitInteger_literal(ctx *gen.Integer_literalContext) {
	rule := ast.Rule{Type: gen.SFeelParserRULE_integer_literal, Literal: ctx.GetText()}
	val, err := strconv.ParseInt(rule.Literal, 10, 64)

	if err != nil {
		s.Errors = append(s.Errors, err)
	} else {
		n := ast.Integer{ParserRule: rule, SignRule: ast.Rule{Type: -1}, Value: val}
		s.stack.Push(n)
	}
}

func (s *SFeelListener) ExitReal_literal(ctx *gen.Real_literalContext) {
	rule := ast.Rule{Type: gen.SFeelParserRULE_real_literal, Literal: ctx.GetText()}
	val, err := strconv.ParseFloat(rule.Literal, 64)

	if err != nil {
		s.Errors = append(s.Errors, err)
	} else {
		n := ast.Float{ParserRule: rule, SignRule: ast.Rule{Type: -1}, Value: val}
		s.stack.Push(n)
	}
}

func (s *SFeelListener) ExitString_literal(ctx *gen.String_literalContext) {
	rule := ast.Rule{Type: gen.SFeelParserRULE_string_literal, Literal: ctx.GetText()}
	val := rule.Literal
	n := ast.String{ParserRule: rule, Value: val}
	s.stack.Push(n)
}

func (s *SFeelListener) ExitBoolean_literal(ctx *gen.Boolean_literalContext) {
	rule := ast.Rule{Type: gen.SFeelParserRULE_boolean_literal, Literal: ctx.GetText()}
	val, err := strconv.ParseBool(rule.Literal)

	if err != nil {
		s.Errors = append(s.Errors, err)
	} else {
		n := ast.Boolean{ParserRule: rule, Value: val}
		s.stack.Push(n)
	}
}

func (s *SFeelListener) ExitDate_time_literal(ctx *gen.Date_time_literalContext) {
	//ToDo Date-Time Datatype
	panic("implement me")
}
