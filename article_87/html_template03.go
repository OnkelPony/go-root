package main

import (
	"os"
	"text/template"
)

const (
	templateFilename = "html_template03.htm"
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.ParseFiles(templateFilename))

	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 0},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "<script>alert('you have been pwned')</script>", 8},
		Role{"Eda", "Wasserfall", 0},
		Role{"Přemysl", "Hájek", 10},
	}

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, roles)
	if err != nil {
		panic(err)
	}
}