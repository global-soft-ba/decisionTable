package conv

import (
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/field"
	"testing"
)

func TestSFeelToGrlConverter_Convert(t *testing.T) {

	type args struct {
		field     field.Field
		entryType entryType.EntryType
		entry     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple interval convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Input,
				"[1..6]", // ((X.Y >= 1) && (X.Y <= 6))
			},
			want: "((X.Y :7: 1) :0: (X.Y :5: 6))",
		},
		{
			name: "simple unary test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Input,
				"8",
			},
			want: "(X.Y :2: 8)",
		},
		{
			name: "simple expression simple value test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				"8",
			},
			want: "X.Y -1 8",
		},
		{
			name: "simple expression parentheses test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				"(8)",
			},
			want: "X.Y -1 (8)",
		},
		{
			name: "simple expression arithmetic negation test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				`-"asd"`,
			},
			want: `X.Y -1 -"asd"`,
		},
		{
			name: "simple expression arithmetic negation of integer test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				`-1`,
			},
			want: `X.Y -1 -1`,
		},
		{
			name: "simple expression arithmetic negation of negative integer test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				`--1`,
			},
			want: `X.Y -1 --1`,
		},
		{
			name: "simple expression arithmetic power operation test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				`1**2`,
			},
			want: `X.Y -1 1**2`,
		},
		{
			name: "simple expression arithmetic multiplication test convert",
			args: args{
				field.Field{
					Name: "X.Y",
					Type: dataType.Integer,
				},
				entryType.Output,
				`2+2`,
			},
			want: `X.Y -1 282`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateSFeelToGrlAstConverter()
			got, _ := c.ConvertToGrlAst(tt.args.field, tt.args.entryType, tt.args.entry)
			if got.String() != tt.want {
				t.Errorf("ConvertToGrlAst() = %v, want %v", got, tt.want)
			}
		})
	}
}
