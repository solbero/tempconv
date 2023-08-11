package convert

import (
	"errors"
	"fmt"

	"github.com/solbero/tempconv/scale"
)

var ErrScaleNotSupported = errors.New("scale not supported")

// InvalidConversionError is an error type for invalid temperature conversions.
type InvalidConversionError struct {
	input  *scale.Scale
	output *scale.Scale
	err    error
}

func (ic InvalidConversionError) Error() string {
	return fmt.Sprintf("invalid conversion from %s to %s: %s", ic.input.Name, ic.output.Name, ic.err.Error())
}

// Convert converts a temperature from a temperature scale to another.
// It returns an error if the conversion is not possible.scale.
func Convert(input, output *scale.Scale) (err error) {
	k := scale.NewKelvin()

	if err = kelvinFrom(input, k); err != nil {
		return err
	}

	if err = kelvinTo(output, k); err != nil {
		return err
	}

	return nil
}

func kelvinFrom(s, k *scale.Scale) (err error) {
	var t float64

	switch s.Type {
	case scale.KELVIN:
		t = s.Temp()
	case scale.CELSIUS:
		t = s.Temp() + 273.15
	case scale.FAHRENHEIT:
		t = (s.Temp()*5 + 459.67*5) / 9
	case scale.RANKINE:
		t = s.Temp() * 5 / 9
	case scale.DELISLE:
		t = (373.15*3 - s.Temp()*2) / 3
	case scale.NEWTON:
		t = (s.Temp()*100 + 273.15*33) / 33
	case scale.REAUMUR:
		t = (s.Temp()*5 + 273.15*4) / 4
	case scale.ROMER:
		t = (s.Temp()*40 - 7.5*40 + 273.15*21) / 21
	default:
		panic(fmt.Errorf("tempconv: %w", InvalidConversionError{input: s, output: k, err: ErrScaleNotSupported}))
	}

	err = k.SetTemp(t)
	if err != nil {
		return fmt.Errorf("tempconv: %w", InvalidConversionError{input: s, output: k, err: errors.Unwrap(err)})
	}

	return nil
}

func kelvinTo(s, k *scale.Scale) (err error) {
	var t float64

	switch s.Type {
	case scale.KELVIN:
		t = k.Temp()
	case scale.CELSIUS:
		t = k.Temp() - 273.15
	case scale.FAHRENHEIT:
		t = (k.Temp()*9 - 459.67*5) / 5
	case scale.RANKINE:
		t = k.Temp() * 9 / 5
	case scale.DELISLE:
		t = (373.15 - k.Temp()) * 3 / 2
	case scale.NEWTON:
		t = (k.Temp() - 273.15) * 33 / 100
	case scale.REAUMUR:
		t = (k.Temp()*4 - 273.15*4) / 5
	case scale.ROMER:
		t = ((k.Temp()*21 - 273.15*21) + 7.5*40) / 40
	default:
		panic(fmt.Errorf("tempconv: %w", InvalidConversionError{input: s, output: k, err: ErrScaleNotSupported}))
	}
	err = s.SetTemp(t)
	if err != nil {
		return fmt.Errorf("tempconv: %w", InvalidConversionError{input: k, output: s, err: errors.Unwrap(err)})
	}

	return nil
}
