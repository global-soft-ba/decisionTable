package visitors

import (
	"decisionTable/converter/grule/expressionlanguages/sfeel/mapper"
	"decisionTable/converter/grule/grlmodel"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestBoolVisitor_BoolInputRules(t *testing.T) {
	type fields struct {
		term grlmodel.Term
		maps mapper.Mapper
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
					Identifier: "credit",
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
			tree := prs.Parse().ValidBoolInput()
			vis := CreateBoolVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerInputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}
