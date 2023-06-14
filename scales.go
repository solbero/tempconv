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

func (b *baseScale) Temp() float64 {
	return b.temp
}

type TempScales interface {
	*Kelvin | *Celsius
	Temp() float64
}

type Kelvin struct {
	baseScale
}

func (k *Kelvin) Init(t float64) *Kelvin {
	k.name, k.temp, k.unit = "Kelvin", t, "K"
	return k
}

func (k *Kelvin) toKelvin() float64 {
	return k.temp
}

func (k *Kelvin) fromKelvin(t float64) {
	k.Init(t)
}

type Celsius struct {
	baseScale
}

func (c *Celsius) Init(t float64) *Celsius {
	c.name, c.temp, c.unit = "Celsius", t, "°C"
	return c
}

func (c *Celsius) toKelvin() float64 {
	return c.temp + 273.15
}

func (c *Celsius) fromKelvin(t float64) {
	c.Init(t - 273.15)
}
