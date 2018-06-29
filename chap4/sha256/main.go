package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("A"))
	fmt.Printf("% x\n% x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	d := bitDiffCount(&c1, &c2)
	fmt.Printf("Bitcount: = %d/%d\n", d, 32*8)

}

func bitDiffCount(aPtr *[32]byte, bPtr *[32]byte) int {
	var bc int

	for i, a := range *aPtr {
		r := a ^ bPtr[i]
		for j := uint(0); j < 8; j++ {
			mask := byte(1 << j)
			if mask&r != 0 {
				bc++
			}
		}
	}
	return bc
}
