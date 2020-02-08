package aula2

const (
	val1 = 32
	val2 = 5 / 9
)

//formula: (32°F − 32) × 5/9 = 0°C
func Ftoc (fahrenheit float32) float32 {
	return (fahrenheit - val1) * val2
}