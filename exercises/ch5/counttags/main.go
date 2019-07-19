package main

import (
	"flag"
	"fmt"
	"io"
	"os"

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
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
