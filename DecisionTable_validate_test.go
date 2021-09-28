package decisionTable

import (
	"errors"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/valid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_InvalidTables(t *testing.T) {
	tests := []struct {
		name  string
		table DecisionTable
		want  bool
		err   error
	}{
		{
			name: "Happy Path create valid2 table",
			table: CreateDecisionTable().
				SetDefinitionKey("test1").
				SetName("ABC").
				SetHitPolicy(data.First).
				SetCollectOperator(data.List).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "O1", Key: "L1", Typ: data.Integer}).
				AddRule("R1").
				AddInputEntry("<3", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: true,
			err:  error(nil),
		},

		{name: "table name is missing",
			table: CreateDecisionTable().
				SetDefinitionKey("Key").
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableNameEmpty,
		},
		{name: "table key is missing",
			table: CreateDecisionTable().
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableKeyEmpty,
		},
		{name: "HitPolicy not valid2 for notation standard",
			table: CreateDecisionTable().
				SetHitPolicy(data.Any).
				SetNotationStandard(data.GRULE).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableHitPolicy,
		},
		{name: "CollectOperator not valid2 for notation standard",
			table: CreateDecisionTable().
				SetHitPolicy(data.Collect).
				SetNotationStandard(data.DMN).
				SetCollectOperator("XYZ").
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableCollectOperator,
		},
		{name: "InputField datatype does not match notation standard",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Float}).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableFieldTypInvalid,
		},
		{name: "InputField is missing",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableInputEmpty,
		},
		{name: "InputField ID is missing ",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				AddInputField(data.TestField{Typ: data.Float}).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableFieldIdIsEmpty,
		},
		{name: "OutputField ID is missing",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				AddOutputField(data.TestField{Typ: data.Float}).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableFieldIdIsEmpty,
		},
		{name: "OutputField is  missing",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableOutputEmpty,
		},
		{name: "OutputField datatype does not match notation standard",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Float}).
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrDTableFieldTypInvalid,
		},
		{name: "Amount of rules does not match amount of input fields",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Float}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Float}).
				AddRule("Rule").
				AddInputEntry("-", data.SFEEL).
				AddInputEntry("-", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrRuleHaveDifferentAmountOfInputFields,
		},
		{name: "Amount of rules does not match amount of output fields",
			table: CreateDecisionTable().
				SetNotationStandard(data.DMN).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Float}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Float}).
				AddRule("Rule").
				AddInputEntry("-", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
			err:  valid.ErrRuleHaveDifferentAmountOfOutputFields,
		},
		{name: "Input rule entry contains wrong expression",
			table: CreateDecisionTable().
				SetName("Test").SetDefinitionKey("1").
				SetHitPolicy(data.First).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddRule("Rule").
				AddInputEntry("xyz", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
			err:  errors.New("couldn't find qualified name xyz in field list"),
		},
		{name: "Output rule entry contains wrong expression",
			table: CreateDecisionTable().
				SetName("Test").SetDefinitionKey("1").
				SetHitPolicy(data.First).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddRule("Rule").
				AddInputEntry("-", data.SFEEL).
				AddOutputEntry("x.s", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
			err:  errors.New("couldn't find qualified name x.s in field list"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.table.Validate()
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("validate() got = %v, want %v", result, tt.want)
			}
			if err != nil && !assert.Contains(t, err, tt.err) {
				t.Errorf("validate() got1 = %v, want %v", err, tt.err)
			}
		})
	}
}

func TestValidate_InvalidTableEntries(t *testing.T) {
	tests := []struct {
		name  string
		table DecisionTable
		want  bool
	}{

		{name: "Input rule entry contains wrong expression",
			table: CreateDecisionTable().
				SetName("Test").SetDefinitionKey("1").
				SetHitPolicy(data.First).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddRule("Rule").
				AddInputEntry("xyz", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
		},
		{name: "Input rule entry contains wrong data typ expression",
			table: CreateDecisionTable().
				SetName("Test").SetDefinitionKey("1").
				SetHitPolicy(data.First).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddRule("Rule").
				AddInputEntry("\"ABC\"", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
		}, //TODO - Output Evaluation is not completed, the test case must be adapted/checked to provoke "ErrDTableEntryReferencedFieldTypInvalid"
		{name: "Output rule entry contains wrong entry reference field",
			table: CreateDecisionTable().
				SetName("Test").SetDefinitionKey("1").
				SetHitPolicy(data.First).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddRule("Rule").
				AddInputEntry(">2", data.SFEEL).
				AddOutputEntry("I2.L1", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := tt.table.Validate()
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("validate() got = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestValidate_CheckForInterferences(t *testing.T) {
	tests := []struct {
		name  string
		table DecisionTable
		want  bool
	}{
		{
			name: "valid2 table without interferences",
			table: CreateDecisionTable().
				SetDefinitionKey("test1").
				SetName("ABC").
				SetHitPolicy(data.First).
				SetCollectOperator(data.List).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I2", Key: "L1", Typ: data.Integer}).
				AddRule("R1").
				AddInputEntry("<3", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: false,
		},
		{
			name: "valid2 table with interferences",
			table: CreateDecisionTable().
				SetDefinitionKey("test1").
				SetName("ABC").
				SetHitPolicy(data.First).
				SetCollectOperator(data.List).
				SetNotationStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Typ: data.Integer}).
				AddRule("R1").
				AddInputEntry("<3", data.SFEEL).
				AddOutputEntry("-", data.SFEEL).
				BuildRule().
				BuildWithoutValidation(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := tt.table.Validate()
			result = tt.table.Interferences()
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("valid2ateContainsInterferences() got = %v, want %v", result, tt.want)
			}
		})
	}
}
