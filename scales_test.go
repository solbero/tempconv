package tempconv

import (
	"fmt"
	"testing"
)

type stringTemp[T TempScales] struct {
	temp T
	want string
}

func TestTempString(t *testing.T) {
	kelvinCases := []stringTemp[*Kelvin]{
		{new(Kelvin).Init(0), "0 K"},
		{new(Kelvin).Init(0.0), "0 K"},
		{new(Kelvin).Init(.0), "0 K"},
	}
	assertString(t, kelvinCases)

	celsiusCases := []stringTemp[*Celsius]{
		{new(Celsius).Init(0), "0 °C"},
		{new(Celsius).Init(0.0), "0 °C"},
		{new(Celsius).Init(.0), "0 °C"},
	}
	assertString(t, celsiusCases)

	fahrenheitCases := []stringTemp[*Fahrenheit]{
		{new(Fahrenheit).Init(0), "0 °F"},
		{new(Fahrenheit).Init(0.0), "0 °F"},
		{new(Fahrenheit).Init(.0), "0 °F"},
	}
	assertString(t, fahrenheitCases)
}

func assertString[T TempScales](t *testing.T, cases []stringTemp[T]) {
	t.Helper()
	for _, c := range cases {
		got := fmt.Sprint(c.temp)
		want := c.want

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
