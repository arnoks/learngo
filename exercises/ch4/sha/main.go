package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	algo := flag.String("a", "sha256", "Select either of the hash functions")
	flag.Parse()
	b := readInput(os.Stdin)
	fmt.Printf("hash(%s) = %v\n", *algo, hash(b, algo))
}

func readInput(rd io.Reader) *[]byte {
	br := bufio.NewReader(rd) // wrap the unbuffered reader with a buffer
	// readall will read the whole file and return no error in case of eof
	b, _ := ioutil.ReadAll(br)
	return &b
}

func hash(b *[]byte, algo *string) []byte {
	switch {
	case *algo == "sha512":
		{
			h := sha512.Sum512(*b)
			return h[:] // return a slice rather than [48]byte
		}
	case *algo == "sha384":
		{
			h := sha512.Sum384(*b)
			return h[:]
		}
	default:
		{
			h := sha256.Sum256(*b)
			return h[:]
		}
	}
}
