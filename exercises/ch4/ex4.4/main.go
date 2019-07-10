package main

import "fmt"

func main() {
	// Rotate array by two elements:
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}  //
	s2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //
	fmt.Printf("%v\n", s)
	rotate(s, 3)
	fmt.Printf("%v\n", s)
	rotate2(s2, 3)
	fmt.Printf("%v\n", s2)

}

// rotate the slice s by offset o using storage
func rotate(s []int, o int) {
	var z = make([]int, o)
	copy(z, s[0:o])
	copy(s[0:], s[o:])
	copy(s[len(s)-o:], z)
}

// rotate in-place by offset o
func rotate2(s []int, o int) {
	l := len(s) - 1
	for i, j := 0, o; i < l; i++ {
		s[i], s[j] = s[j], s[i] // swap i,j
		if j < l {
			j++
		}
	}
}
