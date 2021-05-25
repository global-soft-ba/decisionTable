package grl

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/model"
	"reflect"
	"testing"
)

func TestConverter_converting(t *testing.T) {
	type args struct {
		ruleSet grlmodel.RuleSet
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Valid Convert Multiple Expressions",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{{"O1", "L1", model.Integer, "4", model.SFEEL}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Single Expressions",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
						},
						[]grlmodel.Term{{"O1", "L1", model.Integer, "4", model.SFEEL}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3 \n then \n  L1.O1 = 4; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Multiple Expressions and Assignments",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Multiple Rules, Expressions and Assignments with FirstPolicy",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						2,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
				},
			}},
			want: []string{"rule row_0 \"R1\" salience 2 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
				"rule row_1 \"R2\" salience 1 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
				"rule row_2 \"R3\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Multiple Rules, Expressions and Assignments with PriorityPolicy",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.Priority,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						2,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "==3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "=4", model.SFEEL},
							{"O1", "L1", model.Integer, "=4", model.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "==3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "=4", model.SFEEL},
							{"O1", "L1", model.Integer, "=4", model.SFEEL},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "==3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "=4", model.SFEEL},
							{"O1", "L1", model.Integer, "=4", model.SFEEL},
						},
					},
				},
			}},
			want: []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
				"rule row_1 \"R2\" salience 1 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
				"rule row_2 \"R3\" salience 2 {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Multiple Rules, Expressions and Assignments with wrong Policy",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.Any,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						2,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, "> 3", model.SFEEL},
							{"I3", "L1", model.Integer, "> 4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Integer, "3", model.SFEEL},
							{"I2", "L1", model.Integer, ">3", model.SFEEL},
							{"I3", "L1", model.Integer, ">4", model.SFEEL},
						},
						[]grlmodel.Term{
							{"O1", "L1", model.Integer, "4", model.SFEEL},
							{"O1", "L1", model.Integer, "4", model.SFEEL},
						},
					},
				},
			}},
			want: []string{"rule row_0 \"R1\"  {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
				"rule row_1 \"R2\"  {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}",
				"rule row_2 \"R3\"  {\n when \n   L1.I1 == 3\n   && L1.I2 > 3\n   && L1.I3 > 4 \n then \n  L1.O1 = 4;\n  L1.O1 = 4; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Single Float Expressions",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Float, "3.3", model.SFEEL},
						},
						[]grlmodel.Term{{"O1", "L1", model.String, `"4"`, model.SFEEL}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == 3.3 \n then \n  L1.O1 = \"4\"; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Single Boolean Expressions",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.Boolean, "true", model.SFEEL},
						},
						[]grlmodel.Term{{"O1", "L1", model.Boolean, "false", model.SFEEL}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == true \n then \n  L1.O1 = false; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert Single String Expressions",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.String, `"true"`, model.SFEEL},
						},
						[]grlmodel.Term{{"O1", "L1", model.String, `"false"`, model.SFEEL}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == \"true\" \n then \n  L1.O1 = \"false\"; \n  Complete();\n}"},
			wantErr: false,
		},
		{name: "Valid Convert DatTime String Expressions",
			args: args{grlmodel.RuleSet{
				Key:             "test1",
				Name:            "TableOne",
				HitPolicy:       model.First,
				CollectOperator: model.List,
				Interference:    false,
				Rules: []grlmodel.Rule{
					{
						"0",
						"R1",
						0,
						0,
						[]grlmodel.Term{
							{"I1", "L1", model.DateTime, `DateAndTime("2021-01-01T12:00:00")`, model.SFEEL},
						},
						[]grlmodel.Term{{"O1", "L1", model.DateTime, `DateAndTime("2021-01-01T13:00:00")`, model.SFEEL}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0 {\n when \n   L1.I1 == MakeTime(2021,1,1,12,0,0) \n then \n  L1.O1 = MakeTime(2021,1,1,13,0,0); \n  Complete();\n}"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateDTableToGrlConverter()
			got, err := c.convertRuleSetIntoGRL(tt.args.ruleSet)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertRuleSetIntoGRL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertRuleSetIntoGRL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
