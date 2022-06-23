package ExpandGo

import "testing"

type addExpand struct {
	arg1     float64
	expected float64
}

var addExpands = []addExpand{
	addExpand{4, 27},
	addExpand{0, -1},
	addExpand{-4, -125},
}

func TestExpand(t *testing.T) {

	for _, test := range addExpands {
		if output := expand(test.arg1); output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
