package valid

import (
	"github.com/global-soft-ba/decisionTable/data"
	sfeel2 "github.com/global-soft-ba/decisionTable/lang/sfeel"
	"reflect"
	"testing"
)

func TestDTableValidator_SFeel_Validate(t *testing.T) {
	tests := []struct {
		name  string
		table data.Table
		want  bool
		want1 []error
	}{
		{
			name: "Valid Grule Table",
			table: data.Table{
				Key:              "test1",
				Name:             "TableOne",
				HitPolicy:        data.First,
				CollectOperator:  data.List,
				NotationStandard: data.GRULE,
				InputFields: []data.FieldInterface{data.TestField{
					Name: "I1",
					Key:  "L1",
					Typ:  data.Integer,
				},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "O1",
					Key:  "L1",
					Typ:  data.Float,
				}},
				Rules: []data.Rule{
					{
						Description: "R1",
						InputEntries: []data.EntryInterface{
							sfeel2.CreateInputEntry("<3"),
						},
						OutputEntries: []data.EntryInterface{
							sfeel2.CreateOutputEntry("-"),
						},
					},
					{
						Description: "R2",
						InputEntries: []data.EntryInterface{
							sfeel2.CreateInputEntry("<=3"),
						},
						OutputEntries: []data.EntryInterface{
							sfeel2.CreateOutputEntry("1.2"),
						},
					},
					{
						Description: "R3",
						InputEntries: []data.EntryInterface{
							sfeel2.CreateInputEntry("not(47)"),
						},
						OutputEntries: []data.EntryInterface{
							sfeel2.CreateOutputEntry("1"),
						},
					},
				},
			},
			want:  true,
			want1: nil,
		},
		{
			name: "Wrong hit policy for the NotationStandard",
			table: data.Table{
				Key:              "test1",
				Name:             "TableOne",
				HitPolicy:        data.Any,
				NotationStandard: data.GRULE,
				InputFields: []data.FieldInterface{data.TestField{
					Name: "I1",
					Key:  "L1",
					Typ:  data.String,
				},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "O1",
					Key:  "L1",
					Typ:  data.Integer,
				}},
				Rules: []data.Rule{{
					Description: "R1",
					InputEntries: []data.EntryInterface{
						sfeel2.CreateInputEntry(`"3"`)},
					OutputEntries: []data.EntryInterface{
						sfeel2.CreateOutputEntry("4"),
					},
				}},
			},
			want:  false,
			want1: []error{ErrDTableHitPolicy},
		},
		{
			name: "Wrong CollectOperators DMN",
			table: data.Table{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        data.Collect,
				CollectOperator:  "XYZ",
				NotationStandard: data.DMN,
				InputFields: []data.FieldInterface{data.TestField{
					Name: "I1",
					Key:  "L1",
					Typ:  data.String,
				},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "O1",
					Key:  "L1",
					Typ:  data.Integer,
				}},
				Rules: []data.Rule{{
					Description: "R1",
					InputEntries: []data.EntryInterface{
						sfeel2.CreateInputEntry(`"3"`)},
					OutputEntries: []data.EntryInterface{
						sfeel2.CreateOutputEntry("4")},
				}},
			},
			want:  false,
			want1: []error{ErrDTableCollectOperator, ErrDTableEntryExpressionLangInvalid, ErrDTableEntryExpressionLangInvalid},
		},
		{
			name: "Empty input and output fields",
			table: data.Table{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        data.First,
				NotationStandard: data.GRULE,
				InputFields:      []data.FieldInterface{},
				OutputFields:     []data.FieldInterface{},
				Rules: []data.Rule{{
					Description:   "R1",
					InputEntries:  []data.EntryInterface{},
					OutputEntries: []data.EntryInterface{},
				}},
			},
			want:  false,
			want1: []error{ErrDTableInputEmpty, ErrDTableOutputEmpty},
		},
		{
			name: "Wrong TestField definition",
			table: data.Table{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        data.First,
				NotationStandard: data.GRULE,
				InputFields: []data.FieldInterface{data.TestField{
					Name: "I1",
					Key:  "L1",
					Typ:  data.Double,
				},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "",
					Key:  "L1",
					Typ:  data.Integer,
				}},
				Rules: []data.Rule{{
					Description: "R1",
					InputEntries: []data.EntryInterface{
						sfeel2.CreateInputEntry(">3")},
					OutputEntries: []data.EntryInterface{
						sfeel2.CreateOutputEntry("3"),
					},
				}},
			},
			want:  false,
			want1: []error{ErrDTableFieldTypInvalid, ErrDTableFieldIdIsEmpty},
		},
		{
			name: "Wrong mount of input entries regarding data schemas",
			table: data.Table{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        data.First,
				NotationStandard: data.GRULE,
				InputFields: []data.FieldInterface{
					data.TestField{
						Name: "I1",
						Key:  "L1",
						Typ:  data.String,
					},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "O2",
					Key:  "L1",
					Typ:  data.Integer,
				}},
				Rules: []data.Rule{{
					Description: "R1",
					InputEntries: []data.EntryInterface{
						sfeel2.CreateInputEntry(`"String"`),
						sfeel2.CreateInputEntry(">3")},
					OutputEntries: []data.EntryInterface{
						sfeel2.CreateOutputEntry("-"),
					},
				}},
			},
			want:  false,
			want1: []error{ErrRuleHaveDifferentAmountOfInputFields},
		},
		{
			name: "Wrong mount of output entries regarding data schemas",
			table: data.Table{
				Key:              "K1",
				Name:             "N1",
				HitPolicy:        data.First,
				NotationStandard: data.GRULE,
				InputFields: []data.FieldInterface{
					data.TestField{
						Name: "I1",
						Key:  "L1",
						Typ:  data.String,
					},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "O2",
					Key:  "L1",
					Typ:  data.Integer,
				}},
				Rules: []data.Rule{{
					Description: "R1",
					InputEntries: []data.EntryInterface{
						sfeel2.CreateInputEntry(`"String"`),
					},
					OutputEntries: []data.EntryInterface{
						sfeel2.CreateOutputEntry("-"),
						sfeel2.CreateOutputEntry("-"),
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
		field data.Table
		want  bool
	}{
		{
			name: "Valid Table Without Interferences",
			field: data.Table{
				Key:              "test1",
				Name:             "TableOne",
				HitPolicy:        data.First,
				CollectOperator:  data.List,
				NotationStandard: data.GRULE,
				InputFields: []data.FieldInterface{data.TestField{
					Name: "I1",
					Key:  "L1",
					Typ:  data.String,
				},
				},
				OutputFields: []data.FieldInterface{data.TestField{
					Name: "O1",
					Key:  "L1",
					Typ:  data.Float,
				}},
				Rules: []data.Rule{{
					Description: "R1",
					InputEntries: []data.EntryInterface{
						sfeel2.CreateInputEntry(" =3")},
					OutputEntries: []data.EntryInterface{
						sfeel2.CreateOutputEntry("4")},
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
