// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 7:
//     Vylepšená varianta předchozího příkladu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/07_filled_image_alpha_correct.html

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("07.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for x := 0; x < width; x++ {
		alpha := byte(x)
		c := color.NRGBA{0, 255, 0, alpha}
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}
	png.Encode(outfile, img)
}
