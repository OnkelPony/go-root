// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 6:
//    Implementace jednoduchého rozhraní nazvaného OpenShape.

package main

import (
	"fmt"
	"math"
)

type OpenShape interface {
	length() float64
}

type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length(shape OpenShape) float64 {
	return shape.length()
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)

	line_length := length(line1)
	fmt.Println(line_length)
}
