package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

var dirs = map[string]Vector{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func solvePuzzle01() {
	input := getInput()
	lines := strings.Split(input, "\n")

	points := []Vector{{0, 0}}

	totalDist := 0
	for _, line := range lines {
		data := strings.Split(line, " ")
		dir := dirs[data[0]]
		dist, _ := strconv.Atoi(data[1])
		totalDist += dist

		point := Vector{
			x: points[len(points)-1].x + dir.x*dist,
			y: points[len(points)-1].y + dir.y*dist,
		}
		points = append(points, point)
	}

	sum := 0
	for i := range points {
		prev := i - 1
		next := i + 1

		if prev < 0 {
			prev = len(points) - 1
		}
		if next >= len(points) {
			next = 0
		}

		sum += points[i].x * (points[prev].y - points[next].y)
	}
	a := abs(sum) / 2
	i := a - totalDist/2 + 1

	total := i + totalDist
	fmt.Printf("Total cube meters: %d\n", total)
}

func main() {
	solvePuzzle01()
}
