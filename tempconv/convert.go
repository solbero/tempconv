package tempconv

type Converter[T TempScales] interface {
	TempScales
	toKelvin() float64
	fromKelvin(float64) (T, error)
}

func Convert[T Converter[T], S Converter[S]](input T, output S) (S, error) {
	temp := input.toKelvin()
	return output.fromKelvin(temp)
}
