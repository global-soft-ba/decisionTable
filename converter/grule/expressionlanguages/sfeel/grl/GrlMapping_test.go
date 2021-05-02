package grl

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func Test_test(t *testing.T) {
	var tmpl *template.Template

	tmpl, _ = template.New("Expressions").Parse(`{}`)
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, "")

	fmt.Println("Template=", tpl.String())

}
