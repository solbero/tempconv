package scale

import (
	"fmt"
	"math"
)

const EqualityThresholdFloat64 = 1e-12

const (
	absoluteZeroK  float64 = 0.0
	absoluteZeroC  float64 = -273.15
	absoluteZeroF  float64 = -459.67
	absoluteZeroR  float64 = 0.0
	absoluteZeroDe float64 = 559.725
	absoluteZeroN  float64 = -90.1395
	absoluteZeroRé float64 = -218.52
	absolutezeroRø float64 = -135.90375
)

// AbsoluteZeroError is an error type for temperatures below absolute zero.
var ErrAbsoluteZero = fmt.Errorf("temperature below absolute zero")

// Scale is an interface for temperature scales.
type Scale interface {
	Name() string
	Alias() string
	Temp() float64
	SetTemp(float64) error
	Unit() string
}

// NewKelvin returns a new Kelvin scale.
func NewKelvin() *Kelvin {
	return &Kelvin{baseScale{name: "kelvin", unit: "K"}}
}

// NewCelsius returns a new Celsius scale.
func NewCelsius() *Celsius {
	return &Celsius{baseScale{name: "celsius", unit: "°C"}}
}

// NewFahrenheit returns a new Fahrenheit scale.
func NewFahrenheit() *Fahrenheit {
	return &Fahrenheit{baseScale{name: "fahrenheit", unit: "°F"}}
}

// NewRankine returns a new Rankine scale.
func NewRankine() *Rankine {
	return &Rankine{baseScale{name: "rankine", unit: "°R"}}
}

// NewDelisle returns a new Delisle scale.
func NewDelisle() *Delisle {
	return &Delisle{baseScale{name: "delisle", unit: "°De"}}
}

// NewNewton returns a new Newton scale.
func NewNewton() *Newton {
	return &Newton{baseScale{name: "newton", unit: "°N"}}
}

// NewReaumur returns a new Réaumur scale.
func NewReaumur() *Reaumur {
	return &Reaumur{baseScale{name: "réaumur", alias: "reaumur", unit: "°Ré"}}
}

// NewRomer returns a new Rømer scale.
func NewRomer() *Roemer {
	return &Roemer{baseScale{name: "rømer", alias: "romer", unit: "°Rø"}}
}

type baseScale struct {
	name  string
	alias string
	temp  float64
	unit  string
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

func (b *baseScale) Alias() string {
	return b.alias
}

type Kelvin struct {
	baseScale
}

func (k *Kelvin) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroK)
	if err != nil {
		return err
	}
	k.temp = t
	return nil
}

type Celsius struct {
	baseScale
}

func (c *Celsius) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroC)
	if err != nil {
		return err
	}
	c.temp = t
	return nil
}

type Fahrenheit struct {
	baseScale
}

func (f *Fahrenheit) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroF)
	if err != nil {
		return err
	}
	f.temp = t
	return nil
}

type Rankine struct {
	baseScale
}

func (r *Rankine) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroR)
	if err != nil {
		return err
	}
	r.temp = t
	return nil
}

type Delisle struct {
	baseScale
}

func (d *Delisle) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(-t, -absoluteZeroDe) // Delisle scale is inverted
	if err != nil {
		return err
	}
	d.temp = -t // Revert inversion
	return nil
}

type Newton struct {
	baseScale
}

func (n *Newton) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroN)
	if err != nil {
		return err
	}
	n.temp = t
	return nil
}

type Reaumur struct {
	baseScale
}

func (r *Reaumur) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroRé)
	if err != nil {
		return err
	}
	r.temp = t
	return nil
}

type Roemer struct {
	baseScale
}

func (r *Roemer) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absolutezeroRø)
	if err != nil {
		return err
	}
	r.temp = t
	return nil
}

func ScaleNames() (names [][]string) {
	scales := []Scale{
		NewKelvin(),
		NewCelsius(),
		NewFahrenheit(),
		NewRankine(),
		NewDelisle(),
		NewNewton(),
		NewReaumur(),
		NewRomer(),
	}
	for _, s := range scales {
		if !(s.Alias() == "") {
			names = append(names, []string{s.Name(), s.Alias()})
		} else {
			names = append(names, []string{s.Name()})
		}
	}
	return names
}

func checkAbsoluteZero(t, zero float64) (float64, error) {
	if math.Signbit(t) != math.Signbit(zero) && math.Abs(t-zero) < EqualityThresholdFloat64 {
		return zero, nil
	} else if t < zero {
		return 0, fmt.Errorf("tempconv: %w", ErrAbsoluteZero)
	}
	return t, nil
}
