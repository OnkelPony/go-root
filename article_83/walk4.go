package main

import (
	"fmt"
	"log"

	"go/ast"
	"go/parser"
)

// výraz, který se má naparsovat
const source = `
1 + 2 * (3 + x) + y * (z - 1)
`

func walk(node ast.Node) {
	// dosáhli jsme koncového uzlu?
	if node == nil {
		return
	}

	// tisk hodnoty uzlu a rekurzivní průchod poduzly
	switch x := node.(type) {
	case *ast.BasicLit:
		fmt.Printf("%s ", x.Value)
	case *ast.Ident:
		fmt.Printf("%s ", x.Name)
	case *ast.UnaryExpr:
		fmt.Printf("%s ", x.Op.String())
		walk(x.X)
	case *ast.BinaryExpr:
		walk(x.X)
		fmt.Printf("%s ", x.Op.String())
		walk(x.Y)
	case *ast.ParenExpr:
		fmt.Printf("(")
		walk(x.X)
		fmt.Printf("\b) ")
	}
}

func main() {
	// konstrukce parseru a parsing zdrojového kódu
	node, err := parser.ParseExpr(source)
	if err != nil {
		log.Fatal(err)
	}

	// zahájení průchodu abstraktním syntaktickým stromem
	walk(node)
}
