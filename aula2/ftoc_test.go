package aula2

import (
	"math"
	"testing"
)

func Test32FahrenheitToCelsius_Expected0_success(t *testing.T) {
	conversion := math.Round(Ftoc(32))
	if conversion != 0 {
		t.Errorf("calc of celsius was incorrect, got : %v, want: 0.", conversion)
	} else {
		t.Logf("calc was correct, got : %v, want: 0.", conversion)
	}
}

func Test100FahrenheitToCelsius_Expected37_77_success (t *testing.T) {
	conversion := math.Round(Ftoc(100))
	if conversion != 38 {
		t.Errorf("calc of celsius was incorrect, got : %v, want: 38.", conversion)
	} else {
		t.Logf("calc was correct, got : %v, want: 38.", conversion)
	}
}

func Test70FahrenheitToCelsius_Expected21_11_success (t *testing.T) {
	conversion := math.Round(Ftoc(70))
	if conversion != 21 {
		t.Errorf("calc of celsius was incorrect, got : %v, want: 21.", conversion)
	}else {
		t.Logf("calc was correct, got : %v, want: 21.", conversion)
	}
}

