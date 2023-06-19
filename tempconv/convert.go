package tempconv

import (
	"fmt"
	"math"
)

type InvalidConversionError struct {
	input  TempScale
	output TempScale
}

func (e InvalidConversionError) Error() string {
	return fmt.Sprintf("invalid conversion from '%v' to '%v'", e.input.Name(), e.output.Name())
}

func Convert(input, output TempScale) error {
	k := NewKelvin()
	var err error
	err = kelvinFrom(k, input)
	if err != nil {
		return err
	}
	err = scaleFrom(k, output)
	if err != nil {
		return err
	}
	return err
}

func kelvinFrom(k *kelvin, ts TempScale) error {
	var t float64
	switch ts.(type) {
	case *kelvin:
		t = ts.Temp()
	case *celsius:
		t = ts.Temp() + math.Abs(absoluteZeroC)
	case *fahrenheit:
		t = (ts.Temp() + math.Abs(absoluteZeroF)) * 5 / 9
	default:
		return fmt.Errorf("tempconv: %w", &InvalidConversionError{input: ts, output: k})
	}
	mustSetTemp(k, t)
	return nil
}

func scaleFrom(k *kelvin, ts TempScale) error {
	var t float64
	switch ts.(type) {
	case *kelvin:
		t = k.Temp()
	case *celsius:
		t = k.Temp() - math.Abs(absoluteZeroC)
	case *fahrenheit:
		t = (k.Temp()*9 - math.Abs(absoluteZeroF)*5) / 5
	default:
		return fmt.Errorf("tempconv: %w", &InvalidConversionError{input: k, output: ts})
	}
	mustSetTemp(ts, t)
	return nil
}

func mustSetTemp(ts TempScale, t float64) {
	err := ts.SetTemp(t)
	if err != nil {
		panic(err)
	}
}
