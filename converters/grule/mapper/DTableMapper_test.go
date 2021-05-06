package mapper

import (
	"decisionTable/converters/grule/grlmodel"
	"decisionTable/model"
	"reflect"
	"testing"
)

func TestDTableToGrlMapper_MapToRuleSet(t *testing.T) {
	type args struct {
		data model.TableData
	}
	tests := []struct {
		name    string
		args    args
		want    grlmodel.RuleSet
		wantErr bool
	}{
		{name: "Valid Table",
			args: args{
				data: model.TableData{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        model.Priority,
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
						InputEntries: []model.Entry{
							model.CreateEntry("==3", model.SFEEL),
						},
						OutputEntries: []model.Entry{
							model.CreateEntry("4", model.SFEEL),
						},
					}},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.Priority,
				CollectOperator: model.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{{"I1", "L1", model.String, "==3", model.SFEEL}},
						[]grlmodel.Term{{"O1", "L1", model.Float, "4", model.SFEEL}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Multi Row Table",
			args: args{
				data: model.TableData{
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
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
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
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.String, "==3", model.SFEEL},
							{"I2", "L1", model.String, "==3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "4", model.SFEEL},
							{"O2", "L1", model.Float, "4", model.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Multi Row and Rule Table with First Policy",
			args: args{
				data: model.TableData{
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
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
							},
						},
						{
							Description: "R2",
							InputEntries: []model.Entry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("5", model.SFEEL),
								model.CreateEntry("5", model.SFEEL),
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
						1,
						[]grlmodel.Term{
							{"I1", "L1", model.String, "==3", model.SFEEL},
							{"I2", "L1", model.String, "==3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "4", model.SFEEL},
							{"O2", "L1", model.Float, "4", model.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.String, ">3", model.SFEEL},
							{"I2", "L1", model.String, ">3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "5", model.SFEEL},
							{"O2", "L1", model.Float, "5", model.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Multi Row Table",
			args: args{
				data: model.TableData{
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
						InputEntries: []model.Entry{
							model.CreateEntry("==3", model.SFEEL),
							model.CreateEntry("==3", model.SFEEL),
						},
						OutputEntries: []model.Entry{
							model.CreateEntry("4", model.SFEEL),
							model.CreateEntry("4", model.SFEEL)},
					}},
				}},
			want:    grlmodel.RuleSet{},
			wantErr: true,
		},
		{
			name: "Valid Multi Row and Multi Rule Table with Priority Policy",
			args: args{
				data: model.TableData{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        model.Priority,
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
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
							},
						},
						{
							Description: "R2",
							InputEntries: []model.Entry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("5", model.SFEEL),
								model.CreateEntry("5", model.SFEEL),
							},
						},
					},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.Priority,
				CollectOperator: model.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						1,
						[]grlmodel.Term{
							{"I1", "L1", model.String, "==3", model.SFEEL},
							{"I2", "L1", model.String, "==3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "4", model.SFEEL},
							{"O2", "L1", model.Float, "4", model.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.String, ">3", model.SFEEL},
							{"I2", "L1", model.String, ">3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "5", model.SFEEL},
							{"O2", "L1", model.Float, "5", model.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Multi Row and Multi Rule Table with First Policy",
			args: args{
				data: model.TableData{
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
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
							},
						},
						{
							Description: "R2",
							InputEntries: []model.Entry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("5", model.SFEEL),
								model.CreateEntry("5", model.SFEEL),
							},
						},
						{
							Description: "R3",
							InputEntries: []model.Entry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.Entry{
								model.CreateEntry("5", model.SFEEL),
								model.CreateEntry("5", model.SFEEL),
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
						2,
						[]grlmodel.Term{
							{"I1", "L1", model.String, "==3", model.SFEEL},
							{"I2", "L1", model.String, "==3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "4", model.SFEEL},
							{"O2", "L1", model.Float, "4", model.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Term{
							{"I1", "L1", model.String, ">3", model.SFEEL},
							{"I2", "L1", model.String, ">3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "5", model.SFEEL},
							{"O2", "L1", model.Float, "5", model.SFEEL},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.String, ">3", model.SFEEL},
							{"I2", "L1", model.String, ">3", model.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", model.Float, "5", model.SFEEL},
							{"O2", "L1", model.Float, "5", model.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := DTableMapper{}
			got, err := c.MapDTableToRuleSet(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapDTableToRuleSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapDTableToRuleSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
