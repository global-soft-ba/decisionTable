package generate

import (
	"github.com/global-soft-ba/decisionTable/ast"
	grl "github.com/global-soft-ba/decisionTable/conv/grule/grl/ast"
	"testing"
)

func TestGrlGeneratorListener_GetCode(t *testing.T) {

	type fields struct {
		node ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple comparison test",
			fields: fields{
				node: grl.ComparisonOperations{
					Left:     grl.Integer{Val: 1},
					Operator: grl.ComparisonOperatorGREATEREQ,
					Right:    grl.Integer{Val: 2},
				},
			},
			want: "(1 >= 2)",
		},
		{
			name: "wrong operator ID test",
			fields: fields{
				node: grl.ComparisonOperations{
					Left:     grl.Integer{Val: 1},
					Operator: -1,
					Right:    grl.Integer{Val: 2},
				},
			},
			want: "(1 2)",
		},
		{
			name: "nested addition test",
			fields: fields{
				node: grl.ComparisonOperations{
					Left:     grl.Integer{Val: 1},
					Operator: grl.MathADD,
					Right: grl.ComparisonOperations{
						Left:     grl.Integer{Val: 1},
						Operator: grl.MathADD,
						Right:    grl.Integer{Val: 3},
					},
				},
			},
			want: "(1 + (1 + 3))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen, _ := CreateGrlGeneratorListener()
			walker := grl.CreateGRLTreeWalker(&gen)
			walker.Walk(tt.fields.node)
			if got := gen.GetCode(); got != tt.want {
				t.Errorf("GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
