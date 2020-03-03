package popcount

import (
	"math/rand"
	"testing"
)

// TestPopCounts verifies the correctness of the implementations
// both versions should always return the same values.
func TestPopCounts(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := rand.Uint64()
		p, l, s, e, r := PopCount(v), Loop(v), Shift(v), Eliminate(v), countSetBits(v)
		if p != l || p != s || l != s || s != e {
			t.Errorf("Error, PopCounts yield different results for %v : %v != %v != %v != %v != %v ",
				v, p, l, s, e, r)
		}
	}
}
