package mapper

import (
	converter "decisionTable/converter/grule/model"
	"decisionTable/model"
	"reflect"
	"testing"
)

func TestDTableToGrlMapper_MapToRuleSet(t *testing.T) {
	type args struct {
		data model.DTableData
	}
	tests := []struct {
		name    string
		args    args
		want    converter.RuleSet
		wantErr bool
	}{
		{name: "Valid Table",
			args: args{
				model.DTableData{
					"test1",
					"TableOne",
					model.First,
					model.List,
					model.GRULE,
					[]model.Field{{
						Name:  "I1",
						Label: "L1",
						Typ:   model.String,
					},
					},
					[]model.Field{{
						Name:  "O1",
						Label: "L1",
						Typ:   model.Float,
					}},
					[]model.Rule{{
						Description: "R1",
						InputEntries: []model.Entry{{
							Expression:         "==3",
							ExpressionLanguage: model.GRL,
						}},
						OutputEntries: []model.Entry{{
							Expression:         "4",
							ExpressionLanguage: model.GRL,
						}},
					}},
				}},
			want: converter.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Rules: []converter.Rule{
					{
						"0",
						"R1",
						0,
						[]converter.Expression{{"I1", "L1", "==3"}},
						[]converter.Assignment{{"O1", "L1", "4"}},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := DTableToGrlMapper{}
			got, err := c.MapToRuleSet(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToRuleSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToRuleSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
