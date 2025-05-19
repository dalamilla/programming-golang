package euler

import "testing"

func TestEuler002(t *testing.T) {
	var Euler002Tests = []struct {
		input  int
		result int
	}{
		{8, 10},
		{10, 10},
		{34, 44},
		{60, 44},
		{1000, 798},
		{100000, 60696},
		{4000000, 4613732},
	}

	for _, eu := range Euler002Tests {
		got := Euler002(eu.input)
		if got != eu.result {
			t.Errorf("Euler002(%d) => %d, should return %d", eu.input, got, eu.result)
		}
	}
}
