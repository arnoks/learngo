package main

import (
	"fmt"
	"math"
)

func main() {
	b := 1.0
	for a := 2.0; a < 5.0; a++ {
		fmt.Printf("Hypothenusis(%2.2f,%2.2f) = %2.2f\n", a, b, hypot(a, b))
		b = a
	}
}

func hypot(a, b float64) float64 {
	return (math.Sqrt(a*a + b*b))
}

func growBig() {
	var stack []string
	for i := 0; i < 1000; i++ {
		stack = append(stack, "Hallo")
	}
}
