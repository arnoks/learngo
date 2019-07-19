package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	var filnam string
	flag.StringVar(&filnam, "file", "../../../golang.org.html", "-file filename")
	flag.Parse()
	var r io.Reader
	r = os.Stdin

	if filnam != "" {
		var err error
		r, err = os.Open(filnam)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %v", err)
			os.Exit(1)
		}
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	walk(doc)
	fmt.Println(&counter)
}

func walk(n *html.Node) {
	if n.Type == html.ElementNode {
		counter.Add(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(c)
	}
}

// ElementCounter provides a map for counting strings
type ElementCounter struct {
	m map[string]int
}

// Add the element and increments the counter
func (c *ElementCounter) Add(s string) {
	if c.m == nil {
		c.m = make(map[string]int)
	}
	if c.m[s] == 0 {
		c.m[s] = 1
	} else {
		c.m[s]++
	}
}

// String is the Stringer implementation for the counter
func (c *ElementCounter) String() string {
	var b bytes.Buffer
	var ordered []string
	for elm := range c.m {
		ordered = append(ordered, elm)
	}
	sort.Strings(ordered)

	for _, elm := range ordered {
		fmt.Fprintf(&b, "%10s: %2d\n", elm, c.m[elm])
	}
	return b.String()
}

var counter ElementCounter
