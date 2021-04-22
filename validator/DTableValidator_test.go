package validator

import (
	"decisionTable/constant"
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
				"FIRST",
				"",
				constant.DecisionTableStandardGrule,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "STRING",
				},
				},
				[]model.Field{{
					Name:  "O1",
					Label: "L1",
					Typ:   "FLOAT",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "GRL",
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: "GRL",
					}},
				}},
			},
			want:  true,
			want1: nil,
		},
		{
			name: "Empty Fields",
			table: model.DTableData{
				"",
				"",
				"",
				"",
				"",
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "STRING",
				},
				},
				[]model.Field{{
					Name:  "O1",
					Label: "L1",
					Typ:   "FLOAT",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "GRL",
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: "GRL",
					}},
				}},
			},
			want: false,
			want1: []error{ErrDTableNameEmpty, ErrDTableKeyEmpty, ErrDTableStandardInvalid,
				ErrDTableHitPolicy, ErrDTableFieldTypInvalid, ErrDTableFieldTypInvalid,
				ErrDTableEntryExpressionLangInvalid, ErrDTableEntryExpressionLangInvalid},
		},
		{
			name: "Empty CollectOperators DMN",
			table: model.DTableData{
				"K1",
				"N1",
				constant.CollectOperatorPolicy,
				"",
				constant.DecisionTableStandardDMN,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "STRING",
				},
				},
				[]model.Field{{
					Name:  "O1",
					Label: "L1",
					Typ:   "INTEGER",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "FEEL",
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: "JUEL",
					}},
				}},
			},
			want:  false,
			want1: []error{ErrDTableEmptyCollectOperator},
		},
		{
			name: "Wrong CollectOperators DMN",
			table: model.DTableData{
				"K1",
				"N1",
				constant.CollectOperatorPolicy,
				"XYZ",
				constant.DecisionTableStandardDMN,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "STRING",
				},
				},
				[]model.Field{{
					Name:  "O1",
					Label: "L1",
					Typ:   "INTEGER",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "FEEL",
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: "JUEL",
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
				"FIRST",
				"",
				constant.DecisionTableStandardGrule,
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
				"FIRST",
				"",
				constant.DecisionTableStandardGrule,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "DOUBLE",
				},
				},
				[]model.Field{{
					Name:  "",
					Label: "L1",
					Typ:   "INTEGER",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "GRL",
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: "GRL",
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
				"FIRST",
				"",
				constant.DecisionTableStandardGrule,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "STRING",
				},
				},
				[]model.Field{{
					Name:  "O2",
					Label: "L1",
					Typ:   "INTEGER",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "GRL",
					}, {
						Expression:         " =3",
						ExpressionLanguage: "GRL",
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
				"FIRST",
				"",
				constant.DecisionTableStandardGrule,
				[]model.Field{{
					Name:  "I1",
					Label: "L1",
					Typ:   "STRING",
				},
				},
				[]model.Field{{
					Name:  "O2",
					Label: "L1",
					Typ:   "INTEGER",
				}},
				[]model.Rule{{
					Description: "R1",
					InputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "FEEL",
					}},
					OutputEntries: []model.Entry{{
						Expression:         " =3",
						ExpressionLanguage: "GRL",
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
