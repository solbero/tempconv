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

func TestKelvinToRankine(t *testing.T) {
	kelvinToRankineCases := []conversionCases{
		{NewKelvin(), 0, NewRankine(), 0},
		{NewKelvin(), 273.15, NewRankine(), 491.67},
	}
	assertConversion(t, kelvinToRankineCases)
}

func TestRankineToKelvin(t *testing.T) {
	rankineToKelvinCases := []conversionCases{
		{NewRankine(), 0, NewKelvin(), 0},
		{NewRankine(), 491.67, NewKelvin(), 273.15},
	}
	assertConversion(t, rankineToKelvinCases)
}

func TestKelvinToDelisle(t *testing.T) {
	kelvinToDelisleCases := []conversionCases{
		{NewKelvin(), 0, NewDelisle(), 559.725},
		{NewKelvin(), 373.15, NewDelisle(), 0},
	}
	assertConversion(t, kelvinToDelisleCases)
}

func TestDelisleToKelvin(t *testing.T) {
	delisleToKelvinCases := []conversionCases{
		{NewDelisle(), 0, NewKelvin(), 373.15},
		{NewDelisle(), 559.725, NewKelvin(), 0},
	}
	assertConversion(t, delisleToKelvinCases)
}

func TestKelvinToNewton(t *testing.T) {
	kelvinToNewtonCases := []conversionCases{
		{NewKelvin(), 0, NewNewton(), -90.1395},
		{NewKelvin(), 273.15, NewNewton(), 0},
	}
	assertConversion(t, kelvinToNewtonCases)
}

func TestNewtonToKelvin(t *testing.T) {
	newtonToKelvinCases := []conversionCases{
		{NewNewton(), 0, NewKelvin(), 273.15},
		{NewNewton(), -90.1395, NewKelvin(), 0},
	}
	assertConversion(t, newtonToKelvinCases)
}

func TestKelvinToReaumur(t *testing.T) {
	kelvinToReaumurCases := []conversionCases{
		{NewKelvin(), 0, NewReaumur(), -218.52},
		{NewKelvin(), 273.15, NewReaumur(), 0},
	}
	assertConversion(t, kelvinToReaumurCases)
}

func TestReaumurToKelvin(t *testing.T) {
	reaumurToKelvinCases := []conversionCases{
		{NewReaumur(), 0, NewKelvin(), 273.15},
		{NewReaumur(), -218.52, NewKelvin(), 0},
	}
	assertConversion(t, reaumurToKelvinCases)
}

func TestKelvinToRoemer(t *testing.T) {
	kelvinToRoemerCases := []conversionCases{
		{NewKelvin(), 0, NewRoemer(), -135.90375},
		{NewKelvin(), 258.8642857142857, NewRoemer(), 0},
	}
	assertConversion(t, kelvinToRoemerCases)
}

func TestRoemerToKelvin(t *testing.T) {
	roemerToKelvinCases := []conversionCases{
		{NewRoemer(), 0, NewKelvin(), 258.8642857142857},
		{NewRoemer(), -135.90375, NewKelvin(), 0},
	}
	assertConversion(t, roemerToKelvinCases)
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
