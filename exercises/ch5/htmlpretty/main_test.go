// Develop startElement and endElement into a general HTML

// pretty-printer. Print comment nodes, text nodes and attributes

// of each element (<a href '...'>). Use short forms like <img/>

// instead of <img></img> when an element has no children. Write

// a test to ensure that the output can be parsed Successfully.

package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_forEachNode(t *testing.T) {
	r, err := os.Open("../../../golang.org.html")
	if err != nil {
		t.Errorf("Error opening sample data %v", err)
		return
	}
	doc, err := html.Parse(r)
	if err != nil {
		t.Errorf("Error parsing sample data %v", err)
		return
	}

	type args struct {
		n    *html.Node
		pre  func(n *html.Node, w io.Writer)
		post func(n *html.Node, w io.Writer)
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test", args{doc, startElement, endElement}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			forEachNode(tt.args.n, w, tt.args.pre, tt.args.post)
			if _, err := html.Parse(w); err != nil {
				t.Errorf("forEachNode() returns invalid html")
			}
		})
	}
}
