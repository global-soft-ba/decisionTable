package visitors

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/convert/grule/termconverter/sfeel/mapper"
	"decisionTable/data"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestNumberVisitor_FloatInputRules(t *testing.T) {
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
		{"number comparison LESS",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "<1.2",
				},
				mapping,
			},
			"credit.score < 1.2"},
		{"number comparison LESSEQ",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "<=1.1",
				},
				mapping,
			},
			"credit.score <= 1.1"},
		{"number comparison GR",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: ">1.2",
				},
				mapping,
			},
			"credit.score > 1.2"},
		{"number comparison GREQ",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: ">=1",
				},
				mapping,
			},
			"credit.score >= 1"},
		{"empty number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "-",
				},
				mapping,
			},
			""},
		{"equal number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "1.1",
				},
				mapping,
			},
			"credit.score == 1.1"},
		{"range IN OUT number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "[1.1..90]",
				},
				mapping,
			},
			"((credit.score >= 1.1) && (credit.score <= 90))"},
		{"range IN IN number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "[1.1..90.1[",
				},
				mapping,
			},
			"((credit.score >= 1.1) && (credit.score < 90.1))"},
		{"range OUT IN number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "]1.1..90[",
				},
				mapping,
			},
			"((credit.score > 1.1) && (credit.score < 90))"},
		{"range OUT IN number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Integer,
					Expression: "]1..90]",
				},
				mapping,
			},
			"((credit.score > 1) && (credit.score <= 90))"},
		{"disjunctions number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Integer,
					Expression: "<1,>=2,[3..8],4,7",
				},
				mapping,
			},
			"((credit.score < 1) || (credit.score >= 2) || (((credit.score >= 3) && (credit.score <= 8))) || (credit.score == 4) || (credit.score == 7))",
		},
		{"Negation number input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Integer,
					Expression: "not(1.1)",
				},
				mapping,
			},
			"!(credit.score == 1.1)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidNumberInput()
			vis := CreateNumberVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerInputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberVisitor_NumberOutputRules(t *testing.T) {
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
		{"float assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "1",
				},
				mapping,
			},
			"credit.score = 1"},
		{"float empty assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Float,
					Expression: "-",
				},
				mapping,
			},
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidNumberOutput()
			vis := CreateNumberVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in NumberOutputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}
