package tempconv

import (
	"fmt"
	"math"
)

const equalityThresholdFloat64 = 1e-12

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

// NewRankine returns a new Rankine scale.
func NewRankine() *rankine {
	return &rankine{baseScale{name: "rankine", unit: "°R"}}
}

// NewDelisle returns a new Delisle scale.
func NewDelisle() *delisle {
	return &delisle{baseScale{name: "delisle", unit: "°De"}}
}

// NewNewton returns a new Newton scale.
func NewNewton() *newton {
	return &newton{baseScale{name: "newton", unit: "°N"}}
}

// NewReaumur returns a new Réaumur scale.
func NewReaumur() *reaumur {
	return &reaumur{baseScale{name: "réaumur", alias: "reaumur", unit: "°Ré"}}
}

// NewRomer returns a new Rømer scale.
func NewRomer() *roemer {
	return &roemer{baseScale{name: "rømer", alias: "romer", unit: "°Rø"}}
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

type kelvin struct {
	baseScale
}

func (k *kelvin) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroK)
	if err != nil {
		return err
	}
	k.temp = t
	return nil
}

type celsius struct {
	baseScale
}

func (c *celsius) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroC)
	if err != nil {
		return err
	}
	c.temp = t
	return nil
}

type fahrenheit struct {
	baseScale
}

func (f *fahrenheit) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroF)
	if err != nil {
		return err
	}
	f.temp = t
	return nil
}

type rankine struct {
	baseScale
}

func (r *rankine) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroR)
	if err != nil {
		return err
	}
	r.temp = t
	return nil
}

type delisle struct {
	baseScale
}

func (d *delisle) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(-t, -absoluteZeroDe) // Delisle scale is inverted
	if err != nil {
		return err
	}
	d.temp = -t // Revert inversion
	return nil
}

type newton struct {
	baseScale
}

func (n *newton) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroN)
	if err != nil {
		return err
	}
	n.temp = t
	return nil
}

type reaumur struct {
	baseScale
}

func (r *reaumur) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absoluteZeroRé)
	if err != nil {
		return err
	}
	r.temp = t
	return nil
}

type roemer struct {
	baseScale
}

func (r *roemer) SetTemp(t float64) error {
	t, err := checkAbsoluteZero(t, absolutezeroRø)
	if err != nil {
		return err
	}
	r.temp = t
	return nil
}

func checkAbsoluteZero(t, zero float64) (float64, error) {
	if math.Signbit(t) != math.Signbit(zero) && math.Abs(t-zero) < equalityThresholdFloat64 {
		return zero, nil
	} else if t < zero {
		return 0, fmt.Errorf("tempconv: %w", ErrAbsoluteZero)
	}
	return t, nil
}
