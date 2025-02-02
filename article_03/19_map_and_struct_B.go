// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 19:
//    Mapa struktur.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/19_map_and_struct_B.html

package main

import "fmt"

// Key je uživatelsky definovaná datová struktura
type Key struct {
	id   uint32
	role string
}

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[Key]User)
	fmt.Println(m1)

	m1[Key{1, "admin"}] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	m1[Key{2, "user"}] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	m1[Key{3, "boss"}] = User{
		id:      3,
		name:    "Pepi",
		surname: "Wüẞkotsch",
	}

	fmt.Println(m1)
}
