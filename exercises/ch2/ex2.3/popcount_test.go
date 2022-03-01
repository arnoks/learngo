package popcount

import (
	"math/rand"
	"testing"
)

// TestPopCounts verifies the correctness of the implementations
// all versions should always return the same values.
func TestPopCounts(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := rand.Uint64()
		p, l, s, e, r := PopCount(v), Loop(v), Shift(v), Eliminate(v), countSetBits(v)
		if p != l || p != s || l != s || s != e || p != r {
			t.Errorf("Error, PopCounts yield different results for %v : %v != %v != %v != %v != %v ",
				v, p, l, s, e, r)
		}
	}
}

func TestSWAR(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := rand.Uint64()
		r, c := PopCount(v), numberOfSetBits(v)
		if r != c {
			t.Errorf("results do not match - reference: %v  SWAR: %v", r, c)
		}
	}
}
