// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 15:
//     HTTP server vracející statický obsah
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/15_file_server.html

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	http.ListenAndServe(":8000", nil)
}
