// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 3:
//    Metody s parametry, předání struktury odkazem.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/03_methods_ptr.html

package main

import (
	"fmt"
	"math"
)

// Line je uživatelsky definovaná datová struktura
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

// Metoda definovaná pro strukturu Line
func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

// Metoda definovaná pro strukturu Line
func (line *Line) translate(dx, dy float64) {
	fmt.Printf("Translating line %v by %f %f\n", *line, dx, dy)
	(*line).x1 += dx
	(*line).y1 += dy
	(*line).x2 += dx
	(*line).y2 += dy
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)
	line1.translate(5, 5)
	fmt.Println(line1)

	lineLength := line1.length()
	fmt.Println(lineLength)
}
