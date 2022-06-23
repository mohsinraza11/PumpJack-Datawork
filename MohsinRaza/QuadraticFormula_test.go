package NewTask

import "testing"

var QTests = []struct {
	in1,
	in2,
	in3,
	out1,
	out2 float64
}{
	{1, 2, -15, 3, -5},
}

func TestQuadratic(t *testing.T) {
	for _, tt := range QTests {
		s, p := Quadratic_formula(tt.in1, tt.in2, tt.in3)
		if (s != tt.out1) || (p != tt.out2) {
			t.Errorf("Output %f not equal to expected %f", s, p)
		}
	}
}
