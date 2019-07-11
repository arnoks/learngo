package main

import (
	"testing"
)

func Test_sort(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Sort(values)
}

func Test_appendValues(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tr tree
	appendValues(values, &tr)
}
