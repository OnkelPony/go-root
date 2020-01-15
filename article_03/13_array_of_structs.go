// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 13:
//    Pole struktur.

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users [3]User
	fmt.Println(users)

	user1 := User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	var user2 User
	user2.id = 1
	user2.name = "Pepek"
	user2.surname = "Vyskoč"

	user3 := User{
		id:      1,
		name:    "Josef",
		surname: "Vyskočil"}

	users[0] = user1
	users[1] = user2
	users[2] = user3
	fmt.Println(users)
}
