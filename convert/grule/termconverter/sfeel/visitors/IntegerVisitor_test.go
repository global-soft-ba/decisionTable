package visitors

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/convert/grule/termconverter/sfeel/mapper"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestIntegerVisitor_IntegerInputRules(t *testing.T) {
	type fields struct {
		term grlmodel.Term
		maps mapper.TermMapper
	}

	mapping := mapper.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"integer comparison LESS",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "<1",
				},
				mapping,
			},
			"credit.score < 1"},
		{"integer comparison LESSEQ",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "<=1",
				},
				mapping,
			},
			"credit.score <= 1"},
		{"integer comparison GR",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: ">1",
				},
				mapping,
			},
			"credit.score > 1"},
		{"integer comparison GREQ",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: ">=1",
				},
				mapping,
			},
			"credit.score >= 1"},
		{"empty integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "-",
				},
				mapping,
			},
			""},
		{"equal integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "1",
				},
				mapping,
			},
			"credit.score == 1"},
		{"range IN OUT integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "[1..90]",
				},
				mapping,
			},
			"((credit.score >= 1) && (credit.score <= 90))"},
		{"range IN IN integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "[1..90[",
				},
				mapping,
			},
			"((credit.score >= 1) && (credit.score < 90))"},
		{"range OUT IN integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "]1..90[",
				},
				mapping,
			},
			"((credit.score > 1) && (credit.score < 90))"},
		{"range OUT IN integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "]1..90]",
				},
				mapping,
			},
			"((credit.score > 1) && (credit.score <= 90))"},
		{"disjunctions integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "<1,>=2,[3..8],4,7",
				},
				mapping,
			},
			"((credit.score < 1) || (credit.score >= 2) || (((credit.score >= 3) && (credit.score <= 8))) || (credit.score == 4) || (credit.score == 7))",
		},
		{"Negation integer input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "not(1)",
				},
				mapping,
			},
			"!(credit.score == 1)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidIntegerInput()
			vis := CreateIntegerVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerInputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegerVisitor_IntegerOutputRules(t *testing.T) {
	type fields struct {
		term grlmodel.Term
		maps mapper.TermMapper
	}

	mapping := mapper.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"integer assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "1",
				},
				mapping,
			},
			"credit.score = 1"},
		{"integer empty assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "-",
				},
				mapping,
			},
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidIntegerOutput()
			vis := CreateIntegerVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerOutputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}
