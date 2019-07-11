// charcount computes counts of Unicode characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type charclass int

const (
	digit = charclass(iota)
	letter
	punkt
	space
	lower
	upper
	maxcc
)

func main() {
	counts := make(map[rune]int)
	ccname := [...]string{"digits", "letters", "punkts", "spaces", "lows", "upps"}
	var utflen [utf8.UTFMax + 1]int
	var cc [maxcc]charclass

	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount-2: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// switch work only if character classes are non overlapping
		// switch is just another form of if .. else if ...
		switch {
		case unicode.IsDigit(r) == true:
			cc[digit]++
		case unicode.IsLetter(r) == true:
			cc[letter]++
		case unicode.IsPunct(r) == true:
			cc[punkt]++
		case unicode.IsSpace(r) == true:
			cc[space]++
		}
		if unicode.IsLower(r) == true {
			cc[lower]++
		}
		if unicode.IsUpper(r) == true {
			cc[upper]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Print("rune\tcount\n")
	for i, n := range counts {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("len\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	for i, c := range cc {
		fmt.Printf("\n%s\t%d", ccname[i], c)
	}
}
