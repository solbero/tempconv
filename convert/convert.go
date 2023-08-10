package convert

import (
	"errors"
	"fmt"

	"github.com/solbero/tempconv/scale"
)

var ErrScaleNotSupported = errors.New("scale not supported")

// InvalidConversionError is an error type for invalid temperature conversions.
type InvalidConversionError struct {
	input  scale.Scale
	output scale.Scale
	err    error
}

func (e InvalidConversionError) Error() string {
	return fmt.Sprintf("invalid conversion from %s to %s: %s", e.input.Name(), e.output.Name(), e.err.Error())
}

// Convert converts a temperature from a temperature scale to another.
// It returns an error if the conversion is not possible.scale.
func Convert(fromScale, toScale scale.Scale) error {
	k := scale.NewKelvin()
	var err error
	if err = kelvinFrom(fromScale, k); err != nil {
		return err
	}
	if err = kelvinTo(toScale, k); err != nil {
		return err
	}
	return nil
}

func kelvinFrom(ts scale.Scale, k *scale.Kelvin) error {
	var t float64
	switch ts.(type) {
	case *scale.Kelvin:
		t = ts.Temp()
	case *scale.Celsius:
		t = ts.Temp() + 273.15
	case *scale.Fahrenheit:
		t = (ts.Temp()*5 + 459.67*5) / 9
	case *scale.Rankine:
		t = ts.Temp() * 5 / 9
	case *scale.Delisle:
		t = (373.15*3 - ts.Temp()*2) / 3
	case *scale.Newton:
		t = (ts.Temp()*100 + 273.15*33) / 33
	case *scale.Reaumur:
		t = (ts.Temp()*5 + 273.15*4) / 4
	case *scale.Roemer:
		t = (ts.Temp()*40 - 7.5*40 + 273.15*21) / 21
	default:
		panic(fmt.Errorf("tempconv: %w", InvalidConversionError{input: ts, output: k, err: ErrScaleNotSupported}))
	}
	err := k.SetTemp(t)
	if err != nil {
		return fmt.Errorf("tempconv: %w", InvalidConversionError{input: ts, output: k, err: errors.Unwrap(err)})
	}
	return nil
}

func kelvinTo(ts scale.Scale, k *scale.Kelvin) error {
	var t float64
	switch ts.(type) {
	case *scale.Kelvin:
		t = k.Temp()
	case *scale.Celsius:
		t = k.Temp() - 273.15
	case *scale.Fahrenheit:
		t = (k.Temp()*9 - 459.67*5) / 5
	case *scale.Rankine:
		t = k.Temp() * 9 / 5
	case *scale.Delisle:
		t = (373.15 - k.Temp()) * 3 / 2
	case *scale.Newton:
		t = (k.Temp() - 273.15) * 33 / 100
	case *scale.Reaumur:
		t = (k.Temp()*4 - 273.15*4) / 5
	case *scale.Roemer:
		t = ((k.Temp()*21 - 273.15*21) + 7.5*40) / 40
	default:
		panic(fmt.Errorf("tempconv: %w", InvalidConversionError{input: ts, output: k, err: ErrScaleNotSupported}))
	}
	err := ts.SetTemp(t)
	if err != nil {
		return fmt.Errorf("tempconv: %w", InvalidConversionError{input: k, output: ts, err: errors.Unwrap(err)})
	}
	return nil
}
