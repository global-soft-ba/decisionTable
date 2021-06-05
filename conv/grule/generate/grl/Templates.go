package grl

/*
	rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}

*/

const RULE = `rule {{ template "RULENAME" . }} {{template "SALIENCE" . }} {
 when {{ template "WHEN" .}} 
 then {{ template "THEN" .}} 
 {{ template "INTERFERENCE" }}
}`

const RULENAME = `{{define "RULENAME"}}row_{{ .Field }} "{{ .Description }}"{{end}}`

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

// HitPolicies
/* We assume that the table has unique non overlapping conditions. So maximal hit conditions can be one (or nothing).
   We cannot check the uniqueness. In case that the table is not unique, multi hits are solved with a priority hit policy
   (highest priority wins).
*/
/* Would be correct for real unique => No Salience needed */
/* const UNIQUE = `{{define "SALIENCE"}} {{end}}` */

// UNIQUE Case with fallback in case of overlapping expressions
const (
	HitPolicyDefault  = `{{define "SALIENCE"}}{{end}}`
	HitPolicyUnique   = `{{define "SALIENCE"}}salience {{.Salience}} {{end}}`
	FIRSTHitPolicy    = `{{define "SALIENCE"}}salience {{.InvSalience}}{{end}}`
	PRIORITYHitPolicy = `{{define "SALIENCE"}}salience {{.Salience}}{{end}}`
)

const (
	INTERFERENCE    = `{{define "INTERFERENCE"}}{{end}}`
	NONINTERFERENCE = `{{define "INTERFERENCE"}} Complete();{{end}}`
)

/*

func buildPolicyTemplate(hitPolicy data.HitPolicy) string {
	switch hitPolicy {
	case data.Unique:
		return grl.UNIQUE
	case data.First:
		return grl.FIRST
	case data.Priority:
		return grl.PRIORITY
	default:
		return grl.DEFAULT
	}
}

func buildInterferenceTemplate(interference bool) string {
	switch interference {
	case true:
		return grl.INTERFERENCE
	default:
		return grl.NONINTERFERENCE
	}
}

*/
