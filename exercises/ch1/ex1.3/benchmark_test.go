// Benchmark the two implementations

package main

import (
	"bytes"
	"testing"
)

// run benchmark against the loop based implementation
func BenchmarkLoop(b *testing.B) {
	b.Log("loop benchmark")
	args := []string{"eins", "zwei", "drei", "vier", "fuenf", "eins", "zwei", "drei", "vier", "fuenf"}
	for i := 0; i < b.N; i++ {
		var w bytes.Buffer
		Loopargs(args, &w)
	}
}

// run benchmark against the library based implementation
func BenchmarkJoin(b *testing.B) {
	b.Log("join benchmark")
	args := []string{"eins", "zwei", "drei", "vier", "fuenf", "eins", "zwei", "drei", "vier", "fuenf"}
	for i := 0; i < b.N; i++ {
		var w bytes.Buffer
		Joinargs(args, &w)
	}
}
