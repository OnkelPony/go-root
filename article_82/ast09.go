package main

import (
	"fmt"
	"log"
	"strings"

	"go/ast"
	"go/parser"
)

// výraz, který se má naparsovat
const source = `
(1 + x) * (2 + y)
`

// nový datový typ implementující rozhraní ast.Visitor
type visitor int

// implementace (jediné) funkce předepsané v rozhraní ast.Visitor
func (v visitor) Visit(n ast.Node) ast.Visitor {
	// dosáhli jsme koncového uzlu?
	if n == nil {

		// tisk pozice a typu uzlu
		fmt.Printf("%3d\t", v)
		fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
		return nil
	}

	// tisk pozice a typu uzlu
	fmt.Printf("%3d\t", v)
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func main() {
	// konstrukce parseru a parsing zdrojového kódu
	f, err := parser.ParseExpr(source)
	if err != nil {
		log.Fatal(err)
	}

	var v visitor

	// zahájení průchodu abstraktním syntaktickým stromem
	ast.Walk(v, f)
}
