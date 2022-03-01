// Execise 4.3 demonstrates array pointers
package main

import "fmt"

func main() {
	// Array
	s1 := [...]int{0, 1, 2, 3, 4, 5}
	// Slice
	s2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s1)
	reverse(s1[:])
	reverse(s2)
	fmt.Println(s1)
}

// reverse the input array which is passed in as a pointer
func reverse(s []int) { // we expect a slice , hence arrays will not work
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		(s)[i], (s)[j] = (s)[j], (s)[i]
	}
}
