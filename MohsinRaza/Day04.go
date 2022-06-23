package Task

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{1, 2, 3},
	addTest{4, 5, 9},
	addTest{6, 7, 13},
	addTest{3, 10, 13},
}

func TestSum(t *testing.T) {

	for _, test := range addTests {
		if output := sum(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
