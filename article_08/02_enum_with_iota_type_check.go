// Seriál "Programovací jazyk Go"
//
// Osmá část
//
// Demonstrační příklad číslo 2:
//    Náhrada za enum: běžné konstanty a identifikátor iota.

package main

import "fmt"

type Enum int

const (
	Pondeli Enum = iota
	Utery
	Streda
	Ctvrtek
	Patek
	Sobota
	Nedele
)

func reservation(day int) {
	fmt.Printf("Reservation for %d\n", day)
}

func main() {
	reservation(Pondeli)
	reservation(Sobota)
	reservation(Nedele)

	reservation(3)

	day := Pondeli
	day--
	reservation(day)
}
