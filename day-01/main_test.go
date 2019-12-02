package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	cases := []struct {
		in int; want int
	} {
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range cases {
		got := CalculateFuel(c.in)
		if got != c.want {
			t.Errorf("CalculateFuel(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCalculateFuelRecursive(t *testing.T) {
	cases := []struct {
		in int; want int
	} {
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range cases {
		got := CalculateFuelRecursive(c.in)
		if got != c.want {
			t.Errorf("CalculateFuel(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}