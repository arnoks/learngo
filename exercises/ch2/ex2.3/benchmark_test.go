package popcount

import (
	"math/rand"
	"testing"
)

// BenchmarkPopCount meassures PopCount implementation
func BenchmarkPopCount(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	r = rand.Uint64()
	for n := 0; n < b.N; n++ {
		PopCount(r)
	}
}

// BenchmarkLoop  meassure the loop implemention
func BenchmarkLoop(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	r = rand.Uint64()
	for n := 0; n < b.N; n++ {
		Loop(r)
	}
}

// BenchmarkShift meassure the loop implemention
func BenchmarkShift(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	r = rand.Uint64()
	for n := 0; n < b.N; n++ {
		Shift(r)
	}
}

// BenchmarkEliminate
func BenchmarkEliminate(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	r = rand.Uint64()
	for n := 0; n < b.N; n++ {
		Eliminate(r)
	}
}

func BenchmarkRecursion(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	r = rand.Uint64()
	for n := 0; n < b.N; n++ {
		countSetBits(r)
	}
}

func BenchmarkSWAR(b *testing.B) {
	// run the Fib function b.N times
	var r uint64
	r = rand.Uint64()
	for n := 0; n < b.N; n++ {
		numberOfSetBits(r)
	}
}
