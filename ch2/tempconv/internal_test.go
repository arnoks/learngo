// test the converison
package tempconv

import (
	"fmt"
	"os"
	"testing"
)

func TestConversion(t *testing.T) {
	var f = []Fahrenheit{32.0, 212.0, 68.0, 77}
	var c = []Celsius{0.0, 100.0, 20.0, 25}

	for i := 0; i < len(f); i++ {
		if CToF(c[i]) != f[i] {
			fmt.Printf("Failed to converted from Celsius to Fahrenheit: %g°C != %g°F\n", c[i], f[i])
			os.Exit(1)
		} else {
			fmt.Printf("Success converted from Celsius to Fahrenheit: %g°C == %g°F\n", c[i], f[i])
		}
		if FToC(f[i]) != c[i] {
			fmt.Printf("Failed to converted from Fahrenheit to Celsius %g°F != %g°C\n", f[i], c[i])
			os.Exit(1)
		} else {
			fmt.Printf("Success converted from Fahrenheit to Celsius %g°F == %g°C\n", f[i], c[i])
		}
	}
}
