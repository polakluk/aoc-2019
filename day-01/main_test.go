package main

import "testing"

type testCase struct {
	value    int
	expected int
}

func TestCalculateFuel(t *testing.T) {
	cases := []testCase{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range cases {
		got := CalculateFuel(c.value)
		if got != c.expected {
			t.Errorf("CalculateFuel(%d) == %d, want %d", c.value, got, c.expected)
		}
	}
}

func TestCalculateFuelRecursive(t *testing.T) {
	cases := []testCase{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range cases {
		got := CalculateFuelRecursive(c.value)
		if got != c.expected {
			t.Errorf("CalculateFuel(%d) == %d, want %d", c.value, got, c.expected)
		}
	}
}
