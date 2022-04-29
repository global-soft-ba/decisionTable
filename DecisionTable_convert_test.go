package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
	"reflect"
	"testing"
)

func TestDecisionTable_ConvertWithExpressionOutput(t *testing.T) {
	testTable, _ := NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Employee").
		SetHitPolicy(data.Unique).
		SetExpressionLanguage(data.SFEEL).
		SetStandard(data.GRULE).
		AddInputField(data.TestField{Name: "claim", Key: "TypeOfClaim", Type: data.String}).
		AddInputField(data.TestField{Name: "claim", Key: "ExpenditureOfClaim", Type: data.Integer}).
		AddOutputField(data.TestField{Name: "Employee", Key: "ResponsibleEmployee", Type: data.String}).
		AddOutputField(data.TestField{Name: "Employee", Key: "4EyesPrinciple", Type: data.Boolean}).
		AddOutputField(data.TestField{Name: "Score", Key: "score", Type: data.Integer}).
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
		name   string
		fields DecisionTable

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

			got, err := d.Convert(data.GRULE)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToGrlAst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGrlAst() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestDecisionTable_Convert(t *testing.T) {

	testTable, _ := NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Employee").
		SetHitPolicy(data.Unique).
		SetExpressionLanguage(data.SFEEL).
		SetStandard(data.GRULE).
		AddInputField(data.TestField{Name: "claim", Key: "TypeOfClaim", Type: data.String}).
		AddInputField(data.TestField{Name: "claim", Key: "ExpenditureOfClaim", Type: data.Integer}).
		AddOutputField(data.TestField{Name: "Employee", Key: "ResponsibleEmployee", Type: data.String}).
		AddOutputField(data.TestField{Name: "Employee", Key: "4EyesPrinciple", Type: data.Boolean}).
		AddRule(NewRuleBuilder().
			SetAnnotation("R1").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("< claim.ExpenditureOfClaim").
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
		name   string
		fields DecisionTable

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

			got, err := d.Convert(data.GRULE)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToGrlAst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGrlAst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecisionTable_ConvertQualifiedName(t *testing.T) {

	testTable, _ := NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Name").
		SetHitPolicy(data.Unique).
		SetExpressionLanguage(data.SFEEL).
		SetStandard(data.GRULE).
		AddInputField(data.TestField{Name: "name", Key: "lastname", Type: data.String}).
		AddInputField(data.TestField{Name: "name", Key: "surname", Type: data.String}).
		AddOutputField(data.TestField{Name: "Name", Key: "equal", Type: data.Boolean}).
		AddRule(NewRuleBuilder().
			SetAnnotation("R1").
			AddInputEntry(`name.surname`).
			AddInputEntry(`"xyz"`).
			AddOutputEntry(`true`).
			Build(),
		).
		Build()

	tests := []struct {
		name   string
		fields DecisionTable

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

			got, err := d.Convert(data.GRULE)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToGrlAst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGrlAst() got = %v, want %v", got, tt.want)
			}
		})
	}
}
