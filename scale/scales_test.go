package scale

import (
	"errors"
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	cases := []struct {
		scale *Scale
		name  string
		alias string
		temp  float64
		unit  string
	}{
		{NewKelvin(), "kelvin", "", 0, "K"},
		{NewCelsius(), "celsius", "", 0, "°C"},
		{NewFahrenheit(), "fahrenheit", "", 0, "°F"},
		{NewRankine(), "rankine", "", 0, "°R"},
		{NewDelisle(), "delisle", "", 0, "°De"},
		{NewNewton(), "newton", "", 0, "°N"},
		{NewReaumur(), "réaumur", "reaumur", 0, "°Ré"},
		{NewRomer(), "rømer", "romer", 0, "°Rø"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.scale.Name), func(t *testing.T) {
			if c.scale.Name != c.name {
				t.Errorf("got %v want %v", c.scale.Name, c.name)
			}
			if c.scale.Alias != c.alias {
				t.Errorf("got %v want %v", c.scale.Alias, c.alias)
			}
			if c.scale.temp != c.temp {
				t.Errorf("got %v want %v", c.scale.temp, c.temp)
			}
			if c.scale.Unit != c.unit {
				t.Errorf("got %v want %v", c.scale.Unit, c.unit)
			}
		})
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		scale *Scale
		want  string
	}{
		{NewKelvin(), "0 K"},
		{NewCelsius(), "0 °C"},
		{NewFahrenheit(), "0 °F"},
		{NewRankine(), "0 °R"},
		{NewDelisle(), "0 °De"},
		{NewNewton(), "0 °N"},
		{NewReaumur(), "0 °Ré"},
		{NewRomer(), "0 °Rø"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.scale.Name), func(t *testing.T) {
			got := fmt.Sprint(c.scale)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	cases := []struct {
		scale *Scale
		want  string
	}{
		{NewKelvin(), "kelvin"},
		{NewCelsius(), "celsius"},
		{NewFahrenheit(), "fahrenheit"},
		{NewRankine(), "rankine"},
		{NewDelisle(), "delisle"},
		{NewNewton(), "newton"},
		{NewReaumur(), "réaumur"},
		{NewRomer(), "rømer"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.scale.Name), func(t *testing.T) {
			got := c.scale.Name

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestTemp(t *testing.T) {
	cases := []struct {
		scale *Scale
		want  float64
	}{
		{NewKelvin(), 0},
		{NewCelsius(), 0},
		{NewFahrenheit(), 0},
		{NewRankine(), 0},
		{NewDelisle(), 0},
		{NewNewton(), 0},
		{NewReaumur(), 0},
		{NewRomer(), 0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.scale.Name), func(t *testing.T) {
			got := c.scale.Temp()

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestSetTemp(t *testing.T) {
	cases := []struct {
		scale *Scale
		temp  float64
	}{
		{NewKelvin(), 100},
		{NewCelsius(), 100},
		{NewFahrenheit(), 100},
		{NewRankine(), 100},
		{NewDelisle(), 100},
		{NewNewton(), 100},
		{NewReaumur(), 100},
		{NewRomer(), 100},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.scale.Name), func(t *testing.T) {
			err := c.scale.SetTemp(c.temp)
			if err != nil {
				t.Errorf("got %v want nil", err)
			}

			got := c.scale.Temp()
			want := c.temp

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestUnit(t *testing.T) {
	cases := []struct {
		tempscale *Scale
		want      string
	}{
		{NewKelvin(), "K"},
		{NewCelsius(), "°C"},
		{NewFahrenheit(), "°F"},
		{NewRankine(), "°R"},
		{NewDelisle(), "°De"},
		{NewNewton(), "°N"},
		{NewReaumur(), "°Ré"},
		{NewRomer(), "°Rø"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.tempscale.Name), func(t *testing.T) {
			got := c.tempscale.Unit

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestAbsoluteZeroError(t *testing.T) {
	cases := []struct {
		tempscale *Scale
		temp      float64
	}{
		{NewKelvin(), absoluteZeroK - 1},
		{NewCelsius(), absoluteZeroC - 1},
		{NewFahrenheit(), absoluteZeroF - 1},
		{NewRankine(), absoluteZeroR - 1},
		{NewDelisle(), absoluteZeroDe + 1},
		{NewNewton(), absoluteZeroN - 1},
		{NewReaumur(), absoluteZeroRé - 1},
		{NewRomer(), absolutezeroRø - 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.tempscale.Name), func(t *testing.T) {
			err := c.tempscale.SetTemp(c.temp)
			target := ErrAbsoluteZero

			if !errors.Is(err, target) {
				t.Errorf("got %T want %T", err, target)
			}
		})
	}
}
