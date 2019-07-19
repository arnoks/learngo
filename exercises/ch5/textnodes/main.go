// textnodes prints the contents of all text nodes
// in an HTML document tree. \<script\ or \<style\>
// are not followed, since their contents are not visisble
// in web browser

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
	flag.StringVar(&filnam, "file", "", "-file filename")
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
}

func walk(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode { // \<script\ or \<style\>
			switch c.Data {
			case "script":
				continue
			case "style":
				continue
			default:
			}
		}
		walk(c)
	}
}
