// Extend the visit function so that it extracts other kinds
// of links from the document, such as images, scripts and style
// sheets.
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

// visit append to links each link found in and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a": // links
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img": // images
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		case "link": // Style sheets
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "script": // Style sheets
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
