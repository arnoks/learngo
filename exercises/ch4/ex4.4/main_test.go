package main

import "testing"

func Benchmark_rotate(b *testing.B) {
	var s [50]int
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	for n := 0; n < b.N; n++ {
		rotate(s[:], 25)
	}
}

func Benchmark_rotate2(b *testing.B) {
	var s [50]int
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	for n := 0; n < b.N; n++ {
		rotate2(s[:], 25)
	}
}
