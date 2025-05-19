package euler

import "testing"

func TestEuler005(t *testing.T) {
	var Euler005Tests = []struct {
		input  int
		result int
	}{
		{5, 60},
		{7, 420},
		{10, 2520},
		{13, 360360},
		{20, 232792560},
	}

	for _, eu := range Euler005Tests {
		got := Euler005(eu.input)
		if got != eu.result {
			t.Errorf("Euler005(%d) => %d, should return %d", eu.input, got, eu.result)
		}
	}
}
