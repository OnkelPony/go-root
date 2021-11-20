// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Přístup k tabulce kurzů a použití kurzů při výpočtu převodu měn.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/mocking/kurzy.html
//

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ExchangeDataGetter func(code string) float64

type ExchangeGetter struct {
	getExchangeRate ExchangeDataGetter
}

func NewExchangeGetter(g ExchangeDataGetter) *ExchangeGetter {
	return &ExchangeGetter{getExchangeRate: g}
}

func getExchangeRateFromUrl(code string) float64 {
	const URL = "https://www.cnb.cz/cs/financni_trhy/devizovy_trh/kurzy_devizoveho_trhu/denni_kurz.txt"

	response, err := http.Get(URL)
	if err != nil {
		panic("Connection refused")
	}
	defer response.Body.Close()

	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Content length: %d\n", response.ContentLength)

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		if len(s) == 5 {
			codeStr := s[3]
			rateStr := strings.Replace(s[4], ",", ".", 1)
			if code == codeStr {
				rateF, err := strconv.ParseFloat(rateStr, 64)
				if err != nil {
					panic(err)
				}
				return rateF
			}
		}
	}

	return 0
}

func getExchangeRateFromFile(code string) float64 {
	const FILENAME = "kurzy.txt"

	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		if len(s) == 5 {
			codeStr := s[3]
			rateStr := strings.Replace(s[4], ",", ".", 1)
			if code == codeStr {
				rateF, err := strconv.ParseFloat(rateStr, 64)
				if err != nil {
					panic(err)
				}
				return rateF
			}
		}
	}

	return 0
}

func mockedGetExchangeRate(code string) float64 {
	return 21.5
}

func (g *ExchangeGetter) exchange(amount float64, code string) float64 {
	rate := g.getExchangeRate(code)
	return rate * amount
}

func main() {
	g := NewExchangeGetter(getExchangeRateFromUrl)
	fmt.Printf("%5.2f\n", g.exchange(10, "USD"))
	g2 := NewExchangeGetter(mockedGetExchangeRate)
	fmt.Printf("%5.2f\n", g2.exchange(10, "USD"))
	g3 := NewExchangeGetter(getExchangeRateFromFile)
	fmt.Printf("%5.2f\n", g3.exchange(108, "ILS"))
}
