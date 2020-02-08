package aula2

const (
	val1 = 32
	val2 float32 = 5.0/9.0
)

//formula: (32°F − 32) × 5/9 = 0°C
func Ftoc (fahrenheit float32) float64 {
	return float64((fahrenheit - val1) * val2)
}