// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 8:
//     Implementace jednotkových testů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/test08/add_test.html

package main

import (
	"fmt"
	"testing"
)

type AddTest struct {
	x        int32
	y        int32
	expected int32
}

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}

func CheckAdd(t *testing.T, testInputs []AddTest) {
	for _, i := range testInputs {
		result := add(i.x, i.y)
		if result != i.expected {
			msg := fmt.Sprintf("%d + %d should be %d, got %d instead",
				i.x, i.y, i.expected, result)
			t.Error(msg)
		}
	}
}
