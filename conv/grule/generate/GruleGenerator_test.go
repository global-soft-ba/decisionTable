package generate

import (
	grule "decisionTable/conv/grule/data"
	"decisionTable/conv/grule/grl"
	dtable "decisionTable/data"
	"decisionTable/lang/sfeel"
	"reflect"
	"testing"
)

func CreateTestExpression(field dtable.FieldInterface, entry dtable.EntryInterface) grule.ExpressionInterface {
	res, _ := grl.CreateExpression(field, entry)
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
				targetFormat: string(grule.GRL),
				rules: grule.RuleSet{
					Key:             "test1",
					Name:            "TableOne",
					HitPolicy:       dtable.First,
					CollectOperator: dtable.List,
					Rules: []grule.Rule{
						{
							Name:        "0",
							Description: "R1",
							Expressions: []grule.Term{
								{
									Field:              dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String},
									Expression:         CreateTestExpression(dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String}, sfeel.CreateInputEntry("8")),
									ExpressionLanguage: dtable.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float},
									Expression:         CreateTestExpression(dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float}, sfeel.CreateOutputEntry("4")),
									ExpressionLanguage: dtable.SFEEL,
								},
							},
						},
					},
				},
			},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n    (I1.L1 == 8) \n then \n   (O1.L1 == 4) ;\n  Complete(); \n}"},
			wantErr: false,
		},
		{
			name: "integer unary with first hit policies",
			args: args{
				targetFormat: string(grule.GRL),
				rules: grule.RuleSet{
					Key:             "test1",
					Name:            "TableOne",
					HitPolicy:       dtable.First,
					Interference:    true,
					CollectOperator: dtable.List,
					Rules: []grule.Rule{
						{
							Name:        "0",
							Description: "R1",
							Salience:    0,
							InvSalience: 1,
							Expressions: []grule.Term{
								{
									Field:              dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String},
									Expression:         CreateTestExpression(dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String}, sfeel.CreateInputEntry("8")),
									ExpressionLanguage: dtable.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float},
									Expression:         CreateTestExpression(dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float}, sfeel.CreateOutputEntry("4")),
									ExpressionLanguage: dtable.SFEEL,
								},
							},
						},
						{
							Name:        "1",
							Description: "R2",
							Salience:    1,
							InvSalience: 0,
							Expressions: []grule.Term{
								{
									Field:              dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String},
									Expression:         CreateTestExpression(dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String}, sfeel.CreateInputEntry("10")),
									ExpressionLanguage: dtable.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float},
									Expression:         CreateTestExpression(dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float}, sfeel.CreateOutputEntry("100")),
									ExpressionLanguage: dtable.SFEEL,
								},
							},
						},
					},
				},
			},
			want: []string{
				"rule row_0 \"R1\" salience 1 {\n when \n    (I1.L1 == 8) \n then \n   (O1.L1 == 4) ;\n \n}",
				"rule row_1 \"R2\" salience 0 {\n when \n    (I1.L1 == 10) \n then \n   (O1.L1 == 100) ;\n \n}",
			},
			wantErr: false,
		},
		{
			name: "integer interval test",
			args: args{
				targetFormat: string(grule.GRL),
				rules: grule.RuleSet{
					Key:             "test1",
					Name:            "TableOne",
					HitPolicy:       dtable.Priority,
					CollectOperator: dtable.List,
					Rules: []grule.Rule{
						{
							Name:        "0",
							Description: "R1",
							Expressions: []grule.Term{
								{
									Field:              dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String},
									Expression:         CreateTestExpression(dtable.TestField{Name: "I1", Key: "L1", Typ: dtable.String}, sfeel.CreateInputEntry("[1..6)")),
									ExpressionLanguage: dtable.SFEEL},
							},
							Assignments: []grule.Term{
								{
									Field:              dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float},
									Expression:         CreateTestExpression(dtable.TestField{Name: "O1", Key: "L1", Typ: dtable.Float}, sfeel.CreateOutputEntry("4")),
									ExpressionLanguage: dtable.SFEEL,
								},
							},
						},
					},
				},
			},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n    ((I1.L1 >= 1) && (I1.L1 < 6)) \n then \n   (O1.L1 == 4) ;\n  Complete(); \n}"},
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
