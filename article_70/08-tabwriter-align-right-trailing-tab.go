// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá část
//    Vykreslení tabulek do terminálu v jazyce Go
//    https://www.root.cz/clanky/vykresleni-tabulek-do-terminalu-v-jazyce-go/

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const (
	TabWidth         = 0
	Padding          = 1
	PaddingCharacter = ' '
	Flags            = tabwriter.AlignRight
)

func main() {
	for minWidth := 0; minWidth < 10; minWidth++ {
		fmt.Printf("Min width = %d\n", minWidth)
		w := tabwriter.NewWriter(os.Stdout, minWidth, TabWidth, Padding, PaddingCharacter, Flags)
		fmt.Fprintln(w, "1\t1\t1\t")
		fmt.Fprintln(w, "22\t22\t22\t")
		fmt.Fprintln(w, "333\t333\t333\t")
		fmt.Fprintln(w, "4444\t4444\t4444\t")
		w.Flush()
		fmt.Println()
	}
}
