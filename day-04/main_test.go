package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	value    string
	maxSame  int
	expected bool
}

func TestCases1(t *testing.T) {
	cases := []testCase{
		// {"111111", 0, true},
		// {"223450", 0, false},
		// {"123789", 0, false},
		// {"122345", 0, true},
		// {"122343", 0, false},
		// {"112233", 2, true},
		{"123444", 2, false},
		// {"111122", 2, true},
	}
	for _, c := range cases {
		result := VerifyPassword(c.value, c.maxSame)
		if c.expected != result {
			t.Errorf("Test case %s (%d) ==> Expected %t, got %t", c.value, c.maxSame, c.expected, result)
		}
		fmt.Println(c)
	}
}
