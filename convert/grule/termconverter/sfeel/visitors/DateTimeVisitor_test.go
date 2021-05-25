package visitors

import (
	"decisionTable/convert/grule/grlmodel"
	"decisionTable/convert/grule/termconverter/sfeel/mapper"
	"decisionTable/model"
	"decisionTable/parser/sfeel/parser"
	"reflect"
	"testing"
)

func TestDateTimeVisitor_DateTimeInputRules(t *testing.T) {
	type fields struct {
		term grlmodel.Term
		maps mapper.TermMapper
	}

	mapping := mapper.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"equal datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `DateAndTime("2021-02-01T13:01:00")`,
				},
				mapping,
			},
			"credit.time == MakeTime(2021,2,1,13,1,0)"},
		{"datetime comparison LESS",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `<DateAndTime("2021-01-01T12:00:00")`,
				},
				mapping,
			},
			"credit.time < MakeTime(2021,1,1,12,0,0)"},
		{"integer comparison LESSEQ",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `<=DateAndTime("2021-01-01T12:00:00")`,
				},
				mapping,
			},
			"credit.time <= MakeTime(2021,1,1,12,0,0)"},
		{"integer comparison GR",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `>DateAndTime("2021-01-01T12:00:00")`,
				},
				mapping,
			},
			"credit.time > MakeTime(2021,1,1,12,0,0)"},
		{"datetime comparison GREQ",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `>=DateAndTime("2021-01-01T12:00:00")`,
				},
				mapping,
			},
			"credit.time >= MakeTime(2021,1,1,12,0,0)"},
		{"empty datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: "-",
				},
				mapping,
			},
			""},
		{"range IN IN datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `[DateAndTime("2021-01-01T12:00:00")..DateAndTime("2021-02-01T12:00:00")]`,
				},
				mapping,
			},
			"((credit.time >= MakeTime(2021,1,1,12,0,0)) && (credit.time <= MakeTime(2021,2,1,12,0,0)))"},
		{"range IN OUT datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `[DateAndTime("2021-01-01T12:00:00")..DateAndTime("2021-02-01T12:00:00")[`,
				},
				mapping,
			},
			`((credit.time >= MakeTime(2021,1,1,12,0,0)) && (credit.time < MakeTime(2021,2,1,12,0,0)))`},
		{"range OUT OUT datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `]DateAndTime("2021-01-01T12:00:00")..DateAndTime("2021-02-01T12:00:00")[`,
				},
				mapping,
			},
			`((credit.time > MakeTime(2021,1,1,12,0,0)) && (credit.time < MakeTime(2021,2,1,12,0,0)))`},
		{"range OUT IN datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `]DateAndTime("2021-01-01T12:00:00")..DateAndTime("2021-02-01T12:00:00")]`,
				},
				mapping,
			},
			`((credit.time > MakeTime(2021,1,1,12,0,0)) && (credit.time <= MakeTime(2021,2,1,12,0,0)))`},
		{"disjunctions datetime input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `DateAndTime("2021-01-01T12:00:00"),DateAndTime("2021-02-01T12:00:00"),DateAndTime("2021-03-01T12:00:00")`,
				},
				mapping,
			},
			"((credit.time == MakeTime(2021,1,1,12,0,0)) || (credit.time == MakeTime(2021,2,1,12,0,0)) || (credit.time == MakeTime(2021,3,1,12,0,0)))"},
		{"Negation integer input",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `not(DateAndTime("2021-01-01T12:00:00"))`,
				},
				mapping,
			},
			"!(credit.time == MakeTime(2021,1,1,12,0,0))"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidDateTimeInput()
			vis := CreateDateTimeVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in IntegerInputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTimeVisitor_DateTimeOutputRules(t *testing.T) {
	type fields struct {
		term grlmodel.Term
		maps mapper.TermMapper
	}

	mapping := mapper.SettingsGRL

	tests := []struct {
		name string
		args fields
		want string
	}{
		{"datetime assignment",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: `DateAndTime("2021-01-01T12:00:00")`,
				},
				mapping,
			},
			"credit.time = MakeTime(2021,1,1,12,0,0)"},
		{"datetime empty assignment",
			fields{
				grlmodel.Term{
					Name:       "time",
					Key:        "credit",
					Typ:        model.DateTime,
					Expression: "-",
				},
				mapping,
			},
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.CreateSfeelParser(tt.args.term.Expression)
			tree := prs.Parse().ValidDateTimeOutput()
			vis := CreateDateTimeVisitor(tt.args.term, tt.args.maps)

			if got := tree.Accept(vis).(string); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error in StringOutputRule() => %v, want %v", got, tt.want)
			}
		})
	}
}
