package main

import "fmt"

func main() {
	// Rotate array by two elements:
	s := [...]int{0, 1, 2, 3, 4, 5} // create an array [6]int
	fmt.Println(s)
	reverse(s[:2])
	reverse(s[2:])
	reverse(s[:]) // reverse expects a slice, s is an array but s[:] converts into a slice
	fmt.Println(s)

	// Rotate slice by two elements:
	t := []int{11, 22, 33, 44, 55, 66} // create a slice of ints []int
	fmt.Println(t)
	reverse(t[:2])
	reverse(t[2:])
	reverse(t) // this work!
	fmt.Println(t)
}

// reverse a slice of ints in place

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
