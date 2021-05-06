package templates

/*
	rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}

*/

const RULE = `rule {{ template "RULENAME" . }} {{template "SALIENCE" . }}
when {{ template "WHEN" .}}
then {{ template "THEN" .}}
{{ template "INTERFERENCE" }}`

const RULENAME = `{{define "RULENAME"}}row_{{ .Name }} "{{ .Description }}"{{end}}`

const WHEN = `{{define "WHEN" }}
{{- range $index, $val := .Expressions }}
  {{- if eq $index 0}}
   {{template "ENTRY" $val}}
   {{- else}}
   && {{template "ENTRY" $val}}
  {{- end}}
 {{- end}}
{{- end}}`

const THEN = `{{define "THEN"}}
 {{- range $index, $val := .Assignments }}
  {{template "ENTRY" $val}};
 {{- end}}
{{- end}}`

const ENTRY = `{{define "ENTRY"}}{{.Expression}}{{end}}`
