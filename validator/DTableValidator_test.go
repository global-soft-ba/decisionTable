package validator

import (
	"decisionTable/model"
	"reflect"
	"testing"
)

func TestDTableValidator_Validate(t *testing.T) {
	tests := []struct {
		name  string
		table model.DTableData
		want  bool
		want1 []error
	}{
		{
			name: "Valid Grule Table",
			table: model.DTableData{
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
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: model.GRL,
					}},
				}},
			},
			want:  true,
			want1: nil,
		},
		{
			name: "Wrong hit policy for the NotationStandard",
			table: model.DTableData{
				"test1",
				"TableOne",
				model.Any,
				"",
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
					Typ:   model.Integer,
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: model.GRL,
					}},
				}},
			},
			want:  false,
			want1: []error{ErrDTableHitPolicy},
		},
		{
			name: "Wrong CollectOperators DMN",
			table: model.DTableData{
				"K1",
				"N1",
				model.Collect,
				"XYZ",
				model.DMN,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.String,
				},
				},
				[]model.Field{{
					Name:  "O1",
					Label: "L1",
					Typ:   model.Integer,
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: model.FEEL,
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: model.Juel,
					}},
				}},
			},
			want:  false,
			want1: []error{ErrDTableCollectOperator},
		},
		{
			name: "Empty input and output fields",
			table: model.DTableData{
				"K1",
				"N1",
				model.First,
				"",
				model.GRULE,
				[]model.Field{},
				[]model.Field{},
				[]model.Rule{{
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
			table: model.DTableData{
				"K1",
				"N1",
				model.First,
				"",
				model.GRULE,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.Double,
				},
				},
				[]model.Field{{
					Name:  "",
					Label: "L1",
					Typ:   model.Integer,
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: model.GRL,
					}},
				}},
			},
			want:  false,
			want1: []error{ErrDTableFieldTypInvalid, ErrDTableFieldNameEmpty},
		},
		{
			name: "Wrong rule data schemas",
			table: model.DTableData{
				"K1",
				"N1",
				model.First,
				"",
				model.GRULE,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.String,
				},
				},
				[]model.Field{{
					Name:  "O2",
					Label: "L1",
					Typ:   model.Integer,
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}, {
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}},
					OutputEntries: []model.Entry{},
				}},
			},
			want:  false,
			want1: []error{ErrRuleHaveDifferentAmountOfInputFields, ErrRuleHaveDifferentAmountOfOutputFields},
		},
		{
			name: "Wrong rule expression language",
			table: model.DTableData{
				"K1",
				"N1",
				model.First,
				"",
				model.GRULE,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   model.String,
				},
				},
				[]model.Field{{
					Name:  "O2",
					Label: "L1",
					Typ:   model.Integer,
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: model.FEEL,
					}},
					OutputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}},
				}},
			},
			want:  false,
			want1: []error{ErrDTableEntryExpressionLangInvalid},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DTableValidator{
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
