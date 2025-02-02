// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmnáctá část
//     Knihovny určené pro tvorbu testů v programovacím jazyce Go
//     https://www.root.cz/clanky/knihovny-urcene-pro-tvorbu-testu-v-programovacim-jazyce-go/
//
// Demonstrační příklad číslo 5:
//     Testy pro balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_18/05_factorial_ogletest2/factorial_test.html

package factorial_test

import (
	"github.com/OnkelPony/go-root/article_18/05_factorial_ogletest2"
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"testing"
)

func TestOgletest(t *testing.T) {
	RunTests(t)
}

type FactorialTest struct{}

func init() {
	RegisterTestSuite(&FactorialTest{})
}

func (t *FactorialTest) FactorialForZero() {
	result := factorial.Factorial(0)
	ExpectThat(result, Equals(1))
}

func (t *FactorialTest) FactorialForOne() {
	result := factorial.Factorial(1)
	ExpectThat(result, Equals(1))
}

func (t *FactorialTest) TestFactorialSmallNumber() {
	result := factorial.Factorial(5)
	ExpectThat(result, AllOf(
		GreaterThan(10),
		LessThan(10000)))
}

func (t *FactorialTest) TestFactorialSmallNumberNegative() {
	result := factorial.Factorial(20)
	ExpectThat(result, AllOf(
		GreaterThan(10),
		LessThan(10000)))
}

func (t *FactorialTest) TestFactorialForTen() {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	ExpectThat(result, Equals(expected))
}

func (t *FactorialTest) TestFactorialBigNumber() {
	result := factorial.Factorial(20)
	ExpectThat(result, GreaterThan(0))
}

func (t *FactorialTest) TestFactorialEvenBiggerNumber() {
	result := factorial.Factorial(30)
	ExpectThat(result, GreaterThan(0))
}
