package main

import (
	"github.com/global-soft-ba/decisionTable/converters/grule/tableconverter/grl"
	"github.com/global-soft-ba/decisionTable/converters/interfaces"
	"github.com/global-soft-ba/decisionTable/model"
	"reflect"
	"testing"
)

func TestDecisionTable(t *testing.T) {
	testTable, err := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("Type of claim", "claim", model.String).
		AddInputField("Expenditure of claim", "claim", model.Integer).
		AddOutputField("Responsible employee", "Employee", model.String).
		AddOutputField("4 eyes principle", "Employee", model.Boolean).
		AddRule("R1").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("<1000", model.SFEEL).
		AddOutputEntry(`"Müller"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).BuildRule().
		AddRule("R2").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("[1000..10000]", model.SFEEL).
		AddOutputEntry(`"Meier"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).BuildRule().
		AddRule("R3").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry(">=10000", model.SFEEL).
		AddOutputEntry(`"Schmidt"`, model.SFEEL).
		AddOutputEntry("true", model.SFEEL).BuildRule().
		Build()

	if got := testTable.valid; got != true {
		t.Errorf("Example Decision Table is invalid, got = %v, want true, with error %v", got, err)
	}
}

func TestDecisionTable_Convert(t *testing.T) {

	testTable, _ := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("TypeOfClaim", "claim", model.String).
		AddInputField("ExpenditureOfClaim", "claim", model.Integer).
		AddOutputField("ResponsibleEmployee", "Employee", model.String).
		AddOutputField("4EyesPrinciple", "Employee", model.Boolean).
		AddRule("R1").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("<1000", model.SFEEL).
		AddOutputEntry(`"Müller"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).BuildRule().
		AddRule("R2").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("[1000..10000]", model.SFEEL).
		AddOutputEntry(`"Meier"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).BuildRule().
		AddRule("R3").
		AddInputEntry("-", model.SFEEL).
		AddInputEntry(">=10000", model.SFEEL).
		AddOutputEntry("-", model.SFEEL).
		AddOutputEntry("true", model.SFEEL).BuildRule().
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
			args:   args{grl.CreateDTableToGrlConverter()},
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
