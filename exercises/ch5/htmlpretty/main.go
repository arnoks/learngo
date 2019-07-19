// Develop startElement and endElement into a general HTML
// pretty-printer. Print comment nodes, text nodes and attributes
// of each element (<a href '...'>). Use short forms like <img/>
// instead of <img></img> when an element has no children. Write
// a test to ensure that the output can be parsed Successfully.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func getInput() io.Reader {
	fn := flag.String("f", "", "html file to pretty printed")
	flag.Parse()
	if *fn == "" {
		return os.Stdin
	}
	r, err := os.Open(*fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input %v", err)
		os.Exit(1)
	}
	return r
}

func main() {

	doc, err := html.Parse(getInput())
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	var b bytes.Buffer
	forEachNode(doc, &b, startElement, endElement)
	fmt.Println(b.String())
}

func forEachNode(n *html.Node, w io.Writer, pre, post func(n *html.Node, w io.Writer)) {
	if pre != nil {
		pre(n, w)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode { // do not follow <script and <style>
			switch c.Data {
			case "script":
				continue
			case "style":
				continue
			}
		}
		forEachNode(c, w, pre, post)
	}
	if post != nil {
		post(n, w)
	}
}

var depth int

func startElement(n *html.Node, w io.Writer) {
	hasChildren := n.FirstChild == nil
	switch n.Type {
	case html.ElementNode:
		if hasChildren {
			fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		} else {
			fmt.Fprintf(w, "%*s<%s/", depth*2, "", n.Data)
		}
		for _, attr := range n.Attr {
			fmt.Fprintf(w, " %s=\"%s\"", strings.TrimSpace(attr.Key), strings.TrimSpace(attr.Val))
		}
		fmt.Fprintln(w, ">")
		depth++
	case html.TextNode:
		if txt := strings.TrimSpace(n.Data); txt != "" {
			fmt.Fprintln(w, txt)
		}
	case html.CommentNode:
		if comment := strings.TrimSpace(n.Data); comment != "" {
			fmt.Fprintf(w, "<!--%s-->\n", comment)
		}
	}
}

func endElement(n *html.Node, w io.Writer) {
	hasChildren := n.FirstChild == nil
	if n.Type == html.ElementNode {
		depth--
		if hasChildren {
			fmt.Fprintf(w, "%*s<%s>\n", depth*2, "", strings.TrimSpace(n.Data))
		}
	}
}
