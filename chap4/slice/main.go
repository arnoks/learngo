package main

import (
	"fmt"
)

func main() {
	process()
}

func process() {
	var a [10]int
	fmt.Println(len(a))
	for i, _ := range a {
		a[i] = i * i
	}
	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	var q = [...]int{1, 2, 3, 4, 5, 6}
	r := [...]int{1, 3, 27, 81}
	fmt.Printf("Length of q: %d , r: %d\n", len(q), len(r))

	// Currency Example
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "\u00A3", RMB: "\u00A5"}
	for _, s := range symbol {
		fmt.Println(s)
	}
}
