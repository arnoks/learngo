package popcount

import (
	"math/rand"
	"testing"
)

// BenchmarkPopCount meassures PopCount implementation
func BenchmarkPopCount(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	for n := 0; n < b.N; n++ {
		r = rand.Uint64()
		PopCount(r)
	}
}

// BenchmarkLoop  meassure the loop implemention
func BenchmarkLoop(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	for n := 0; n < b.N; n++ {
		r = rand.Uint64()
		Loop(r)
	}
}
