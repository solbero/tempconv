package tempconv

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

type conversionCases struct {
	input  TempScale
	temp   float64
	output TempScale
	want   float64
}

func TestKelvinToCelsius(t *testing.T) {
	cases := []conversionCases{
		{NewKelvin(), 0, NewCelsius(), -273.15},
		{NewKelvin(), 273.15, NewCelsius(), 0.0},
	}
	assertConversion(t, cases)
}

func TestCelsiusToKelvin(t *testing.T) {
	celsiusToKelvinCases := []conversionCases{
		{NewCelsius(), 0, NewKelvin(), 273.15},
		{NewCelsius(), -273.15, NewKelvin(), 0.0},
	}
	assertConversion(t, celsiusToKelvinCases)
}

func TestKelvinToFahrenheit(t *testing.T) {
	kelvinToFahrenheitCases := []conversionCases{
		{NewKelvin(), 0, NewFahrenheit(), -459.67},
		{NewKelvin(), 255.3722222222222, NewFahrenheit(), 0},
	}
	assertConversion(t, kelvinToFahrenheitCases)
}

func TestFahrenheitToKelvin(t *testing.T) {
	fahrenheitToKelvinCases := []conversionCases{
		{NewFahrenheit(), 0, NewKelvin(), 255.3722222222222},
		{NewFahrenheit(), -459.67, NewKelvin(), 0},
	}
	assertConversion(t, fahrenheitToKelvinCases)
}

func assertConversion(t *testing.T, cases []conversionCases) {
	t.Helper()
	for _, c := range cases {
		err := c.input.SetTemp(c.temp)
		if err != nil {
			t.Errorf("got %v want %v", err, nil)
		}

		err = Convert(c.input, c.output)
		if err != nil {
			t.Errorf("got %v want %v", err, nil)
		}

		got := c.output.Temp()

		if !assertAlmostEqual(got, c.want, float64EqualityThreshold) {
			t.Errorf("got %v want %v", got, c.want)
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
