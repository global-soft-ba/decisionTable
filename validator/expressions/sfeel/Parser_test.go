package sfeel

import (
	"decisionTable/model"
	"fmt"
	"testing"
)

func TestParser_ValidMultipleRequest(t *testing.T) {
	parser := CreateParser()

	InputEntries := []model.Entry{
		model.CreateEntry(">3", model.SFEEL),
		model.CreateEntry("3", model.SFEEL),
		model.CreateEntry("3x", model.SFEEL),
	}

	for _, v := range InputEntries {
		ok, err := parser.ValidateInputEntry(v)
		fmt.Println(ok, err)
	}
}

func TestParser_ValidateInputEntry(t *testing.T) {

	tests := []struct {
		name string
		args model.Entry
		want bool
	}{
		{
			name: "Valid Expression",
			args: model.CreateEntry(`<x10`, model.FEEL),
			want: false,
		},
		{
			name: "Invalid Expression",
			args: model.CreateEntry(`<10`, model.FEEL),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateParser()
			got, _ := p.ValidateInputEntry(tt.args)
			if got != tt.want {
				t.Errorf("ValidateInputEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ValidateOutputEntry(t *testing.T) {

	tests := []struct {
		name string
		args model.Entry
		want bool
	}{
		{
			name: "Valid Expression",
			args: model.CreateEntry(`10`, model.FEEL),
			want: true,
		},
		{
			name: "Valid Expression",
			args: model.CreateEntry(`10`, model.FEEL),
			want: true,
		},
		{
			name: "Invalid Expression",
			args: model.CreateEntry(`<10`, model.FEEL),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateParser()
			got, err := p.ValidateOutputEntry(tt.args)

			if got != tt.want {
				t.Errorf("ValidateInputEntry() got = %v, want %v with %v", got, tt.want, err)
			}

		})
	}
}
