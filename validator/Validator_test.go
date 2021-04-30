package validator

import (
	"decisionTable/model"
	"decisionTable/validator/expression"
	"reflect"
	"testing"
)

func TestDTableValidator_Validate(t *testing.T) {
	tests := []struct {
		name  string
		table model.TableData
		want  bool
		want1 []error
	}{
		{
			name: "Valid Grule Table",
			table: model.TableData{
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
				Rules: []model.Rule{
					{
						Description: "R1",
						InputEntries: []model.Entry{
							model.CreateEntry("<3", model.SFEEL),
						},
						OutputEntries: []model.Entry{
							model.CreateEntry("-", model.SFEEL),
						},
					},
					{
						Description: "R2",
						InputEntries: []model.Entry{
							model.CreateEntry("<=3", model.SFEEL),
						},
						OutputEntries: []model.Entry{
							model.CreateEntry(`"yes"`, model.SFEEL),
						},
					},
					{
						Description: "R3",
						InputEntries: []model.Entry{
							model.CreateEntry("not(47)", model.SFEEL),
						},
						OutputEntries: []model.Entry{
							model.CreateEntry(`"no"`, model.SFEEL),
						},
					},
				},
			},
			want:  true,
			want1: nil,
		},
		{
			name: "Wrong hit policy for the NotationStandard",
			table: model.TableData{
				Key:              "test1",
				Name:             "TableOne",
				HitPolicy:        model.Any,
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
					Typ:   model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{
						model.CreateEntry("3", model.SFEEL)},
					OutputEntries: []model.Entry{
						model.CreateEntry("4", model.SFEEL),
					},
				}},
			},
			want:  false,
			want1: []error{ErrDTableHitPolicy},
		},
		{
			name: "Wrong CollectOperators DMN",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.Collect,
				CollectOperator:  "XYZ",
				NotationStandard: model.DMN,
				InputFields: []model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.String,
				},
				},
				OutputFields: []model.Field{{
					Name:  "O1",
					Label: "L1",
					Typ:   model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{
						model.CreateEntry("3", model.SFEEL)},
					OutputEntries: []model.Entry{
						model.CreateEntry("4", model.SFEEL)},
				}},
			},
			want:  false,
			want1: []error{ErrDTableCollectOperator, ErrDTableEntryExpressionLangInvalid, ErrDTableEntryExpressionLangInvalid},
		},
		{
			name: "Empty input and output fields",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.First,
				NotationStandard: model.GRULE,
				InputFields:      []model.Field{},
				OutputFields:     []model.Field{},
				Rules: []model.Rule{{
					Description:   "R1",
					InputEntries:  []model.Entry{},
					OutputEntries: []model.Entry{},
				}},
			},
			want:  false,
			want1: []error{ErrDTableInputEmpty, ErrDTableOutputEmpty},
		},
		{
			name: "Wrong Field definition",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.First,
				NotationStandard: model.GRULE,
				InputFields: []model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.Double,
				},
				},
				OutputFields: []model.Field{{
					Name:  "",
					Label: "L1",
					Typ:   model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{
						model.CreateEntry(">3", model.SFEEL)},
					OutputEntries: []model.Entry{
						model.CreateEntry("3", model.SFEEL),
					},
				}},
			},
			want:  false,
			want1: []error{ErrDTableFieldTypInvalid, ErrDTableFieldNameEmpty},
		},
		{
			name: "Wrong rule data schemas",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.First,
				NotationStandard: model.GRULE,
				InputFields: []model.Field{
					{
						Name:  "I1",
						Label: "L1",
						Typ:   model.String,
					},
				},
				OutputFields: []model.Field{{
					Name:  "O2",
					Label: "L1",
					Typ:   model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{
						model.CreateEntry("<3", model.SFEEL),
						model.CreateEntry(">3", model.SFEEL)},
					OutputEntries: []model.Entry{},
				}},
			},
			want:  false,
			want1: []error{ErrRuleHaveDifferentAmountOfInputFields, ErrRuleHaveDifferentAmountOfOutputFields},
		},
		{
			name: "Wrong rule expression language",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.First,
				NotationStandard: model.GRULE,
				InputFields: []model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.String,
				},
				},
				OutputFields: []model.Field{{
					Name:  "O2",
					Label: "L1",
					Typ:   model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{
						model.CreateEntry("3", model.FEEL)},
					OutputEntries: []model.Entry{
						model.CreateEntry("3", model.SFEEL)},
				}},
			},
			want:  false,
			want1: []error{ErrDTableEntryExpressionLangInvalid, expression.ErrDTableNoParserFoundForExpressionLanguage},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Validator{
				dTable: tt.table,
			}
			got, got1 := d.Validate()
			if got != tt.want {
				t.Errorf("Validate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Validate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDTableValidator_ValidateInterferences(t *testing.T) {

	tests := []struct {
		name  string
		field model.TableData
		want  bool
	}{
		{
			name: "Valid Table Without Interferences",
			field: model.TableData{
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
					InputEntries: []model.Entry{
						model.CreateEntry(" =3", model.SFEEL)},
					OutputEntries: []model.Entry{
						model.CreateEntry("4", model.SFEEL)},
				}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Validator{
				dTable: tt.field,
			}
			if got := d.ValidateContainsInterferences(); got != tt.want {
				t.Errorf("ValidateContainsInterferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
