package tempconv

import "fmt"

// CToF converts from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9.0/5.0 + 32.0) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// FToC converts from Fahrenheit to  Celsius
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32.0) * 5.0 / 9.0) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
