package antlr

import (
	gen "decisionTable/SFeel/gen"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"testing"
)

func TestASTRepresentation_InputEntries(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "interval integer",
			args: args{"[1..5]"},
			want: "[1..5]",
		},
		{
			name: "interval float",
			args: args{"[1.1..55.1]"},
			want: "[1.1E+00..5.51E+01]",
		},
		{
			name: "equal unary integer",
			args: args{"1"},
			want: "1",
		},
		{
			name: "equal unary string",
			args: args{`"1"`},
			want: `"1"`,
		},
		{
			name: "equal unary boolean",
			args: args{`true`},
			want: `true`,
		},
		{
			name: "less unary statements",
			args: args{`<1,<3`},
			want: `<1,<3`,
		},
		{
			name: "not unary statements",
			args: args{`not(<1,<3)`},
			want: `not(<1,<3)`,
		},
		{
			name: "empty unary statement",
			args: args{`-`},
			want: "",
		},
		{
			name: "unary statement with qualified_nam",
			args: args{`<a`},
			want: "<a",
		},
		{
			name: "unary statement with multiple qualified_names",
			args: args{`<a.a1.a2`},
			want: "<a.a1.a2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := CreateSFeelParser(tt.args.expr)
			tree := prs.Parser().SyntaxCheck()
			//Up-Front Syntax Check
			base := gen.BaseSFeelListener{}
			antlr.ParseTreeWalkerDefault.Walk(&base, tree)

			if len(prs.Errors()) == 0 {
				// Finally parse the expression
				conv := CreateSFeelListener()
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
