package tempconv

// Tempconv defines Fahrenheit Celsius as types
type Fahrenheit float64
type Celsius float64

//Temperature constants
const (
	AbsoluteZeroC = -273.15
	FreezingC     = 0.0   // The temperature of freezing water under standard pressure
	BoilingC      = 100.0 // The temperature of boiling water under standard pressure
)
