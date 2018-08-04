package main

import (
	"os"
	"text/template"
)

const ExampleTmpl = `
{{define "A"}}AAAAAAAAAAAAAAAA {{end}}
{{define "B"}}BBBBBBBBBBBBBBBB {{end}}
{{define "C"}}CCCCCCCCCCCCCCCC {{end}}

{{define "Main"}}
{{template "A"}}
{{template "B"}}
{{template "C"}}
{{end}}
`

func main() {
	template.Must(template.New("Example").
		Parse(ExampleTmpl)).
		ExecuteTemplate(os.Stdout, "Main", nil)
}
