package euler

import "testing"

func TestEuler003(t *testing.T) {
	var Euler003Tests = []struct {
		input  int
		result int
	}{
		{2, 2},
		{3, 3},
		{5, 5},
		{7, 7},
		{8, 2},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, eu := range Euler003Tests {
		got := Euler003(eu.input)
		if got != eu.result {
			t.Errorf("Euler003(%d) => %d, should return %d", eu.input, got, eu.result)
		}
	}
}
