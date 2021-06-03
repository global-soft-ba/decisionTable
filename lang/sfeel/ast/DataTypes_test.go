package ast

import (
	"reflect"
	"testing"
)

func TestCheckDataTypePrecedence(t *testing.T) {
	type args struct {
		types []Node
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			name: "precedence single value",
			args: args{[]Node{Integer{}, Integer{}, Integer{}, Integer{}, Integer{}, Integer{}}},
			want: reflect.TypeOf(Integer{}),
		},
		{
			name: "precedence value float",
			args: args{[]Node{Integer{}, Integer{}, Float{}, Integer{}, Integer{}, Integer{}}},
			want: reflect.TypeOf(Float{}),
		},
		{
			name: "empty precedence",
			args: args{[]Node{Float{}, Boolean{}, Integer{}, Float{}, Integer{}, Integer{}, Integer{}}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDataTypePrecedences(tt.args.types...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkDataTypePrecedence() = %v, want %v", got, tt.want)
			}
		})
	}
}
