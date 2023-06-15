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
		{mustInit(new(Kelvin).Init(0)), "0 K"},
		{mustInit(new(Kelvin).Init(0.0)), "0 K"},
		{mustInit(new(Kelvin).Init(.0)), "0 K"},
	}
	assertString(t, kelvinCases)

	celsiusCases := []stringTemp[*Celsius]{
		{mustInit(new(Celsius).Init(0)), "0 °C"},
		{mustInit(new(Celsius).Init(0.0)), "0 °C"},
		{mustInit(new(Celsius).Init(.0)), "0 °C"},
	}
	assertString(t, celsiusCases)

	fahrenheitCases := []stringTemp[*Fahrenheit]{
		{mustInit(new(Fahrenheit).Init(0)), "0 °F"},
		{mustInit(new(Fahrenheit).Init(0.0)), "0 °F"},
		{mustInit(new(Fahrenheit).Init(.0)), "0 °F"},
	}
	assertString(t, fahrenheitCases)
}

func TestInitError(t *testing.T) {
	k := new(Kelvin)
	c := new(Celsius)
	f := new(Fahrenheit)

	if _, err := k.Temp(); err == nil {
		t.Errorf("got %v want error", err)
	}
	if _, err := c.Temp(); err == nil {
		t.Errorf("got %v want error", err)
	}
	if _, err := f.Temp(); err == nil {
		t.Errorf("got %v want error", err)
	}
}

func TestAbsoluteZeroError(t *testing.T) {
	k := new(Kelvin)
	c := new(Celsius)
	f := new(Fahrenheit)

	if _, err := k.Init(absoluteZeroK - 1); err == nil {
		t.Errorf("got %v want error", err)
	}
	if _, err := c.Init(absoluteZeroC - 1); err == nil {
		t.Errorf("got %v want error", err)
	}
	if _, err := f.Init(absoluteZeroF - 1); err == nil {
		t.Errorf("got %v want error", err)
	}
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

func mustInit[T TempScales](s T, err error) T {
	if err != nil {
		panic(err)
	}
	return s
}
