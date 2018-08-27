// Report all html elements and their usage for the given html file via stdin
// Output the elemets in alphabthical order
package main

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int)
	visit(doc, elements)

	var keys []string
	for k, _ := range elements {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%-20s \t %3d\n", k, elements[k])
	}
}

func visit(n *html.Node, elements map[string]int) {
	if n.Type == html.ElementNode {
		elements[n.Data] += 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, elements)
	}
}
