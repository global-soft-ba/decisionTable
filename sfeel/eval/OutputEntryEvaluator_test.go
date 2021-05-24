package eval

import (
	"decisionTable/sfeel/ast"
	"testing"
)

func TestOutputEntryEvaluator_Eval(t *testing.T) {
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
			name:    "incorrect unary evaluation",
			args:    args{ast.UnaryTests{}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := CreateOutputEntryEvaluator()
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
