// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá pátá část
//    Použití databáze Redis v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/pouziti-databaze-redis-v-aplikacich-naprogramovanych-v-go/

package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

func main() {
	// vytvoření nového klienta s předáním konfiguračních parametrů
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// neměli bychom zapomenout na ukončení práce s klientem
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	// pouze zobrazíme textovou podobu struktury Client
	fmt.Println("Redis client:", client)
}
