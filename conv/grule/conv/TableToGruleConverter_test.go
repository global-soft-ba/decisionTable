package conv

import (
	grule "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/grl"
	dtable "decisionTable/data"
	"decisionTable/lang/sfeel"
	"reflect"
	"testing"
)

func CreateTestExpression(fieldName string, entry dtable.EntryInterface) grule.ExpressionInterface {
	res, _ := grl.CreateExpression(fieldName, entry)
	return res
}

func TestTableToGruleConverter_Convert(t *testing.T) {

	type args struct {
		table dtable.Table
	}
	tests := []struct {
		name    string
		args    args
		want    grule.RuleSet
		wantErr bool
	}{
		{name: "Valid Table",
			args: args{
				table: dtable.Table{
					Key:              "test1",
					Name:             "TableOne",
					HitPolicy:        dtable.Priority,
					CollectOperator:  dtable.List,
					NotationStandard: dtable.GRULE,
					InputFields: []dtable.FieldInterface{dtable.Field{
						Name: "I1",
						Key:  "L1",
						Typ:  dtable.String,
					},
					},
					OutputFields: []dtable.FieldInterface{dtable.Field{
						Name: "O1",
						Key:  "L1",
						Typ:  dtable.Float,
					}},
					Rules: []dtable.Rule{{
						Description: "R1",
						InputEntries: []dtable.EntryInterface{
							sfeel.CreateInputEntry("3"),
						},
						OutputEntries: []dtable.EntryInterface{
							sfeel.CreateOutputEntry("4"),
						},
					}},
				}},
			want: grule.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       dtable.Priority,
				CollectOperator: dtable.List,
				Rules: []grule.Rule{
					{
						Name:        "0",
						Description: "R1",
						Expressions: []grule.Term{
							{
								dtable.Field{"I1", "L1", dtable.String},
								CreateTestExpression("I1.L1", sfeel.CreateInputEntry("3")),
								dtable.SFEEL},
						},
						Assignments: []grule.Term{
							{
								Field:              dtable.Field{"O1", "L1", dtable.Float},
								Expression:         CreateTestExpression("I1.L1", sfeel.CreateOutputEntry("4")),
								ExpressionLanguage: dtable.SFEEL,
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := TableToGruleConverter{}
			got, err := c.Convert(tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
{name: "Valid Table",
			args: args{
				data: model.TableData{
					Key:              "test1",
					Field:             "TableOne",
					HitPolicy:        model.Priority,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{{
						Field: "I1",
						Key:  "L1",
						Typ:  model.String,
					},
					},
					OutputFields: []model.Field{{
						Field: "O1",
						Key:  "L1",
						Typ:  model.Float,
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
				Field:            "TableOne",
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
					Field:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{
						{
							Field: "I1",
							Key:  "L1",
							Typ:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Typ:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Typ:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Typ:  model.Float,
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
				Field:            "TableOne",
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
					Field:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{
						{
							Field: "I1",
							Key:  "L1",
							Typ:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Typ:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Typ:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Typ:  model.Float,
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
				Field:            "TableOne",
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
					Field:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{{
						Field: "I1",
						Key:  "L1",
						Typ:  model.String,
					}},
					OutputFields: []model.Field{{
						Field: "O1",
						Key:  "L1",
						Typ:  model.Float,
					}, {
						Field: "O2",
						Key:  "L1",
						Typ:  model.Float,
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
					Field:             "TableOne",
					HitPolicy:        model.Priority,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{
						{
							Field: "I1",
							Key:  "L1",
							Typ:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Typ:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Typ:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Typ:  model.Float,
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
				Field:            "TableOne",
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
					Field:             "TableOne",
					HitPolicy:        model.First,
					CollectOperator:  model.List,
					NotationStandard: model.GRULE,
					InputFields: []model.Field{
						{
							Field: "I1",
							Key:  "L1",
							Typ:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Typ:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Typ:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Typ:  model.Float,
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
				Field:            "TableOne",
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
*/
