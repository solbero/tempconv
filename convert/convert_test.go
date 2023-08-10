package convert

import (
	"fmt"
	"math"
	"testing"

	"github.com/solbero/tempconv/scale"
)

type conversionCases struct {
	input  scale.Scale
	temp   float64
	output scale.Scale
	want   float64
}

func TestTemperatureConversion(t *testing.T) {
	cases := []conversionCases{
		{scale.NewKelvin(), 0, scale.NewCelsius(), -273.15},
		{scale.NewKelvin(), 273.15, scale.NewCelsius(), 0.0},
		{scale.NewCelsius(), 0, scale.NewKelvin(), 273.15},
		{scale.NewCelsius(), -273.15, scale.NewKelvin(), 0.0},
		{scale.NewKelvin(), 0, scale.NewFahrenheit(), -459.67},
		{scale.NewKelvin(), 255.3722222222222, scale.NewFahrenheit(), 0},
		{scale.NewFahrenheit(), 0, scale.NewKelvin(), 255.3722222222222},
		{scale.NewFahrenheit(), -459.67, scale.NewKelvin(), 0},
		{scale.NewKelvin(), 0, scale.NewRankine(), 0},
		{scale.NewKelvin(), 273.15, scale.NewRankine(), 491.67},
		{scale.NewRankine(), 0, scale.NewKelvin(), 0},
		{scale.NewRankine(), 491.67, scale.NewKelvin(), 273.15},
		{scale.NewKelvin(), 0, scale.NewDelisle(), 559.725},
		{scale.NewKelvin(), 373.15, scale.NewDelisle(), 0},
		{scale.NewDelisle(), 0, scale.NewKelvin(), 373.15},
		{scale.NewDelisle(), 559.725, scale.NewKelvin(), 0},
		{scale.NewKelvin(), 0, scale.NewNewton(), -90.1395},
		{scale.NewKelvin(), 273.15, scale.NewNewton(), 0},
		{scale.NewNewton(), 0, scale.NewKelvin(), 273.15},
		{scale.NewNewton(), -90.1395, scale.NewKelvin(), 0},
		{scale.NewKelvin(), 0, scale.NewReaumur(), -218.52},
		{scale.NewKelvin(), 273.15, scale.NewReaumur(), 0},
		{scale.NewReaumur(), 0, scale.NewKelvin(), 273.15},
		{scale.NewReaumur(), -218.52, scale.NewKelvin(), 0},
		{scale.NewKelvin(), 0, scale.NewRomer(), -135.90375},
		{scale.NewKelvin(), 258.8642857142857, scale.NewRomer(), 0},
		{scale.NewRomer(), 0, scale.NewKelvin(), 258.8642857142857},
		{scale.NewRomer(), -135.90375, scale.NewKelvin(), 0},
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
				t.Fatalf("%v", err)
			}

			err = Convert(c.input, c.output)
			if err != nil {
				t.Fatalf("%v", err)
			}

			got := c.output.Temp()

			if !assertAlmostEqual(got, c.want, scale.EqualityThresholdFloat64) {
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
