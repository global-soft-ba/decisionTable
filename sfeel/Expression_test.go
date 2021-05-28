package sfeel

import (
	"decisionTable/model"
	"decisionTable/sfeel/ast"
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
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateInputExpression(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("CreateInputExpression() got = %v, want %v", got, tt.want)
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
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateOutputExpression(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("CreateOutputExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpression_ValidateDataType(t *testing.T) {
	type fields struct {
		ast        ast.Node
		expression string
	}
	type args struct {
		varType model.DataTyp
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
			fields:  fields{ast: ast.Integer{}},
			args:    args{model.Integer},
			want:    true,
			wantErr: false,
		},
		{
			name:    "correct validation string",
			fields:  fields{ast: ast.String{}},
			args:    args{model.String},
			want:    true,
			wantErr: false,
		},
		{
			name:    "incorrect validation string",
			fields:  fields{ast: ast.String{}},
			args:    args{model.Float},
			want:    false,
			wantErr: true,
		},
		{
			name: "correct validation interval",
			fields: fields{
				ast: ast.Interval{StartValue: ast.Integer{}, EndValue: ast.Float{}},
			},
			args:    args{model.Float},
			want:    true,
			wantErr: false,
		},
		{
			name: "incorrect validation interval",
			fields: fields{
				ast: ast.Interval{StartValue: ast.Integer{}, EndValue: ast.Float{}},
			},
			args:    args{model.Integer},
			want:    false,
			wantErr: true,
		},
		{
			name: "correct validation unary test",
			fields: fields{
				ast: ast.UnaryTest{Value: ast.Integer{}},
			},
			args:    args{model.Integer},
			want:    true,
			wantErr: false,
		},
		{
			name: "incorrect validation unary test",
			fields: fields{
				ast: ast.UnaryTest{Value: ast.Integer{}},
			},
			args:    args{model.Boolean},
			want:    false,
			wantErr: true,
		},
		{
			name: "correct validation unary tests",
			fields: fields{
				ast: ast.UnaryTests{
					UnaryTests: []ast.Node{
						ast.UnaryTest{Value: ast.Boolean{}},
						ast.UnaryTest{Value: ast.Boolean{}},
						ast.UnaryTest{Value: ast.Boolean{}},
					},
				}},
			args:    args{model.Boolean},
			want:    true,
			wantErr: false,
		},
		{
			name: "empty unary validation",
			fields: fields{
				ast: ast.EmptyStatement{}},
			args:    args{model.Boolean},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Expression{
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
		ast        ast.Node
		expression string
	}
	type args struct {
		fields []model.Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Field
		wantErr bool
	}{
		{
			name: "simple test",
			fields: fields{
				ast: ast.QualifiedName{Value: []string{"b", "c"}},
			},
			args: args{fields: []model.Field{
				model.Field{Name: "a", Key: "b", Typ: model.Integer},
				model.Field{Name: "b", Key: "c", Typ: model.Integer},
				model.Field{Name: "a", Key: "c", Typ: model.Integer},
			}},
			want:    []model.Field{model.Field{Name: "b", Key: "c", Typ: model.Integer}},
			wantErr: false,
		},
		{
			name: "nested 1 level test",
			fields: fields{
				ast: ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"a", "b"}}},
			},
			args: args{fields: []model.Field{
				model.Field{Name: "a", Key: "b", Typ: model.Integer},
				model.Field{Name: "b", Key: "c", Typ: model.Integer},
				model.Field{Name: "a", Key: "c", Typ: model.Integer},
			}},
			want:    []model.Field{model.Field{Name: "a", Key: "b", Typ: model.Integer}},
			wantErr: false,
		},
		{
			name: "nested 2 level test",
			fields: fields{
				ast: ast.UnaryTests{
					UnaryTests: []ast.Node{
						ast.UnaryTest{Value: ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"b", "c"}}}},
						ast.UnaryTest{Value: ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"a", "c"}}}},
						ast.UnaryTest{Value: ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"a", "b"}}}},
						ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"a", "b"}}},
						ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"a", "c"}}},
					},
				},
			},
			args: args{fields: []model.Field{
				model.Field{Name: "a", Key: "b", Typ: model.Integer},
				model.Field{Name: "b", Key: "c", Typ: model.Integer},
				model.Field{Name: "a", Key: "c", Typ: model.Integer},
			}},
			want: []model.Field{
				model.Field{Name: "b", Key: "c", Typ: model.Integer},
				model.Field{Name: "a", Key: "c", Typ: model.Integer},
				model.Field{Name: "a", Key: "b", Typ: model.Integer},
				model.Field{Name: "a", Key: "b", Typ: model.Integer},
				model.Field{Name: "a", Key: "c", Typ: model.Integer},
			},
			wantErr: false,
		},
		{
			name: "invalid qualified name",
			fields: fields{
				ast: ast.UnaryTest{Value: ast.QualifiedName{Value: []string{"x", "x"}}},
			},
			args: args{fields: []model.Field{
				model.Field{Name: "a", Key: "b", Typ: model.Integer},
				model.Field{Name: "b", Key: "c", Typ: model.Integer},
				model.Field{Name: "a", Key: "c", Typ: model.Integer},
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Expression{
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
