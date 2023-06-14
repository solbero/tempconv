package tempconv

type Converter interface {
	TempScales
	toKelvin() float64
	fromKelvin(float64)
}

func Convert[T, S Converter](input T, output S) S {
	temp := input.toKelvin()
	output.fromKelvin(temp)
	return output

}
