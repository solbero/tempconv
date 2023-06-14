package tempconv

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

type convertTemp[T, S TempScales] struct {
	input  T
	output S
	want   float64
}

func TestConvertKelvinToCelsius(t *testing.T) {
	kelvinToCelsiusCases := []convertTemp[*Kelvin, *Celsius]{
		{new(Kelvin).Init(0), new(Celsius), -273.15},
		{new(Kelvin).Init(273.15), new(Celsius), 0.0},
	}
	assertConvert(t, kelvinToCelsiusCases)
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	celsiusToKelvinCases := []convertTemp[*Celsius, *Kelvin]{
		{new(Celsius).Init(0), new(Kelvin), 273.15},
		{new(Celsius).Init(-273.15), new(Kelvin), 0.0},
	}
	assertConvert(t, celsiusToKelvinCases)
}

func TestConvertKelvinToFahrenheit(t *testing.T) {
	kelvinToFahrenheitCases := []convertTemp[*Kelvin, *Fahrenheit]{
		{new(Kelvin).Init(0), new(Fahrenheit), -459.67},
		{new(Kelvin).Init(255.3722222222222), new(Fahrenheit), 0},
	}
	assertConvert(t, kelvinToFahrenheitCases)
}

func TestConvertFahrenheitToKelvin(t *testing.T) {
	fahrenheitToKelvinCases := []convertTemp[*Fahrenheit, *Kelvin]{
		{new(Fahrenheit).Init(0), new(Kelvin), 255.3722222222222},
		{new(Fahrenheit).Init(-459.67), new(Kelvin), 0},
	}
	assertConvert(t, fahrenheitToKelvinCases)
}

func assertConvert[T, S Converter](t *testing.T, cases []convertTemp[T, S]) {
	t.Helper()
	for _, c := range cases {
		out := Convert(c.input, c.output)
		got := out.Temp()
		want := c.want

		if !assertAlmostEqual(got, want, float64EqualityThreshold) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func assertAlmostEqual(got, want, epsilon float64) bool {
	sum := math.Abs(got + want)
	diff := math.Abs(got - want)

	if got == want {
		return true
	} else if want == 0 || got == 0 || sum < math.SmallestNonzeroFloat64 {
		return diff < epsilon*math.SmallestNonzeroFloat64
	}

	return diff/math.Min(sum, math.MaxFloat64) < epsilon
}