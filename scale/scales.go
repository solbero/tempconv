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

const (
	KELVIN = iota
	CELSIUS
	FAHRENHEIT
	RANKINE
	DELISLE
	NEWTON
	REAUMUR
	ROMER
)

// AbsoluteZeroError is an error type for temperatures below absolute zero.
var ErrAbsoluteZero = fmt.Errorf("temperature below absolute zero")

// NewKelvin returns a new Kelvin scale.
func NewKelvin() *Scale {
	return &Scale{Type: KELVIN, Name: "kelvin", Unit: "K"}
}

// NewCelsius returns a new Celsius scale.
func NewCelsius() *Scale {
	return &Scale{Type: CELSIUS, Name: "celsius", Unit: "°C"}
}

// NewFahrenheit returns a new Fahrenheit scale.
func NewFahrenheit() *Scale {
	return &Scale{Type: FAHRENHEIT, Name: "fahrenheit", Unit: "°F"}
}

// NewRankine returns a new Rankine scale.
func NewRankine() *Scale {
	return &Scale{Type: RANKINE, Name: "rankine", Unit: "°R"}
}

// NewDelisle returns a new Delisle scale.
func NewDelisle() *Scale {
	return &Scale{Type: DELISLE, Name: "delisle", Unit: "°De"}
}

// NewNewton returns a new Newton scale.
func NewNewton() *Scale {
	return &Scale{Type: NEWTON, Name: "newton", Unit: "°N"}
}

// NewReaumur returns a new Réaumur scale.
func NewReaumur() *Scale {
	return &Scale{Type: REAUMUR, Name: "réaumur", Alias: "reaumur", Unit: "°Ré"}
}

// NewRomer returns a new Rømer scale.
func NewRomer() *Scale {
	return &Scale{Type: ROMER, Name: "rømer", Alias: "romer", Unit: "°Rø"}
}

type Scale struct {
	Type  int
	Name  string
	Alias string
	temp  float64
	Unit  string
}

func (b Scale) String() string { return fmt.Sprintf("%g %v", b.temp, b.Unit) }
func (b *Scale) Temp() float64 { return b.temp }
func (b *Scale) SetTemp(t float64) (err error) {
	switch b.Type {
	case KELVIN:
		t, err = checkAbsoluteZero(t, absoluteZeroK)
	case CELSIUS:
		t, err = checkAbsoluteZero(t, absoluteZeroC)
	case FAHRENHEIT:
		t, err = checkAbsoluteZero(t, absoluteZeroF)
	case RANKINE:
		t, err = checkAbsoluteZero(t, absoluteZeroR)
	case DELISLE:
		t, err = checkAbsoluteZero(-t, -absoluteZeroDe)
		t = -t // Delisle scale is inverted
	case NEWTON:
		t, err = checkAbsoluteZero(t, absoluteZeroN)
	case REAUMUR:
		t, err = checkAbsoluteZero(t, absoluteZeroRé)
	case ROMER:
		t, err = checkAbsoluteZero(t, absolutezeroRø)
	}

	if err != nil {
		return err
	}

	b.temp = t
	return nil
}

func ScaleNames() (names [][]string) {
	scales := []*Scale{
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
		if !(s.Alias == "") {
			names = append(names, []string{s.Name, s.Alias})
		} else {
			names = append(names, []string{s.Name})
		}
	}

	return names
}

func checkAbsoluteZero(t, absoluteZero float64) (float64, error) {
	if math.Signbit(t) != math.Signbit(absoluteZero) && math.Abs(t-absoluteZero) < EqualityThresholdFloat64 {
		return absoluteZero, nil
	} else if t < absoluteZero {
		return 0, fmt.Errorf("tempconv: %w", ErrAbsoluteZero)
	}

	return t, nil
}
