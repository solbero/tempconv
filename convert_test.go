package tempconv

import (
	"fmt"
	"testing"
)

type convertTemp[T, S TempScales] struct {
	input  T
	output S
	want   string
}

func TestConvertKelvinToCelsius(t *testing.T) {

	kelvinToCelsiusCases := []convertTemp[*Kelvin, *Celsius]{
		{new(Kelvin).Init(0), new(Celsius), "-273.15 °C"},
		{new(Kelvin).Init(273.15), new(Celsius), "0 °C"},
		{new(Kelvin).Init(373.15), new(Celsius), "100 °C"},
	}
	assertConvert(t, kelvinToCelsiusCases)
}

func TestConvertCelsiusToKelvin(t *testing.T) {

	celsiusToKelvinCases := []convertTemp[*Celsius, *Kelvin]{
		{new(Celsius).Init(0), new(Kelvin), "273.15 K"},
		{new(Celsius).Init(-273.15), new(Kelvin), "0 K"},
		{new(Celsius).Init(100), new(Kelvin), "373.15 K"},
	}
	assertConvert(t, celsiusToKelvinCases)
}

func assertConvert[T, S Converter](t *testing.T, cases []convertTemp[T, S]) {
	t.Helper()
	for _, c := range cases {
		got := fmt.Sprint(Convert(c.input, c.output))
		want := c.want

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
