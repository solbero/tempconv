package tempconv

import (
	"errors"
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	cases := []struct {
		tempscale TempScale
		want      string
	}{
		{NewKelvin(), "0 K"},
		{NewCelsius(), "0 °C"},
		{NewFahrenheit(), "0 °F"},
	}

	for _, c := range cases {
		got := fmt.Sprint(c.tempscale)

		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestName(t *testing.T) {
	cases := []struct {
		tempscale TempScale
		want      string
	}{
		{NewKelvin(), "kelvin"},
		{NewCelsius(), "celsius"},
		{NewFahrenheit(), "fahrenheit"},
	}

	for _, c := range cases {
		got := c.tempscale.Name()

		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestTemp(t *testing.T) {
	cases := []struct {
		tempscale TempScale
		want      float64
	}{
		{NewKelvin(), 0},
		{NewCelsius(), 0},
		{NewFahrenheit(), 0},
	}

	for _, c := range cases {
		got := c.tempscale.Temp()

		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetTemp(t *testing.T) {
	cases := []struct {
		tempscale TempScale
		temp      float64
		want      string
	}{
		{NewKelvin(), 100, "100 K"},
		{NewCelsius(), 100, "100 °C"},
		{NewFahrenheit(), 100, "100 °F"},
	}

	for _, c := range cases {
		err := c.tempscale.SetTemp(c.temp)
		if err != nil {
			t.Errorf("got %v want nil", err)
		}

		got := fmt.Sprint(c.tempscale)

		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestUnit(t *testing.T) {
	cases := []struct {
		tempscale TempScale
		want      string
	}{
		{NewKelvin(), "K"},
		{NewCelsius(), "°C"},
		{NewFahrenheit(), "°F"},
	}

	for _, c := range cases {
		got := c.tempscale.Unit()

		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestAbsoluteZeroError(t *testing.T) {
	cases := []struct {
		tempscale TempScale
		temp      float64
	}{
		{NewKelvin(), absoluteZeroK - 1},
		{NewCelsius(), absoluteZeroC - 1},
		{NewFahrenheit(), absoluteZeroF - 1},
	}

	for _, c := range cases {
		err := c.tempscale.SetTemp(c.temp)
		var target *AbsoluteZeroError

		if !errors.As(err, &target) {
			t.Errorf("got %T want %T", err, target)
		}
	}
}
