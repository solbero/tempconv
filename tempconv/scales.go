package tempconv

import "fmt"

const (
	absoluteZeroK float64 = 0.0
	absoluteZeroC float64 = -273.15
	absoluteZeroF float64 = -459.67
)

// AbsoluteZeroError is an error type for temperatures below absolute zero.
type AbsoluteZeroError struct {
}

func (e AbsoluteZeroError) Error() string {
	return "temperature is below absolute zero"
}

// TempScale is an interface for temperature scales.
type TempScale interface {
	Name() string
	Temp() float64
	SetTemp(float64) error
	Unit() string
}

// NewKelvin returns a new Kelvin scale.
func NewKelvin() *kelvin {
	return &kelvin{baseScale{name: "kelvin", unit: "K"}}
}

// NewCelsius returns a new Celsius scale.
func NewCelsius() *celsius {
	return &celsius{baseScale{name: "celsius", unit: "°C"}}
}

// NewFahrenheit returns a new Fahrenheit scale.
func NewFahrenheit() *fahrenheit {
	return &fahrenheit{baseScale{name: "fahrenheit", unit: "°F"}}
}

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

func (b *baseScale) Unit() string {
	return b.unit
}

func (b *baseScale) Name() string {
	return b.name
}

type kelvin struct {
	baseScale
}

func (k *kelvin) SetTemp(t float64) error {
	if t < absoluteZeroK {
		return fmt.Errorf("tempconv: %w", &AbsoluteZeroError{})
	}
	k.temp = t
	return nil
}

type celsius struct {
	baseScale
}

func (c *celsius) SetTemp(t float64) error {
	if t < absoluteZeroC {
		return fmt.Errorf("tempconv: %w", &AbsoluteZeroError{})
	}
	c.temp = t
	return nil
}

type fahrenheit struct {
	baseScale
}

func (f *fahrenheit) SetTemp(t float64) error {
	if t < absoluteZeroF {
		return fmt.Errorf("tempconv: %w", &AbsoluteZeroError{})
	}
	f.temp = t
	return nil
}
