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

func main() {
	// parametry předávané vyhodnocovanému výrazu
	parameters := map[string]interface{}{
		"x": 10,
		"y": 20,
	}

	// vyhodnocení výrazu
	value, err := gval.Evaluate("2*x + y", parameters)
	if err != nil {
		// kód pro zpracování chyby při vyhodnocování výrazu
		fmt.Println(err)
		os.Exit(1)
	}

	// výpis výsledku výrazu
	fmt.Print(value)
}
