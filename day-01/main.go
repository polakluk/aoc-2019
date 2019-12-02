package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Printf("Original fuel: %d", resultOriginal)
	fmt.Printf("Recursive fuel: %d", resultRecursive)
}
