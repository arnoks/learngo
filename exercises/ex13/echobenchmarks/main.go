// echobenchmark compares the two implementations of join.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Loopargs(a []string, w io.Writer) {
	var s, sep string
	for i := 0; i < len(a); i++ {
		s += sep + a[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func Joinargs(a []string, w io.Writer) {
	fmt.Fprintln(w, strings.Join(a, " "))
}

func main() {
	Loopargs(os.Args[1:], os.Stdout)
	Joinargs(os.Args[1:], os.Stdout)
}
