package euler

import "testing"

func TestEuler001(t *testing.T) {

	var Euler001Tests = []struct {
		input  int
		result int
	}{
		{1000, 233168},
		{49, 543},
		{8456, 16687353},
		{19564, 89301183},
	}

	for _, eu := range Euler001Tests {
		got := Euler001(eu.input)
		if got != eu.result {
			t.Errorf("Euler001(%d) => %d, should return %d", eu.input, got, eu.result)
		}
	}
}
