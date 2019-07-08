package popcount

import (
	"math/rand"
	"testing"
)

// TestPopCounts verifies the correctness of the implementations
// both versions should always return the same values.
func TestPopoCounts(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := rand.Uint64()
		p, l := PopCount(v), PopCountLoop(v)
		if PopCount(v) != PopCountLoop(v) {
			t.Errorf("Error, PopCounts yield different results for %v : %v != %v ", v, p, l)
		}
	}
}
