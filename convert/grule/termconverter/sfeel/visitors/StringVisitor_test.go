package visitors

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/convert/grule/termconverter/sfeel/mapper"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestStringVisitor_StringInputRules(t *testing.T) {
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
		{"equal string input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.String,
					Expression: `"1"`,
				},
				mapping,
			},
			`credit.score == "1"`},

		{"disjunctions string input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: `"1","2","ABC"`,
				},
				mapping,
			},
			`((credit.score == "1") || (credit.score == "2") || (credit.score == "ABC"))`,
		},
		{"Negation string input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: `not("1")`,
				},
				mapping,
			},
			`!(credit.score == "1")`},
		{"Empty string input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Integer,
					Expression: "-",
				},
				mapping,
			},
			``},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidStringInput()
			vis := CreateStringVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerInputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringVisitor_StringOutputRules(t *testing.T) {
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
		{"string assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.String,
					Expression: `"1"`,
				},
				mapping,
			},
			`credit.score = "1"`},
		{"string empty assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.String,
					Expression: "-",
				},
				mapping,
			},
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidStringOutput()
			vis := CreateStringVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in StringOutputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}
