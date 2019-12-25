package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cmds := ParseProgramCode(scanner.Text())
		newCmds := RestartCommands(cmds)
		RunCommands(newCmds)
		fmt.Println("At position 0, we got ", newCmds[0])

		pair := FindNounVerb(cmds)
		valuePair := 100*pair.noun + pair.verb

		fmt.Printf(
			"Pair calculation: 100 * noun + verb = 100 * %d + %d = %d",
			pair.noun,
			pair.verb,
			valuePair,
		)
	}
}
