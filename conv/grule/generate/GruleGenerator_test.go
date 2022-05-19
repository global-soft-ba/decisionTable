package generate

import (
	grule "github.com/global-soft-ba/decisionTable/conv/grule/data"
	"github.com/global-soft-ba/decisionTable/conv/grule/grl"
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/entryType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
	"reflect"
	"testing"
)

func CreateTestExpression(field field.Field, expressionLanguage expressionLanguage.ExpressionLanguage, entryType entryType.EntryType, entry string) grule.ExpressionInterface {
	res, _ := grl.CreateExpression(field, expressionLanguage, entryType, entry)
	return res
}

func TestGruleGenerator_Generate(t *testing.T) {
	type args struct {
		rules        grule.RuleSet
		targetFormat string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "integer unary test",
			args: args{
				targetFormat: string(standard.GRULE),
				rules: grule.RuleSet{
					Key:             "test1",
					Name:            "TableOne",
					HitPolicy:       hitPolicy.First,
					CollectOperator: collectOperator.List,
					Rules: []grule.Rule{
						{
							Name:       "0",
							Annotation: "R1",
							Expressions: []grule.Term{
								{
									Field:              field.Field{Name: "I1.L1", Type: dataType.String},
									Expression:         CreateTestExpression(field.Field{Name: "I1.L1", Type: dataType.String}, expressionLanguage.SFEEL, entryType.Input, "8"),
									ExpressionLanguage: expressionLanguage.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              field.Field{Name: "O1.L1", Type: dataType.Float},
									Expression:         CreateTestExpression(field.Field{Name: "O1.L1", Type: dataType.Float}, expressionLanguage.SFEEL, entryType.Output, "4"),
									ExpressionLanguage: expressionLanguage.SFEEL,
								},
							},
						},
					},
				},
			},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n\t(I1.L1 == 8)\n then \n\tO1.L1 = 4;\n Complete();\n}"},
			wantErr: false,
		},
		{
			name: "integer unary with first hit policies",
			args: args{
				targetFormat: string(standard.GRULE),
				rules: grule.RuleSet{
					Key:             "test1",
					Name:            "TableOne",
					HitPolicy:       hitPolicy.First,
					Interference:    true,
					CollectOperator: collectOperator.List,
					Rules: []grule.Rule{
						{
							Name:        "0",
							Annotation:  "R1",
							Salience:    0,
							InvSalience: 1,
							Expressions: []grule.Term{
								{
									Field:              field.Field{Name: "I1.L1", Type: dataType.String},
									Expression:         CreateTestExpression(field.Field{Name: "I1.L1", Type: dataType.String}, expressionLanguage.SFEEL, entryType.Input, "8"),
									ExpressionLanguage: expressionLanguage.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              field.Field{Name: "O1.L1", Type: dataType.Float},
									Expression:         CreateTestExpression(field.Field{Name: "O1.L1", Type: dataType.Float}, expressionLanguage.SFEEL, entryType.Output, "4"),
									ExpressionLanguage: expressionLanguage.SFEEL,
								},
							},
						},
						{
							Name:        "1",
							Annotation:  "R2",
							Salience:    1,
							InvSalience: 0,
							Expressions: []grule.Term{
								{
									Field:              field.Field{Name: "I1.L1", Type: dataType.String},
									Expression:         CreateTestExpression(field.Field{Name: "I1.L1", Type: dataType.String}, expressionLanguage.SFEEL, entryType.Input, "10"),
									ExpressionLanguage: expressionLanguage.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              field.Field{Name: "O1.L1", Type: dataType.Float},
									Expression:         CreateTestExpression(field.Field{Name: "O1.L1", Type: dataType.Float}, expressionLanguage.SFEEL, entryType.Output, "100"),
									ExpressionLanguage: expressionLanguage.SFEEL,
								},
							},
						},
					},
				},
			},
			want: []string{
				"rule row_0 \"R1\" salience 1 {\n when \n\t(I1.L1 == 8)\n then \n\tO1.L1 = 4;\n \n}",
				"rule row_1 \"R2\" salience 0 {\n when \n\t(I1.L1 == 10)\n then \n\tO1.L1 = 100;\n \n}",
			},
			wantErr: false,
		},
		{
			name: "integer interval test",
			args: args{
				targetFormat: string(standard.GRULE),
				rules: grule.RuleSet{
					Key:             "test1",
					Name:            "TableOne",
					HitPolicy:       hitPolicy.Priority,
					CollectOperator: collectOperator.List,
					Rules: []grule.Rule{
						{
							Name:       "0",
							Annotation: "R1",
							Expressions: []grule.Term{
								{
									Field:              field.Field{Name: "I1.L1", Type: dataType.String},
									Expression:         CreateTestExpression(field.Field{Name: "I1.L1", Type: dataType.String}, expressionLanguage.SFEEL, entryType.Input, "[1..6)"),
									ExpressionLanguage: expressionLanguage.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              field.Field{Name: "O1.L1", Type: dataType.Float},
									Expression:         CreateTestExpression(field.Field{Name: "O1.L1", Type: dataType.Float}, expressionLanguage.SFEEL, entryType.Output, "4"),
									ExpressionLanguage: expressionLanguage.SFEEL,
								},
							},
						},
					},
				},
			},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n\t((I1.L1 >= 1) && (I1.L1 < 6))\n then \n\tO1.L1 = 4;\n Complete();\n}"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateGruleGenerator()
			got, err := g.Generate(tt.args.rules, tt.args.targetFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

/* MIgrate test cases

{name: "Validate ConvertToGrlAst Multiple Expressions",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{{"O1", "L1", data.Integer, "4", data.SFEEL}},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Single Expressions",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
					},
					[]grlmodel2.Term{{"O1", "L1", data.Integer, "4", data.SFEEL}},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3 \n then \n  L1.O1 = 4; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Multiple Expressions and Assignments",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Multiple Rules, Expressions and Assignments with FirstPolicy",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					2,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
				{
					"1",
					"R2",
					1,
					1,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
				{
					"2",
					"R3",
					2,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
			},
		}},
		want: []string{"rule row_0 \"R1\" salience 2 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
			"rule row_1 \"R2\" salience 1 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
			"rule row_2 \"R3\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Multiple Rules, Expressions and Assignments with PriorityPolicy",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.Priority,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					2,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "==3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "=4", data.SFEEL},
						{"O1", "L1", data.Integer, "=4", data.SFEEL},
					},
				},
				{
					"1",
					"R2",
					1,
					1,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "==3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "=4", data.SFEEL},
						{"O1", "L1", data.Integer, "=4", data.SFEEL},
					},
				},
				{
					"2",
					"R3",
					2,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "==3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "=4", data.SFEEL},
						{"O1", "L1", data.Integer, "=4", data.SFEEL},
					},
				},
			},
		}},
		want: []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
			"rule row_1 \"R2\" salience 1 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
			"rule row_2 \"R3\" salience 2 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Multiple Rules, Expressions and Assignments with wrong Policy",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.Any,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					2,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
				{
					"1",
					"R2",
					1,
					1,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, "> 3", data.SFEEL},
						{"I3", "L1", data.Integer, "> 4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
				{
					"2",
					"R3",
					2,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Integer, "3", data.SFEEL},
						{"I2", "L1", data.Integer, ">3", data.SFEEL},
						{"I3", "L1", data.Integer, ">4", data.SFEEL},
					},
					[]grlmodel2.Term{
						{"O1", "L1", data.Integer, "4", data.SFEEL},
						{"O1", "L1", data.Integer, "4", data.SFEEL},
					},
				},
			},
		}},
		want: []string{"rule row_0 \"R1\"  {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
			"rule row_1 \"R2\"  {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
			"rule row_2 \"R3\"  {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Single Float Expressions",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Float, "3.3", data.SFEEL},
					},
					[]grlmodel2.Term{{"O1", "L1", data.String, `"4"`, data.SFEEL}},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3.3 \n then \n  L1.O1 = \"4\"; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Single Boolean Expressions",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.Boolean, "true", data.SFEEL},
					},
					[]grlmodel2.Term{{"O1", "L1", data.Boolean, "false", data.SFEEL}},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == true \n then \n  L1.O1 = false; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst Single String Expressions",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.String, `"true"`, data.SFEEL},
					},
					[]grlmodel2.Term{{"O1", "L1", data.String, `"false"`, data.SFEEL}},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == \"true\" \n then \n  L1.O1 = \"false\"; \n  Complete();\n}"},
		wantErr: false,
	},
	{name: "Validate ConvertToGrlAst DatTime String Expressions",
		args: args{grlmodel2.RuleSet{
			Key:             "test1",
			Name:            "TableOne",
			HitPolicy:       data.First,
			CollectOperator: data.List,
			Interference:    false,
			Rules: []grlmodel2.Rule{
				{
					"0",
					"R1",
					0,
					0,
					[]grlmodel2.Term{
						{"I1", "L1", data.DateTime, `DateAndTime("2021-01-01T12:00:00")`, data.SFEEL},
					},
					[]grlmodel2.Term{{"O1", "L1", data.DateTime, `DateAndTime("2021-01-01T13:00:00")`, data.SFEEL}},
				},
			},
		}},
		want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == MakeTime(2021,1,1,12,0,0) \n then \n  L1.O1 = MakeTime(2021,1,1,13,0,0); \n  Complete();\n}"},
		wantErr: false,
	},
*/
