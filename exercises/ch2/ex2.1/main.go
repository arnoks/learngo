// ex2.1 is using tempconv to convert a given parameter into 3 different scales
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arnoks/learngo/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		} else {
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			k := tempconv.Kelvin(t)
			fmt.Printf("%v == %v\n%v == %v\n%v == %v\n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
		}
	}
}
