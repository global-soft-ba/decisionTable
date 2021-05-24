package sfeel

import (
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
			got, err := CreateInputEntry(tt.args.exp)
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
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateOutputEntry(tt.args.exp)
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
