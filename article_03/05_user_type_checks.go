// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 5:
//    Kontrola uživatelsky definovaných typů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/05_user_type_checks.html

package main

import "fmt"

// ID of user
type ID uint32

// Name of user
type Name string

// Surname of user
type Surname string

func main() {
	var i ID
	i = 0
	fmt.Println(i)

	var str = "common string"

	var n Name = str
	var s Surname = str

	n = "Jan"
	s = "Novák"

	fmt.Println(n)
	fmt.Println(s)
}
