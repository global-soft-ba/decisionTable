package json

import (
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"text/template"
)

func GenerateTemplates(hp hitPolicy.HitPolicy, interference bool) (*template.Template, error) {
	/*
		t, err := template.New(RULE).Funcs(
				template.FuncMap{
					"getFormat": func() data.Standard {
						return data.JSON
					},
				},
			).Parse(RULE)
	*/
	panic("implement me")
}
