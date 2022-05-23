package conv

import (
	grule "github.com/global-soft-ba/decisionTable/conv/grule/data"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl"
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/decisionTable"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/rule"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func CreateTestExpression(field field.Field, expressionLanguage expressionLanguage.ExpressionLanguage, entryType entryType.EntryType, entry string) grule.ExpressionInterface {
	res, _ := grl.CreateExpression(field, expressionLanguage, entryType, entry)
	return res
}

func TestTableToGruleConverter_Convert(t *testing.T) {

	type args struct {
		table decisionTable.DecisionTable
	}
	tests := []struct {
		name    string
		args    args
		want    grule.RuleSet
		wantErr bool
	}{
		{name: "Validate Table",
			args: args{
				table: decisionTable.DecisionTable{
					ID:                 "test1",
					Name:               "TableOne",
					HitPolicy:          hitPolicy.Priority,
					CollectOperator:    collectOperator.List,
					ExpressionLanguage: expressionLanguage.SFEEL,
					Standard:           standard.GRULE,
					InputFields: []field.Field{{
						Name: "I1.L1",
						Type: dataType.String,
					},
					},
					OutputFields: []field.Field{{
						Name: "O1.L1",
						Type: dataType.Float,
					}},
					Rules: []rule.Rule{{
						Annotation: "R1",
						InputEntries: []string{
							"3",
						},
						OutputEntries: []string{
							"4",
						},
					}},
				}},
			want: grule.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       hitPolicy.Priority,
				CollectOperator: collectOperator.List,
				Rules: []grule.Rule{
					{
						Name:       "0",
						Annotation: "R1",
						Expressions: []grule.Term{
							{
								Field:              field.Field{Name: "I1.L1", Type: dataType.String},
								Expression:         CreateTestExpression(field.Field{Name: "I1.L1", Type: dataType.String}, expressionLanguage.SFEEL, entryType.Input, "3"),
								ExpressionLanguage: expressionLanguage.SFEEL},
						},
						Assignments: []grule.Term{
							{
								Field:              field.Field{Name: "O1.L1", Type: dataType.Float},
								Expression:         CreateTestExpression(field.Field{Name: "O1.L1", Type: dataType.Float}, expressionLanguage.SFEEL, entryType.Output, "4"),
								ExpressionLanguage: expressionLanguage.SFEEL,
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

func TestTableToGruleConverter_CheckForInterferences(t *testing.T) {
	tests := []struct {
		name          string
		decisionTable decisionTable.DecisionTable
		want          bool
	}{
		{
			name: "Valid table without interferences",
			decisionTable: decisionTable.DecisionTable{
				InputFields: []field.Field{
					{
						Name: "I1.L1",
						Type: dataType.Integer,
					},
				},
				OutputFields: []field.Field{
					{
						Name: "I2.L1",
						Type: dataType.Integer,
					},
				},
			},
			want: false,
		},
		{
			name: "Valid table with interferences",
			decisionTable: decisionTable.DecisionTable{
				InputFields: []field.Field{
					{
						Name: "I1.L1",
						Type: dataType.Integer,
					},
				},
				OutputFields: []field.Field{
					{
						Name: "I1.L1",
						Type: dataType.Integer,
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := TableToGruleConverter{}
			result := c.checkIfContainsInterferences(tt.decisionTable)
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("CheckIfContainsInterferences() got = %v, want %v", result, tt.want)
			}
		})
	}
}

//ToDo Migrate testcases
/*
{name: "Validate Table",
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
						Type:  model.String,
					},
					},
					OutputFields: []model.Field{{
						Field: "O1",
						Key:  "L1",
						Type:  model.Float,
					}},
					Rules: []model.Rule{{
						Annotation: "R1",
						InputEntries: []model.DummyEntry{
							model.CreateEntry("==3", model.SFEEL),
						},
						OutputEntries: []model.DummyEntry{
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
			name: "Validate Multi Row Table",
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
							Type:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Type:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Type:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Type:  model.Float,
						},
					},
					Rules: []model.Rule{
						{
							Annotation: "R1",
							InputEntries: []model.DummyEntry{
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
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
			name: "Validate Multi Row and Rule Table with First Policy",
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
							Type:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Type:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Type:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Type:  model.Float,
						},
					},
					Rules: []model.Rule{
						{
							Annotation: "R1",
							InputEntries: []model.DummyEntry{
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
							},
						},
						{
							Annotation: "R2",
							InputEntries: []model.DummyEntry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
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
						Type:  model.String,
					}},
					OutputFields: []model.Field{{
						Field: "O1",
						Key:  "L1",
						Type:  model.Float,
					}, {
						Field: "O2",
						Key:  "L1",
						Type:  model.Float,
					}},
					Rules: []model.Rule{{
						Annotation: "R1",
						InputEntries: []model.DummyEntry{
							model.CreateEntry("==3", model.SFEEL),
							model.CreateEntry("==3", model.SFEEL),
						},
						OutputEntries: []model.DummyEntry{
							model.CreateEntry("4", model.SFEEL),
							model.CreateEntry("4", model.SFEEL)},
					}},
				}},
			want:    grlmodel.RuleSet{},
			wantErr: true,
		},
		{
			name: "Validate Multi Row and Multi Rule Table with Priority Policy",
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
							Type:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Type:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Type:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Type:  model.Float,
						},
					},
					Rules: []model.Rule{
						{
							Annotation: "R1",
							InputEntries: []model.DummyEntry{
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
							},
						},
						{
							Annotation: "R2",
							InputEntries: []model.DummyEntry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
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
			name: "Validate Multi Row and Multi Rule Table with First Policy",
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
							Type:  model.String,
						},
						{
							Field: "I2",
							Key:  "L1",
							Type:  model.String,
						},
					},
					OutputFields: []model.Field{
						{
							Field: "O1",
							Key:  "L1",
							Type:  model.Float,
						},
						{
							Field: "O2",
							Key:  "L1",
							Type:  model.Float,
						},
					},
					Rules: []model.Rule{
						{
							Annotation: "R1",
							InputEntries: []model.DummyEntry{
								model.CreateEntry("==3", model.SFEEL),
								model.CreateEntry("==3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
								model.CreateEntry("4", model.SFEEL),
								model.CreateEntry("4", model.SFEEL),
							},
						},
						{
							Annotation: "R2",
							InputEntries: []model.DummyEntry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
								model.CreateEntry("5", model.SFEEL),
								model.CreateEntry("5", model.SFEEL),
							},
						},
						{
							Annotation: "R3",
							InputEntries: []model.DummyEntry{
								model.CreateEntry(">3", model.SFEEL),
								model.CreateEntry(">3", model.SFEEL),
							},
							OutputEntries: []model.DummyEntry{
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
