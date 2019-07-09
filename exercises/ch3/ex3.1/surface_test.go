package main

import (
	"fmt"
	"math"
	"testing"
)

// TestNan the Nan return code
func TestInfinits(t *testing.T) {
	v := -1.
	r := math.Sqrt(v)
	if math.IsNaN(r) {
		t.Errorf("Sqrt (-1) = %v", r)
	}
	fmt.Printf("srt(%v) = %v\n", v, r)

	for i := 10; i > -10; i-- {
		r = math.Pi / float64(i)
		if math.IsInf(r, 0) {
			t.Errorf("Division by zero pi/%v -> %v", i, r)
			continue
		}
	}
}
