package main

import (
	"fmt"
	"strings"
)

type Galaxy struct {
	x int
	y int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func findExpandedRows(input []string) []bool {
	res := make([]bool, len(input))

	for row := range input {
		foundGalaxy := false
		for col := range input[0] {
			if input[row][col] == '#' {
				foundGalaxy = true
				break
			}
		}
		res[row] = !foundGalaxy
	}

	return res
}

func findExpandedCols(input []string) []bool {
	res := make([]bool, len(input))

	for col := range input[0] {
		foundGalaxy := false
		for row := range input {
			if input[row][col] == '#' {
				foundGalaxy = true
				break
			}
		}
		res[col] = !foundGalaxy
	}

	return res
}

func parseGalaxies(input []string, expandedRows, expandedCols []bool, expandAmount int) []Galaxy {
	res := []Galaxy{}

	x := 0
	y := 0
	for row := range input {
		x = 0
		if expandedRows[row] {
			y += expandAmount
			continue
		}
		for col := range input[row] {
			if expandedCols[col] {
				x += expandAmount
				continue
			}

			if input[row][col] == '#' {
				galaxy := Galaxy{
					x: x,
					y: y,
				}
				res = append(res, galaxy)
			}

			x++
		}
		y++
	}

	return res
}

func calcDistance(a, b Galaxy) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func calcTotalDistance(galaxies []Galaxy) int {
	res := 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			diff := calcDistance(galaxies[i], galaxies[j])
			res += diff
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	lines := strings.Split(input, "\n")

	expandedRows := findExpandedRows(lines)
	expandedCols := findExpandedCols(lines)
	galaxies := parseGalaxies(lines, expandedRows, expandedCols, 2)

	totalDistance := calcTotalDistance(galaxies)

	fmt.Printf("Total distance: %d\n", totalDistance)
}

func solvePuzzle02() {
	input := getInput()
	lines := strings.Split(input, "\n")

	expandedRows := findExpandedRows(lines)
	expandedCols := findExpandedCols(lines)
	galaxies := parseGalaxies(lines, expandedRows, expandedCols, 1000000)

	totalDistance := calcTotalDistance(galaxies)

	fmt.Printf("Total distance: %d\n", totalDistance)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
