// Retreive all links from a website and print them to stdout
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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild.NextSibling)
	}

	fmt.Printf("Scanning Node\n")
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				fmt.Printf("Appending %s\n", a.Val)
			}
		}
	}
	return links
}
