package main

// find duplictes in input file by read whole files at once and splitting
// lines using strings.Split.

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	for _, filename := range files {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		} else {
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%03d:\t%s\n", n, line)
		}
	}
}
