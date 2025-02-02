// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 2:
//    Příkaz return.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/02_return_statement.html

package main

func f1() {
	println("f1")
	return
}

func main() {
	println("Hello world!")
	f1()
}
