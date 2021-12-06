package main

import (
	"os"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = "The {{.}} language is often referred to as Golang, but the proper name is '{{.}}'."
)

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New(templateName).Parse(templateFormat))

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, "Go")
	if err != nil {
		panic(err)
	}
}
