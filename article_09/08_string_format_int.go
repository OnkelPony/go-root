// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 8:
//    Základní funkce pro formátování řetězců: FormatInt.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/08_string_format_int.html

package main

import (
	"strconv"
	"strings"
)

func main() {
	value := int64(0xcafebabe)
	for base := 2; base < 36; base++ {
		println(base, strings.ToUpper(strconv.FormatInt(value, base)))
	}
}
