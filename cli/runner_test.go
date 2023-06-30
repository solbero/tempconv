package cli

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"testing"

	"github.com/solbero/tempconv/tempconv"
)

func TestRun(t *testing.T) {
	var cases = []struct {
		config *config
		want   string
	}{
		{&config{temp: 0, fromScale: tempconv.NewCelsius(), toScale: tempconv.NewFahrenheit(), decimal: 0}, "32"},
		{&config{temp: 0, fromScale: tempconv.NewCelsius(), toScale: tempconv.NewFahrenheit(), decimal: 2}, "32.00"},
		{&config{temp: 0, fromScale: tempconv.NewCelsius(), toScale: tempconv.NewFahrenheit(), decimal: 2, unit: true}, "32.00 °F"},
		{&config{temp: -273.15, fromScale: tempconv.NewCelsius(), toScale: tempconv.NewKelvin(), decimal: 2}, "0.00"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v->%s", *c.config, c.want), func(t *testing.T) {
			w := new(bytes.Buffer)
			flags := flag.NewFlagSet("test", flag.ContinueOnError)
			err := Run(w, c.config, flags, flags.Name())
			if err != nil {
				t.Errorf("got %v want %v", err, nil)
			}
			if w.String() != c.want {
				t.Errorf("got %v want %v", w.String(), c.want)
			}
		})
	}

}

func TestRunHelp(t *testing.T) {
	conf := &config{help: true}

	w := new(bytes.Buffer)
	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	err := Run(w, conf, flags, flags.Name())
	if err != nil {
		t.Errorf("got %v want %v", err, nil)
	}
}

func TestRunVersion(t *testing.T) {
	conf := &config{version: true}

	w := new(bytes.Buffer)
	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	err := Run(w, conf, flags, flags.Name())

	if err != nil {
		t.Errorf("got %v want %v", err, nil)
	}

	if w.String() != flags.Name() {
		t.Errorf("got %v want %v", w.String(), flags.Name())
	}
}

func TestRunError(t *testing.T) {
	var cases = []struct {
		name   string
		config *config
		want   error
	}{
		{"absolute zero error",
			&config{temp: -300, fromScale: tempconv.NewCelsius(), toScale: tempconv.NewKelvin(), decimal: 2},
			tempconv.ErrAbsoluteZero},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := new(bytes.Buffer)
			flags := flag.NewFlagSet("test", flag.ContinueOnError)
			err := Run(w, c.config, flags, flags.Name())
			if err == nil {
				t.Errorf("got %v want error", err)
			}
			if !errors.Is(err, c.want) {
				t.Errorf("got %v want %v", err, c.want)
			}
		})
	}
}
