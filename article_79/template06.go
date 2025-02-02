// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá devátá část:
//    Standardní šablonovací systém jazyka Go
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_79/template06.html

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
