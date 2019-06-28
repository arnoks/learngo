// dup1 is a simple program for finding duplicate lines in the input
// it searches for lines in Stdin and prints the frequency to Stdout
// the output is no sorted an will differ for each run
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d:\t%s\n", n, line)
	}
}
