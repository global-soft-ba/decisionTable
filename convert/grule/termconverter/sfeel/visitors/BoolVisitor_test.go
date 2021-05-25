package visitors

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/convert/grule/termconverter/sfeel/mapper"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestBoolVisitor_BoolInputRules(t *testing.T) {
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
		{"equal bool input",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Boolean,
					Expression: "false",
				},
				mapping,
			},
			`credit.score == false`},
		{"Empty bool input",
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
		term grlmodel.Term
		maps mapper.TermMapper
	}

	mapping := mapper.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"boolean assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Boolean,
					Expression: "true",
				},
				mapping,
			},
			"credit.score = true"},
		{"boolean empty assignment",
			fields{
				grlmodel.Term{
					Name:       "score",
					Key:        "credit",
					Typ:        model.Boolean,
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
