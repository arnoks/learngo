// Modify dup2 to print the names of all files in which each duplicated line
// occurs.
// The implementation is useing a set to keep track of the files

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type counter struct {
	n     int
	files map[string]bool // variable for holding a set.
	// boolean maps are a simple approach to implement a set.
}

func main() {
	files := os.Args[1:]
	counts := make(map[string]*counter)

	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else { // loop over the files
		for _, file := range files {
			fp, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(fp, counts, file)
		}
	}
	for line, c := range counts {
		fmt.Printf("%d:[%s]", c.n, line)
		for f, _ := range c.files {
			fmt.Printf("%s ", f)
		}
		fmt.Println()
	}
}

func countLines(f *os.File, counts map[string]*counter, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var t string = strings.TrimSpace(input.Text())
		c, ok := counts[t] // test if we have the line already
		if !ok {           // create a counter if we do not have one
			c = new(counter)
			c.files = make(map[string]bool) // alllocate the file set
		}
		c.n++
		c.files[filename] = true
		counts[t] = c
	}
}
