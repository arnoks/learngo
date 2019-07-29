package main

import (
	"bufio"
)

// ByteCounter implements the Writer interafe in order to count
// the number of bytes in the intput
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// WordCounter implements the Writer Interface  in order to count
// the words of the intput
type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	i := 0
	av, _, e := bufio.ScanWords(p, true)
	for ; av > 0 || e != nil; av, _, e = bufio.ScanWords(p[i:], true) {
		i += av
		*w++
	}
	return int(*w), e
}

// LineCounter implements Writer to count the number of lines of the provided stream
type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	i := 0
	av, _, e := bufio.ScanLines(p, true)
	for ; av > 0 || e != nil; av, _, e = bufio.ScanLines(p[i:], true) {
		i += av
		*l++
	}
	return int(*l), e
}

func main() {

}
