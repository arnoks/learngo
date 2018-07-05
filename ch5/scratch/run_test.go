package main

import (
	"testing"
)

func TestGrowbig(t *testing.T) {
	n := 100
	t.Log("Running Test..")
	for i := 0; i < n; i++ {
		grow_big()
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

/*func TestMain(m *testing.M) {
	main()
}
*/
