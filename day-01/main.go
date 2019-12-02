package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var resultOriginal, resultRecursive int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		resultOriginal += CalculateFuel(val)
		resultRecursive += CalculateFuelRecursive(val)
	}

	fmt.Println("Original fuel: %d", resultOriginal)
	fmt.Println("Recursive fuel: %d", resultRecursive)
}
