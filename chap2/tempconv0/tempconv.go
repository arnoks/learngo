// tempconv prives funtions for the confervsion between °F and °C
package tempconv

type Fahrenheit float64
type Celsius float64

const (
	AbsoluteZeroC = -273.15
	FreezingC     = 0.0
	BoilingC      = 100.0
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
