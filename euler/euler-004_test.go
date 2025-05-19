package euler

import "testing"

func TestEuler004(t *testing.T) {
	var Euler004Tests = []struct {
		input  int
		result int
	}{
		{2, 9009},
		{3, 906609},
	}

	for _, eu := range Euler004Tests {
		got := Euler004(eu.input)
		if got != eu.result {
			t.Errorf("Euler004(%d) => %d, should return %d", eu.input, got, eu.result)
		}
	}
}
