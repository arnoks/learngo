// parallelmdelbrot uses a go routine per stride. Each stride get updated
// independently. Parallel updates of the strides does not seem to be an issue
// for the image data structure.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 4096, 4096
)

var img *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))

type stripe struct {
	start, end int
}

var wq sync.WaitGroup

var row_index = make(chan stripe, 8)

func main() {
	start := time.Now()
	if len(os.Args) == 1 {
		single()
	} else {
		switch os.Args[1] {
		case "single":
			single()
		case "parallel":
			w := 4
			if len(os.Args) == 3 {
				c, err := strconv.Atoi(os.Args[2])
				if err == nil {
					w = c
				}
			}
			parallel(w)
		default:
			single()
		}
	}
	fmt.Fprintf(os.Stderr, "Calculation Time: %v\n", time.Since(start))
	png.Encode(os.Stdout, img)
	fmt.Fprintf(os.Stderr, "Encoding Time: %v\n", time.Since(start))
}

func single() {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, madelbrot_color(z))
		}
	}
}

func parallel(worker int) {
	setup(worker)
	var s stripe
	s_height := height / worker
	for s.start = 0; s.start < height; s.start += s_height {
		s.end = s.start + s_height
		fmt.Fprintln(os.Stderr, s)
		row_index <- s
	}
	close(row_index)
	wq.Wait()
}

func setup(c int) {
	for i := 0; i < c; i++ {
		wq.Add(1)
		go rowCalculator(row_index)
	}
}

func rowCalculator(row <-chan stripe) {
	s := <-row
	for py := s.start; py < s.end; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, madelbrot_color(z))
		}
	}
	wq.Done()
}

func madelbrot_color(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{uint8(255 - contrast*n), uint8(100 - n*3), uint8(contrast * n), 0xfF}
		}
	}
	return color.RGBA{0x0F, 0x00, 0x3F, 0xff}
}
