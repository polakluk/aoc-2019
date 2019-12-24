package main

import (
	"strconv"
	"strings"
)

// ParseProgramCode Parses line of commands into separate commands
func ParseProgramCode(program string) []int {
	strCommands := strings.Split(program, ",")
	commands := make([]int, len(strCommands))
	for idx, val := range strCommands {
		commands[idx], _ = strconv.Atoi(val)
	}
	return commands
}

// RestartCommands Restart commands to initial state
func RestartCommands(commands []int) []int {
	newCommands := commands
	newCommands[1] = 12
	newCommands[2] = 2

	return newCommands
}

// RunCommands Runs program
func RunCommands(commands []int) {
	idx := 0
	for commands[idx] != 99 {
		command := commands[idx]
		op1 := commands[commands[idx+1]]
		op2 := commands[commands[idx+2]]
		destination := commands[idx+3]

		switch command {
		case 1:
			commands[destination] = op1 + op2
		case 2:
			commands[destination] = op1 * op2
		}
		idx += 4
	}
}
