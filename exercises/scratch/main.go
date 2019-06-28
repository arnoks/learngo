package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(r)
	if len(b) > 0 {
		fmt.Print(string(b))
	}
}
