// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 2:
//     Implementace jednotkových testů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/test02/add_test.html

package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(108, 666)
	if result != 774 {
		t.Error("1+2 should be 774, got ", result, "instead")
	}
}
