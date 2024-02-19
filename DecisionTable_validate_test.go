package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
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
			name: "Happy path - create valid table",
			decisionTable: NewDecisionTableBuilder().
				SetID("test1").
				SetName("Test 1").
				SetHitPolicy(hitPolicy.First).
				SetCollectOperator(collectOperator.List).
				SetExpressionLanguage(expressionLanguage.SFEEL).
				SetStandard(standard.GRULE).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddOutputField(field.Field{Name: "O1.L1", Type: dataType.Integer}).
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
			name:          "Name is missing",
			decisionTable: NewDecisionTableBuilder().BuildWithoutValidation(),
			wantErr:       true,
			errContains:   validator.ErrDecisionTableNameIsRequired,
		},
		{
			name:          "ID is missing",
			decisionTable: NewDecisionTableBuilder().BuildWithoutValidation(),
			wantErr:       true,
			errContains:   validator.ErrDecisionTableIdIsRequired,
		},
		{
			name: "HitPolicy not valid for standard",
			decisionTable: NewDecisionTableBuilder().
				SetHitPolicy(hitPolicy.Any).
				SetStandard(standard.GRULE).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableHitPolicyIsInvalid,
		},
		{
			name: "CollectOperator not valid for standard",
			decisionTable: NewDecisionTableBuilder().
				SetHitPolicy(hitPolicy.Collect).
				SetStandard(standard.DMN).
				SetCollectOperator("XYZ").
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableCollectOperatorIsInvalid,
		},
		{
			name:          "InputField is missing",
			decisionTable: NewDecisionTableBuilder().BuildWithoutValidation(),
			wantErr:       true,
			errContains:   validator.ErrDecisionTableInputFieldIsRequired,
		},
		{
			name:          "OutputField is missing",
			decisionTable: NewDecisionTableBuilder().BuildWithoutValidation(),
			wantErr:       true,
			errContains:   validator.ErrDecisionTableOutputFieldIsRequired,
		},
		{
			name: "InputField name is missing ",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(standard.DMN).
				AddInputField(field.Field{Type: dataType.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldNameIsRequired,
		},
		{
			name: "OutputField bname is missing",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(standard.DMN).
				AddOutputField(field.Field{Type: dataType.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldNameIsRequired,
		},
		{
			name: "InputField datatype does not match standard",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(standard.DMN).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldTypeIsInvalid,
		},
		{
			name: "OutputField datatype does not match standard",
			decisionTable: NewDecisionTableBuilder().
				SetStandard(standard.DMN).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Float}).
				BuildWithoutValidation(),
			wantErr:     true,
			errContains: validator.ErrDecisionTableFieldTypeIsInvalid,
		},
		{
			name: "Amount of rules does not match amount of input fields",
			decisionTable: NewDecisionTableBuilder().
				SetExpressionLanguage(expressionLanguage.SFEEL).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Float}).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Float}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
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
				SetExpressionLanguage(expressionLanguage.SFEEL).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Float}).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Float}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
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
				SetHitPolicy(hitPolicy.First).
				SetExpressionLanguage(expressionLanguage.SFEEL).
				SetStandard(standard.GRULE).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
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
				SetHitPolicy(hitPolicy.First).
				SetExpressionLanguage(expressionLanguage.SFEEL).
				SetStandard(standard.GRULE).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
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
				t.Errorf("validate() got = %v, wantErr = %v", err, tt.wantErr)
			}
			if tt.wantErr == true && !assert.Contains(t, err.Error(), tt.errContains) {
				t.Errorf("validate() got = %v, want = %s", err, tt.errContains)
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
				SetID("test1").
				SetName("Test 1").
				SetHitPolicy(hitPolicy.First).
				SetExpressionLanguage(expressionLanguage.SFEEL).
				SetStandard(standard.GRULE).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("Rule").
					AddInputEntry("xyz").
					AddOutputEntry("-").
					Build(),
				).
				BuildWithoutValidation(),
			wantErr: true,
		},
		//{  //TODO not implemented yet
		//	name: "Input rule entry contains wrong data type expression",
		//	decisionTable: NewDecisionTableBuilder().
		//		SetID("1").
		//		SetName("Test").
		//		SetHitPolicy(hitPolicy.First).
		//		SetExpressionLanguage(expressionLanguage.SFEEL).
		//		SetStandard(standard.GRULE).
		//		AddInputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
		//		AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
		//		AddRule(NewRuleBuilder().
		//			SetAnnotation("Rule").
		//			AddInputEntry("\"ABC\"").
		//			AddOutputEntry("-").
		//			Build(),
		//		).
		//		BuildWithoutValidation(),
		//	wantErr: true,
		//},
		{
			// TODO: Output evaluation is not completed, the test case must be adapted/checked to provoke "ErrDTableEntryReferencedFieldTypInvalid"
			name: "Output data entry contains wrong entry reference field",
			decisionTable: NewDecisionTableBuilder().
				SetID("1").
				SetName("Test").
				SetHitPolicy(hitPolicy.First).
				SetExpressionLanguage(expressionLanguage.SFEEL).
				SetStandard(standard.GRULE).
				AddInputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddOutputField(field.Field{Name: "I1.L1", Type: dataType.Integer}).
				AddRule(NewRuleBuilder().
					SetAnnotation("R1").
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
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() got = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
