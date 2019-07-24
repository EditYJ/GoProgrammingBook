package tempconv

import "fmt"

type Celsius float64
type Fahreheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	Fdegree Fahreheit = 100

	// 绝对零度
	KelvinZero Kelvin = 273.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}
func (f Fahreheit) String() string {
	return fmt.Sprintf("%g °F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g °K", k)
}
