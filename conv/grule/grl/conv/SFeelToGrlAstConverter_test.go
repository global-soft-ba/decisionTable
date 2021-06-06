package conv

import (
	"decisionTable/data"
	"decisionTable/lang/sfeel"
	"testing"
)

func TestSFeelToGrlConverter_Convert(t *testing.T) {

	type args struct {
		field      data.FieldInterface
		sfeelEntry data.EntryInterface
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple interval convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Typ:  data.Integer,
				},
				sfeel.CreateInputEntry("[1..6]"),
			},
			want: "((X.Y :6: 1) :0: (X.Y :4: 6))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateSFeelToGrlAstConverter()
			got, _ := c.ConvertToGrlAst(tt.args.field, tt.args.sfeelEntry)
			if got.String() != tt.want {
				t.Errorf("ConvertToGrlAst() = %v, want %v", got, tt.want)
			}
		})
	}
}
