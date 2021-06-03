package conv

import (
	"decisionTable/data"
	"decisionTable/lang/sfeel"
	"testing"
)

func TestSFeelConverter_Convert(t *testing.T) {
	type fields struct {
		listener SFeelListener
	}
	type args struct {
		e data.EntryInterface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple integer test",
			fields: fields{listener: SFeelListener{data: "Init"}},
			args:   args{sfeel.CreateInputEntry("123456789")},
		},
		{
			name:   "simple unary test",
			fields: fields{listener: SFeelListener{data: "Init"}},
			args:   args{sfeel.CreateInputEntry("<1")},
		},
		{
			name:   "simple unary tests",
			fields: fields{listener: SFeelListener{data: "Init"}},
			args:   args{sfeel.CreateInputEntry("<1,>=4,>5")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := SFeelConverter{
				listener: tt.fields.listener,
			}
			c.Convert(tt.args.e)
		})
	}
}
