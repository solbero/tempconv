package cli

import (
	"bytes"
	"flag"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/solbero/tempconv/scale"
)

func TestParseArgs(t *testing.T) {
	var cases = []struct {
		args []string
		conf *config
	}{
		{[]string{"0", "celsius", "kelvin"},
			&config{temp: 0, input: scale.NewCelsius(), output: scale.NewFahrenheit(), decimal: 2}},
		{[]string{"--", "-0", "celsius", "kelvin"},
			&config{temp: 0, input: scale.NewCelsius(), output: scale.NewFahrenheit(), decimal: 2}},
		{[]string{"0", "c", "k"},
			&config{temp: 0, input: scale.NewCelsius(), output: scale.NewFahrenheit(), decimal: 2}},
		{[]string{"-u", "0", "celsius", "kelvin"},
			&config{temp: 0, input: scale.NewCelsius(), output: scale.NewFahrenheit(), decimal: 2, unit: true}},
		{[]string{"-u", "-d", "4", "0", "celsius", "kelvin"},
			&config{temp: 0, input: scale.NewCelsius(), output: scale.NewFahrenheit(), decimal: 4, unit: true}},
		{[]string{"-h"},
			&config{decimal: 2, help: true}},
		{[]string{"-h", "0", "celsius", "kelvin"},
			&config{decimal: 2, help: true}},
		{[]string{"-v"},
			&config{decimal: 2, version: true}},
		{[]string{"-v", "0", "celsius", "kelvin"},
			&config{decimal: 2, version: true}},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.args, " "), func(t *testing.T) {
			w := new(bytes.Buffer)
			flags := flag.NewFlagSet("test", flag.ContinueOnError)
			conf, err := ParseArgs(w, c.args, flags)
			if err != nil {
				t.Errorf("got %v want %v", err, nil)
			}

			if reflect.DeepEqual(*conf, c.conf) {
				t.Errorf("got %v want %v", *conf, c.conf)
			}
		})
	}
}

func TestParseArgsError(t *testing.T) {
	var cases = []struct {
		args []string
	}{
		{[]string{}},
		{[]string{"0"}},
		{[]string{"0", "kelvin"}},
		{[]string{"fifty", "celsius", "kelvin"}},
		{[]string{"0", "celsius", "wedgwood"}},
		{[]string{"-10", "celsius", "kelvin"}},
		{[]string{"0", "celsius", "kelvin", "extra"}},
		{[]string{"-d", "0", "celsius", "kelvin"}},
		{[]string{"-d", "13", "0", "celsius", "kelvin"}},
		{[]string{"-d", "-1", "0", "celsius", "kelvin"}},
		{[]string{"-f", "0", "celsius", "kelvin"}},
		{[]string{"-h", "-v"}},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.args, " "), func(t *testing.T) {
			w := new(bytes.Buffer)
			flags := flag.NewFlagSet("test", flag.ContinueOnError)
			_, err := ParseArgs(w, c.args, flags)

			if err == nil {
				t.Errorf("got %v want error", err)
			}
		})
	}
}

func TestMatchAll(t *testing.T) {
	slice := []string{
		"abc",
		"abb",
	}
	cases := []struct {
		pattern string
		want    []string
	}{
		{"ab", []string{"abc", "abb"}},
		{"abc", []string{"abc"}},
		{"def", []string{}},
		{"", []string{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s:%s->%s", c.pattern, strings.Join(slice, " "), strings.Join(c.want, " ")), func(t *testing.T) {
			got := matchAll(c.pattern, slice)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %q want %q", got, c.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	cases := []struct {
		slice [][]string
		want  []string
	}{
		{[][]string{{"a", "b"}}, []string{"a", "b"}},
		{[][]string{{"a", "b"}, {"c", "d"}}, []string{"a", "b", "c", "d"}},
		{[][]string{{}}, []string{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v->%v", c.slice, c.want), func(t *testing.T) {
			got := flatten(c.slice)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %q want %q", got, c.want)
			}
		})
	}
}
