package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"reflect"
	"testing"
)

func TestDecisionTable_ConvertWithExpressionOutput(t *testing.T) {
	testTable, _ := NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Employee").
		SetHitPolicy(hitPolicy.Unique).
		SetExpressionLanguage(expressionLanguage.SFEEL).
		SetStandard(standard.GRULE).
		AddInputField(field.Field{Name: "claim.TypeOfClaim", Type: dataType.String}).
		AddInputField(field.Field{Name: "claim.ExpenditureOfClaim", Type: dataType.Integer}).
		AddOutputField(field.Field{Name: "Employee.ResponsibleEmployee", Type: dataType.String}).
		AddOutputField(field.Field{Name: "Employee.4EyesPrinciple", Type: dataType.Boolean}).
		AddOutputField(field.Field{Name: "Score.score", Type: dataType.Integer}).
		AddRule(NewRuleBuilder().
			SetAnnotation("R1").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("<1000").
			AddOutputEntry(`"Müller"`).
			AddOutputEntry("false").
			AddOutputEntry("1+2").
			Build(),
		).
		AddRule(NewRuleBuilder().
			SetAnnotation("R2").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("[1000..10000]").
			AddOutputEntry(`"Meier"`).
			AddOutputEntry("false").
			AddOutputEntry("1-2").
			Build(),
		).
		AddRule(NewRuleBuilder().
			SetAnnotation("R3").
			AddInputEntry("-").
			AddInputEntry(">=10000").
			AddOutputEntry("-").
			AddOutputEntry("true").
			AddOutputEntry("1**2").
			Build(),
		).
		Build()

	tests := []struct {
		name    string
		fields  DecisionTable
		want    []string
		wantErr bool
	}{
		{
			name:   "DecisionTable To GruleRuleSet",
			fields: testTable,
			want: []string{
				"rule row_0 \"R1\" salience 0 {\n when \n\t(claim.TypeOfClaim == \"Car Accident\")\n\t&& (claim.ExpenditureOfClaim < 1000)\n then \n\tEmployee.ResponsibleEmployee = \"Müller\";\n\tEmployee.4EyesPrinciple = false;\n\tScore.score = (1 + 2);\n Complete();\n}",
				"rule row_1 \"R2\" salience 1 {\n when \n\t(claim.TypeOfClaim == \"Car Accident\")\n\t&& ((claim.ExpenditureOfClaim >= 1000) && (claim.ExpenditureOfClaim <= 10000))\n then \n\tEmployee.ResponsibleEmployee = \"Meier\";\n\tEmployee.4EyesPrinciple = false;\n\tScore.score = (1 - 2);\n Complete();\n}",
				"rule row_2 \"R3\" salience 2 {\n when \n\t(claim.ExpenditureOfClaim >= 10000)\n then \n\tEmployee.4EyesPrinciple = true;\n\tScore.score = Pow(1, 2);\n Complete();\n}",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields

			got, err := d.Convert(standard.GRULE)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToGrlAst() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGrlAst() got = %v, want = %v", got, tt.want)
			}
		})
	}

}

func TestDecisionTable_Convert(t *testing.T) {

	testTable, _ := NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Employee").
		SetHitPolicy(hitPolicy.Unique).
		SetExpressionLanguage(expressionLanguage.SFEEL).
		SetStandard(standard.GRULE).
		AddInputField(field.Field{Name: "claim.TypeOfClaim", Type: dataType.String}).
		AddInputField(field.Field{Name: "claim.ExpenditureOfClaim", Type: dataType.Integer}).
		AddOutputField(field.Field{Name: "Employee.ResponsibleEmployee", Type: dataType.String}).
		AddOutputField(field.Field{Name: "Employee.4EyesPrinciple", Type: dataType.Boolean}).
		AddRule(NewRuleBuilder().
			SetAnnotation("R1").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("< 1000").
			AddOutputEntry(`"Müller"`).
			AddOutputEntry("false").
			Build(),
		).
		AddRule(NewRuleBuilder().
			SetAnnotation("R2").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("[1000..10000]").
			AddOutputEntry(`"Meier"`).
			AddOutputEntry("false").
			Build(),
		).
		AddRule(NewRuleBuilder().
			SetAnnotation("R3").
			AddInputEntry("-").
			AddInputEntry(">=10000").
			AddOutputEntry("-").
			AddOutputEntry("true").
			Build(),
		).
		Build()

	tests := []struct {
		name    string
		fields  DecisionTable
		want    []string
		wantErr bool
	}{
		{
			name:   "DecisionTable To GruleRuleSet",
			fields: testTable,
			want: []string{
				"rule row_0 \"R1\" salience 0 {\n when \n\t(claim.TypeOfClaim == \"Car Accident\")\n\t&& (claim.ExpenditureOfClaim < 1000)\n then \n\tEmployee.ResponsibleEmployee = \"Müller\";\n\tEmployee.4EyesPrinciple = false;\n Complete();\n}",
				"rule row_1 \"R2\" salience 1 {\n when \n\t(claim.TypeOfClaim == \"Car Accident\")\n\t&& ((claim.ExpenditureOfClaim >= 1000) && (claim.ExpenditureOfClaim <= 10000))\n then \n\tEmployee.ResponsibleEmployee = \"Meier\";\n\tEmployee.4EyesPrinciple = false;\n Complete();\n}",
				"rule row_2 \"R3\" salience 2 {\n when \n\t(claim.ExpenditureOfClaim >= 10000)\n then \n\tEmployee.4EyesPrinciple = true;\n Complete();\n}",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields

			got, err := d.Convert(standard.GRULE)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToGrlAst() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGrlAst() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDecisionTable_ConvertQualifiedName(t *testing.T) {

	testTable, _ := NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Name").
		SetHitPolicy(hitPolicy.Unique).
		SetExpressionLanguage(expressionLanguage.SFEEL).
		SetStandard(standard.GRULE).
		AddInputField(field.Field{Name: "name.lastname", Type: dataType.String}).
		AddInputField(field.Field{Name: "name.surname", Type: dataType.String}).
		AddOutputField(field.Field{Name: "Name.equal", Type: dataType.Boolean}).
		AddRule(NewRuleBuilder().
			SetAnnotation("R1").
			AddInputEntry(`name.surname`).
			AddInputEntry(`"xyz"`).
			AddOutputEntry(`true`).
			Build(),
		).
		Build()

	tests := []struct {
		name    string
		fields  DecisionTable
		want    []string
		wantErr bool
	}{
		{
			name:   "DecisionTable To GruleRuleSet",
			fields: testTable,
			want: []string{
				"rule row_0 \"R1\" salience 0 {\n when \n\t(name.lastname == name.surname)\n\t&& (name.surname == \"xyz\")\n then \n\tName.equal = true;\n Complete();\n}",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields

			got, err := d.Convert(standard.GRULE)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToGrlAst() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGrlAst() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
