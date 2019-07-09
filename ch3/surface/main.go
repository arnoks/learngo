// Surface computes an SVG rendering of a 3-D surface function.package main

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 300            // convas size
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // Axis ranges (-.yxrange.. +xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x,y axes 30°
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style = 'stroke:grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, c, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, c, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, c, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, c, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, string, bool) {
	// Find point (x,y) at corner of cells(i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	//z := sinus(x, y)                                // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	z := eggbox(x, y)                               // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale             // 30° projection
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale // 30° projection
	c := "ff0000"

	if math.IsInf(z, 0) {
		return 0., 0., c, false
	}
	if z < 0 {
		c = "0000ff"
	}

	return sx, sy, c, true
}

// sinus
func sinus(x, y float64) float64 {
	r := math.Hypot(x, y) //distance from (0,0)
	return math.Sin(r) / r
}

// sphere
func eggbox(x, y float64) float64 {
	z := math.Sin(x/1.5) * math.Sin(y/1.5) * 0.2
	return z
}
