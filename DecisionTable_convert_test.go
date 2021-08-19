package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data"
	"reflect"
	"testing"
)

func TestDecisionTable_Convert(t *testing.T) {

	testTable, _ := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(data.GRULE).
		SetHitPolicy(data.Unique).
		AddInputField(data.TestField{Name: "claim", Key: "TypeOfClaim", Typ: data.String}).
		AddInputField(data.TestField{Name: "claim", Key: "ExpenditureOfClaim", Typ: data.Integer}).
		AddOutputField(data.TestField{Name: "Employee", Key: "ResponsibleEmployee", Typ: data.String}).
		AddOutputField(data.TestField{Name: "Employee", Key: "4EyesPrinciple", Typ: data.Boolean}).
		AddRule("R1").
		AddInputEntry(`"Car Accident"`, data.SFEEL).
		AddInputEntry("<1000", data.SFEEL).
		AddOutputEntry(`"Müller"`, data.SFEEL).
		AddOutputEntry("false", data.SFEEL).BuildRule().
		AddRule("R2").
		AddInputEntry(`"Car Accident"`, data.SFEEL).
		AddInputEntry("[1000..10000]", data.SFEEL).
		AddOutputEntry(`"Meier"`, data.SFEEL).
		AddOutputEntry("false", data.SFEEL).BuildRule().
		AddRule("R3").
		AddInputEntry("-", data.SFEEL).
		AddInputEntry(">=10000", data.SFEEL).
		AddOutputEntry("-", data.SFEEL).
		AddOutputEntry("true", data.SFEEL).BuildRule().
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

			got, err := d.Convert(string(data.GRULE))
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
