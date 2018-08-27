package main

import (
	"flag"
	"fmt"
	"strings"
)

var nl = flag.Bool("n", false, "Print skip newline")
var sep = flag.String("s", " ", "Field separator")

func main() {
	flag.Parse()
	var s = strings.Join(flag.Args(), *sep)
	if *nl {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}
