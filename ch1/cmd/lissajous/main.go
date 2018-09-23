// lissajous is a cmd frontend for the the lissajous package
package main

import (
	"flag"
	"image/color"
	"os"

	"github.com/arnoks/learngo/ch1/lissajous"
)

var freq = flag.Float64("freq", 1.55, "Frequency as float 0.0 ... ")
var fg = flag.Uint("fg", 0xff, "Drawing color")
var bg = flag.Uint("bg", 0x00, "Background color")

func main() {
	flag.Parse()
	bgc := color.Gray{uint8(*bg)}
	fgc := color.Gray{uint8(*fg)}

	l := lissajous.Parameter{Freq: *freq, Palette: []color.Color{bgc, fgc}}
	l.Draw(os.Stdout)
	//	lissajous.Lissajous(os.Stdout)
}
