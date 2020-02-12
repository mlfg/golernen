package aula1

func Divide(x float64, y float64 ) float64 {
	var xpositive, ypositive float64
	var negative float64  = 1

	if x == 0 || y == 0 {
		return 0
	}

	xpositive = x
	ypositive = y

	if x < 0 {
		xpositive = x * -1
		negative = -1
	}

	if y < 0 {
		ypositive = y * -1
		negative = -1
	}
		return (xpositive / ypositive) * negative
}