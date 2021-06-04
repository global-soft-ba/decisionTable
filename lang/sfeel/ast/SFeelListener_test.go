package ast

import (
	"decisionTable/ast"
	"testing"
)

func TestSFeelListener_Walk(t *testing.T) {
	type fields struct {
		listener SFeelListenerInterface
	}
	type args struct {
		node Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple integer test",
			fields: fields{&SFeelListener{}},
			args:   args{Integer{}},
		},
		{
			name:   "simple unary tests",
			fields: fields{&SFeelListener{}},
			args:   args{UnaryTests{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			walker := ast.CreateTreeWalker(&BaseListener{tt.fields.listener})
			walker.Walk(tt.args.node)

		})
	}
}
