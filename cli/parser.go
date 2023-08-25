package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/solbero/tempconv/scale"
)

const usageMsg = "try 'tempconv -h' for more information"

type config struct {
	temp    float64
	input   *scale.Scale
	output  *scale.Scale
	decimal int
	unit    bool
	version bool
	help    bool
}

func ParseArgs(w io.Writer, args []string, flags *flag.FlagSet) (conf *config, err error) {
	flags.SetOutput(w)
	flags.Usage = func() {}

	// Parse flags
	conf = &config{}
	flags.IntVar(&conf.decimal, "d", 2, "Number of decimal places [default: 2, min: 0, max: 12]")
	flags.BoolVar(&conf.unit, "u", false, "Include temperature unit")
	flags.BoolVar(&conf.version, "v", false, "Show version and exit")
	flags.BoolVar(&conf.help, "h", false, "Show help and exit")

	err = flags.Parse(args)
	if err != nil {
		fmt.Fprint(w, usageMsg) // flag.Parse already prints an error
		return nil, err
	}

	// Check temp decimal places
	var msg string
	min, max := 0, 12
	if conf.decimal < min || conf.decimal > max {
		msg = fmt.Sprintf("invalid value for -d flag: %v, must be between %v and %v", conf.decimal, min, max)
		fprinte(w, msg)
		return nil, fmt.Errorf(msg)
	}

	// Check mutually exclusive flags
	if conf.version && conf.help {
		msg = "mutually exclusive flags: -h, -v"
		fprinte(w, msg)
		return nil, errors.New("mutually exclusive flags: -h, -v")
	}

	// Print version or help
	if conf.version || conf.help {
		return conf, nil
	}

	// Check non-flag arguments
	nonFlagArgs := flags.Args()
	err = checkNonFlagArgs(nonFlagArgs)
	if err != nil {
		fprinte(w, err.Error())
		return nil, err
	}

	// Parse non-flag arguments
	temp := nonFlagArgs[0]
	input := nonFlagArgs[1]
	output := nonFlagArgs[2]
	conf.temp, err = strconv.ParseFloat(temp, 64)

	if err != nil {
		msg = fmt.Sprintf("invalid value for temp argument: %s", nonFlagArgs[0])
		fprinte(w, msg)
		return nil, fmt.Errorf(msg)
	}
	conf.input, err = parseScale(input)

	if err != nil {
		fprinte(w, err.Error())
		return nil, err
	}
	conf.output, err = parseScale(output)

	if err != nil {
		fprinte(w, err.Error())
		return nil, err
	}

	return conf, nil
}

func parseScale(name string) (*scale.Scale, error) {
	scales := flatten(scale.ScaleNames())
	matches := matchAll(name, scales)

	if len(matches) == 0 {
		return nil, fmt.Errorf("unknown temperature scale: %s", name)
	} else if len(matches) > 1 {
		return nil, fmt.Errorf("ambiguous temperature scale: %s, matches: %s", name, strings.Join(matches, ", "))
	}

	switch matches[0] {
	case "kelvin":
		return scale.NewKelvin(), nil
	case "celsius":
		return scale.NewCelsius(), nil
	case "fahrenheit":
		return scale.NewFahrenheit(), nil
	case "rankine":
		return scale.NewRankine(), nil
	case "delisle":
		return scale.NewDelisle(), nil
	case "newton":
		return scale.NewNewton(), nil
	case "réaumur":
		return scale.NewReaumur(), nil
	case "reaumur": // alias
		return scale.NewReaumur(), nil
	case "rømer":
		return scale.NewRomer(), nil
	case "romer": // alias
		return scale.NewRomer(), nil
	default:
		panic(errors.New("unable to parse scale"))
	}
}

func checkNonFlagArgs(args []string) error {
	required := []string{"temp", "from scale", "to scale"} // required args

	if len(args) == 0 {
		return fmt.Errorf("missing required arguments: %s", strings.Join(required[0:], ", "))
	} else if len(args) == 1 {
		return fmt.Errorf("missing required arguments: %s", strings.Join(required[1:], ", "))
	} else if len(args) == 2 {
		return fmt.Errorf("missing required argument: %s", strings.Join(required[2:], ", "))
	} else if len(args) > 3 {
		return fmt.Errorf("supplied too many arguments: %v", strings.Join(args[3:], ", "))
	}

	return nil
}

func matchAll(pattern string, slice []string) []string {
	matches := []string{}

	if pattern == "" {
		return matches
	}

	for i, s := range slice {
		if strings.HasPrefix(s, strings.ToLower(pattern)) {
			matches = append(matches, slice[i])
		}
	}

	return matches
}

func flatten(slice [][]string) []string {
	flat := []string{}

	for _, s := range slice {
		flat = append(flat, s...)
	}

	return flat
}

func fprinte(w io.Writer, msg string) {
	w.Write([]byte(msg + "\n" + usageMsg))
}
