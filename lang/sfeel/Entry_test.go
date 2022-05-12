package sfeel

import (
	"github.com/global-soft-ba/decisionTable/ast"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/field"
	sfeel "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	"reflect"
	"testing"
)

func TestEvalInputEntry(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Correct input expression",
			args:    args{exp: "<1,<2"},
			want:    "<1,<2",
			wantErr: false,
		},
		{
			name:    "Incorrect input expression",
			args:    args{exp: "1+1"},
			want:    "1+1",
			wantErr: true,
		},
		{
			name:    "Incorrect input expression",
			args:    args{exp: "1<<<<+1"},
			want:    "1<<<<+1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateInputEntry(tt.args.exp)
			_, err := got.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("CreateInputEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOutputEntry(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "incorrect output expression",
			args:    args{exp: "<1,<2"},
			want:    "<1,<2",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateOutputEntry(tt.args.exp)
			_, err := got.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("CreateOutputEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpression_ValidateDataType(t *testing.T) {
	type fields struct {
		ast        sfeel.Node
		expression string
	}
	type args struct {
		varType dataType.DataType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "correct validation integer",
			fields:  fields{ast: sfeel.Integer{}},
			args:    args{dataType.Integer},
			want:    true,
			wantErr: false,
		},
		{
			name:    "correct validation string",
			fields:  fields{ast: sfeel.String{}},
			args:    args{dataType.String},
			want:    true,
			wantErr: false,
		},
		{
			name:    "incorrect validation string",
			fields:  fields{ast: sfeel.String{}},
			args:    args{dataType.Float},
			want:    false,
			wantErr: true,
		},
		{
			name: "correct validation interval",
			fields: fields{
				ast: sfeel.Interval{StartValue: sfeel.Integer{}, EndValue: sfeel.Float{}},
			},
			args:    args{dataType.Float},
			want:    true,
			wantErr: false,
		},
		{
			name: "incorrect validation interval",
			fields: fields{
				ast: sfeel.Interval{StartValue: sfeel.Integer{}, EndValue: sfeel.Float{}},
			},
			args:    args{dataType.Integer},
			want:    false,
			wantErr: true,
		},
		{
			name: "correct validation unary test",
			fields: fields{
				ast: sfeel.UnaryTest{Value: sfeel.Integer{}},
			},
			args:    args{dataType.Integer},
			want:    true,
			wantErr: false,
		},
		{
			name: "incorrect validation unary test",
			fields: fields{
				ast: sfeel.UnaryTest{Value: sfeel.Integer{}},
			},
			args:    args{dataType.Boolean},
			want:    false,
			wantErr: true,
		},
		{
			name: "correct validation unary tests",
			fields: fields{
				ast: sfeel.UnaryTests{
					UnaryTests: []ast.Node{
						sfeel.UnaryTest{Value: sfeel.Boolean{}},
						sfeel.UnaryTest{Value: sfeel.Boolean{}},
						sfeel.UnaryTest{Value: sfeel.Boolean{}},
					},
				}},
			args:    args{dataType.Boolean},
			want:    true,
			wantErr: false,
		},
		{
			name: "empty unary validation",
			fields: fields{
				ast: sfeel.EmptyStatement{}},
			args:    args{dataType.Boolean},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entry{
				ast:        tt.fields.ast,
				expression: tt.fields.expression,
			}
			got, err := e.ValidateDataTypeOfExpression(tt.args.varType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDataTypeOfExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateDataTypeOfExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpression_ValidateFieldReferences(t *testing.T) {
	type fields struct {
		ast        sfeel.Node
		expression string
	}
	type args struct {
		fields []field.Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []field.Field
		wantErr bool
	}{
		{
			name: "simple test",
			fields: fields{
				ast: sfeel.QualifiedName{Value: []string{"b", "c"}},
			},
			args: args{fields: []field.Field{
				{Name: "a.b", Type: dataType.Integer},
				{Name: "b.c", Type: dataType.Integer},
				{Name: "a.c", Type: dataType.Integer},
			}},
			want:    []field.Field{{Name: "b.c", Type: dataType.Integer}},
			wantErr: false,
		},
		{
			name: "nested 1 level test",
			fields: fields{
				ast: sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"a", "b"}}},
			},
			args: args{fields: []field.Field{
				{Name: "a.b", Type: dataType.Integer},
				{Name: "b.c", Type: dataType.Integer},
				{Name: "a.c", Type: dataType.Integer},
			}},
			want:    []field.Field{{Name: "a.b", Type: dataType.Integer}},
			wantErr: false,
		},
		{
			name: "nested 2 level test",
			fields: fields{
				ast: sfeel.UnaryTests{
					UnaryTests: []ast.Node{
						sfeel.UnaryTest{Value: sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"b", "c"}}}},
						sfeel.UnaryTest{Value: sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"a", "c"}}}},
						sfeel.UnaryTest{Value: sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"a", "b"}}}},
						sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"a", "b"}}},
						sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"a", "c"}}},
					},
				},
			},
			args: args{fields: []field.Field{
				{Name: "a.b", Type: dataType.Integer},
				{Name: "b.c", Type: dataType.Integer},
				{Name: "a.c", Type: dataType.Integer},
			}},
			want: []field.Field{
				{Name: "b.c", Type: dataType.Integer},
				{Name: "a.c", Type: dataType.Integer},
				{Name: "a.b", Type: dataType.Integer},
				{Name: "a.b", Type: dataType.Integer},
				{Name: "a.c", Type: dataType.Integer},
			},
			wantErr: false,
		},
		{
			name: "invalid qualified name",
			fields: fields{
				ast: sfeel.UnaryTest{Value: sfeel.QualifiedName{Value: []string{"x", "x"}}},
			},
			args: args{fields: []field.Field{
				{Name: "a.b", Type: dataType.Integer},
				{Name: "b.c", Type: dataType.Integer},
				{Name: "a.c", Type: dataType.Integer},
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entry{
				ast:        tt.fields.ast,
				expression: tt.fields.expression,
			}
			got, err := e.ValidateExistenceOfFieldReferencesInExpression(tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExistenceOfFieldReferencesInExpression() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateExistenceOfFieldReferencesInExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}
