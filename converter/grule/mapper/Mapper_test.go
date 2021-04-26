package mapper

import (
	"decisionTable/converter/grule/grlmodel"
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
		want    grlmodel.RuleSet
		wantErr bool
	}{
		{name: "Valid Table",
			args: args{
				data: model.DTableData{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{{
						Name:  "I1",
						Label: "L1",
						Typ:   model.String,
					},
					},
					OutputFields: []model.Field{{
						Name:  "O1",
						Label: "L1",
						Typ:   model.Float,
					}},
					Rules: []model.Rule{{
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
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						[]grlmodel.Expression{{"I1", "L1", "==3"}},
						[]grlmodel.Expression{{"O1", "L1", "4"}},
					},
				},
			},
			wantErr: false,
		},
		{name: "Valid Multi Row Table",
			args: args{
				data: model.DTableData{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{
						{
							Name:  "I1",
							Label: "L1",
							Typ:   model.String,
						},
						{
							Name:  "I2",
							Label: "L1",
							Typ:   model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Name:  "O1",
							Label: "L1",
							Typ:   model.Float,
						},
						{
							Name:  "O2",
							Label: "L1",
							Typ:   model.Float,
						},
					},
					Rules: []model.Rule{
						{
							Description: "R1",
							InputEntries: []model.Entry{
								{
									Expression:         "==3",
									ExpressionLanguage: model.GRL,
								},
								{
									Expression:         "==3",
									ExpressionLanguage: model.GRL,
								},
							},
							OutputEntries: []model.Entry{
								{
									Expression:         "4",
									ExpressionLanguage: model.GRL,
								},
								{
									Expression:         "4",
									ExpressionLanguage: model.GRL,
								},
							},
						}},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						[]grlmodel.Expression{
							{"I1", "L1", "==3"},
							{"I2", "L1", "==3"}},
						[]grlmodel.Expression{
							{"O1", "L1", "4"},
							{"O2", "L1", "4"},
						},
					},
				},
			},
			wantErr: false,
		},
		{name: "Valid Multi Row and Rule Table",
			args: args{
				data: model.DTableData{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{
						{
							Name:  "I1",
							Label: "L1",
							Typ:   model.String,
						},
						{
							Name:  "I2",
							Label: "L1",
							Typ:   model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Name:  "O1",
							Label: "L1",
							Typ:   model.Float,
						},
						{
							Name:  "O2",
							Label: "L1",
							Typ:   model.Float,
						},
					},
					Rules: []model.Rule{
						{
							Description: "R1",
							InputEntries: []model.Entry{
								{
									Expression:         "==3",
									ExpressionLanguage: model.GRL,
								},
								{
									Expression:         "==3",
									ExpressionLanguage: model.GRL,
								},
							},
							OutputEntries: []model.Entry{
								{
									Expression:         "4",
									ExpressionLanguage: model.GRL,
								},
								{
									Expression:         "4",
									ExpressionLanguage: model.GRL,
								},
							},
						},
						{
							Description: "R2",
							InputEntries: []model.Entry{
								{
									Expression:         ">3",
									ExpressionLanguage: model.GRL,
								},
								{
									Expression:         ">3",
									ExpressionLanguage: model.GRL,
								},
							},
							OutputEntries: []model.Entry{
								{
									Expression:         "5",
									ExpressionLanguage: model.GRL,
								},
								{
									Expression:         "5",
									ExpressionLanguage: model.GRL,
								},
							},
						},
					},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						[]grlmodel.Expression{
							{"I1", "L1", "==3"},
							{"I2", "L1", "==3"}},
						[]grlmodel.Expression{
							{"O1", "L1", "4"},
							{"O2", "L1", "4"},
						},
					},
					{
						"1",
						"R2",
						1,
						[]grlmodel.Expression{
							{"I1", "L1", ">3"},
							{"I2", "L1", ">3"}},
						[]grlmodel.Expression{
							{"O1", "L1", "5"},
							{"O2", "L1", "5"},
						},
					},
				},
			},
			wantErr: false,
		},
		{name: "Invalid Multi Row Table",
			args: args{
				data: model.DTableData{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{{
						Name:  "I1",
						Label: "L1",
						Typ:   model.String,
					}},
					OutputFields: []model.Field{{
						Name:  "O1",
						Label: "L1",
						Typ:   model.Float,
					}, {
						Name:  "O2",
						Label: "L1",
						Typ:   model.Float,
					}},
					Rules: []model.Rule{{
						Description: "R1",
						InputEntries: []model.Entry{{
							Expression:         "==3",
							ExpressionLanguage: model.GRL,
						},
							{
								Expression:         "==3",
								ExpressionLanguage: model.GRL,
							}},
						OutputEntries: []model.Entry{{
							Expression:         "4",
							ExpressionLanguage: model.GRL,
						},
							{
								Expression:         "4",
								ExpressionLanguage: model.GRL,
							}},
					}},
				}},
			want:    grlmodel.RuleSet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Mapper{}
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
