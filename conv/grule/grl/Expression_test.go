package grl

import (
	"github.com/global-soft-ba/decisionTable/conv/grule/data"
	dTable "github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/lang/sfeel"
	"reflect"
	"testing"
)

func TestCreateExpression(t *testing.T) {
	type args struct {
		field field.Field
		entry dTable.EntryInterface
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "convert sfeel.interval  to ast",
			args: args{
				field: field.Field{Name: "X.Y", Type: dataType.Integer},
				entry: sfeel.CreateInputEntry("[1..6]"),
			},
			want:    "((X.Y :7: 1) :0: (X.Y :5: 6))",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateExpression(tt.args.field, tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("CreateExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpression_Convert(t *testing.T) {
	type fields struct {
		field field.Field
		entry dTable.EntryInterface
	}
	type args struct {
		targetFormat data.OutputFormat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "convert interval to grl",
			fields: fields{
				field: field.Field{Name: "X.Y", Type: dataType.Integer},
				entry: sfeel.CreateInputEntry("[1..6]"),
			},
			args:    args{targetFormat: data.GRL},
			want:    "((X.Y >= 1) && (X.Y <= 6))",
			wantErr: false,
		},
		{
			name: "convert interval to grl",
			fields: fields{
				field: field.Field{Name: "X.Y", Type: dataType.Integer},
				entry: sfeel.CreateInputEntry("[1..6]"),
			},
			args:    args{targetFormat: data.JSON},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := CreateExpression(tt.fields.field, tt.fields.entry)
			got, err := e.Convert(tt.args.targetFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
