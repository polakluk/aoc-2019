package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Point structure holding a single point
type Point struct {
	row int
	col int
}

// GenerateWireCoords parses commands into wire points
func GenerateWireCoords(wire string) ([]Point, []int) {
	// MovementMap simplifies generating points in 2D grid
	MovementMap := map[string]Point{
		"R": Point{0, 1},
		"L": Point{0, -1},
		"U": Point{-1, 0},
		"D": Point{1, 0},
	}

	commands := strings.Split(wire, ",")
	var coordinates = make([]Point, len(commands)+1)
	var distances = make([]int, len(commands)+1)
	actRow, actCol, actDist := 0, 0, 0

	coordinates[0] = Point{0, 0}

	for idx, val := range commands {
		size, _ := strconv.Atoi(val[1:])
		actRow += MovementMap[string(val[0])].row * size
		actCol += MovementMap[string(val[0])].col * size

		coordinates[idx+1] = Point{actRow, actCol}
		actDist += size
		distances[idx+1] = actDist
	}

	return coordinates, distances
}

// WithinSegment tests whether provided coordinates are within segment
func WithinSegment(startPoint Point, stopPoint Point, row float64, col float64) bool {
	return math.Min(float64(startPoint.row), float64(stopPoint.row)) <= row &&
		math.Max(float64(startPoint.row), float64(stopPoint.row)) >= row &&
		math.Min(float64(startPoint.col), float64(stopPoint.col)) <= col &&
		math.Max(float64(startPoint.col), float64(stopPoint.col)) >= col
}

// calculateIntersection calculates intestection point of segment defined by points AB & CD (given the first flag is True in result)
func calculateIntersection(pointA Point, pointB Point, pointC Point, pointD Point) (bool, int, int) {
	a1 := pointB.row - pointA.row
	a2 := pointD.row - pointC.row

	b1 := pointA.col - pointB.col
	b2 := pointC.col - pointD.col

	c1 := a1*pointA.col + b1*pointA.row
	c2 := a2*pointC.col + b2*pointC.row

	determinant := a1*b2 - a2*b1

	if determinant != 0 {
		col := (b2*c1 - b1*c2) / determinant
		row := (a1*c2 - a2*c1) / determinant

		return true, row, col
	}
	return false, 0, 0
}

// CalculateClosestIntersectionPoint returns closest intersection point to origin
func CalculateClosestIntersectionPoint(wire1 []Point, wire2 []Point) (Point, int) {
	idxWire1 := 0
	var closestDistance float64 = -1
	var closestPoint Point = Point{0, 0}

	for idxWire1 < len(wire1)-1 {
		idxWire2 := 1
		for idxWire2 < len(wire2)-1 {
			found, row, col := calculateIntersection(wire1[idxWire1], wire1[idxWire1+1], wire2[idxWire2], wire2[idxWire2+1])

			if found &&
				WithinSegment(wire1[idxWire1], wire1[idxWire1+1], float64(row), float64(col)) &&
				WithinSegment(wire2[idxWire2], wire2[idxWire2+1], float64(row), float64(col)) {

				currentDistanct := math.Abs(float64(col)) + math.Abs(float64(row))
				if currentDistanct < float64(closestDistance) || closestDistance == -1 {
					closestDistance = currentDistanct
					closestPoint = Point{row, col}
					fmt.Printf("closestDist = %f\n", closestDistance)
				}
			}

			idxWire2++
		}

		idxWire1++
	}

	return closestPoint, int(closestDistance)
}

//  distFromPoint calculates distance of given coordinates from the given point in manhattan coordinates
func distFromPoint(point Point, row float64, col float64) int {
	return int(math.Abs(math.Abs(col)-math.Abs(float64(point.col))) + math.Abs(math.Abs(row)-math.Abs(float64(point.row))))
}

// CalculateShortestPaths returns shortests paths to an intersection point
func CalculateShortestPaths(wire1 []Point, wire2 []Point, wire1Dists []int, wire2Dists []int) (int, int) {
	idxWire1 := 0
	shortestWire1, shortestWire2 := 100000000, 100000000

	for idxWire1 < len(wire1)-1 {
		idxWire2 := 1
		for idxWire2 < len(wire2)-1 {
			found, row, col := calculateIntersection(wire1[idxWire1], wire1[idxWire1+1], wire2[idxWire2], wire2[idxWire2+1])

			if found &&
				WithinSegment(wire1[idxWire1], wire1[idxWire1+1], float64(row), float64(col)) &&
				WithinSegment(wire2[idxWire2], wire2[idxWire2+1], float64(row), float64(col)) {

				currentWire1Dist := wire1Dists[idxWire1] + distFromPoint(wire1[idxWire1], float64(row), float64(col))
				currentWire2Dist := wire2Dists[idxWire2] + distFromPoint(wire2[idxWire2], float64(row), float64(col))
				if currentWire1Dist+currentWire2Dist < shortestWire1+shortestWire2 {
					shortestWire1 = currentWire1Dist
					shortestWire2 = currentWire2Dist
				}
			}

			idxWire2++
		}

		idxWire1++
	}

	return shortestWire1, shortestWire2
}
