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
		fmt.Println("At position 0, we got ", cmds[0])
	}
}
