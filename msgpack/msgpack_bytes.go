package main

import (
	"log"
	"os"

	"github.com/ugorji/go/codec"
)

const filename = "/tmp/bytes.bin"

func main() {
	// vytvořit soubor s binárními daty
	fout, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	log.Print("Output file created")

	// handler
	var handler codec.MsgpackHandle

	// objekt realizující zakódování dat
	encoder := codec.NewEncoder(fout, &handler)

	log.Print("Encoder created")

	const N = 1000
	var values [N]byte

	for i := 0; i < N; i++ {
		values[i] = byte(i)
	}

	// zakódování dat
	err = encoder.Encode(values)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done")
}
