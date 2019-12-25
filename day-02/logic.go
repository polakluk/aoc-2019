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
	newCommands := make([]int, len(commands))
	copy(newCommands, commands)
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

		if idx+3 >= len(commands) || idx >= len(commands) {
			return
		}
	}
}

//NounVerbPair Pair of noun and verb
type NounVerbPair struct {
	noun int
	verb int
}

// FindNounVerb Tries to find a noun and a verb
func FindNounVerb(commands []int) NounVerbPair {
	result := NounVerbPair{0, 0}
	keepLooking := true

	for noun := 0; noun < 100 && keepLooking; noun++ {
		for verb := 0; verb < 100 && keepLooking; verb++ {
			newCommands := make([]int, len(commands))
			copy(newCommands, commands)

			newCommands[1] = noun
			newCommands[2] = verb
			RunCommands(newCommands)
			if newCommands[0] == 19690720 {
				keepLooking = false
				result.noun = noun
				result.verb = verb
			}
		}
	}

	return result
}
