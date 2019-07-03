package main

// echo command printing both the index and the argument itself

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("args[%d] = %s\n", i, arg)
	}
}
