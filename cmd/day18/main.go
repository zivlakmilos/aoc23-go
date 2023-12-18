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

var dirs2 = map[byte]Vector{
	'0': {1, 0},
	'1': {0, 1},
	'2': {-1, 0},
	'3': {0, -1},
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func calcArea(points []Vector, totalDist int) int {
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

	return i + totalDist
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

	total := calcArea(points, totalDist)
	fmt.Printf("Total cube meters: %d\n", total)
}

func solvePuzzle02() {
	input := getInput()
	lines := strings.Split(input, "\n")

	points := []Vector{{0, 0}}

	totalDist := 0
	for _, line := range lines {
		data := strings.Split(line, " ")
		dir := dirs2[data[2][len(data[2])-2]]
		dist, _ := strconv.ParseInt(data[2][2:len(data[2])-2], 16, 32)
		totalDist += int(dist)

		point := Vector{
			x: points[len(points)-1].x + dir.x*int(dist),
			y: points[len(points)-1].y + dir.y*int(dist),
		}
		points = append(points, point)
	}

	total := calcArea(points, totalDist)
	fmt.Printf("Total cube meters: %d\n", total)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
