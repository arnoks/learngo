package msort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort1(t *testing.T) {
	var a = []string{"C", "F", "H", "U", "A", "T", "B", "D", "Z", "X", "Y"}
	sort.Sort(StringSlice(a))
	for _, s := range a {
		fmt.Println(s)
	}
}

func TestSort2(t *testing.T) {
	var a = []string{"C", "F", "H", "U", "A", "T", "B", "D", "Z", "X", "Y"}
	sort.Strings(a)
	for _, s := range a {
		fmt.Println(s)
	}
}
