package tempconv

// CtoF converts from Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// CtoK converts from Celsius to Kelvin
func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FtoC converts from Fahrenheit to  Celsius
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// FtoK converts from Fahrenheit to  Celsius
func FtoK(f Fahrenheit) Kelvin { return Kelvin(((f - 32.0) * 5.0 / 9.0) + 273.15) }

// KtoC converts from Celsius to Kelvin
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KtoF converts from Kelvin to Fahrenheit
func KtoF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9.0/5.0 + 32.0) }
