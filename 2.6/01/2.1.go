package main

import (
	. "GOPLProblemSets/2.6/00/tempconv"
	"fmt"
)

type Kelvin float64

func (k Kelvin) String() string { return fmt.Sprintf("%g K", k) }

func CToK(c Celsius) Kelvin    { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius    { return Celsius(k - 273.15) }
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func main() {
	var (
		c Celsius    = AbsoluteZeroC
		f Fahrenheit = 2333
		k Kelvin     = 233
	)

	fmt.Println(CToK(c))
	fmt.Println(KToC(k))
	fmt.Println(KToF(k))
	fmt.Println(FToK(f))
}
