package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

// String Method for the *tree type reveals
// the sequence of values in the tree.
func (t *tree) String() string {
	if t == nil {
		return ""
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%v ", t.value)
	fmt.Fprint(&buf, t.left.String())
	fmt.Fprint(&buf, t.right.String())
	return buf.String()
}

// Sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root)

}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add a value to a node
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return & tree(value:value)
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
