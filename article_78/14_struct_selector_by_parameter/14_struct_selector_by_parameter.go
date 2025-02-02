// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá osmá část:
//    Jazyk Go a vyhodnocování výrazů v době běhu aplikace
//    https://www.root.cz/clanky/jazyk-go-a-vyhodnocovani-vyrazu-v-dobe-behu-aplikace/

package main

import (
	"fmt"
	"os"

	"github.com/PaesslerAG/gval"
)

// User je uživatelsky definovaná datová struktura s viditelnými atributy
type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	user := User{
		1,
		"Pepek",
		"Vyskoč"}

	// parametry předávané vyhodnocovanému výrazu
	parameters := map[string]interface{}{
		"selector1": "Name",
		"selector2": "Surname",
		"arr":       []int{10, 20, 30},
		"user":      user,
	}

	// vyhodnocení výrazu
	value, err := gval.Evaluate(`"Jméno:    " + user[selector1] + "\nPříjmení: " + user[selector2]`, parameters)
	if err != nil {
		// kód pro zpracování chyby při vyhodnocování výrazu
		fmt.Println(err)
		os.Exit(1)
	}

	// výpis výsledku výrazu
	fmt.Print(value)
}
