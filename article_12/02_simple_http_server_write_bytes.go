// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 2:
//    Alternativní způsob vytvoření sekvence bajtů z řetězce
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/02_simple_http_server_write_bytes.html

package main

import (
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	response := "Hello world (čau světe)!\n"
	writer.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
