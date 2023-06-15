package tempconv

import "fmt"

const (
	absoluteZeroK float64 = 0.0
	absoluteZeroC float64 = -273.15
	absoluteZeroF float64 = -459.67
)

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

func (k *Kelvin) Init(t float64) (*Kelvin, error) {
	n := "Kelvin"
	u := "K"
	if t < absoluteZeroK {
		return nil, absoluteZeroError(t, absoluteZeroK)
	}
	k.name, k.temp, k.unit = n, t, u
	return k, nil
}

func (k *Kelvin) toKelvin() float64 {
	return k.temp
}

func (k *Kelvin) fromKelvin(t float64) (*Kelvin, error) {
	return k.Init(t)
}

type Celsius struct {
	baseScale
}

func (c *Celsius) Init(t float64) (*Celsius, error) {
	n := "Celsius"
	u := "°C"
	if t < absoluteZeroC {
		return nil, absoluteZeroError(t, absoluteZeroC)
	}
	c.name, c.temp, c.unit = n, t, u
	return c, nil
}

func (c *Celsius) toKelvin() float64 {
	return c.temp + 273.15
}

func (c *Celsius) fromKelvin(t float64) (*Celsius, error) {
	return c.Init(t - 273.15)
}

type Fahrenheit struct {
	baseScale
}

func (f *Fahrenheit) Init(t float64) (*Fahrenheit, error) {
	n := "Fahrenheit"
	u := "°F"
	if t < absoluteZeroF {
		return nil, absoluteZeroError(t, absoluteZeroF)
	}
	f.name, f.temp, f.unit = n, t, u
	return f, nil
}

func (f *Fahrenheit) toKelvin() float64 {
	return (f.temp + 459.67) * 5 / 9
}

func (f *Fahrenheit) fromKelvin(t float64) (*Fahrenheit, error) {
	return f.Init((t*9 - 459.67*5) / 5)
}

func absoluteZeroError(temp, zero float64) error {
	return fmt.Errorf("tempconv: input temperature %g is less than absolute zero %g", temp, absoluteZeroC)
}
