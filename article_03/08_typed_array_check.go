// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 8:
//    Pole s prvky definovaného typu + kontrola typů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/08_typed_array_check.html

package main

import "fmt"

// Mesic je nový datový typ se jménem viditelným mimo balíček
type Mesic string

func main() {
	mesice := [12]Mesic{
		"Leden",
		"Únor",
		"Březen",
		"Duben",
		"Květen",
		"Červen",
		"Červenec",
		"Srpen",
		"Září",
		"Říjen",
		"Listopad",
		"Prosinec"}

	//var mesice2 [12]string = mesice
	var mesice2 [12]string
	for i, mesic := range mesice {
		mesice2[i] = string(mesic)
	}
	fmt.Println(mesice)
	fmt.Println(mesice2)
}
