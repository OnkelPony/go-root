// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/factorial_convey/behaviour_test.html
//

package factorial

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFactorial(t *testing.T) {
	Convey("0! should be equal 1", t, func() {
		So(Factorial(0), ShouldEqual, 1)
	})
}

func TestFactorial2(t *testing.T) {
	Convey("10! should be greater than 1", t, func() {
		So(Factorial(10), ShouldBeGreaterThan, 1)
	})
	Convey("10! should be between 1 and 10000000", t, func() {
		So(Factorial(10), ShouldBeBetween, 1, 10000000)
	})
}
