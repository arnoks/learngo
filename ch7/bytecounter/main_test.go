package main

import (
	"fmt"
	"testing"
)

func TestBc(t *testing.T) {
	var bc ByteCounter
	var wc WordCounter

	fmt.Printf("byte count: %d \n", bc)
	bc.Write([]byte("Hello world"))
	fmt.Printf("byte count: %d \n", bc)
	fmt.Fprintln(&bc, "Hello World")
	fmt.Printf("byte count: %d \n", bc)
	fmt.Fprintln(&wc, "Hello World")
	fmt.Printf("Word count: %v \n", wc)
}
