package TrigGO

import "testing"

func Testtrig(t *testing.T) {

	got := trig(1, 3)
	want := 0.141120080598672
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
