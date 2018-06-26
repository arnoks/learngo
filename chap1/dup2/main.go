// dup2 will find duplicates in all its input files
//
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else { // loop over the files
		for _, file := range files {
			fp, err := os.Open(file)
			if err == nil {
				countLines(fp, counts)
			} else {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
		}
	}
	for line, n := range counts {
		fmt.Printf("%d:\t%s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
