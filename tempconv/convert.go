package tempconv

import (
	"fmt"
	"math"
)

// InvalidConversionError is an error type for invalid temperature conversions.
type InvalidConversionError struct {
	input  TempScale
	output TempScale
}

func (e InvalidConversionError) Error() string {
	return fmt.Sprintf("invalid conversion from '%v' to '%v'", e.input.Name(), e.output.Name())
}

// Convert converts a temperature from a temperature scale to another.
// It returns an error if the conversion is not possible.
func Convert(fromScale, toScale Scale) error {
	k := NewKelvin()
	var err error
	if err = kelvinFrom(fromScale, k); err != nil {
		return err
	}
	if err = kelvinTo(toScale, k); err != nil {
		return err
	}
	return nil
}

func kelvinFrom(ts Scale, k *kelvin) error {
	var t float64
	switch ts.(type) {
	case *kelvin:
		t = ts.Temp()
	case *celsius:
		t = ts.Temp() + math.Abs(absoluteZeroC)
	case *fahrenheit:
		t = (ts.Temp()*5 + math.Abs(absoluteZeroF)*5) / 9
	case *rankine:
		t = ts.Temp() * 5 / 9
	case *delisle:
		t = (373.15*3 - ts.Temp()*2) / 3
	case *newton:
		t = (ts.Temp()*100 + math.Abs(absoluteZeroC)*33) / 33
	case *reaumur:
		t = (ts.Temp()*5 + math.Abs(absoluteZeroC)*4) / 4
	case *roemer:
		t = (ts.Temp()*40 - 7.5*40 + math.Abs(absoluteZeroC)*21) / 21
	default:
		panic(fmt.Errorf("tempconv: %w", &InvalidConversionError{input: ts, output: k}))
	}
	return k.SetTemp(t)
}

func kelvinTo(ts Scale, k *kelvin) error {
	var t float64
	switch ts.(type) {
	case *kelvin:
		t = k.Temp()
	case *celsius:
		t = k.Temp() - math.Abs(absoluteZeroC)
	case *fahrenheit:
		t = (k.Temp()*9 - math.Abs(absoluteZeroF)*5) / 5
	case *rankine:
		t = k.Temp() * 9 / 5
	case *delisle:
		t = (373.15 - k.Temp()) * 3 / 2
	case *newton:
		t = (k.Temp()*33 - math.Abs(absoluteZeroC)*33) / 100
	case *reaumur:
		t = (k.Temp()*4 - math.Abs(absoluteZeroC)*4) / 5
	case *roemer:
		t = ((k.Temp()*21 - math.Abs(absoluteZeroC)*21) + 7.5*40) / 40
	default:
		panic(fmt.Errorf("tempconv: %w", &InvalidConversionError{input: k, output: ts}))
	}
	return ts.SetTemp(t)
}
