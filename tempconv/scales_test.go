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
		{NewRankine(), "0 °R"},
		{NewDelisle(), "0 °De"},
		{NewNewton(), "0 °N"},
		{NewReaumur(), "0 °Ré"},
		{NewRoemer(), "0 °Rø"},
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
		{NewRankine(), "rankine"},
		{NewDelisle(), "delisle"},
		{NewNewton(), "newton"},
		{NewReaumur(), "réaumur"},
		{NewRoemer(), "rømer"},
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
		{NewRankine(), 0},
		{NewDelisle(), 0},
		{NewNewton(), 0},
		{NewReaumur(), 0},
		{NewRoemer(), 0},
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
		{NewRankine(), 100, "100 °R"},
		{NewDelisle(), 100, "100 °De"},
		{NewNewton(), 100, "100 °N"},
		{NewReaumur(), 100, "100 °Ré"},
		{NewRoemer(), 100, "100 °Rø"},
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
		{NewRankine(), "°R"},
		{NewDelisle(), "°De"},
		{NewNewton(), "°N"},
		{NewReaumur(), "°Ré"},
		{NewRoemer(), "°Rø"},
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
		{NewRankine(), absoluteZeroR - 1},
		{NewDelisle(), absoluteZeroDe + 1},
		{NewNewton(), absoluteZeroN - 1},
		{NewReaumur(), absoluteZeroRé - 1},
		{NewRoemer(), absolutezeroRø - 1},
	}

	for _, c := range cases {
		err := c.tempscale.SetTemp(c.temp)
		var target *AbsoluteZeroError

		if !errors.As(err, &target) {
			t.Errorf("got %T want %T", err, target)
		}
	}
}
