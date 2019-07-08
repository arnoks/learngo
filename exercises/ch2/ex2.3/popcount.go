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

// PopCountLoop calculates the population count using a loop
func PopCountLoop(x uint64) int {
	var c byte
	for i := uint(0); i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}
