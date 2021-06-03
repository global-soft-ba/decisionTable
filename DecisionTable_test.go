package main

import (
	"decisionTable/conv/grule"
	"decisionTable/convert/interfaces"
	"decisionTable/data"
	"reflect"
	"testing"
)

func TestDecisionTable(t *testing.T) {
	testTable, err := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(data.GRULE).
		SetHitPolicy(data.Unique).
		AddInputField("Type of claim", "claim", data.String).
		AddInputField("Expenditure of claim", "claim", data.Integer).
		AddOutputField("Responsible employee", "Employee", data.String).
		AddOutputField("4 eyes principle", "Employee", data.Boolean).
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
		AddInputEntry(`"Car Accident"`, data.SFEEL).
		AddInputEntry(">=10000", data.SFEEL).
		AddOutputEntry(`"Schmidt"`, data.SFEEL).
		AddOutputEntry("true", data.SFEEL).BuildRule().
		Build()

	if got := testTable.valid; got != true {
		t.Errorf("Example Decision Table is invalid, got = %v, want true, with error %v", got, err)
	}
}

func TestDecisionTable_Convert(t *testing.T) {

	testTable, _ := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(data.GRULE).
		SetHitPolicy(data.Unique).
		AddInputField("TypeOfClaim", "claim", data.String).
		AddInputField("ExpenditureOfClaim", "claim", data.Integer).
		AddOutputField("ResponsibleEmployee", "Employee", data.String).
		AddOutputField("4EyesPrinciple", "Employee", data.Boolean).
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

	type args struct {
		converter interfaces.ConverterInterface
	}
	tests := []struct {
		name    string
		fields  DecisionTable
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:   "DecisionTable To GruleRuleSet",
			fields: testTable,
			args:   args{grule.CreateDTableToGrlConverter()},
			want: []string{
				"rule row_0 \"R1\" salience 0  {\n when \n   claim.TypeOfClaim == \"Car Accident\"\n   && claim.ExpenditureOfClaim < 1000 \n then \n  Employee.ResponsibleEmployee = \"Müller\";\n  Employee.4EyesPrinciple = false; \n  Complete();\n}",
				"rule row_1 \"R2\" salience 1  {\n when \n   claim.TypeOfClaim == \"Car Accident\"\n   && ((claim.ExpenditureOfClaim >= 1000) && (claim.ExpenditureOfClaim <= 10000)) \n then \n  Employee.ResponsibleEmployee = \"Meier\";\n  Employee.4EyesPrinciple = false; \n  Complete();\n}",
				"rule row_2 \"R3\" salience 2  {\n when \n   claim.ExpenditureOfClaim >= 10000 \n then \n  Employee.4EyesPrinciple = true; \n  Complete();\n}",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields

			got, err := d.Convert(tt.args.converter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
