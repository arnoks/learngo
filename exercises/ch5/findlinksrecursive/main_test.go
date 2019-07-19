package main

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_visit(t *testing.T) {
	r, err := os.Open("../../../golang.org.html")
	if err != nil {
		t.Errorf("Error opening test file %v", err)
		return
	}
	doc, err := html.Parse(r)
	if err != nil {
		t.Errorf("Error parsing test data file %v", err)
		return
	}
	var stack []string
	type args struct {
		cur   *html.Node
		stack []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Full golang home page", args{doc, stack}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visit(tt.args.cur, tt.args.stack)
			for _, s := range stack {
				fmt.Printf(s)
			}
		})
	}
}
