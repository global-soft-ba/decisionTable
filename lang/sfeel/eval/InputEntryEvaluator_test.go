package eval

import (
	ast "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	"testing"
)

func TestInputEntryEvaluator_Eval(t *testing.T) {
	type args struct {
		node ast.Node
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "incorrect datatype evaluation",
			args:    args{ast.Float{}},
			want:    false,
			wantErr: true,
		},
		{
			name:    "correct unary evaluation",
			args:    args{ast.UnaryTests{}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "correct empty unary evaluation",
			args:    args{ast.EmptyStatement{}},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := CreateInputEntryEvaluator()
			got, err := e.Eval(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eval() got = %v, want %v with error %v", got, tt.want, err)
			}
		})
	}
}
