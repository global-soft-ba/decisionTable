package grule

import (
	grlmodel2 "decisionTable/conv/grule/data"
	"decisionTable/data"
	"reflect"
	"testing"
)

func TestConverter_converting(t *testing.T) {
	type args struct {
		ruleSet grlmodel2.RuleSet
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Valid Convert Multiple Expressions",
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
		{name: "Valid Convert Single Expressions",
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
		{name: "Valid Convert Multiple Expressions and Assignments",
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
		{name: "Valid Convert Multiple Rules, Expressions and Assignments with FirstPolicy",
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
		{name: "Valid Convert Multiple Rules, Expressions and Assignments with PriorityPolicy",
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
		{name: "Valid Convert Multiple Rules, Expressions and Assignments with wrong Policy",
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
		{name: "Valid Convert Single Float Expressions",
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
		{name: "Valid Convert Single Boolean Expressions",
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
		{name: "Valid Convert Single String Expressions",
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
		{name: "Valid Convert DatTime String Expressions",
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
