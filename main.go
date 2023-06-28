package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/solbero/tempconv/tempconv"
)

const helpTemplateStr = `tempconv converts temperatures between different temperature scales.

Usage:
  tempconv <temp> <from scale> <to scale> [OPTIONS]
  tempconv -h
  tempconv -v

Scales:
{{- range .Scales }}
  {{ . }}
{{- end }}

It is possible to use shorthand as long as it uniquely identifies a scale.

Options:
{{- range .Flags }}
  -{{ printf "%-2s" .Name}} {{.Usage}}
{{- end }}

Examples:
  tempconv 0 celsius kelvin
  tempconv 0 c k
  tempconv 0 c k -d 1`

var helpTemplate *template.Template

func helpInfo() (info struct {
	Scales []string
	Flags  []flag.Flag
}) {
	info.Scales = supportedTempScales()
	info.Flags = func() (l []flag.Flag) {
		flag.VisitAll(func(f *flag.Flag) {
			l = append(l, *f)
		})
		return l
	}()
	return info
}

func supportedTempScales() []string {
	scales := []tempconv.Scale{
		tempconv.NewKelvin(),
		tempconv.NewCelsius(),
		tempconv.NewFahrenheit(),
	}
	var list []string
	for _, s := range scales {
		list = append(list, s.Name())
	}
	return list
}

func matchAll(s string, l []string) []string {
	matches := []string{}
	if s == "" {
		return matches
	}
	for i, str := range l {
		if strings.HasPrefix(str, strings.ToLower(s)) {
			matches = append(matches, l[i])
		}
	}
	return matches
}

func parseTempScale(s string, scales []string) (tempconv.Scale, error) {
	matches := matchAll(s, scales)

	if len(matches) == 0 {
		return nil, fmt.Errorf("unknown scale '%v'", s)
	} else if len(matches) > 1 {
		return nil, fmt.Errorf("ambiguous scale '%v', matches '%v'", s, strings.Join(matches, "', '"))
	}

	switch matches[0] {
	case "kelvin":
		return tempconv.NewKelvin(), nil
	case "celsius":
		return tempconv.NewCelsius(), nil
	case "fahrenheit":
		return tempconv.NewFahrenheit(), nil
	default:
		panic("failed to parse scale")
	}
}

func printErr(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
	flag.Usage()
}

func init() {
	helpTemplate = template.Must(template.New("tempconv").Parse(helpTemplateStr))
}

func main() {
	unit := flag.Bool("u", false, "Include temperature unit")
	version := flag.Bool("v", false, "Show version")
	decimal := flag.Int("d", 2, "Number of decimal places [default: 2]")
	help := flag.Bool("h", false, "Show help")

	errWriter := flag.CommandLine.Output()

	flag.Usage = func() {
		fmt.Fprintln(errWriter, "Try 'tempconv -h' for more information.")
	}

	if len(os.Args) == 1 {
		printErr(errWriter, "missing required arguments")
		os.Exit(2)

	} else if len(os.Args) == 2 {
		flag.CommandLine.Parse(os.Args[1:])
		if *help {
			helpTemplate.Execute(os.Stdout, helpInfo())
			os.Exit(0)
		} else if *version {
			fmt.Fprintln(os.Stdout, "0.1.0")
			os.Exit(0)
		}
		printErr(errWriter, "missing required arguments")
		os.Exit(2)

	} else if len(os.Args) > 2 && len(os.Args) < 4 {
		printErr(errWriter, "missing required arguments")
		os.Exit(2)

	} else if len(os.Args) >= 4 {
		args := os.Args[1:4]
		flag.CommandLine.Parse(os.Args[4:])

		tempStr := args[0]
		inputStr := args[1]
		outputStr := args[2]

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			printErr(errWriter, "argument 'temp' is not a number")
			os.Exit(2)
		}

		scales := supportedTempScales()
		fromTempScale, err := parseTempScale(inputStr, scales)
		if err != nil {
			printErr(errWriter, err.Error())
			os.Exit(2)
		}

		toTempScale, err := parseTempScale(outputStr, scales)
		if err != nil {
			printErr(errWriter, err.Error())
			os.Exit(2)
		}

		err = fromTempScale.SetTemp(temp)
		if err != nil {
			printErr(errWriter, err.Error())
			os.Exit(2)
		}

		err = tempconv.Convert(fromTempScale, toTempScale)
		if err != nil {
			printErr(errWriter, err.Error())
			os.Exit(2)
		}

		if *unit {
			fmt.Fprintf(os.Stdout, "%.*f %s\n", *decimal, toTempScale.Temp(), toTempScale.Unit())
		} else {
			fmt.Fprintf(os.Stdout, "%.*f\n", *decimal, toTempScale.Temp())
		}
	}
}
