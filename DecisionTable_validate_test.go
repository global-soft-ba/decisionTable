package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_InvalidTables(t *testing.T) {
	tests := []struct {
		name          string
		decisionTable DecisionTable
		wantErr       bool
		errContains   string
	}{
		{
			name: "Happy Path create valid table",
			decisionTable: NewDecisionTableBuilder().
				SetID("test1").
				SetName("ABC").
				SetHitPolicy(data.First).
				SetCollectOperator(data.List).
				SetStandard(data.GRULE).
				SetExpressionLanguage(data.SFEEL).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "O1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
					AddInputEntry("<3").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: false,
		},
		{
			name: "Name is missing",
			decisionTable: NewDecisionTableBuilder().
				SetID("id").
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableNameIsRequired,
		},
		{
			name: "ID is missing",
			decisionTable: NewDecisionTableBuilder().
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableIdIsRequired,
		},
		{
			name: "HitPolicy not valid for standard",
			decisionTable: NewDecisionTableBuilder().
				SetHitPolicy(data.Any).
				SetStandard(data.GRULE).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableHitPolicyIsInvalid,
		},
		{
			name: "CollectOperator not valid for standard",
			decisionTable: NewDecisionTableBuilder().
				SetHitPolicy(data.Collect).
				SetStandard(data.DMN).
				SetCollectOperator("XYZ").
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableCollectOperatorIsInvalid,
		},
		{
			name: "InputField datatype does not match standard",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldTypeIsInvalid,
		},
		{
			name: "InputField is missing",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableInputFieldIsRequired,
		},
		{
			name: "InputField ID is missing ",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				AddInputField(data.TestField{Type: data.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldIdIsRequired,
		},
		{
			name: "OutputField ID is missing",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				AddOutputField(data.TestField{Type: data.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldIdIsRequired,
		},
		{
			name: "OutputField is missing",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableOutputFieldIsRequired,
		},
		{
			name: "OutputField datatype does not match standard",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldTypeIsInvalid,
		},
		{
			name: "Amount of rules does not match amount of input fields",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Float}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Float}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("-").
					AddInputEntry("-").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrInputCountMismatch,
		},
		{
			name: "Amount of rules does not match amount of output fields",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(data.DMN).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Float}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Float}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("-").
					AddOutputEntry("-").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrOutputCountMismatch,
		},
		{
			name: "Input rule entry contains wrong expression",
			decisionTable: NewDecisionTableBuilder().
				SetID("1").
				SetName("Test").
				SetHitPolicy(data.First).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("xyz").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: true,
		},
		{
			name: "Output rule entry contains wrong expression",
			decisionTable: NewDecisionTableBuilder().
				SetID("1").
				SetName("Test").
				SetHitPolicy(data.First).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("-").
					AddOutputEntry("x.s").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.decisionTable.Validate(tt.decisionTable.Standard())
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() got = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == true && !assert.Contains(t, err.Error(), tt.errContains) {
				t.Errorf("validate() got = %v, want nil", err)
			}
		})
	}
}

func TestValidate_InvalidTableEntries(t *testing.T) {
	tests := []struct {
		name          string
		decisionTable DecisionTable
		wantErr       bool
	}{
		{
			name: "Input rule entry contains wrong expression",
			decisionTable: NewDecisionTableBuilder().
				SetID("1").
				SetName("Test").
				SetHitPolicy(data.First).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("xyz").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: true,
		},
		{
			name: "Input rule entry contains wrong data type expression",
			decisionTable: NewDecisionTableBuilder().
				SetID("1").
				SetName("Test").
				SetHitPolicy(data.First).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("\"ABC\"").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: true,
		},
		{
			// TODO: Output Evaluation is not completed, the test case must be adapted/checked to provoke "ErrDTableEntryReferencedFieldTypInvalid"
			name: "Output rule entry contains wrong entry reference field",
			decisionTable: NewDecisionTableBuilder().
				SetID("1").
				SetName("Test").
				SetHitPolicy(data.First).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry(">2").
					AddOutputEntry("I2.L1").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.decisionTable.Validate(tt.decisionTable.Standard())
			if err == nil && tt.wantErr == true {
				t.Errorf("validate() got = nil, want error{}")
			}
			if err != nil && tt.wantErr == false {
				t.Errorf("validate() got = %v, want nil", err)
			}
		})
	}
}

func TestValidate_CheckForInterferences(t *testing.T) {
	tests := []struct {
		name          string
		decisionTable DecisionTable
		want          bool
	}{
		{
			name: "valid2 table without interferences",
			decisionTable: NewDecisionTableBuilder().
				SetID("test1").
				SetName("ABC").
				SetHitPolicy(data.First).
				SetCollectOperator(data.List).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I2", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
					AddInputEntry("<3").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			want: false,
		},
		{
			name: "valid2 table with interferences",
			decisionTable: NewDecisionTableBuilder().
				SetID("test1").
				SetName("ABC").
				SetHitPolicy(data.First).
				SetCollectOperator(data.List).
				SetStandard(data.GRULE).
				AddInputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddOutputField(data.TestField{Name: "I1", Key: "L1", Type: data.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
					AddInputEntry("<3").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.decisionTable.CheckIfContainsInterferences()
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("CheckIfContainsInterferences() got = %v, want %v", result, tt.want)
			}
		})
	}
}
