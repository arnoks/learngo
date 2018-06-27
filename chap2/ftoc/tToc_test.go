package main

import (
	"fmt"
	"os"
	"testing"
)

func TestWarm(t *testing.T) {
	val := 24.0
	resC := FToC(val)
	expC := (val - 32.0) * 5.0 / 9.
	if resC-expC > 1e-6 {
		fmt.Printf("%g != %g\n", resC, expC)
		os.Exit(1)
	}
	os.Exit(0)
}
