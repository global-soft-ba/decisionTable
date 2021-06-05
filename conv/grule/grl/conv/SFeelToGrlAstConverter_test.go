package conv

import (
	"decisionTable/data"
	"decisionTable/lang/sfeel"
	"testing"
)

func TestSFeelToGrlConverter_Convert(t *testing.T) {

	type args struct {
		fieldName  string
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
				"X",
				sfeel.CreateInputEntry("[1..6]"),
			},
			want: "X opId:6 1 opId:0 X opId:4 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateSFeelToGrlAstConverter()
			got, _ := c.ConvertToGrlAst(tt.args.fieldName, tt.args.sfeelEntry)
			if got.String() != tt.want {
				t.Errorf("ConvertToGrlAst() = %v, want %v", got, tt.want)
			}
		})
	}
}
