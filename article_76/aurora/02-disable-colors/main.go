// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá šestá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani/

package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	var colorizer aurora.Aurora
	colorizer = aurora.NewAurora(false)

	fmt.Println(colorizer.Red("Test"))
	fmt.Println(colorizer.Green("Test"))
	fmt.Println(colorizer.Blue("Test"))
	fmt.Println(colorizer.Cyan("Test"))
	fmt.Println(colorizer.Magenta("Test"))
	fmt.Println(colorizer.Yellow("Test"))
}
