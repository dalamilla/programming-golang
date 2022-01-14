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
