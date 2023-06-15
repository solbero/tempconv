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

func (b *baseScale) Temp() (float64, error) {
	if b.name == "" || b.unit == "" {
		return 0, fmt.Errorf("tempconv: %T is not initalized", b)
	}
	return b.temp, nil
}

type TempScales interface {
	*Kelvin | *Celsius | *Fahrenheit
	Temp() (float64, error)
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

type Fahrenheit struct {
	baseScale
}

func (f *Fahrenheit) Init(t float64) *Fahrenheit {
	f.name, f.temp, f.unit = "Fahrenheit", t, "°F"
	return f
}

func (f *Fahrenheit) toKelvin() float64 {
	return (f.temp + 459.67) * 5 / 9
}

func (f *Fahrenheit) fromKelvin(t float64) {
	f.Init((t*9 - 459.67*5) / 5)
}
