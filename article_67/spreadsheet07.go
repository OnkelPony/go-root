// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá sedmá část
//    Tvorba sešitů pro tabulkové procesory v programovacím jazyku Go – formát xlsx
//    https://www.root.cz/clanky/tvorba-sesitu-pro-tabulkove-procesory-v-programovacim-jazyku-go-format-xlsx/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_67/spreadsheet07.html

package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test07.xlsx"

func main() {
	// konstrukce struktury typu File
	worksheet := xlsx.NewFile()

	fmt.Println(worksheet)

	// přidání listu do sešitu
	sheet, err := worksheet.AddSheet("Tabulka1")
	if err != nil {
		panic(err)
	}
	defer sheet.Close()

	fmt.Println(sheet)

	// přidání řádku do tabulky
	row := sheet.AddRow()

	// přidání buňky na řádek a naplnění hodnotou
	row.AddCell().SetFloatWithFormat(1/3.0, "#0")
	row.AddCell().SetFloatWithFormat(1/3.0, "#0.0")
	row.AddCell().SetFloatWithFormat(1/3.0, "#0.00")
	row.AddCell().SetFloatWithFormat(1/3.0, "#0.000")
	row.AddCell().SetFloatWithFormat(1/3.0, "#0.0000")
	row.AddCell().SetFloatWithFormat(1/3.0, "#0.00000")
	row.AddCell().SetFloatWithFormat(1/3.0, "#0.000000")

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
