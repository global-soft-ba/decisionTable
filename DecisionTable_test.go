package main

import (
	"decisionTable/converter"
	"decisionTable/converter/grule"
	"decisionTable/model"
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
		AddInputEntry("Car Accident", model.GRL).
		AddInputEntry("<1000", model.GRL).
		AddOutputEntry("Müller", model.GRL).
		AddOutputEntry("false", model.GRL).BuildRule().
		AddRule("R2").
		AddInputEntry("Car Accident", model.GRL).
		AddInputEntry("[1000..10000]", model.GRL).
		AddOutputEntry("Meier", model.GRL).
		AddOutputEntry("false", model.GRL).BuildRule().
		AddRule("R3").
		AddInputEntry("Car Accident", model.GRL).
		AddInputEntry(">=10000", model.GRL).
		AddOutputEntry("Schmidt", model.GRL).
		AddOutputEntry("true", model.GRL).BuildRule().
		Build()

	if got := testTable.valid; got != true {
		t.Errorf("Example Decision Table is invalid, got = %v, want true, with error %v", got, err)
	}
}

func TestDecisionTableEmptyEntry(t *testing.T) {
	testTable, err := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("Type of claim", "claim", model.String).
		AddOutputField("Responsible employee", "Employee", model.String).
		AddRule("R1").
		AddInputEntry("Car Accident", model.GRL).
		AddOutputEntry("-", model.GRL).
		BuildRule().
		Build()

	if got := testTable.Rules()[0].OutputEntries[0].EmptyExpression(); got != true {
		t.Errorf("EmptyEntry in decision table is invalid, got = %v, want = %v true, with error %v", got, true, err)
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
		AddInputEntry("==\"Car Accident\"", model.GRL).
		AddInputEntry("< 1000", model.GRL).
		AddOutputEntry("=Müller", model.GRL).
		AddOutputEntry("=false", model.GRL).BuildRule().
		AddRule("R2").
		AddInputEntry("==\"Car Accident\"", model.GRL).
		AddInputEntry("[1000..10000]", model.GRL).
		AddOutputEntry("=Meier", model.GRL).
		AddOutputEntry("=false", model.GRL).BuildRule().
		AddRule("R3").
		AddInputEntry("-", model.GRL).
		AddInputEntry(">= 10000", model.GRL).
		AddOutputEntry("-", model.GRL).
		AddOutputEntry("=true", model.GRL).BuildRule().
		Build()

	type args struct {
		converter converter.DTableConverterInterface
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
				"rule row_0 \"R1\" salience 0 \nwhen \n   claim.TypeOfClaim ==\"Car Accident\"\n   && claim.ExpenditureOfClaim < 1000\nthen \n  Employee.ResponsibleEmployee =Müller;\n  Employee.4EyesPrinciple =false;\n  complete();",
				"rule row_1 \"R2\" salience 1 \nwhen \n   claim.TypeOfClaim ==\"Car Accident\"\n   && claim.ExpenditureOfClaim [1000..10000]\nthen \n  Employee.ResponsibleEmployee =Meier;\n  Employee.4EyesPrinciple =false;\n  complete();",
				"rule row_2 \"R3\" salience 2 \nwhen \n   claim.ExpenditureOfClaim >= 10000\nthen \n  Employee.4EyesPrinciple =true;\n  complete();",
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
