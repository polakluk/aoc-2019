package main

func CalculateFuel(fuel int) int {
	return int(fuel/3) - 2
}

func CalculateFuelRecursive(fuel int) int {
	newFuel := CalculateFuel(fuel)
	if newFuel <= 0 {
		return 0
	} else {
		return newFuel + CalculateFuelRecursive(newFuel)
	}
}
