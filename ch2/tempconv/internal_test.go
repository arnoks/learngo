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
		if CToF(c[i]) != f[i] {
			t.Errorf("Failed converting from Celsius To Fahrenheit: %v != %v\n", c[i], f[i])
		}
		if FToC(f[i]) != c[i] {
			t.Errorf("Failed converting from Fahrenheit To Celsius %v != %v\n", f[i], c[i])
		}
		if CToK(c[i]) != k[i] {
			t.Errorf("Failed converting from Celsius To Kelvin %v != %v\n", c[i], k[i])
		}
		if KToC(k[i]) != c[i] {
			t.Errorf("Failed converting from Kelvin To Celsius %v != %v\n", k[i], c[i])
		}
		if FToK(f[i]) != k[i] {
			t.Errorf("Failed converting from Fahrenheit To Kelvin %v != %v\n", f[i], k[i])
		}
		if KToF(k[i]) != f[i] {
			t.Errorf("Failed converting from Fahrenheit To Kelvin %v != %v\n", k[i], f[i])
		}
	}
}
