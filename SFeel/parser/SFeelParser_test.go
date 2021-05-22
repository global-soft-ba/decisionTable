package parser

import (
	antlr2 "decisionTable/SFeel/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"testing"
)

func TestASTRepresentation(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "interval",
			args: args{"[1..5]"},
			want: "[1..5]",
		},
		{
			name: "interval real",
			args: args{"[1.1..55.1]"},
			want: "[1.1E+00..5.51E+01]",
		},
		{
			name: "equal unary integer",
			args: args{"1"},
			want: "=1",
		},
		{
			name: "equal unary string",
			args: args{`"1"`},
			want: `="1"`,
		},
		{
			name: "equal unary boolean",
			args: args{`true`},
			want: `=true`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := CreateSFeelParser(tt.args.expr)
			tree := prs.Parser().Start()

			if len(prs.Errors()) == 0 {
				// Finally parse the expression
				conv := antlr2.CreateSFeelListener()
				antlr.ParseTreeWalkerDefault.Walk(&conv, tree)
				if got := conv.GetAST().String(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CreateSFeelParser() = %v, want %v", got, tt.want)
				}
			} else {
				t.Errorf("CreateSFeelParser(): Syntax error in expression = %v", prs.Errors())
			}
		})
	}
}
