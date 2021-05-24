package antlr

import (
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
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
			ast, err := CreateParser(tt.args.expr).Parse()
			if err != nil {
				t.Errorf("Parsing Error = %v", err)
			}
			if got := ast.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
