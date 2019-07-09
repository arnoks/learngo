package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {

	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("A"))
	fmt.Printf("% x\n% x\n Hashes equal? =  %t\n", c1, c2, c1 == c2)
	s1 := c1[:]
	s2 := c2[:]
	d := bitDiffCount(&s1, &s2)
	fmt.Printf("Bitcount: = %d/%d\n", d, 32*8)

}

// Count different bits ex 4.1
func bitDiffCount(aPtr *[]byte, bPtr *[]byte) int {
	var bc byte
	for i, a := range *aPtr {
		r := a ^ (*bPtr)[i]
		bc += pc[r]
	}
	return int(bc)
}
