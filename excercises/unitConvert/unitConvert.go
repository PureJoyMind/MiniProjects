package unitconvert

// unitconvert is a command-line tool to convert temprature units.
// It supports all conversions for Celsius, Kelvin, and Farenheit

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	KelvinDif     float64 = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%.4g°C", c) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.4g°K", k) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.4g°F", f) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - Celsius(KelvinDif)) }

func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(KelvinDif)) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(Celsius(k - Kelvin(KelvinDif)))) }

func FToK(f Fahrenheit) Kelvin { return Kelvin(FToC(f) + Celsius(KelvinDif)) }
