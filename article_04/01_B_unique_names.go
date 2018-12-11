// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 1B:
//    Unikátní jména metod.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x1, y1 float64
}

type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length() float64 {
	return 0
}

func (point Point) length() float64 {
	return 0
}

func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)

	line_length := line1.length()
	fmt.Println(line_length)
}
