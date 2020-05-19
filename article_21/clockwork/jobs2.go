// Seriál "Programovací jazyk Go"
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 2:
//     Úloha vytvořená nástrojem clockwork.

package main

import (
	"github.com/whiteShtef/clockwork"
)

func task() {
	println("task/job called")
}

func main() {
	scheduler := clockwork.NewScheduler()
	scheduler.Schedule().Every().Second().Do(task)
	scheduler.Run()
}
