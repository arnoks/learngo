// Mandelbrot2 uses a confined image which is updated by the setPoint go
// routine. This was done to make the image update concurrency proof. The
// implementation using a queue operation for each point has lots of overhead
// and is lower than the single threaded implementation.

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

type stripe struct {
	start, end int
}

type Point struct {
	x, y int
	c    color.Color
}

var wq sync.WaitGroup

var row_index = make(chan stripe, 8)
var imageQ = make(chan Point, 4096)
var Done = make(chan struct{})

func main() {
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
}

func single() {
	start := time.Now()
	var img *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, madelbrot_color(z))
		}
	}
	fmt.Fprintf(os.Stderr, "Calculation Time: %v\n", time.Since(start))
	start = time.Now()
	png.Encode(os.Stdout, img)
	fmt.Fprintf(os.Stderr, "Encoding Time: %v\n", time.Since(start))
}

func parallel(worker int) {
	setup(worker)
	go setPoint(imageQ)
	var s stripe
	s_height := height / worker
	for s.start = 0; s.start < height; s.start += s_height {
		s.end = s.start + s_height
		row_index <- s
	}
	wq.Wait()
	close(imageQ)
	<-Done
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
			imageQ <- Point{px, py, madelbrot_color(z)}
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

func setPoint(points chan Point) {
	var img *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))
	start := time.Now()
	for point := range points {
		img.Set(point.x, point.y, point.c)
	}
	fmt.Fprintf(os.Stderr, "Calculation Time: %v\n", time.Since(start))
	start = time.Now()
	png.Encode(os.Stdout, img)
	fmt.Fprintf(os.Stderr, "Encoding Time: %v\n", time.Since(start))
	Done <- struct{}{}
}
