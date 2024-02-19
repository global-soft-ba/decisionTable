package grl

import (
	"github.com/global-soft-ba/decisionTable/conv/grule/data"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"testing"
)

func TestExpression_Convert(t *testing.T) {
	type fields struct {
		field     field.Field
		entryType entryType.EntryType
		entry     string
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
			name: "convert string to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.String},
				entryType: entryType.Input,
				entry:     `"ABC"`,
			},
			args:    args{targetFormat: data.GRL},
			want:    `(X.Y == "ABC")`,
			wantErr: false,
		},
		{
			name: "convert positive integer to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     `10`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y == 10)",
			wantErr: false,
		},
		{
			name: "convert negative integer to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     `-10`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y == -10)",
			wantErr: false,
		},
		{
			name: "convert positive float to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Float},
				entryType: entryType.Input,
				entry:     `10.1`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y == 10.1)",
			wantErr: false,
		},
		//{ //TODO not implemented yet
		//	name: "convert negative float to grl",
		//	fields: fields{
		//		field:     field.Field{Name: "X.Y", Type: dataType.Float},
		//		entryType: entryType.Input,
		//		entry:     `-10.1`,
		//	},
		//	args:    args{targetFormat: data.GRL},
		//	want:    "(X.Y == -10.1)",
		//	wantErr: false,
		//},
		{
			name: "convert boolean to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Boolean},
				entryType: entryType.Input,
				entry:     `true`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y == true)",
			wantErr: false,
		},
		//{ // TODO not implemented yet
		//	name: "convert date to grl",
		//	fields: fields{
		//		field:     field.Field{Name: "X.Y", Type: dataType.DateTime},
		//		entryType: entryType.Input,
		//		entry:     `date and time(“2016-03-16T12:00:00”)`,
		//	},
		//	args:    args{targetFormat: data.GRL},
		//	want:    "",
		//	wantErr: false,
		//},
		{
			name: "convert comparison (less than) to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     `<10`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y < 10)",
			wantErr: false,
		},
		{
			name: "convert comparison (less than or equal) to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     `<=10`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y <= 10)",
			wantErr: false,
		},
		{
			name: "convert comparison (greater than) to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     `>10`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y > 10)",
			wantErr: false,
		},
		{
			name: "convert comparison (greater than or equal) to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     `>=10`,
			},
			args:    args{targetFormat: data.GRL},
			want:    "(X.Y >= 10)",
			wantErr: false,
		},
		{
			name: "convert interval to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     "[1..6]",
			},
			args:    args{targetFormat: data.GRL},
			want:    "((X.Y >= 1) && (X.Y <= 6))",
			wantErr: false,
		},
		{
			name: "convert interval to grl",
			fields: fields{
				field:     field.Field{Name: "X.Y", Type: dataType.Integer},
				entryType: entryType.Input,
				entry:     "[1..6]",
			},
			args:    args{targetFormat: data.JSON},
			want:    "",
			wantErr: true,
		},
		//{TODO not implemented yet
		//	name: "convert negation to grl",
		//	fields: fields{
		//		field:     field.Field{Name: "X.Y", Type: dataType.Integer},
		//		entryType: entryType.Input,
		//		entry:     "not(10)",
		//	},
		//	args:    args{targetFormat: data.GRL},
		//	want:    "(X.Y != 10)",
		//	wantErr: false,
		//},
		//{ TODO not implemented yet
		//	name: "convert disjunction to grl",
		//	fields: fields{
		//		field:     field.Field{Name: "X.Y", Type: dataType.Integer},
		//		entryType: entryType.Input,
		//		entry:     "<200, -100, 0, 100, >200",
		//	},
		//	args:    args{targetFormat: data.GRL},
		//	want:    "((X.Y < 200) || ((X.Y == -100) || ((X.Y == 0) || ((X.Y == 100) || (X.Y > 200))))), want = ((X.Y == -100) || ((X.Y == 0) || (X.Y == 100)))",
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := CreateExpression(tt.fields.field, expressionLanguage.SFEEL, tt.fields.entryType, tt.fields.entry)
			got, err := e.Convert(tt.args.targetFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
