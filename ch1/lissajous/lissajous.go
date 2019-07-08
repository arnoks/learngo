package lissajous

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const (
	whiteIndex = 0
	blackIndex = 1
	cycles     = 5
	res        = 0.001
	size       = 100
	nframes    = 64
	delay      = 8
)

// Parameter provides the features of the Lissajous
type Parameter struct {
	Freq    float64
	Palette []color.Color
}

// MakeGif generates fills the provided gif with a Lissajous figure
// it uses a standard two color Palette
func (p *Parameter) MakeGif(g *gif.GIF) {
	phase := 0.0
	for i := 0; i < g.LoopCount; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, p.Palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*p.Freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		g.Delay = append(g.Delay, delay)
		g.Image = append(g.Image, img)
	}
}

// MakeColoredGif generasted a colorful lissajous based on a standard palette
func (p *Parameter) MakeColoredGif(g *gif.GIF) {
	phase := 0.0

	for i := 0; i < g.LoopCount; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette.WebSafe)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*p.Freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(200)))
		}
		phase += 0.1
		g.Delay = append(g.Delay, delay)
		g.Image = append(g.Image, img)
	}
}

// Draw draws a Lissojous to the writer
func (p *Parameter) Draw(out io.Writer) {
	g := gif.GIF{LoopCount: 64}
	p.MakeGif(&g)
	gif.EncodeAll(out, &g)
}

// LissajousFun Draws a standard Lissajous
func Lissajous(out io.Writer) {
	var p = Parameter{rand.Float64() * 3.0, []color.Color{color.White, color.Black}}
	p.Draw(out)
}
