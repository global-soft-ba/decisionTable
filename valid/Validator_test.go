package valid

import (
	"decisionTable/model"
	"decisionTable/sfeel"
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
					Name: "I1",
					Key:  "L1",
					Typ:  model.Integer,
				},
				},
				OutputFields: []model.Field{{
					Name: "O1",
					Key:  "L1",
					Typ:  model.Float,
				}},
				Rules: []model.Rule{
					{
						Description: "R1",
						InputEntries: []model.EntryInterface{
							sfeel.CreateInputEntry("<3"),
						},
						OutputEntries: []model.EntryInterface{
							sfeel.CreateOutputEntry("-"),
						},
					},
					{
						Description: "R2",
						InputEntries: []model.EntryInterface{
							sfeel.CreateInputEntry("<=3"),
						},
						OutputEntries: []model.EntryInterface{
							sfeel.CreateOutputEntry("1.2"),
						},
					},
					{
						Description: "R3",
						InputEntries: []model.EntryInterface{
							sfeel.CreateInputEntry("not(47)"),
						},
						OutputEntries: []model.EntryInterface{
							sfeel.CreateOutputEntry("1"),
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
					Name: "I1",
					Key:  "L1",
					Typ:  model.String,
				},
				},
				OutputFields: []model.Field{{
					Name: "O1",
					Key:  "L1",
					Typ:  model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.EntryInterface{
						sfeel.CreateInputEntry(`"3"`)},
					OutputEntries: []model.EntryInterface{
						sfeel.CreateOutputEntry("4"),
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
					Name: "I1",
					Key:  "L1",
					Typ:  model.String,
				},
				},
				OutputFields: []model.Field{{
					Name: "O1",
					Key:  "L1",
					Typ:  model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.EntryInterface{
						sfeel.CreateInputEntry(`"3"`)},
					OutputEntries: []model.EntryInterface{
						sfeel.CreateOutputEntry("4")},
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
					InputEntries:  []model.EntryInterface{},
					OutputEntries: []model.EntryInterface{},
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
					Name: "I1",
					Key:  "L1",
					Typ:  model.Double,
				},
				},
				OutputFields: []model.Field{{
					Name: "",
					Key:  "L1",
					Typ:  model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.EntryInterface{
						sfeel.CreateInputEntry(">3")},
					OutputEntries: []model.EntryInterface{
						sfeel.CreateOutputEntry("3"),
					},
				}},
			},
			want:  false,
			want1: []error{ErrDTableFieldTypInvalid, ErrDTableFieldNameEmpty},
		},
		{
			name: "Wrong mount of input entries regarding data schemas",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.First,
				NotationStandard: model.GRULE,
				InputFields: []model.Field{
					{
						Name: "I1",
						Key:  "L1",
						Typ:  model.String,
					},
				},
				OutputFields: []model.Field{{
					Name: "O2",
					Key:  "L1",
					Typ:  model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.EntryInterface{
						sfeel.CreateInputEntry(`"String"`),
						sfeel.CreateInputEntry(">3")},
					OutputEntries: []model.EntryInterface{
						sfeel.CreateOutputEntry("-"),
					},
				}},
			},
			want:  false,
			want1: []error{ErrRuleHaveDifferentAmountOfInputFields},
		},
		{
			name: "Wrong mount of output entries regarding data schemas",
			table: model.TableData{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        model.First,
				NotationStandard: model.GRULE,
				InputFields: []model.Field{
					{
						Name: "I1",
						Key:  "L1",
						Typ:  model.String,
					},
				},
				OutputFields: []model.Field{{
					Name: "O2",
					Key:  "L1",
					Typ:  model.Integer,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.EntryInterface{
						sfeel.CreateInputEntry(`"String"`),
					},
					OutputEntries: []model.EntryInterface{
						sfeel.CreateOutputEntry("-"),
						sfeel.CreateOutputEntry("-"),
					},
				}},
			},
			want:  false,
			want1: []error{ErrRuleHaveDifferentAmountOfOutputFields},
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
					Name: "I1",
					Key:  "L1",
					Typ:  model.String,
				},
				},
				OutputFields: []model.Field{{
					Name: "O1",
					Key:  "L1",
					Typ:  model.Float,
				}},
				Rules: []model.Rule{{
					Description: "R1",
					InputEntries: []model.EntryInterface{
						sfeel.CreateInputEntry(" =3")},
					OutputEntries: []model.EntryInterface{
						sfeel.CreateOutputEntry("4")},
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
