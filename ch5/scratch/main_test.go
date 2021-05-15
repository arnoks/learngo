package main

import (
	"fmt"
	"testing"
)

func TestGrowbig(t *testing.T) {
	n := 100
	t.Log("Running Test..")
	for i := 0; i < n; i++ {
		growBig()
	}
}

func TestHypot(t *testing.T) {
	var cases = []struct {
		a, b, want float64
	}{
		{3.0, 4.0, 5.0},
		{4.0, 5.0, 6.40312},
	}

	for _, c := range cases {
		got := hypot(c.a, c.b)
		if got != c.want {
			t.Errorf("got %g want %g", got, c.want)
		}
	}
}

func TestScope(t *testing.T) {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y)
}

func f() int {
	return 0
}

func g(x int) int {
	return x + 0
}
