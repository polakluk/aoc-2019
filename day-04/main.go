package main

import (
	"fmt"
	"strconv"
)

// handleSearch controls search through the interval and returns number of valid passwords
func handleSearch(beginInterval int, endInterval int) (int, int) {
	passwords := 0
	passwords2 := 0
	actNumber := beginInterval
	for actNumber <= endInterval {
		if VerifyPassword(strconv.Itoa(actNumber), 0) {
			passwords++
		}
		if VerifyPassword(strconv.Itoa(actNumber), 2) {
			passwords2++
		}
		actNumber++
	}
	return passwords, passwords2
}

func main() {
	passwords, passwords2 := handleSearch(387638, 919123)

	fmt.Printf("potential passwords part 1 = %d\n", passwords)
	fmt.Printf("potential passwords part 2 = %d\n", passwords2)
}
