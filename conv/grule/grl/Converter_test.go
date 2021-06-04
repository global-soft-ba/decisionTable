package grl

import (
	grl "decisionTable/conv/grule/data"
	table "decisionTable/data"
	"decisionTable/lang/sfeel"
	"testing"
)

func TestConverter_Convert(t *testing.T) {
	type args struct {
		fieldName    string
		entry        table.EntryInterface
		targetFormat grl.OutputFormat
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid integer convert",
			args: args{
				fieldName:    "X",
				entry:        sfeel.CreateInputEntry("12345"),
				targetFormat: grl.GRL,
			},
			want:    "12345",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Converter{}
			got, err := c.Convert(tt.args.fieldName, tt.args.entry, tt.args.targetFormat)
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
