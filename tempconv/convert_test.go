package tempconv

import (
	"fmt"
	"math"
	"testing"
)

type conversionCases struct {
	input  Scale
	temp   float64
	output Scale
	want   float64
}

func TestTemperatureConversion(t *testing.T) {
	cases := []conversionCases{
		{NewKelvin(), 0, NewCelsius(), -273.15},
		{NewKelvin(), 273.15, NewCelsius(), 0.0},
		{NewCelsius(), 0, NewKelvin(), 273.15},
		{NewCelsius(), -273.15, NewKelvin(), 0.0},
		{NewKelvin(), 0, NewFahrenheit(), -459.67},
		{NewKelvin(), 255.3722222222222, NewFahrenheit(), 0},
		{NewFahrenheit(), 0, NewKelvin(), 255.3722222222222},
		{NewFahrenheit(), -459.67, NewKelvin(), 0},
		{NewKelvin(), 0, NewRankine(), 0},
		{NewKelvin(), 273.15, NewRankine(), 491.67},
		{NewRankine(), 0, NewKelvin(), 0},
		{NewRankine(), 491.67, NewKelvin(), 273.15},
		{NewKelvin(), 0, NewDelisle(), 559.725},
		{NewKelvin(), 373.15, NewDelisle(), 0},
		{NewDelisle(), 0, NewKelvin(), 373.15},
		{NewDelisle(), 559.725, NewKelvin(), 0},
		{NewKelvin(), 0, NewNewton(), -90.1395},
		{NewKelvin(), 273.15, NewNewton(), 0},
		{NewNewton(), 0, NewKelvin(), 273.15},
		{NewNewton(), -90.1395, NewKelvin(), 0},
		{NewKelvin(), 0, NewReaumur(), -218.52},
		{NewKelvin(), 273.15, NewReaumur(), 0},
		{NewReaumur(), 0, NewKelvin(), 273.15},
		{NewReaumur(), -218.52, NewKelvin(), 0},
		{NewKelvin(), 0, NewRomer(), -135.90375},
		{NewKelvin(), 258.8642857142857, NewRomer(), 0},
		{NewRomer(), 0, NewKelvin(), 258.8642857142857},
		{NewRomer(), -135.90375, NewKelvin(), 0},
	}

	assertConversion(t, cases)
}

func assertConversion(t *testing.T, cases []conversionCases) {
	t.Helper()
	for _, c := range cases {
		msg := fmt.Sprintf("%g %v -> %g %v", c.temp, c.input.Name(), c.want, c.output.Name())
		t.Run(msg, func(t *testing.T) {
			err := c.input.SetTemp(c.temp)
			if err != nil {
				t.Errorf("got %v want %v", err, nil)
			}

			err = Convert(c.input, c.output)
			if err != nil {
				t.Errorf("got %v want %v", err, nil)
			}

			got := c.output.Temp()

			if !assertAlmostEqual(got, c.want, equalityThresholdFloat64) {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
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
