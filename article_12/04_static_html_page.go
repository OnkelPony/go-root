// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 4:
//     Server, který vrací statické HTML stránky načtené ze souborů
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/04_static_html_page.html

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadFile("index.html")
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func missingEndpoint(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadFile("missing.html")
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/missing", missingEndpoint)
	http.ListenAndServe(":8000", nil)
}
