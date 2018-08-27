package main

import (
	"bufio"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	i := 0
	var err error
	for ; err == nil; _, _, err = bufio.ScanWords(p, false) {
		i++
	}
	*c += WordCounter(i)
	return i, nil
}

func main() {

}
