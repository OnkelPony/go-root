// Seriál "Programovací jazyk Go"
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Demonstrační příklad číslo 5:
//     Spuštění dvou gorutin

package main

import (
	"fmt"
	"time"
)

func printChars() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
		time.Sleep(200 * time.Millisecond)
	}
}

func print_dots() {
	for i := 0; i < 30; i++ {
		fmt.Print(".")
		time.Sleep(200 * time.Millisecond)
	}
}

func print_spaces() {
	for i := 0; i < 60; i++ {
		fmt.Print(" ")
		time.Sleep(110 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main begin")
	go printChars()
	go print_spaces()
	go print_dots()
	time.Sleep(6 * time.Second)
	fmt.Println("main end")
}
