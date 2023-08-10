package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"text/template"

	"github.com/solbero/tempconv/convert"
	"github.com/solbero/tempconv/scale"
)

var templateParsed *template.Template

const helpTemplate = `tempconv converts temperatures between different temperature scales.

Usage:
  tempconv [OPTIONS] <temp> <from scale> <to scale>
  tempconv -h
  tempconv -v

If temperature is negative, it must be prefixed with '--' to avoid being interpreted as a flag.

Scales:
{{- range .Scales}}{{- range $i, $v := .}}
{{- if not $i}}
  {{else if $i}}, {{end}}{{$v}}
{{- end}}{{- end}}

It is possible to use abbreviations as long as it uniquely identifies a scale.

Options:
{{- range .Flags }}
  -{{ printf "%-2s" .Name}} {{.Usage}}
{{- end}}

Examples:
  tempconv 0 celsius kelvin
  tempconv 0 c k
  tempconv -u -d 4 0 celsius kelvin
  tempconv -u -- -10 celsius kelvin`

func templateData(scales [][]string, flags *flag.FlagSet) (info struct {
	Scales [][]string
	Flags  []flag.Flag
}) {
	info.Scales = scales
	info.Flags = func() (l []flag.Flag) {
		flags.VisitAll(func(f *flag.Flag) {
			l = append(l, *f)
		})
		return l
	}()
	return info
}

func init() {
	templateParsed = template.Must(template.New("tempconv").Parse(helpTemplate))
}

func Run(w io.Writer, conf *config, flags *flag.FlagSet, version string) (err error) {
	if conf.help {
		data := templateData(scale.ScaleNames(), flags)
		templateParsed.Execute(w, data)
		return nil
	}

	if conf.version {
		fmt.Fprint(w, version)
		return nil
	}

	err = conf.fromScale.SetTemp(conf.temp)
	if err != nil {
		err := errors.Unwrap(err)
		fprinte(w, err.Error())
		return err
	}

	err = convert.Convert(conf.fromScale, conf.toScale)
	if err != nil {
		err := errors.Unwrap(err)
		fprinte(w, err.Error())
		return err
	}

	if conf.unit {
		fmt.Fprintf(w, "%.*f %s", conf.decimal, conf.toScale.Temp(), conf.toScale.Unit())
	} else {
		fmt.Fprintf(w, "%.*f", conf.decimal, conf.toScale.Temp())
	}
	return nil
}
