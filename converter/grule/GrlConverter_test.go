package grule

import (
	"decisionTable/converter/grule/grlmodel"
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
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{{"O1", "L1", model.Integer, "=4"}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  complete();"},
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
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
						},
						[]grlmodel.Expression{{"O1", "L1", model.Integer, "=4"}},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0\nwhen \n   L1.I1 ==3\nthen \n  L1.O1 =4;\n  complete();"},
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
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
				},
			}},
			want:    []string{"rule row_0 \"R1\" salience 0\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();"},
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
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
				},
			}},
			want: []string{"rule row_0 \"R1\" salience 2\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();",
				"rule row_1 \"R2\" salience 1\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();",
				"rule row_2 \"R3\" salience 0\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();"},
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
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
				},
			}},
			want: []string{"rule row_0 \"R1\" salience 0\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();",
				"rule row_1 \"R2\" salience 1\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();",
				"rule row_2 \"R3\" salience 2\nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();"},
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
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
					{
						"1",
						"R2",
						1,
						1,
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
					{
						"2",
						"R3",
						2,
						0,
						[]grlmodel.Expression{
							{"I1", "L1", model.Integer, "==3"},
							{"I2", "L1", model.Integer, ">3"},
							{"I3", "L1", model.Integer, ">4"},
						},
						[]grlmodel.Expression{
							{"O1", "L1", model.Integer, "=4"},
							{"O1", "L1", model.Integer, "=4"},
						},
					},
				},
			}},
			want: []string{"rule row_0 \"R1\" \nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();",
				"rule row_1 \"R2\" \nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();",
				"rule row_2 \"R3\" \nwhen \n   L1.I1 ==3\n   && L1.I2 >3\n   && L1.I3 >4\nthen \n  L1.O1 =4;\n  L1.O1 =4;\n  complete();"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := GrlConverter{}
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
