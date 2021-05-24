package eval

import (
	"decisionTable/SFeel/ast"
	"testing"
)

func TestEvaluator_Eval(t *testing.T) {
	type args struct {
		node ast.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "incorrect datatype evaluation",
			args: args{ast.Float{}},
			want: false,
		},
		{
			name: "correct unary evaluation",
			args: args{ast.UnaryTests{}},
			want: true,
		},
		{
			name: "correct empty unary evaluation",
			args: args{ast.EmptyUnaryTest{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := CreateEvaluator()
			got, err := e.Eval(tt.args.node)
			if got != tt.want {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
				t.Errorf("Error %v", err)
			}
		})
	}
}
