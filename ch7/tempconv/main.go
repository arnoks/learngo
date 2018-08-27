package main

import (
	"flag"
	"fmt"

	"gopl.io/ch2/tempconv"
)

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperatur %q", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var low = CelsiusFlag("low", 18.0, "the lower temperature boundary")
var high = CelsiusFlag("high", 22.0, "the higher temperature boundary")

func main() {
	temp := celsiusFlag{20}
	flag.CommandLine.Var(&temp, "temp", "the temperature")
	flag.Parse()
	fmt.Println(low)
	fmt.Println(temp)
	fmt.Println(high)

}
