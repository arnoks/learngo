package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	var stack []string

	visit(doc, stack)
	for _, s := range stack {
		fmt.Printf(s)
	}
}

func visit(cur *html.Node, stack []string) {
	if cur == nil {
		return
	}
	addlink(cur, stack)

	// So it seems we have to follow both routes in order to catch all
	visit(cur.FirstChild, stack)
	visit(cur.NextSibling, stack)
}

func addlink(n *html.Node, stack []string) {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					stack = append(stack, a.Val)
					fmt.Println(a.Val)
				}
			}
		}
	}
}
