package popcount

// pc[i] is the population count of i

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count - number of set bit within a uint64
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Loop calculates the population count using a loop
func Loop(x uint64) int {
	var c byte
	for i := uint(0); i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}

// Shift calculates the population count using a loop and shift through 64 bits
func Shift(x uint64) int {
	var c uint64
	for i := uint64(0); i < 64; i++ {
		c += x & 1
		x = x >> 1
	}
	return int(c)
}

// Eliminate calculates the population count using bit elimination
func Eliminate(x uint64) int {
	c := 0
	for ; x != 0; c++ {
		x = x & (x - 1)
	}
	return c
}

// Recursive Apporach ~ 300 fold the lookup version
func countSetBits(n uint64) int {
	if n != 0 {
		return 1 + countSetBits(n&(n-1))
	}
	return 0
}
