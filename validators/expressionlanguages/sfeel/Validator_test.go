package sfeel

import (
	"decisionTable/model"
	"fmt"
	"testing"
)

func TestParser_ValidMultipleRequest(t *testing.T) {
	parser := CreateValidator()

	InputEntries := []model.Entry{
		model.CreateEntry(`"valid string"`, model.SFEEL),
		model.CreateEntry(`<"invalid string operation`, model.SFEEL),
		model.CreateEntry("3", model.SFEEL),
	}

	for _, v := range InputEntries {
		ok, err := parser.ValidateInputEntry(model.Field{"", "", model.Integer}, v)
		fmt.Println(ok, err)
	}
}

func TestValidator_ValidateInputEntry(t *testing.T) {
	type args struct {
		field model.Field
		entry model.Entry
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid String Expression",
			args: args{
				model.Field{Name: "", Key: "", Typ: model.String},
				model.CreateEntry(`"Hello"`, model.FEEL),
			},
			want: true,
		},
		{
			name: "Invalid String Expression",
			args: args{
				model.Field{Name: "", Key: "", Typ: model.String},
				model.CreateEntry(`<"Hello"`, model.FEEL),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			got, err := v.ValidateInputEntry(tt.args.field, tt.args.entry)
			if got != tt.want {
				t.Errorf("ValidateInputEntry() got = %v, want %v with error %v", got, tt.want, err)
			}

		})
	}
}

func TestValidator_ValidateOutputEntry(t *testing.T) {
	type args struct {
		field model.Field
		entry model.Entry
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid String Expression",
			args: args{
				model.Field{Name: "", Key: "", Typ: model.String},
				model.CreateEntry(`"Hello"`, model.FEEL),
			},
			want: true,
		},
		{
			name: "Invalid String Expression",
			args: args{
				model.Field{Name: "", Key: "", Typ: model.String},
				model.CreateEntry(`"Hello","Hello"`, model.FEEL),
			},
			want: false,
		},
		{
			name: "Invalid Integer Expression",
			args: args{
				model.Field{Name: "", Key: "", Typ: model.Integer},
				model.CreateEntry("1,2", model.FEEL),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			got, err := v.ValidateOutputEntry(tt.args.field, tt.args.entry)
			if got != tt.want {
				t.Errorf("ValidateOutputEntry() got = %v, want %v with error %v", got, tt.want, err)
			}
		})
	}
}
