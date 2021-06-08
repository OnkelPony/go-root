// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá druhá část
//    Pomůcky při tvorbě jednotkových testů v jazyce Go
//    https://www.root.cz/clanky/pomucky-pri-tvorbe-jednotkovych-testu-v-jazyce-go/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func CaptureLog(function func()) (string, error) {
	// backup of the real stdout
	logOutput := os.Stdout

	// temporary replacement for log output
	reader, writer, err := os.Pipe()
	if err != nil {
		return "", err
	}

	// temporarily replace real log output by the mocked one
	defer func() {
		log.SetOutput(logOutput)
	}()
	log.SetOutput(writer)

	// channel with captured standard output
	captured := make(chan string)

	// synchronization object
	wg := new(sync.WaitGroup)
	// we are going to wait for one goroutine only
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		// goroutine is started -> inform main one via WaitGroup object
		wg.Done()
		io.Copy(&buf, reader)
		captured <- buf.String()
	}()
	// wait for goroutine to start
	wg.Wait()
	// provided function that (probably) prints something to standard output
	function()
	writer.Close()
	return <-captured, nil
}

func main() {
	str, err := CaptureLog(func() { log.Println("Hello world") })
	if err != nil {
		panic(err)
	}

	fmt.Println("Captured logs:")
	fmt.Println("-------------------------------")
	fmt.Println(str)
	fmt.Println("-------------------------------")
}
