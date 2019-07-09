// test the converison
package tempconv

import (
	"testing"
)

func TestCFConversion(t *testing.T) {
	var c = []Celsius{FreezingC, BoilingC, 20.0, 25.}
	var k = []Kelvin{FreezingK, BoilingK, 293.15, 298.15}
	var f = []Fahrenheit{FreezingF, BoilingF, 68.0, 77.}

	for i := 0; i < len(f); i++ {
		if CtoF(c[i]) != f[i] {
			t.Errorf("Failed converting from Celsius to Fahrenheit: %v != %v\n", c[i], f[i])
		}
		if FtoC(f[i]) != c[i] {
			t.Errorf("Failed converting from Fahrenheit to Celsius %v != %v\n", f[i], c[i])
		}
		if CtoK(c[i]) != k[i] {
			t.Errorf("Failed converting from Celsius to Kelvin %v != %v\n", c[i], k[i])
		}
		if KtoC(k[i]) != c[i] {
			t.Errorf("Failed converting from Kelvin to Celsius %v != %v\n", k[i], c[i])
		}
		if FtoK(f[i]) != k[i] {
			t.Errorf("Failed converting from Fahrenheit to Kelvin %v != %v\n", f[i], k[i])
		}
		if KtoF(k[i]) != f[i] {
			t.Errorf("Failed converting from Fahrenheit to Kelvin %v != %v\n", k[i], f[i])
		}
	}
}
