package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	freq := make(map[string]int)
	for s.Scan() { // default scanner is a line scanner
		word := s.Text()
		freq[word]++
		if s.Err() != nil {
			break
		}
	}

	var keys []string
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, word := range keys {
		fmt.Printf("[%v]:\t%v \n", word, freq[word])
	}
}
