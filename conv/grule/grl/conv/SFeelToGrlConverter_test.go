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
			want: "((X >= 1) && (X <= 6))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateSFeelToGrlConverter()
			got, _ := c.Convert(tt.args.fieldName, tt.args.sfeelEntry)
			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
