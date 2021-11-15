// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 3:
//     Server odpovídající klientovi opožděně
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/03_slow_server.html

package main

import (
	"net"
	"time"
)

func main() {
	cnt := byte(0)

	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		println("Can't open the port!")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		//start := time.Now()
		println("connection accepted at", time.Now().Second())
		time.Sleep(3 * time.Second)
		defer conn.Close()
		if err != nil {
			println("Connection refused!")
		} else {
			var b = []byte{cnt}
			cnt++
			conn.Write(b)
		}
		println("connection closed at", time.Now().Second())
	}
}
