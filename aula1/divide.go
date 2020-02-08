package aula1

func Divide(x float32, y float32 ) float32 {
	var xpositive, ypositive float32
	var negative float32  = 1

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

	if y > x {
		return (ypositive / xpositive) * negative
	} else {
		return (xpositive / ypositive) * negative
	}
}