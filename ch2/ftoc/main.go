// fToc Calculates the °C for a given °F temperature here freezing and boiling temperature
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, FToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, FToC(boilingF))
}

func FToC(f float64) float64 {
	// FToC funtion does a conversion form Fahrenheit to Celsius
	return (f - 32) * 5 / 9
}

func CToF(f float64) float64 {
	// CToF funtion does a conversion form Fahrenheit to Celsius
	return (f + 32) / 5 * 9
}
