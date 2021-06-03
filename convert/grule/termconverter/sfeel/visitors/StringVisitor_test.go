package visitors

import (
	grlmodel2 "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/grl/symbols"
	"decisionTable/data"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestStringVisitor_StringInputRules(t *testing.T) {
	type fields struct {
		term grlmodel2.Term
		maps symbols.TermMapper
	}

	mapping := symbols.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"equal string input",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.String,
					Expression: `"1"`,
				},
				mapping,
			},
			`credit.score == "1"`},

		{"disjunctions string input",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Integer,
					Expression: `"1","2","ABC"`,
				},
				mapping,
			},
			`((credit.score == "1") || (credit.score == "2") || (credit.score == "ABC"))`,
		},
		{"Negation string input",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Integer,
					Expression: `not("1")`,
				},
				mapping,
			},
			`!(credit.score == "1")`},
		{"Empty string input",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Integer,
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
		term grlmodel2.Term
		maps symbols.TermMapper
	}

	mapping := symbols.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"string assignment",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.String,
					Expression: `"1"`,
				},
				mapping,
			},
			`credit.score = "1"`},
		{"string empty assignment",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.String,
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
