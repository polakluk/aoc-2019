package main

import (
	"testing"
)

type testCase struct {
	value       string
	expectedIdx int
	expectedVal int
}

func TestCases1(t *testing.T) {
	cases := []testCase{
		{"1,0,0,0,99", 0, 2},
		{"2,3,0,3,99", 3, 6},
		{"2,4,4,5,99,0", 5, 9801},
		{"1,1,1,4,99,5,6,0,99", 0, 30},
	}
	for _, c := range cases {
		cmds := ParseProgramCode(c.value)
		RunCommands(cmds)

		if cmds[c.expectedIdx] != c.expectedVal {
			t.Errorf("RunCommands(%s)[%d] == %d, want %d", c.value, c.expectedIdx, cmds[c.expectedIdx], c.expectedVal)
		}
	}
}
