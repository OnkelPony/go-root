// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/

package main

import (
	"github.com/mingrammer/cfmt"
)

func main() {
	cfmt.Success("Success message")
	cfmt.Info("Info message")
	cfmt.Warning("Warning message")
	cfmt.Error("Error message")
}
