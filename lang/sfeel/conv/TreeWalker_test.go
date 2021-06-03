package conv

import (
	"decisionTable/lang/sfeel/ast"
	"fmt"
	"testing"
)

func TestWalker_Walk(t *testing.T) {
	type fields struct {
		listener SFeelBaseListenerInterface
	}
	type args struct {
		node ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple integer test",
			fields: fields{&SFeelBaseListener{}},
			args:   args{ast.Integer{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			walker := TreeWalker{
				listener: tt.fields.listener,
			}
			walker.Walk(tt.args.node)
			fmt.Println(walker)
		})
	}
}
