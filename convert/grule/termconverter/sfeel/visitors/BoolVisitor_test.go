package visitors

import (
	grlmodel2 "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/grl/symbols"
	"decisionTable/data"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestBoolVisitor_BoolInputRules(t *testing.T) {
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
		{"equal bool input",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Boolean,
					Expression: "false",
				},
				mapping,
			},
			`credit.score == false`},
		{"Empty bool input",
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
			tree := prs.Parse().ValidBoolInput()
			vis := CreateBoolVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerInputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolVisitor_BoolOutputRules(t *testing.T) {
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
		{"boolean assignment",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Boolean,
					Expression: "true",
				},
				mapping,
			},
			"credit.score = true"},
		{"boolean empty assignment",
			fields{
				grlmodel2.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        data.Boolean,
					Expression: "-",
				},
				mapping,
			},
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidBoolOutput()
			vis := CreateBoolVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in StringOutputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}
