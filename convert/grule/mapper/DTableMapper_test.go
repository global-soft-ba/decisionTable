package mapper

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/data"
	"reflect"
	"testing"
)

func TestDTableToGrlMapper_MapToRuleSet(t *testing.T) {
	type args struct {
		data data.Table
	}
	tests := []struct {
		name    string
		args    args
		want    grlmodel.RuleSet
		wantErr bool
	}{
		{name: "Valid Table",
			args: args{
				data: data.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        data.Priority,
					CollectOperator:  data.List,
					NotationStandard: data.GRULE,
					InputFields: []data.Field{{
						Name: "I1",
						Key:  "L1",
						Typ:  data.String,
					},
					},
					OutputFields: []data.Field{{
						Name: "O1",
						Key:  "L1",
						Typ:  data.Float,
					}},
					Rules: []data.Rule{{
						Description: "R1",
						InputEntries: []data.Entry{
							data.CreateEntry("==3", data.SFEEL),
						},
						OutputEntries: []data.Entry{
							data.CreateEntry("4", data.SFEEL),
						},
					}},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       data.Priority,
				CollectOperator: data.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{{"I1", "L1", data.String, "==3", data.SFEEL}},
						[]grlmodel.Term{{"O1", "L1", data.Float, "4", data.SFEEL}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Multi Row Table",
			args: args{
				data: data.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        data.First,
					CollectOperator:  data.List,
					NotationStandard: data.GRULE,
					InputFields: []data.Field{
						{
							Name: "I1",
							Key:  "L1",
							Typ:  data.String,
						},
						{
							Name: "I2",
							Key:  "L1",
							Typ:  data.String,
						},
					},
					OutputFields: []data.Field{
						{
							Name: "O1",
							Key:  "L1",
							Typ:  data.Float,
						},
						{
							Name: "O2",
							Key:  "L1",
							Typ:  data.Float,
						},
					},
					Rules: []data.Rule{
						{
							Description: "R1",
							InputEntries: []data.Entry{
								data.CreateEntry("==3", data.SFEEL),
								data.CreateEntry("==3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("4", data.SFEEL),
								data.CreateEntry("4", data.SFEEL),
							},
						}},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       data.First,
				CollectOperator: data.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", data.String, "==3", data.SFEEL},
							{"I2", "L1", data.String, "==3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "4", data.SFEEL},
							{"O2", "L1", data.Float, "4", data.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Multi Row and Rule Table with First Policy",
			args: args{
				data: data.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        data.First,
					CollectOperator:  data.List,
					NotationStandard: data.GRULE,
					InputFields: []data.Field{
						{
							Name: "I1",
							Key:  "L1",
							Typ:  data.String,
						},
						{
							Name: "I2",
							Key:  "L1",
							Typ:  data.String,
						},
					},
					OutputFields: []data.Field{
						{
							Name: "O1",
							Key:  "L1",
							Typ:  data.Float,
						},
						{
							Name: "O2",
							Key:  "L1",
							Typ:  data.Float,
						},
					},
					Rules: []data.Rule{
						{
							Description: "R1",
							InputEntries: []data.Entry{
								data.CreateEntry("==3", data.SFEEL),
								data.CreateEntry("==3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("4", data.SFEEL),
								data.CreateEntry("4", data.SFEEL),
							},
						},
						{
							Description: "R2",
							InputEntries: []data.Entry{
								data.CreateEntry(">3", data.SFEEL),
								data.CreateEntry(">3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("5", data.SFEEL),
								data.CreateEntry("5", data.SFEEL),
							},
						},
					},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       data.First,
				CollectOperator: data.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						1,
						[]grlmodel.Term{
							{"I1", "L1", data.String, "==3", data.SFEEL},
							{"I2", "L1", data.String, "==3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "4", data.SFEEL},
							{"O2", "L1", data.Float, "4", data.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						0,
						[]grlmodel.Term{
							{"I1", "L1", data.String, ">3", data.SFEEL},
							{"I2", "L1", data.String, ">3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "5", data.SFEEL},
							{"O2", "L1", data.Float, "5", data.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Multi Row Table",
			args: args{
				data: data.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        data.First,
					CollectOperator:  data.List,
					NotationStandard: data.GRULE,
					InputFields: []data.Field{{
						Name: "I1",
						Key:  "L1",
						Typ:  data.String,
					}},
					OutputFields: []data.Field{{
						Name: "O1",
						Key:  "L1",
						Typ:  data.Float,
					}, {
						Name: "O2",
						Key:  "L1",
						Typ:  data.Float,
					}},
					Rules: []data.Rule{{
						Description: "R1",
						InputEntries: []data.Entry{
							data.CreateEntry("==3", data.SFEEL),
							data.CreateEntry("==3", data.SFEEL),
						},
						OutputEntries: []data.Entry{
							data.CreateEntry("4", data.SFEEL),
							data.CreateEntry("4", data.SFEEL)},
					}},
				}},
			want:    grlmodel.RuleSet{},
			wantErr: true,
		},
		{
			name: "Valid Multi Row and Multi Rule Table with Priority Policy",
			args: args{
				data: data.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        data.Priority,
					CollectOperator:  data.List,
					NotationStandard: data.GRULE,
					InputFields: []data.Field{
						{
							Name: "I1",
							Key:  "L1",
							Typ:  data.String,
						},
						{
							Name: "I2",
							Key:  "L1",
							Typ:  data.String,
						},
					},
					OutputFields: []data.Field{
						{
							Name: "O1",
							Key:  "L1",
							Typ:  data.Float,
						},
						{
							Name: "O2",
							Key:  "L1",
							Typ:  data.Float,
						},
					},
					Rules: []data.Rule{
						{
							Description: "R1",
							InputEntries: []data.Entry{
								data.CreateEntry("==3", data.SFEEL),
								data.CreateEntry("==3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("4", data.SFEEL),
								data.CreateEntry("4", data.SFEEL),
							},
						},
						{
							Description: "R2",
							InputEntries: []data.Entry{
								data.CreateEntry(">3", data.SFEEL),
								data.CreateEntry(">3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("5", data.SFEEL),
								data.CreateEntry("5", data.SFEEL),
							},
						},
					},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       data.Priority,
				CollectOperator: data.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						1,
						[]grlmodel.Term{
							{"I1", "L1", data.String, "==3", data.SFEEL},
							{"I2", "L1", data.String, "==3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "4", data.SFEEL},
							{"O2", "L1", data.Float, "4", data.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						0,
						[]grlmodel.Term{
							{"I1", "L1", data.String, ">3", data.SFEEL},
							{"I2", "L1", data.String, ">3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "5", data.SFEEL},
							{"O2", "L1", data.Float, "5", data.SFEEL},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Multi Row and Multi Rule Table with First Policy",
			args: args{
				data: data.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        data.First,
					CollectOperator:  data.List,
					NotationStandard: data.GRULE,
					InputFields: []data.Field{
						{
							Name: "I1",
							Key:  "L1",
							Typ:  data.String,
						},
						{
							Name: "I2",
							Key:  "L1",
							Typ:  data.String,
						},
					},
					OutputFields: []data.Field{
						{
							Name: "O1",
							Key:  "L1",
							Typ:  data.Float,
						},
						{
							Name: "O2",
							Key:  "L1",
							Typ:  data.Float,
						},
					},
					Rules: []data.Rule{
						{
							Description: "R1",
							InputEntries: []data.Entry{
								data.CreateEntry("==3", data.SFEEL),
								data.CreateEntry("==3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("4", data.SFEEL),
								data.CreateEntry("4", data.SFEEL),
							},
						},
						{
							Description: "R2",
							InputEntries: []data.Entry{
								data.CreateEntry(">3", data.SFEEL),
								data.CreateEntry(">3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("5", data.SFEEL),
								data.CreateEntry("5", data.SFEEL),
							},
						},
						{
							Description: "R3",
							InputEntries: []data.Entry{
								data.CreateEntry(">3", data.SFEEL),
								data.CreateEntry(">3", data.SFEEL),
							},
							OutputEntries: []data.Entry{
								data.CreateEntry("5", data.SFEEL),
								data.CreateEntry("5", data.SFEEL),
							},
						},
					},
				}},
			want: grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       data.First,
				CollectOperator: data.List,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						2,
						[]grlmodel.Term{
							{"I1", "L1", data.String, "==3", data.SFEEL},
							{"I2", "L1", data.String, "==3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "4", data.SFEEL},
							{"O2", "L1", data.Float, "4", data.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Term{
							{"I1", "L1", data.String, ">3", data.SFEEL},
							{"I2", "L1", data.String, ">3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "5", data.SFEEL},
							{"O2", "L1", data.Float, "5", data.SFEEL},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Term{
							{"I1", "L1", data.String, ">3", data.SFEEL},
							{"I2", "L1", data.String, ">3", data.SFEEL}},
						[]grlmodel.Term{
							{"O1", "L1", data.Float, "5", data.SFEEL},
							{"O2", "L1", data.Float, "5", data.SFEEL},
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
