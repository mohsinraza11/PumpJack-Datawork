package EulerGo

import "testing"

type addEuler struct {
	arg1     complex128
	expected complex128
}

var addEulers = []addEuler{
	addEuler{1, 0 + 1.2246467991473515e-16i},
	addEuler{1, 0 + 1.2246467991473515e-16i},
}

func Testeuler(t *testing.T) {

	for _, test := range addEulers {
		if output := euler(); output != test.expected {
			t.Errorf("Output %.2f not equal to expected %.2f", output, test.expected)
		}
	}
}
