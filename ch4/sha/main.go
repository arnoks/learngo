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
	hash(b, algo)
}

func readInput(rd io.Reader) *[]byte {
	br := bufio.NewReader(rd)
	b, err := ioutil.ReadAll(br)
	if err != nil {
		b = []byte("Hallo")
		return &b
	}
	return &b
}

func hash(b *[]byte, algo *string) {
	switch {
	case *algo == "sha512":
		{
			fmt.Printf("sha512: %v\n", sha512.Sum512(*b))
		}
	case *algo == "sha384":
		{
			fmt.Printf("sha384: %v\n", sha512.Sum384(*b))
		}

	default:
		{
			fmt.Printf("sha256: %v\n", sha256.Sum256(*b))
		}
	}
}
