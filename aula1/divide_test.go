package aula1

import "testing"

func TestDivide20_per_5_Expected4_success(t *testing.T) {
	divisionresult := Divide(20, 5)
	if divisionresult != 4 {
		t.Errorf("division was incorrect, got : %v, want : 4", divisionresult)
	} else {
		t.Logf("division was corret, got : %v, want: 4", divisionresult)
	}
}

func TestDivide100_per_0_Expected0_success(t *testing.T) {
	divisionresult := Divide(100, 0)
	if divisionresult != 0 {
		t.Errorf("division was incorrect, got : %v, want : 0", divisionresult)
	} else {
		t.Logf("division was corret, got : %v, want: 0", divisionresult)
	}
}

func TestDivide_minus20_per_minus2_Expected10_success(t *testing.T) {
	divisionresult := Divide(-20, -2)
	if divisionresult != -10{
		t.Errorf("division was incorrect, got : %v, want : -10", divisionresult)
	} else {
		t.Logf("division was corret, got : %v, want: -10", divisionresult)
	}
}

func TestDivide9_per_20_Expected0_0_45_success(t *testing.T) {
	divisionresult := Divide(9,20 )
	if divisionresult != 0.45{
		t.Errorf("division was incorrect, got : %v, want : 0.45", divisionresult)
	} else {
		t.Logf("division was corret, got : %v, want: 0.45", divisionresult)
	}
}