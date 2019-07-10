// Execise 4.3 demonstrates array pointers
package main

import "fmt"

func main() {
	s := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}

// reverse the input array which is passed in as a pointer
func reverse(s *[6]int) { // actually this is an pointer to a slice
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
