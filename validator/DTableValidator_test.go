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
			table: model.DTableData{
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

func TestDTableValidator_ValidateInterferences(t *testing.T) {

	tests := []struct {
		name  string
		field model.DTableData
		want  bool
	}{
		{
			name: "Valid Table Without Interferences",
			field: model.DTableData{
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
						Expression:         " =3",
						ExpressionLanguage: model.GRL,
					}},
					OutputEntries: []model.Entry{{
						Expression:         "4",
						ExpressionLanguage: model.GRL,
					}},
				}},
			},
			want: false,
		},
		{
			name: "Valid Table Without Interferences",
			field: model.DTableData{
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
						Name:  "I2",
						Label: "L1",
						Typ:   model.String,
					}},
				Rules: []model.Rule{{
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
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DTableValidator{
				dTable: tt.field,
			}
			if got := d.ValidateInterferences(); got != tt.want {
				t.Errorf("ValidateInterferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
