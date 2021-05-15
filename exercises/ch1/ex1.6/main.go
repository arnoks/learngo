package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var background = color.Black
var forground = color.RGBA{0x00, 0xFF, 0x00, 0xFF}

var palette = []color.Color{background,
	color.RGBA{0x10, 0x10, 0x00, 0xFF},
	color.RGBA{0x20, 0x10, 0x00, 0xFF},
	color.RGBA{0x30, 0x20, 0x00, 0xFF},
	color.RGBA{0x40, 0x30, 0x00, 0xFF},
	color.RGBA{0x50, 0x40, 0x00, 0xFF},
	color.RGBA{0x60, 0x50, 0x00, 0xFF},
	color.RGBA{0x70, 0x60, 0x00, 0xFF},
	color.RGBA{0x80, 0x70, 0x00, 0xFF},
	color.RGBA{0x90, 0x80, 0x00, 0xFF},
	color.RGBA{0xA0, 0x90, 0x00, 0xFF},
	color.RGBA{0xB0, 0xA0, 0x00, 0xFF},
	color.RGBA{0xC0, 0xB0, 0x00, 0xFF},
	color.RGBA{0xD0, 0xC0, 0x00, 0xFF},
	color.RGBA{0xE0, 0xD0, 0x00, 0xFF},
	color.RGBA{0xF0, 0xE0, 0x00, 0xFF},
}

const (
	backgroundIndex = 0
	forgroundIndex  = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0 //
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	c := 1

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c))
		}

		if i%4 == 0 {
			if c+1 < len(palette) {
				c += 1
			} else {
				c = 1
			}
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
