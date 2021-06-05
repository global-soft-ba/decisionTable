package grl

import (
	"decisionTable/conv/grule/data"
	dTable "decisionTable/data"
	"decisionTable/lang/sfeel"
	"reflect"
	"testing"
)

func TestCreateExpression(t *testing.T) {
	type args struct {
		fieldName string
		entry     dTable.EntryInterface
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
				fieldName: "X",
				entry:     sfeel.CreateInputEntry("[1..6]"),
			},
			want:    "X opId:6 1 opId:0 X opId:4 6",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateExpression(tt.args.fieldName, tt.args.entry)
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
		fieldName string
		entry     dTable.EntryInterface
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
				fieldName: "X",
				entry:     sfeel.CreateInputEntry("[1..6]"),
			},
			args:    args{targetFormat: data.GRL},
			want:    "((X >= 1) && (X <= 6))",
			wantErr: false,
		},
		{
			name: "convert interval to grl",
			fields: fields{
				fieldName: "X",
				entry:     sfeel.CreateInputEntry("[1..6]"),
			},
			args:    args{targetFormat: data.JSON},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := CreateExpression(tt.fields.fieldName, tt.fields.entry)
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
