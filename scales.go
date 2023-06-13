package tempconv

import "fmt"

type baseScale struct {
	name string
	temp float64
	unit string
}

func (b baseScale) String() string {
	return fmt.Sprintf("%g %v", b.temp, b.unit)
}

type TempScales interface {
	*Kelvin | *Celsius
}

type Kelvin struct {
	baseScale
}

func (k *Kelvin) Init(t float64) *Kelvin {
	k.name, k.temp, k.unit = "Kelvin", t, "K"
	return k
}

type Celsius struct {
	baseScale
}

func (c *Celsius) Init(t float64) *Celsius {
	c.name, c.temp, c.unit = "Celsius", t, "°C"
	return c
}
