package tempconv

import "fmt"

// Kelvin defines °Klevin as a float
type Kelvin float64

// Kelvin.String prints a float as °Kelvin
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// Celsius defines °Celsius as a float
type Celsius float64

// Celsius.String prints a float as °Celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Fahrenheit defines °Fahrenheit as float
type Fahrenheit float64

// Fahrenheit.String prints a float as °Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// Temperature constants
const (
	// Absolute denotes absolute zero temperatures in 3 scalas
	AbsoluteZeroK = 0.0
	AbsoluteZeroC = -273.15
	AbsoluteZeroF = -523.67

	// The temperature of freezing water under standard pressure
	FreezingK = 273.15
	FreezingC = 0.0
	FreezingF = 32.0

	// Absolute denotes the boiling temperature of water at 1atm for 3 scales
	BoilingK = 373.15
	BoilingC = 100.0 // The temperature of boiling water under standard pressure
	BoilingF = 212.0
)
