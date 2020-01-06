package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	wire1    string
	wire2    string
	expected int
}

func TestCases1(t *testing.T) {
	cases := []testCase{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, c := range cases {
		fmt.Println(c)
		wire1Coordinates, _ := GenerateWireCoords(c.wire1)
		wire2Coordinates, _ := GenerateWireCoords(c.wire2)

		intersectionPoint, closestDistance := CalculateClosestIntersectionPoint(wire1Coordinates, wire2Coordinates)

		fmt.Println(intersectionPoint)
		if c.expected != closestDistance {
			t.Errorf("Expected %d, got  %d", c.expected, closestDistance)
		}

	}
}

func TestCases2(t *testing.T) {
	cases := []testCase{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for _, c := range cases {
		fmt.Println(c)
		wire1Coordinates, wire1Dist := GenerateWireCoords(c.wire1)
		wire2Coordinates, wire2Dist := GenerateWireCoords(c.wire2)

		shortest1, shortest2 := CalculateShortestPaths(wire1Coordinates, wire2Coordinates, wire1Dist, wire2Dist)

		fmt.Printf("wire1 = %d; wire2 = %d\n", shortest1, shortest2)
		if c.expected != shortest1+shortest2 {
			t.Errorf("Expected %d, got  %d", c.expected, shortest1+shortest2)
		}

	}
}
