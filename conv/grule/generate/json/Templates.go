package json

import (
	dTable "github.com/global-soft-ba/decisionTable/data"
	"text/template"
)

func GenerateTemplates(hitPolicy dTable.HitPolicy, interference bool) (*template.Template, error) {
	/*
		t, err := template.New(RULE).Funcs(
				template.FuncMap{
					"getFormat": func() data.OutputFormat {
						return data.JSON
					},
				},
			).Parse(RULE)
	*/
	panic("implement me")
}
