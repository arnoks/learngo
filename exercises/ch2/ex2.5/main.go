package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Uint64()
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", x&(x-1))
}
