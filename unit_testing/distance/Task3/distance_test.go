package Task3

import "testing"

type distTest struct {
	x1, x2, y1, y2, expected float64
}

var distTests = []distTest{
	distTest{3, 7, 2, 8, 7.21},
	distTest{0, 6, 0, 8, 10},
}

func TestDist(t *testing.T) {

	for _, test := range distTests {
		if output := Distance_formula(test.x1, test.x2, test.y1, test.y2); output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
