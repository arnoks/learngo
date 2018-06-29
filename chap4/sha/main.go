package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var sha = flag.String("sha", "256", "Select either of the hash functions 256 384 512")

func main() {
	flag.Parse()
	a := flag.Args()
	for _, s := range a {
		ba := []byte(s)

		var hash *byte

		if *sha == "256" {
			hash = sha256.Sum256(ba)
		}
		if *sha == "384" {
			hash = sha512.Sum384(ba)
		}
		if *sha == "512" {
			hash = sha512.Sum512(ba)
		}
		fmt.Printf("%b\n", hash)
	}
}
