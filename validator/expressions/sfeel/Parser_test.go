package sfeel

import (
	"decisionTable/model"
	"testing"
)

func TestParser_ValidateEntry(t *testing.T) {

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
			got, _ := p.ValidateEntry(tt.args)
			if got != tt.want {
				t.Errorf("ValidateEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}
