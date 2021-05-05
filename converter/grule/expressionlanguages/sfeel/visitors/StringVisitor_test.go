package visitors

import (
	"decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestStringVisitor_StringInputRules(t *testing.T) {
	type fields struct {
		term grlmodel.Term
		maps mapper.Mapper
	}

	mapping := mapper.GrlMapping

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"equal string input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Identifier: "credit",
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
					Identifier: "credit",
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
					Identifier: "credit",
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
					Identifier: "credit",
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
