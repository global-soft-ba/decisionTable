package conv

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/lang/sfeel"
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
					Type: data.Integer,
				},
				sfeel.CreateInputEntry("[1..6]"), // ((X.Y >= 1) && (X.Y <= 6))
			},
			want: "((X.Y :7: 1) :0: (X.Y :5: 6))",
		},
		{
			name: "simple unary test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateInputEntry("8"),
			},
			want: "(X.Y :2: 8)",
		},
		{
			name: "simple expression simple value test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry("8"),
			},
			want: "X.Y -1 8",
		},
		{
			name: "simple expression parentheses test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry("(8)"),
			},
			want: "X.Y -1 (8)",
		},
		{
			name: "simple expression arithmetic negation test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry(`-"asd"`),
			},
			want: `X.Y -1 -"asd"`,
		},
		{
			name: "simple expression arithmetic negation of integer test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry(`-1`),
			},
			want: `X.Y -1 -1`,
		},
		{
			name: "simple expression arithmetic negation of negative integer test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry(`--1`),
			},
			want: `X.Y -1 --1`,
		},
		{
			name: "simple expression arithmetic power operation test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry(`1**2`),
			},
			want: `X.Y -1 1**2`,
		},
		{
			name: "simple expression arithmetic multiplication test convert",
			args: args{
				data.TestField{
					Name: "X",
					Key:  "Y",
					Type: data.Integer,
				},
				sfeel.CreateOutputEntry(`2+2`),
			},
			want: `X.Y -1 282`,
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
