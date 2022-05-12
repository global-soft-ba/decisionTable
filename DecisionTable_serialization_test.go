package decisionTable

import (
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

var decisionTableJson = `{
	"id": "determineEmployee",
	"name": "Determine Employee",
	"hitPolicy": "unique",
	"collectOperator": "",
	"expressionLanguage": "sfeel",
	"standard": "grule",
	"inputFields": [
		{
			"name": "Claim.TypeOfClaim",
			"type": "string"
		},
		{
			"name": "Claim.ExpenditureOfClaim",
			"type": "integer"
		}
	],
	"outputFields": [
		{
			"name": "Employee.ResponsibleEmployee",
			"type": "string"
		},
		{
			"name": "Employee.FourEyesPrinciple",
			"type": "boolean"
		}
	],
	"rules": [
		{
			"annotation": "R1",
			"inputEntries": [
				"\"Car Accident\"",
				"\u003c1000"
			],
			"outputEntries": [
				"\"Müller\"",
				"false"
			]
		},
		{
			"annotation": "R2",
			"inputEntries": [
				"\"Car Accident\"",
				"[1000..10000]"
			],
			"outputEntries": [
				"\"Schulz\"",
				"false"
			]
		},
		{
			"annotation": "R3",
			"inputEntries": [
				"-",
				"\u003e=10000"
			],
			"outputEntries": [
				"-",
				"true"
			]
		}
	]
}`

var decisionTableBuilder = NewDecisionTableBuilder().
	SetID("determineEmployee").
	SetName("Determine Employee").
	SetHitPolicy(hitPolicy.Unique).
	SetExpressionLanguage(expressionLanguage.SFEEL).
	SetStandard(standard.GRULE).
	AddInputField(field.Field{Name: "Claim.TypeOfClaim", Type: dataType.String}).
	AddInputField(field.Field{Name: "Claim.ExpenditureOfClaim", Type: dataType.Integer}).
	AddOutputField(field.Field{Name: "Employee.ResponsibleEmployee", Type: dataType.String}).
	AddOutputField(field.Field{Name: "Employee.FourEyesPrinciple", Type: dataType.Boolean}).
	AddRule(NewRuleBuilder().
		SetAnnotation("R1").
		AddInputEntry(`"Car Accident"`).
		AddInputEntry("<1000").
		AddOutputEntry(`"Müller"`).
		AddOutputEntry("false").
		Build(),
	).
	AddRule(NewRuleBuilder().
		SetAnnotation("R2").
		AddInputEntry(`"Car Accident"`).
		AddInputEntry("[1000..10000]").
		AddOutputEntry(`"Schulz"`).
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
	)

func TestDecisionTable_Serialize(t *testing.T) {
	tests := []struct {
		name string
		dtb  DecisionTableBuilderInterface
		want string
	}{
		{
			name: "Serialization test",
			dtb:  decisionTableBuilder,
			want: decisionTableJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.dtb.Serialize()
			assert.Equalf(t, tt.want, got, "Serialize()")
		})
	}
}

func TestDecisionTable_Unserialize(t *testing.T) {
	tests := []struct {
		name string
		json string
		want DecisionTableBuilderInterface
	}{
		{
			name: "Unserialization test",
			json: decisionTableJson,
			want: decisionTableBuilder,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Unserialize(tt.json)
			assert.Equalf(t, tt.want, got, "Unserialize(%v)", tt.json)
		})
	}
}
