package tempconv

// CToF converts from Celsius To Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// CToK converts from Celsius To Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts from Fahrenheit To  Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// FToK converts from Fahrenheit To  Celsius
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32.0) * 5.0 / 9.0) + 273.15) }

// KToC converts from Celsius To Kelvin
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts from Kelvin To Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9.0/5.0 + 32.0) }
