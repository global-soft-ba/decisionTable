package new_eval_omg_sfeel

import (
	"decisionTable/SFeel/ast"
	"decisionTable/model"
	"testing"
)

func TestEvaluator_EvalDataType(t *testing.T) {
	type fields struct {
		field model.Field
	}
	type args struct {
		data ast.Node
	}
	fieldModel := []model.Field{{Name: "Name", Key: "Label", Typ: model.Integer}, {Name: "Name2", Key: "Label2", Typ: model.String}}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "correct datatype matching",
			fields:  fields{field: model.Field{Name: "Name", Key: "Label", Typ: model.Integer}},
			args:    args{data: ast.Integer{}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "incorrect datatype matching",
			fields:  fields{field: model.Field{Name: "Name", Key: "Label", Typ: model.Integer}},
			args:    args{data: ast.String{}},
			want:    false,
			wantErr: true,
		},
		{
			name:    "correct qualified name matching",
			fields:  fields{field: model.Field{Name: "Name", Key: "Label", Typ: model.Integer}},
			args:    args{data: ast.QualifiedName{Value: []string{"Name", "Label"}}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "not existent qualified name",
			fields:  fields{field: model.Field{Name: "Name", Key: "Label", Typ: model.Integer}},
			args:    args{data: ast.QualifiedName{Value: []string{"X", "Label"}}},
			want:    false,
			wantErr: true,
		},
		{
			name:    "not correct qualified name typ matching",
			fields:  fields{field: model.Field{Name: "Name", Key: "Label", Typ: model.Integer}},
			args:    args{data: ast.QualifiedName{Value: []string{"Name2", "Label2"}}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := FieldsEvaluator{
				field:  tt.fields.field,
				fields: fieldModel,
			}
			got, err := e.evalDataType(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalDataType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EvalDataType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
