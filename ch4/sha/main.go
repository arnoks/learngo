package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var v string
	algo := flag.String("a", "sha256", "Select either of the hash functions")
	flag.Parse()
	if len(os.Args) > 1 {
		v = os.Args[len(os.Args)-1]
	}
	hash(v, algo)
}

func hash(s string, algo *string) {
	if s == "" {
		s = "a"
	}
	switch {
	case *algo == "sha512":
		{
			fmt.Printf("sha512(%s) = %v\n", s, sha512.Sum512([]byte(s)))
		}
	case *algo == "sha384":
		{
			fmt.Printf("sha384(%s) = %v\n", s, sha512.Sum384([]byte(s)))
		}

	default:
		{
			fmt.Printf("sha256(%s) = %v\n", s, sha256.Sum256([]byte(s)))
		}
	}
}
